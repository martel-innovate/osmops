Local clusters
--------------
> Run OSM GitOps on your box, building everything from scratch.

Build a local OSM release 10 cluster and a separate Kind cluster to
host OSM Ops. Then simulate commits to this repository to watch OSM
Ops create and update an OpenLDAP KNF. All that in the comfort of
your laptop.



## Build & deploy

So we're going to start from scratch and work our way up to run a
fully-fledged OSM GitOps pipeline on your box. Roll up your sleeves,
we'll build two clusters in this section! Here's what our testbed will
look like when we're done:

![Deployment diagram.][dia.depl]

The one cluster is a Kubernetes Kind cluster that hosts FluxCD's own
Source Controller and our OSM Ops service, both in the `flux-system`
namespace. The other cluster is a release 10 OSM cluster configured
with a Kubernetes VIM, all inside a Multipass (QEMU) VM. Both beasts
run on your box and connect through a local bridge network. Source
Controller monitors this very OSM Ops repository you're looking at
on GitHub. The OSM Ops service connects both to Source Controller,
within the Kind cluster, and to the OSM north-bound interface (NBI)
running inside the Multipass VM.


### Before you start...

Clone this repo locally, then install

* Nix - https://nixos.org/guides/install-nix.html
* Docker >= 19.03
* Multipass >= 1.6.2

Keep in mind you're going to need a beefy box to run this demo smoothly.
With lots of effort and patience, I've managed to run it on my 4 core, 8
GB RAM laptop but my guess is that you'd need a box with at least double
that horse power.


### OSM cluster

Spin up an Ubuntu 18.04 VM with Multipass and install OSM release 10
in it:

```bash
$ multipass launch --name osm --cpus 2 --mem 6G --disk 40G 18.04
$ multipass shell osm
% wget https://osm-download.etsi.org/ftp/osm-10.0-ten/install_osm.sh
% chmod +x install_osm.sh
% ./install_osm.sh 2>&1 | tee install.log
% exit
```

Notice I couldn't actually install OSM release 10 because of a few
issues with the installer. By the time you try this, hopefully the
OSM guys will have fixed those bugs and you'll have a smooth ride.
But if it gets bumpy, you can try my patched OSM install scripts by
following the steps in [multipass.install.sh][osm-install].

Once you've got a base OSM 10 cluster up and running, you've got to
configure [KNF infra for an isolated Kubernetes cluster][osm.knf-setup]:

```bash
$ multipass shell osm
% wget https://osm-download.etsi.org/ftp/osm-10.0-ten/install_osm.sh
% osm vim-create --name mylocation1 --user u --password p --tenant p \
    --account_type dummy --auth_url http://localhost/dummy
% osm k8scluster-add cluster --creds .kube/config --vim mylocation1 \
    --k8s-nets '{k8s_net1: null}' --version "v1.15.12" \
    --description="Isolated K8s cluster at mylocation1"
% exit
```

Also, you've got to add some repos where OSM can fetch Helm charts
from:

```bash
$ multipass shell osm
% osm repo-add --type helm-chart --description "Bitnami repo" bitnami https://charts.bitnami.com/bitnami
% osm repo-add --type helm-chart --description "Cetic repo" cetic https://cetic.github.io/helm-charts
% osm repo-add --type helm-chart --description "Elastic repo" elastic https://helm.elastic.co
% exit
```

When done, upload the OSM OpenLDAP packages we're going to use to create
NS instances. To do that, open a terminal in this repo's root dir, then:

```bash
$ cd _tmp/osm-pkgs
$ multipass mount ./ osm:/mnt/osm-pkgs
$ multipass shell osm
% cd /mnt/osm-pkgs
% osm nfpkg-create openldap_knf.tar.gz
% osm nspkg-create openldap_ns.tar.gz
% exit
```

Note down the VM's IPv4 address where the OSM NBI can be accessed:

```bash
$ multipass info osm
```

It should be the first one on the list, the `192.168.*` one.


### Kind cluster

Open a terminal in this repo's root dir, then create a Kind cluster
and deploy Source Controller in it:

```bash
# nix-shell starts a Nix shell with all the tools we're going to need.
$ nix-shell

$ kind create cluster --name dev
$ flux check --pre
$ flux install \
    --namespace=flux-system \
    --network-policy=false \
    --components=source-controller
```

Next, build and deploy OSM Ops. First off, build the Go code, create
a Docker image for the service, and upload it to Kind's own local
Docker registry:

```bash
$ make docker-build
$ kind load docker-image ghcr.io/c0c0n3/osmops:latest --name dev
```

We need to tell OSM Ops how to connect to the OSM NBI. Create an
`nbi-connection.yaml` file with the content below

```yaml
hostname: 192.168.64.19:80
project: admin
user: admin
password: admin
```

but replace `192.168.64.19` with the OSM IP address you noted down
earlier. (The username and password are those of the default OSM
admin user that gets created automatically for you during the OSM
installation.) Since we've got a password there, we'll stash this
config away in a Kubernetes secret:

```bash
$ kubectl -n flux-system create secret generic nbi-connection \
    --from-file nbi-connection.yaml
```

Finally, deploy OSM Ops to the Kind cluster:

```bash
$ kubectl apply -f _deployment_/osmops.deploy.yaml
```

If you open up `osmops.deploy.yaml`, you'll see the OSM Ops service
gets deployed to the same namespace of Source Controller, namely
`flux-system`, and runs under the same account. Also, notice our
secret above becomes available to OSM Ops at

    /etc/osmops/nbi-connection.yaml

More about it later.



## Doing GitOps with OSM

After toiling away at prep steps, we're finally ready for some GitOps
action. In fact, we're going to create and then update an OpenLDAP
KNF through OSM Ops YAML files. Specifically, we'll start off with
an initial Git revision of this file

```yaml
kind: NsInstance
name: ldap
description: Demo LDAP NS instance
nsdName: openldap_ns
vnfName: openldap
vimAccountName: mylocation1
kdu:
  name: ldap
  params:
    replicaCount: "2"
```

which says we want the OSM cluster to run a two-replica OpenLDAP KNF
named `ldap` that OSM should instantiate using the NSD, VNFD and KDU
descriptors found in the OSM packages we installed earlier. Also,
notice the VIM account name is that of the VIM we set up earlier for
the KNF infra. We'll watch OSM Ops process the file and the OSM cluster
end up with a brand new NS instance: an OpenLDAP Kubernetes service
with two pods. Then we'll change the number of replicas in the YAML
to `1`, commit the change to the Git repo and watch OSM Ops trigger
a change in the OSM cluster that'll make the pods scale down to one.
How fun!


### Setting up the OSM Ops pipeline

The OSM Ops show starts as soon as you connect a Git repo through
FluxCD. Out of convenience, we're going to use this very repo on
GitHub starting at tag `test.0.0.5`. In fact, that tagged revision
contains a `_deployment_/kdu/ldap.ops.yaml` file with the YAML above.
Open a terminal in your local repo root dir and create a test Git
source within Flux like so:

```bash
$ nix-shell
$ flux create source git test \
    --url=https://github.com/c0c0n3/source-watcher \
    --tag=test.0.0.5
```

This command creates a Kubernetes GitRepository custom resource. As
soon as Source Controller gets notified of this new custom resource,
it'll fetch the content of `test.0.0.5` and make it available to OSM
Ops which will then realise the deployment config declared in the
YAML file in our local OSM cluster running on Multipass. But how can
OSM Ops know what OSM cluster to connect to? Well, OSM Ops looks for
a YAML config file, `osm_ops_config.yaml`, in the root of the repo
directory tree it gets from Source Controller. At `test.0.0.5`, our
repo has an `osm_ops_config.yaml` in the root dir with this content:

```yaml
targetDir: _deployment_
fileExtensions:
  - .ops.yaml
connectionFile: /etc/osmops/nbi-connection.yaml
```

This configuration tells OSM Ops to get the OSM connection details
from `/etc/osmops/nbi-connection.yaml`. Ha! Remember that Kubernetes
secret mounted on the OSM Ops pod? Yep, that's how it happens! The
other fields tell OSM Ops to look for OSM Ops GitOps files in the
`_deployment_` directory (recursively) and only consider files with
an extension of `.ops.yaml`.


### Watching reconciliation as it happens

Now browse to the OSM Web UI at the Multipass IP address you noted
down earlier (e.g. http://192.168.64.19/) and log in with the OSM
admin user—username: `admin`, password: `admin`. You should be able
to see that OSM is busy creating a new NS instance called `ldap`,
similar to what you see on this screenshot:

![OSM busy creating the OpenLDAP instance.][osm-ui.busy]

Depending on how much horse power your box has, this could take a
while—think minutes. Instead of twiddling your thumbs as you wait,
why not have a look at what's happening under the bonnet? Start a
terminal in the repo root dir and have a look at the OSM Ops service
logs as in the example below:

```bash
$ nix-shell

# figure out the name of the OSM Ops pod, it's the one starting with
# 'source-watcher'.
$ kubectl -n flux-system get pods
NAME                                READY   STATUS    RESTARTS   AGE
source-controller-d58957ccd-pj7p8   1/1     Running   0          7m39s
source-watcher-df9cbc8bf-cjxpq      1/1     Running   0          12s

# then get the logs.
$ kubectl -n flux-system logs source-watcher-df9cbc8bf-cjxpq
```

What you see in the logs should be similar to

```bash
2021-10-06T16:39:25.179Z	INFO	controller.gitrepository	New revision detected	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system", "revision": "test.0.0.5/59cc9586c318642d9fd2399fa638adb24649d53c"}
2021-10-06T16:39:25.742Z	INFO	controller.gitrepository	Extracted tarball into /tmp/test120670538: 132 files, 38 dirs (362.0482ms)	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system"}
2021-10-06T16:39:26.140Z	INFO	controller.gitrepository	processing	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system", "file": "/tmp/test120670538/_deployment_/kdu/ldap.ops.yaml"}
```

In plain English: you should be able to see OSM Ops detect a new Git
revision of `test.0.0.5`, download its content and then process the
`ldap.ops.yaml` file. To see what's happening in OSM land, shell into
the OSM VM and you should see two pods being created for the OpenLDAP
service:

```bash
$ multipass shell osm
% kubectl get ns
#              ^ pick the one that looks like an UUID
% kubectl -n fada443a-905c-4241-8a33-4dcdbdac55e7 get pods
NAME                                                READY   STATUS              RESTARTS   AGE
stable-openldap-1-2-7-0046589243-6f9f8b8f6d-n9bz2   0/1     ContainerCreating   0          30s
stable-openldap-1-2-7-0046589243-6f9f8b8f6d-x2mmd   0/1     ContainerCreating   0          30s
```

Then some time later the two pods should be fully operational

```bash
% kubectl -n fada443a-905c-4241-8a33-4dcdbdac55e7 get pods
NAME                                                READY   STATUS    RESTARTS   AGE
stable-openldap-1-2-7-0046589243-6f9f8b8f6d-n9bz2   1/1     Running   0          109s
stable-openldap-1-2-7-0046589243-6f9f8b8f6d-x2mmd   1/1     Running   0          109s
```

and eventually that should be reflected in the OSM Web UI too, as in
the screenshot below.

![OSM done creating the OpenLDAP instance.][osm-ui.done]

**Important**. Wait until the two Kubernetes pods are up and running
and the deployment state got updated in the OSM Web UI before moving
on to the next step.


### Updating the deployment configuration

As promised, we should change the number of replicas in `ldap.ops.yaml`
to `1`, commit the change to the Git repo and watch OSM Ops trigger
a change in the OSM cluster that'll make the pods scale down to one.
But there's a snag: you can't actually commit to this repo. Stop
jeering, we've got a workaround :-) We can manually force FluxCD to
fetch the revision tagged `test.0.0.6` which has the same content
of `test.0.0.5` except for `ldap.ops.yaml` where the number of replicas
is `1` instead of two—[`test.0.0.6` v `test.0.0.5` diff here][repo.tags-diff].
Open a terminal in your local repo root dir and run:

```bash
$ nix-shell
$ flux create source git test \
    --url=https://github.com/c0c0n3/source-watcher \
    --tag=test.0.0.6
```

If you then look at the OSM Ops service logs

```bash
$ kubectl -n flux-system logs source-watcher-df9cbc8bf-cjxpq
```

you should be able to spot OSM Ops process the contents of `test.0.0.6`

```bash
2021-10-06T17:03:45.006Z	INFO	controller.gitrepository	New revision detected	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system", "revision": "test.0.0.6/37ec18d984e7b0e4e0de98ec0061b955c413e4ef"}
2021-10-06T17:03:45.293Z	INFO	controller.gitrepository	Extracted tarball into /tmp/test535104545: 132 files, 38 dirs (126.3411ms)	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system"}
2021-10-06T17:03:45.326Z	INFO	controller.gitrepository	processing	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system", "file": "/tmp/test535104545/_deployment_/kdu/ldap.ops.yaml"}
```

Meanwhile in OSM land...

```bash
$ multipass shell osm
% kubectl get ns
#              ^ pick the one that looks like an UUID
% kubectl -n fada443a-905c-4241-8a33-4dcdbdac55e7 get pods
NAME                                                READY   STATUS        RESTARTS   AGE
stable-openldap-1-2-7-0046589243-6f9f8b8f6d-n9bz2   1/1     Terminating   0          15m
stable-openldap-1-2-7-0046589243-6f9f8b8f6d-x2mmd   1/1     Running       0          15m
```

one of the OpenLDAP pods should get shut down. Eventually the OSM UI
should reflect your NS instance for the OpenLDAP service got scaled
down to one Kubernetes pod. Be patient, unless you've got a beefy
box, this too will take a while. If you then take a look at the NS
operation history for `ldap`, you should see two entries there, one
for the create and the other for the update like in this screenshot:

![OpenLDAP operations history in OSM.][osm-ui.history]



## Clean up

Get rid of the Kind cluster with OSM Ops and Source Controller

```bash
$ kind delete cluster --name dev
```

Zap the Multipass OSM VM

```bash
$ multipass stop osm
$ multipass delete osm
$ multipass purge
```




[dia.depl]: ./demo.local-clusters.png
[osm.knf-setup]: https://osm.etsi.org/docs/user-guide/05-osm-usage.html#adding-kubernetes-cluster-to-osm
[osm-install]: ../../_tmp/osm-install/multipass.install.sh
[osm-ui.busy]: ./osm-ui.1.png
[osm-ui.done]: ./osm-ui.2.png
[osm-ui.history]: ./osm-ui.3.png
[repo.tags-diff]: https://github.com/c0c0n3/source-watcher/compare/test.0.0.5...test.0.0.6

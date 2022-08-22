Affordable5G Malaga Demo
------------------------
> Showcasing OSM Ops at the Affordable5G official review.

Use OSM Ops to set up a GitOps pipeline in the Affordable5G Malaga
environment. Then connect a GitHub repo and watch OSM Ops create a
Nemergent KNF in the Malaga cluster.


## Setup

We'll run our demo in the Affordable5G virtual environment in Malaga.
The virtual servers there are already set up with most of the bits
and bobs we're going to need but we still have to roll out our own
OSM Ops stuff before we can demo anything. Here's what the Malaga
environment will look like when we're done:

![Deployment diagram.][dia.depl]

As you can see the lay of the land is conceptually similar to that
of the [Local clusters demo][demo.local] we know and love. OSM sits
in its own box aptly called `osm` with an IP of `10.11.23.249` and
is configured with a VIM pointing to a Kubernetes cluster (MicroK8s
flavour) made up of two nodes, creatively called `node1` and `node2`,
with IPs of `10.11.23.96` and `10.11.23.97`, respectively. The Kubernetes
cluster hosts FluxCD's own Source Controller and our OSM Ops service,
both in the `flux-system` namespace. Source Controller monitors an
OSM demo repository on GitHub. The OSM Ops service connects both to
Source Controller, within the same cluster, and to the OSM north-bound
interface (NBI) running on `10.11.23.249`, outside the Kubernetes cluster.


### Before you start...

We're not going to build or install anything on your box (phew!),
all the action will take place in the Malaga environment. To be able
to do stuff with the Malaga boxes from your machine, you've got to
set up a VPN. We use OpenVPN with this config file

* pfSense-TCP4-9443-MARTEL2_Affordable5G-config.ovpn

but your set up could be different. Also, if you've cloned this repo
locally and have Nix, there's no need to install OpenVPN. Just `cd`
into your local repo root directory and run

```bash
$ nix-shell
$ sudo openvpn pfSense-TCP4-9443-MARTEL2_Affordable5G-config.ovpn
#              ^ replace w/ your own OpenVPN config file
```

You'll need the VPN tunnel to be on to be able to run the commands
in the rest of this document. So keep this terminal open and leave
OpenVPN run in the foreground until we're done.


### Kubernetes cluster

So here's the good news: the Malaga environment comes with a two-node
Kubernetes cluster pre-configured for the Affordable5G demo. Specifically,
there's a MicroK8s (version `1.21.5`) cluster made up of the two boxes
we mentioned earlier—`node1` (`10.11.23.96`) and `node2` (`10.11.23.97`).

But we still need to take care of our own stuff:

* install and configure the FluxCD CLI on `node2`;
* deploy FluxCD and OSM Ops services to the Kubernetes cluster;
* configure OSM Ops.

So here goes! SSH into `node2`

```bash
$ ssh node2@10.11.23.97
```

and install Nix

```bash
$ curl -L https://nixos.org/nix/install | sh
$ . /home/node2/.nix-profile/etc/profile.d/nix.sh
```

Then download the OSM Ops demo bundle and use it to start a Nix shell
with the tools we'll need for the show

```bash
$ wget https://github.com/c0c0n3/osmops.demo/archive/refs/tags/a5g-0.1.0.tar.gz
$ tar xzf a5g-0.1.0.tar.gz
$ cd osmops.demo-a5g-0.1.0
$ nix-shell
```

Now there's a snag. The FluxCD command (`flux`) won't work with the
`kubectl` version installed on `node2` and it knows zilch about MicroK8s,
so it can't obviously run `microk8s kubectl` instead of plain `kubectl`.
(See [this blog post][flux-mk8s] about it.) But the Nix shell packs
a `kubectl` version compatible with `flux`, so all we need to do is
make plain `kubectl` use the same config as `microk8s kubectl`.

```bash
$ mkdir -p ~/.kube
$ ln -s /var/snap/microk8s/current/credentials/client.config ~/.kube/config
```

With this little hack in place, we can deploy Source Controller

```bash
$ flux check --pre
$ flux install \
    --namespace=flux-system \
    --network-policy=false \
    --components=source-controller
```

Next up, our very own OSM Ops. First off, we need to tell OSM Ops how
to connect to the OSM NBI running on the `osm` box (`10.11.23.249`).
Create an `nbi-connection.yaml` file

```bash
$ nano nbi-connection.yaml
```

with the content below

```yaml
hostname: 10.11.23.249:80
project: admin
user: admin
password: admin
```

Since we've got a password there, we'll stash this config away in a
Kubernetes secret:

```bash
$ kubectl -n flux-system create secret generic nbi-connection \
    --from-file nbi-connection.yaml
```

Finally, deploy the OSM Ops service to the Kubernetes cluster

```bash
$ kubectl apply -f deployment/osmops.deploy.yaml
```

If you open up `osmops.deploy.yaml`, you'll see the OSM Ops service
gets deployed to the same namespace of Source Controller, namely
`flux-system`, and runs under the same account. Also, notice our
secret above becomes available to OSM Ops at

    /etc/osmops/nbi-connection.yaml

More about it later.



### OSM cluster

The OSM cluster has already been set up for us, yay! In fact, the `osm`
node (`10.11.23.249`) hosts a fully-fledged OSM Release 10 instance
configured with a VIM account called `dummyvim` that's tied to the
Kubernetes (MicroK8s) cluster. Also, the OSM config includes the Helm
chart repos below:

- https://charts.bitnami.com/bitnami
- https://cetic.github.io/helm-charts
- https://helm.elastic.co
- http://osm-download.etsi.org/ftp/Packages/vnf-onboarding-tf/helm/
- https://pencinarsanz-atos.github.io/nemergent-chart/

The last one is actually the only one we care about for this demo
since it hosts the Helm chart for the Nemergent services we're going
to deploy through our GitOps pipeline. To create NS instances from
that chart there have to be an NSD and VNFD in OSM. That's been done
for us too. In fact, there's an NSD called `affordable_nsd` pointing
to a VNFD called `affordable_vnfd`. The VNFD declares a KDU (name:
`nemergent`) referencing the above Helm repo.



## Doing GitOps with OSM

After toiling away at prep steps, we're finally ready for some GitOps
action. In fact, we're going to create a Nemergent KNF through OSM Ops
YAML files. Specifically, we'll fetch this YAML from a Git repo

```yaml
kind: NsInstance
name: nemergent
description: Demo Nemergent NS instance
nsdName: affordable_nsd
vnfName: affordable_vnfd
vimAccountName: dummyvim
kdu:
  name: nemergent
```

This OSM Ops deployment descriptor says we want the OSM cluster to
run a Nemergent KNF called `nemergent` that OSM should instantiate
using the Nemergent NSD, VNFD and KDU descriptors mentioned earlier.
Also, notice the VIM account name is that of the VIM connected to the
Kubernetes cluster. We'll watch OSM Ops process the file and the OSM
cluster end up with a brand new NS instance: a Nemergent Kubernetes
stateful set with 14 services each running a single pod.

Ideally we'd demo an update too. That is, show how updating the file
in the Git repo eventually results in a corresponding update to the
cluster state. Unfortunately at the moment the Nemergent Helm chart
doesn't have any KDU params we can tweak, so we can't do the update.
But you can still have a look at the [Local clusters demo][demo.local]
to see how updates work.


### Setting up the OSM Ops pipeline

The OSM Ops show starts as soon as you connect a Git repo through
FluxCD. We published a [repo on GitHub][osmops.demo] that we'll use
for this demo. In fact, we'll start off with the content of the repo
at tag [a5g-0.1.0][a5g-0.1.0]. This Git version contains a `nemergent.ops.yaml`
file in the `deployment/kdu` directory with the YAML above. So go
back to your SSH terminal on `node2` and create a `osmops.demo` Git
source within Flux like so:

```bash
$ flux create source git osmops.demo \
    --url=https://github.com/c0c0n3/osmops.demo \
    --tag=a5g-0.1.0
```

This command creates a Kubernetes GitRepository custom resource. As
soon as Source Controller gets notified of this new custom resource,
it'll fetch the content of `a5g-0.1.0` and make it available to OSM
Ops which will then realise the deployment config declared in the
YAML file in the OSM-managed Kubernetes cluster. But how can OSM Ops
know what OSM cluster to connect to? Well, OSM Ops looks for a YAML
config file, `osm_ops_config.yaml`, in the root of the repo directory
tree it gets from Source Controller. At `a5g-0.1.0`, our demo repo
has an `osm_ops_config.yaml` in the root dir with this content:

```yaml
targetDir: deployment
fileExtensions:
  - .ops.yaml
connectionFile: /etc/osmops/nbi-connection.yaml
```

This configuration tells OSM Ops to get the OSM connection details
from `/etc/osmops/nbi-connection.yaml`. Ha! Remember that Kubernetes
secret mounted on the OSM Ops pod? Yep, that's how it happens! The
other fields tell OSM Ops to look for OSM Ops GitOps files in the
`deployment` directory (recursively) and only consider files with
an extension of `.ops.yaml`.


### Watching reconciliation as it happens

Now browse to the OSM Web UI at http://10.11.23.249/ and log in with
the OSM admin user—username: `admin`, password: `admin`. You should
be able to see that OSM is busy creating a new NS instance called
`nemergent`, similar to what you see on this screenshot:

![OSM busy creating the Nemergent instance.][osm-ui.busy]

This could take a little while. Instead of twiddling your thumbs as
you wait, why not have a look at what's happening under the bonnet?
Go back to your SSH terminal on `node2` and have a look at the OSM
Ops service logs as in the example below:

```bash
# figure out the name of the OSM Ops pod, it's the one starting with
# 'source-watcher'.
$ kubectl -n flux-system get pods
NAME                                READY   STATUS    RESTARTS   AGE
source-controller-d58957ccd-p5994   1/1     Running   0          2d23h
source-watcher-5494d664d5-v66rf     1/1     Running   0          5h27m

# then get the logs.
$ kubectl -n flux-system logs source-watcher-5494d664d5-v66rf
```

What you see in the logs should be similar to

```bash
2021-11-01T14:58:25.770Z	INFO	controller.gitrepository	New revision detected	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "osmops.demo", "namespace": "flux-system", "revision": "a5g-0.1.0/019aefa83f185700ad5c8e11bfd5d91599a5b39a"}
2021-11-01T14:58:25.772Z	INFO	controller.gitrepository	Extracted tarball into /tmp/osmops.demo875847309: 5 files, 3 dirs (546.733µs)	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "osmops.demo", "namespace": "flux-system"}
2021-11-01T14:58:25.773Z	INFO	controller.gitrepository	processing	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "osmops.demo", "namespace": "flux-system", "file": "/tmp/osmops.demo875847309/deployment/kdu/nemergent.ops.yaml"}
```

In plain English: you should be able to see OSM Ops detect a new Git
revision of `a5g-0.1.0`, download its content and then process the
`nemergent.ops.yaml` file. If you look at the Kubernetes pods in the
OSM namespace, you should be able to see 14 pods being created for
the Nemergent NS instance:

```bash
$ kubectl get ns
#              ^ pick the one that looks like an UUID
$ kubectl -n 2b091f50-0555-4296-afe8-d825cc2b19f6 get pods
NAME           READY   STATUS              RESTARTS   AGE
scscf-0        0/1     ContainerCreating   0          11s
http-proxy-0   0/1     ContainerCreating   0          11s
pcscf-0        0/1     ContainerCreating   0          11s
rtp-engine-0   0/1     ContainerCreating   0          11s
pas-0          0/1     ContainerCreating   0          12s
idms-0         0/1     ContainerCreating   0          11s
hss-0          0/1     ContainerCreating   0          10s
icscf-0        0/1     ContainerCreating   0          10s
db-0           0/1     ContainerCreating   0          10s
cas-0          0/1     ContainerCreating   0          11s
cms-0          0/1     ContainerCreating   0          10s
kms-0          0/1     ContainerCreating   0          10s
redis-0        0/1     ContainerCreating   0          10s
enabler-ws-0   0/1     ContainerCreating   0          11s
```

Then some time later the 14 pods should be fully operational

```bash
kubectl -n 2b091f50-0555-4296-afe8-d825cc2b19f6 get pods
NAME           READY   STATUS    RESTARTS   AGE
scscf-0        1/1     Running   0          38s
http-proxy-0   1/1     Running   0          38s
pcscf-0        1/1     Running   0          38s
rtp-engine-0   1/1     Running   0          38s
pas-0          1/1     Running   0          38s
idms-0         1/1     Running   0          38s
hss-0          1/1     Running   0          37s
icscf-0        1/1     Running   0          38s
db-0           1/1     Running   0          37s
cas-0          1/1     Running   0          38s
cms-0          1/1     Running   0          37s
kms-0          1/1     Running   0          38s
redis-0        1/1     Running   0          37s
enabler-ws-0   1/1     Running   0          38s
```

and eventually that should be reflected in the OSM Web UI too, as in
the screenshot below.

![OSM done creating the Nemergent instance.][osm-ui.done]



## Clean up

Use the OSM UI (http://10.11.23.249/) to zap the Nemergent NS instance.
Then go back to your SSH terminal on `node2` and run

```bash
$ kubectl delete -f deployment/osmops.deploy.yaml
$ kubectl -n flux-system delete secret nbi-connection
$ flux uninstall --namespace=flux-system
$ cd ~
$ rm -rf osmops.demo-a5g-0.1.0
$ rm .kube/config
```

Finally, don't forget to kill the OpenVPN process we started at the
beginning of the demo, otherwise your box will stay connected to the
Malaga environment through a VPN tunnel.




[a5g-0.1.0]: https://github.com/c0c0n3/osmops.demo/tree/a5g-0.1.0
[dia.depl]: ./demo.malaga.png
[demo.local]: ./local-clusters.md
[flux-mk8s]: https://boxofcables.dev/using-flux2-with-microk8s/
[osmops.demo]: https://github.com/c0c0n3/osmops.demo
[osm-ui.busy]: ./malaga.osm-ui.1.png
[osm-ui.done]: ./malaga.osm-ui.2.png

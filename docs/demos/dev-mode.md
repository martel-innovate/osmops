Dev mode
--------
> Run OSM GitOps on your box, building everything from scratch.

Build a local OSM release 10 cluster and a separate Kind cluster to
host FluxCD. Build and run OSM Ops directly on your box. Then simulate
commits to this repository to watch OSM Ops create and update an
OpenLDAP KNF.

Notice this demo is almost the same as the [Local clusters][demo.local]
demo, except OSM Ops now runs directly on your box, outside the Kind
cluster. This setup comes in handy if you want to easily debug OSM
Ops to see what it's actually up to.



## Build & deploy

As in the [Local clusters][demo.local] demo, we'll build one Multipass
VM hosting an OSM release 10 cluster configured with a Kubernetes VIM.
We'll also build a Kind cluster, but unlike the [Local clusters][demo.local]
demo, the cluster will only host Source Controller. We'll run OSM
Ops directly on localhost and connect it to Source Controller through
a proxy.


### Before you start...

Same requirements as in the [Local clusters][demo.local] demo.


### OSM cluster

Build and run it just like in the [Local clusters][demo.local] demo.


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

Then have `kubeclt` proxy calls to port `8181` on localhost to the
Source Controller service inside the cluster:

```bash
$ kubectl -n flux-system port-forward svc/source-controller 8181:80
```

Keep this terminal open since the proxy process will have to be up
for the entire duration of the demo. In fact, OSM Ops will run on
localhost and connect to port `8181` to talk to Source Controller.


### OSM Ops

We'll run OSM Ops outside the cluster. Open a new terminal in the
repo root dir and run

```bash
$ nix-shell
$ export SOURCE_HOST=localhost:8181
$ make run
```

The `SOURCE_HOST` environment variable tells OSM Ops to connect to
Source Controller on localhost at port `8181`. Keep this terminal
open since the OSM Ops process will have to be up for the entire
duration of the demo.

Similar to the [Local clusters][demo.local] demo, we're going to have
OSM Ops create and update an NS instance (OpenLDAP KNF) by looking
at the OSM GitOps files in this repo's `_deployment_` dir, but at
tags `test.0.0.3` and `test.0.0.4` instead of `test.0.0.5` and `test.0.0.6`.
The `osm_ops_config.yaml` file is the same for both `test.0.0.3` and
`test.0.0.4` and points to an NBI connection file sitting on the same
box where OSM Ops runs: `/tmp/osm_ops_secret.yaml`. You need to
create this file on your box's `/tmp` directory with the following
content:

```yaml
hostname: 192.168.64.19:80
project: admin
user: admin
password: admin
```

but replace `192.168.64.19` with the OSM IP address you noted down
earlier.



## Doing GitOps with OSM

The workflow here is basically the same as in the [Local clusters][demo.local]
demo. In fact, we're going to create and then update the same OpenLDAP
KNF through OSM Ops YAML files, except we'll use the files at tags
`test.0.0.3` and `test.0.0.4` instead of `test.0.0.5` and `test.0.0.6`.


### Setting up the OSM Ops pipeline

Open yet another terminal in the repo root dir, then create a test
GitHub source within Flux

```bash
$ nix-shell
$ flux create source git test \
    --url=https://github.com/c0c0n3/source-watcher \
    --tag=test.0.0.3
```

This command creates a Kubernetes GitRepository custom resource. As
soon as Source Controller gets notified of this new custom resource,
it'll fetch the content of `test.0.0.3` and make it available to OSM
Ops which will then realise the deployment config declared in the
YAML file in our local OSM cluster running on Multipass.

How does OSM Ops find Source Controller? Like we saw earlier OSM Ops
got configured to connect to Source Controller on localhost at port
`8181`. What about OSM? As explained earlier, both tags `test.0.0.3`
and `test.0.0.4` come with an `osm_ops_config.yaml` file that says
the NBI connection file sits on localhost: `/tmp/osm_ops_secret.yaml`.


### Watching reconciliation as it happens

Now if you switch back to the terminal running OSM Ops, you should
be able to see it processing the files in the `_deployment_` dir as
it was at tag `test.0.0.3`. It should call OSM NBI to create an NS
instance using the OSM OpenLDAP package we uploaded earlier with two
replicas as specified in the `ldap.ops.yaml` in `_deployment_/kdu`.
It's going to take a while for the deployment state to reflect in
the OSM Web UI, but you can check what's going on under the bonnet
by shelling into the OSM VM

```bash
$ multipass shell osm
% kubectl get ns
#              ^ pick the one that looks like an UUID
% kubectl -n fada443a-905c-4241-8a33-4dcdbdac55e7 get pods
# ... you should see two pods being created for the OpenLDAP service
```

**Important**. Wait until the two Kubernetes pods are up and running
and the deployment state got updated in the OSM Web UI before moving
on to the next step.


### Updating the deployment configuration

Just like in the [Local clusters][demo.local] demo, we're going to
simulate a commit to change the number replicas in `ldap.ops.yaml`
to `1` and watch OSM Ops trigger a change in the OSM cluster that'll
make the pods scale down to one. To simulate that commit, just make
Flux switch to tag `test.0.0.4`

```bash
$ flux create source git test \
    --url=https://github.com/c0c0n3/source-watcher \
    --tag=test.0.0.4
```

The content of `ldap.ops.yaml` at tag `test.0.0.4` is the same as that
of tag `test.0.0.3` except for the replica count which is `1`. So you
should see that eventually your NS instance for the OpenLDAP service
gets scaled down to one Kubernetes pod. Be patient, unless you've got
a beefy box, this too will take a while.



### Clean up

Kill all the processes running in your terminals, then zap the two
clusters as explained in the clean up of the [Local clusters][demo.local]
demo.




[demo.local]: ./local-clusters.md

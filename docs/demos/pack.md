Packaging
---------
> Have OSM Ops manage your OSM packages too!

Build a local OSM release 10 cluster and a separate Kind cluster to
host OSM Ops. Then simulate commits to this repository to watch OSM
Ops create OpenLDAP KNF & NS packages, instantiate an OpenLDAP KNF,
and finally update the OpenLDAP KNF & NS packages. It all happens
on your laptop!

Notice this demo builds on the [Local clusters][demo.local] demo to
showcase how OSM Ops can also create and update OSM packages from
sources in a git repo---which, for this demo, is the one you're
reading this page from :-)



## Build & deploy

As in the [Local clusters][demo.local] demo, we'll build one Multipass
VM hosting an OSM release 10 cluster configured with a Kubernetes VIM
and a Kind cluster to host OSM Ops and FluxCD. Have a look at the diagram
and explanation there to get a handle on the lay of the land.


### Before you start...

Same requirements as in the [Local clusters][demo.local] demo.


### OSM cluster

Follow the steps in the [Local clusters][demo.local] demo up to where
it says to run OSM client to upload the two OSM packages---i.e.
`openldap_knf.tar.gz` and `openldap_ns.tar.gz`. **Skip that part**
where it asks you to upload the packages. In fact, we'll make OSM
Ops create and upload those two packages for us.


### Kind cluster

Build and run it just like in the [Local clusters][demo.local] demo.



## Doing GitOps with OSM

After toiling away at prep steps, we're finally ready for some GitOps
action. In fact, we're going make OSM Ops create OpenLDAP KNF & NS
packages, instantiate an OpenLDAP KNF, and finally update the OpenLDAP
KNF & NS packages.

Specifically, we'll start off with the deployment configuration in
this repo at tag [test.0.0.7][test.0.0.7]. At this tag, the repo
contains a [deployment directory][test.0.0.7.deploy] with

* An [OpenLDAP KNF package source][test.0.0.7.knf].
* An [OpenLDAP NS package source][test.0.0.7.ns].
* An [OSM Ops YAML file][test.0.0.7.kdu] requesting the OSM cluster
  to run an OpenLDAP KNF instantiated from NSD, VNFD and KDU descriptors
  found in the above packages.

Read the [Local clusters][demo.local] demo's section about GitOps
for an explanation of the OSM Ops YAML file. The way OSM Ops handles
packages is a bit more involved, you can [read here about it][docs.pkgs],
but for the purpose of this demo all you need to know is that OSM
Ops can create or update OSM packages from source directories you
keep in your GitOps repo. Each source directory contains the files
you'd normally use to make an OSM package tarball, except for the
`checksums.txt` file which OSM Ops generates for you when making
the tarball.

On processing the repo at tag `test.0.0.7`, OSM Ops will create the
OpenLDAP KNF and NS packages in OSM, then make OSM instantiate the
OpenLDAP KNF using the data in the packages just created in OSM.

After that, we'll simulate a commit to the repo by switching over
to tag `test.0.0.8`. The only changes between tag `test.0.0.7` and
`test.0.0.8` are updated version numbers for the source packages,
from `1.0` to `1.1`, as shown in [this diff][tag-diff]. OSM Ops will
pick up the changes and update both packages in OSM.


### Setting up the OSM Ops pipeline

The OSM Ops show starts as soon as you connect a Git repo through
FluxCD. As mentioned earlier, we're going to use this very repo on
GitHub at tag `test.0.0.7`. Open a terminal in your local repo root
dir and create a test Git source within Flux like so:

```bash
$ nix-shell
$ flux create source git test \
    --url=https://github.com/c0c0n3/source-watcher \
    --tag=test.0.0.7
```

This command creates a Kubernetes GitRepository custom resource. As
soon as Source Controller gets notified of this new custom resource,
it'll fetch the content of `test.0.0.7` and make it available to OSM
Ops which will then realise the deployment config in our local OSM
cluster running on Multipass. As explained in the [Local clusters][demo.local]
demo, OSM Ops figures out which OSM cluster to connect to by reading
the `osm_ops_config.yaml` file in the root of the repo directory tree
it gets from Source Controller. At `test.0.0.7`, the content of that
file is

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
an extension of `.ops.yaml`. As for OSM package sources, OSM Ops
looks for them in the `osm-pkgs` dir beneath the target dir, which
in our case is: `_deployment_/osm-pkgs`.


### Watching reconciliation as it happens

Now browse to the OSM Web UI at the Multipass IP address you noted
down earlier (e.g. http://192.168.64.19/) and log in with the OSM
admin user—username: `admin`, password: `admin`. You should be able
to see that OSM now has both an OpenLDAP KNF and NS package and is
busy creating a new NS instance called `ldap`, similar to what you
see on the screenshot in the [Local clusters][demo.local] demo. And
as in the [Local clusters][demo.local] demo, if you grab the OSM Ops
logs, you should see what OSM Ops did. The log file should contain
entries similar to the ones below

```log
2022-06-21T18:34:45.368Z	INFO	controller.gitrepository	New revision detected	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system", "revision": "test.0.0.7/d3a8cbf812447c05cf44814db40f6c6da86ab49f"}
2022-06-21T18:34:45.951Z	INFO	controller.gitrepository	Extracted tarball into /tmp/test988276587: 233 files, 81 dirs (399.195386ms)	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system"}
2022-06-21T18:34:45.997Z	INFO	controller.gitrepository	processing	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system", "osm package": "/tmp/test988276587/_deployment_/osm-pkgs/openldap_knf"}
2022-06-21T18:35:05.808Z	INFO	controller.gitrepository	processing	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system", "osm package": "/tmp/test988276587/_deployment_/osm-pkgs/openldap_ns"}
2022-06-21T18:35:13.872Z	INFO	controller.gitrepository	processing	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system", "file": "/tmp/test988276587/_deployment_/kdu/ldap.ops.yaml"}
```


### Updating the deployment configuration

Now we should make some changes to the source packages in the repo
to see OSM Ops update the packages in OSM. As in the Local clusters
demo, we'll use a shortcut: manually force FluxCD to fetch tag `test.0.0.8`
which has the same content as `test.0.0.7` except for the package
versions which are `1.1`. Again, have a look at the [diff between
these two tags][tag-diff]. Open a terminal in your local repo root
dir and run:

```bash
$ nix-shell
$ flux create source git test \
    --url=https://github.com/c0c0n3/source-watcher \
    --tag=test.0.0.8
```

If you then look at the OSM Ops service logs

```log
2022-06-21T18:50:42.261Z	INFO	controller.gitrepository	New revision detected	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system", "revision": "test.0.0.8/2a8ee4439b06d4ac94c64ec187e88f619d6a97d1"}
2022-06-21T18:50:42.415Z	INFO	controller.gitrepository	Extracted tarball into /tmp/test805307342: 233 files, 81 dirs (115.131087ms)	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system"}
2022-06-21T18:50:42.420Z	INFO	controller.gitrepository	processing	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system", "osm package": "/tmp/test805307342/_deployment_/osm-pkgs/openldap_knf"}
2022-06-21T18:50:42.927Z	INFO	controller.gitrepository	processing	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system", "osm package": "/tmp/test805307342/_deployment_/osm-pkgs/openldap_ns"}
2022-06-21T18:50:43.921Z	INFO	controller.gitrepository	processing	{"reconciler group": "source.toolkit.fluxcd.io", "reconciler kind": "GitRepository", "name": "test", "namespace": "flux-system", "file": "/tmp/test805307342/_deployment_/kdu/ldap.ops.yaml"}
```

you should be able to see OSM Ops having processed tag `test.0.0.8`.
In particular the two source packages OpenLDAP KNF and OpenLDAP NS.
Now if you go back to the OSM Web UI and navigate to the NS Packages
page, you should see the OpenLDAP NS package has now version `1.1`,
i.e. exactly what was in the YAML source file in our repo. Likewise,
if you navigate the the VNF Packages page, you should be able to see
the OpenLDAP KNF package's version is now `1.1` too.



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




[demo.local]: ./local-clusters.md
[docs.pkgs]: ../osm-pkgs.md
[tag-diff]: https://github.com/c0c0n3/source-watcher/compare/test.0.0.7...c0c0n3:test.0.0.8
[test.0.0.7]: https://github.com/c0c0n3/source-watcher/tree/test.0.0.7
[test.0.0.7.deploy]: https://github.com/c0c0n3/source-watcher/tree/test.0.0.7/_deployment_
[test.0.0.7.kdu]: https://github.com/c0c0n3/source-watcher/blob/test.0.0.7/_deployment_/kdu/ldap.ops.yaml
[test.0.0.7.knf]: https://github.com/c0c0n3/source-watcher/tree/test.0.0.7/_deployment_/osm-pkgs/openldap_knf
[test.0.0.7.ns]: https://github.com/c0c0n3/source-watcher/tree/test.0.0.7/_deployment_/osm-pkgs/openldap_ns

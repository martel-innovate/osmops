LCM fails
---------
> epic!

So building & deploying a custom OSM LCM image has been lots of fun!
Here's what didn't work and possible workarounds.


### OSM 11 VM

Had to build it a couple of times. Some of the install script tasks
[failed][failed-osm-install] but the script went ahead. Eventually
I ended up with a broken OSM install in my hands---some OSM services
didn't get deployed to the K8s cluster. Not sure what the cause of
those random failure is, possibly lack of enough compute resources
and the install procedure not being robust enough to cater for slow
boxes?


### LCM build failures - part 1

If you bump into this lovely error about Docker permissions when
building the LCM image artifacts

```console
% devops/tools/local-build.sh --module common,IM,N2VC,LCM,NBI stage-2
Performing Stage 2
Building common
Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Post "http://%2Fvar%2Frun%2Fdocker.sock/v1.24/build?buildargs=%7B%7D&cachefrom=%5B%5D&cgroupparent=&cpuperiod=0&cpuquota=0&cpusetcpus=&cpusetmems=&cpushares=0&dockerfile=Dockerfile&labels=%7B%7D&memory=0&memswap=0&networkmode=default&rm=1&shmsize=0&t=common-stage2&target=&ulimits=null&version=1": dial unix /var/run/docker.sock: connect: permission denied
docker: Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Post "http://%2Fvar%2Frun%2Fdocker.sock/v1.24/containers/create": dial unix /var/run/docker.sock: connect: permission denied.
See 'docker run --help'.
Failed to build common
```

It could be because you've got to log out the VM after installing
OSM. To be on the safe side, it's actually best to shut down and
then restart the VM after installation.


### LCM build failures - part 2

So going past the first hurdle was relatively easy. But then the
command to build the LCM image artifacts took about 50 mins and
I didn't get a clean build in the end:

```console
% devops/tools/local-build.sh --module common,IM,N2VC,RO,LCM,NBI stage-2
...
dpkg-deb: building package 'python3-n2vc' in '../python3-n2vc_11.0.0rc1.post36+g23c4455-1_all.deb'.
 dpkg-genbuildinfo
 dpkg-genchanges  >../n2vc_11.0.0rc1.post36+g23c4455-1_amd64.changes
dpkg-genchanges: info: including full source code in upload
 dpkg-source --after-build .
dpkg-source: info: using options from n2vc-11.0.0rc1.post36+g23c4455/debian/source/options: --extend-diff-ignore=\.egg-info$
dpkg-buildpackage: info: full upload (original source is included)
dist run-test: commands[3] | sh -c 'rm n2vc/requirements.txt'
____________________________________________________________ summary _____________________________________________________________
  dist: commands succeeded
  congratulations :)
renamed './deb_dist/python3-n2vc_11.0.0rc1.post36+g23c4455-1_all.deb' -> '/home/ubuntu/snap/qhttp/common/python3-n2vc_11.0.0rc1.post36+g23c4455-1_all.deb'
Directory /home/ubuntu/workspace/RO does not exist
% echo $?
1
```

Going ahead to the next step anyway, just in case the build failure
wasn't critical...

```console
% devops/tools/local-build.sh --module LCM stage-3
...
Step 14/57 : RUN curl $PYTHON3_OSM_LCM_URL -o osm_lcm.deb
 ---> Running in 92e6b11d10dc
curl: no URL specified!
curl: try 'curl --help' or 'curl --manual' for more information
The command '/bin/sh -c curl $PYTHON3_OSM_LCM_URL -o osm_lcm.deb' returned a non-zero code: 2
Failed to build lcm
```

Oh deary, deary. Maybe I shouldn't have gone ahead.


### LCM build failures - part 3

So it turns out the reason for this error message

> Directory /home/ubuntu/workspace/RO does not exist

is that the command

```console
% devops/tools/local-build.sh --module common,IM,N2VC,RO,LCM,NBI stage-2
```

tries to build an OSM component called RO. In fact there's an RO repo.
Since the command also tries building NBI, we're going to clone and set
up these two repos too:

```console
% git clone https://osm.etsi.org/gerrit/osm/RO
% git clone https://osm.etsi.org/gerrit/osm/NBI
% for r in IM LCM N2VC NBI RO common devops; do cp commit-msg $r/.git/hooks/; done
```

Now running again the build command got me past the directory error,
but the build seems to get into an infinite loop when installing RO
deps

```console
...
dist_ro_vim_vmware installdeps: -r/build/requirements.txt, -r/build/requirements-dist.txt
```

it just sits there for half an hour seemingly making no progress.
Could it be an issue with VMWare deps? Well, I killed the process
and ran the command again. And again the process got stuck on installing
deps

```console
...
dist_ro_vim_vmware installdeps: -r/build/requirements.txt, -r/build/requirements-dist.txt
...
dist_ro_sdn_odl_of installdeps: -r/build/requirements.txt, -r/build/requirements-dist.txt
```

Notice how this time the VMWare deps step succeeded while the build
got stuck on another component. As a last ditch attempt I tried building
just the RO component

```console
% devops/tools/local-build.sh --module RO stage-2
```

But the build got stuck again on another `installdeps` step

```console
...
dist_ro_vim_vmware installdeps: -r/build/requirements.txt, -r/build/requirements-dist.txt
...
dist_ro_sdn_odl_of installdeps: -r/build/requirements.txt, -r/build/requirements-dist.txt
...
dist_ro_sdn_floodlight_of installdeps: -r/build/requirements.txt, -r/build/requirements-dist.txt
```

I ran it once more and it got stack on an `installdeps` step of a
component that worked in all previous runs

```console
...
dist_ro_vim_openvim installdeps: -r/build/requirements.txt, -r/build/requirements-dist.txt
```

Deadlock bug?


### LCM build failures - part 4

So there's no way we can build RO. All we can do is exclude it from
the build and hope we can still build LCM

```console
% devops/tools/local-build.sh --module common,IM,N2VC,LCM,NBI stage-2
```

With this tweak the build succeeds and we can also go past the LCM
URL failure in creating the Docker image

```console
% devops/tools/local-build.sh --module LCM stage-3
```

The command runs cleanly and tags `opensourcemano/lcm:devel`.




[failed-osm-install]: ./osm-install/install.failed.log

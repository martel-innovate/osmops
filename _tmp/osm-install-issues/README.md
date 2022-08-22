OSM installation issues
-----------------------
> ...documenting the struggle for posterity.

I've captured some of the OSM install sessions that resulted in broken
installs and parked them in this dir just in case I ran into the same
issues again I've got an idea how to fix stuff quickly. One sore point
was the PGP keys---quite a few of them. Here's an example of how to fix
the K8s ones, the others are similar:

- https://stackoverflow.com/questions/49877401

Also notice the install script was broken. See the `multipass*` and
`patched*` scripts in the `osm-install` dir. It looks like the OSM
guys fixed it though:

- https://osm.etsi.org/gitlab/osm/devops/-/commit/fdbe776e9bb9e43f7d4dc0f8c023b93d258666e2


### June 2022 Update

So I rebuilt another OSM 10 VM at the beginning of Jun 2022. Here's
how

```console
$ multipass launch --name osm --cpus 2 --mem 6G --disk 40G 18.04
$ multipass shell osm
% wget https://osm-download.etsi.org/ftp/osm-10.0-ten/install_osm.sh
% chmod +x install_osm.sh
% ./install_osm.sh 2>&1 | tee install.log
% exit
```

The good news is that this time I didn't have to patch the install
script---yay! But...the `osm` client didn't install properly and it
looks like something else may be broken too. In fact, you can't run
the client b/c `pycurl` is missing and when logging into the UI I
couldn't manage to create repos and K8s clusters. Install log saved
to:

- [broken-osm10.install-log.jun2022.log][jun2022-log]




[jun2022-log]: ./broken-osm10.install-log.jun2022.log
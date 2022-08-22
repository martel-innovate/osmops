OSM Packages for demo
=====================

Original packages downloaded from:

- https://osm-download.etsi.org/ftp/Packages/examples/

You need to also add OSM repos before you can create an NS instance from
these packages.


OpenLDAP
--------

### Issue
The OpenLDAP Helm chart in the original is 1.2.3 which doesn't work when
upgrading the NS instance---see `clusterIP` issue documented in the [OSM
client HTTP message flows][msg-flows].

### Fix
We modified the VNFD to use version 1.2.7 instead which doesn't have this
problem. The `openldap_knf.tar.gz` file in this dir contains the fix.

Here's what I did.

1. Extract original package to `openldap_knf` dir.
2. Change Helm chart version to `1.2.7`.
3. Repackage.

Here are the commands for the repackage step

```bash
$ md5sum openldap_knf/openldap_vnfd.yaml > openldap_knf/checksums.txt
$ rm openldap_knf.tar.gz
$ tar -czvf openldap_knf.tar.gz openldap_knf
```




[msg-flows]: ../osm-mitm/message-flows.md
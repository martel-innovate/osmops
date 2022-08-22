Get a token.

```console
curl -v 192.168.64.22/osm/admin/v1/tokens \
  -H 'Accept: application/json' -H 'Content-Type: application/yaml' \
  -d '{"username": "admin", "password": "admin", "project_id": "admin"}'

export OSM_TOKEN=wa7FDNWma96ODtC0PofsQoi1GBAi7Ah6
```

Create OpenLDAP KNF package using original OSM package.

```console
curl -v 192.168.64.22/osm/vnfpkgm/v1/vnf_packages_content \
  -H "Authorization: Bearer ${OSM_TOKEN}" \
  -H 'Accept: application/json' -H 'Content-Type: application/gzip' \
  -H 'Content-Filename: openldap_knf.tar.gz' \
  -H 'Content-File-MD5: 2a7d74587151e9fd0c1fd727003b8a1b' \
  --data-binary @../osm-pkgs/openldap_knf.tar.gz
```

List all KNF packages in YAML format.

```console
curl -v 192.168.64.22/osm/vnfpkgm/v1/vnf_packages_content \
  -H "Authorization: Bearer ${OSM_TOKEN}"
```

Delete OpenLDAP KNF package.
NOTE: can't use ID declared in the package (`openldap_knf`); you've
got to use OSM's own ID (`_id` field).

```console
curl -v 192.168.64.22/osm/vnfpkgm/v1/vnf_packages_content/cc10f9ff-64d2-44c1-a096-95ce17b32b70 \
  -X DELETE \
  -H "Authorization: Bearer ${OSM_TOKEN}"
```


Create OpenLDAP KNF package using OSMOps-generated package.

```console
curl -v 192.168.64.22/osm/vnfpkgm/v1/vnf_packages_content \
  -H "Authorization: Bearer ${OSM_TOKEN}" \
  -H 'Accept: application/json' -H 'Content-Type: application/gzip' \
  -H 'Content-Filename: openldap_knf.tar.gz' \
  -H 'Content-File-MD5: 92821dce2b09c67cc17c780037f3ff03' \
  --data-binary @osmops-generated/openldap_knf.tar.gz
```

Update OpenLDAP KNF package using OSMOps-generated package.

```console
curl -v 192.168.64.22/osm/vnfpkgm/v1/vnf_packages_content/openldap_knf \
  -X PUT \
  -H "Authorization: Bearer ${OSM_TOKEN}" \
  -H 'Accept: application/json' -H 'Content-Type: application/gzip' \
  -H 'Content-Filename: openldap_knf.tar.gz' \
  -H 'Content-File-MD5: 92821dce2b09c67cc17c780037f3ff03' \
  --data-binary @osmops-generated/openldap_knf.tar.gz
```

Apparently you can't PUT the tgz. Notice you get the same error if
you use the OSM KNF package ID:

- /osm/vnfpkgm/v1/vnf_packages_content/943a86dc-a90e-4add-be34-571f3e90f41b

```log
2022-06-17T09:09:58 INFO nbi.server _cplogging.py:213 [17/Jun/2022:09:09:58]  CRITICAL: Exception 'RequestBody' object has no attribute 'get'
Traceback (most recent call last):
  File "/usr/lib/python3/dist-packages/osm_nbi/nbi.py", line 1585, in default
    op_id = self.engine.edit_item(
  File "/usr/lib/python3/dist-packages/osm_nbi/engine.py", line 372, in edit_item
    return self.map_topic[topic].edit(session, _id, indata, kwargs)
  File "/usr/lib/python3/dist-packages/osm_nbi/base_topic.py", line 630, in edit
    indata = self._remove_envelop(indata)
  File "/usr/lib/python3/dist-packages/osm_nbi/descriptor_topics.py", line 628, in _remove_envelop
    if clean_indata.get("etsi-nfv-vnfd:vnfd"):
AttributeError: 'RequestBody' object has no attribute 'get'
2022-06-17T09:09:58 INFO nbi.access _cplogging.py:283 10.244.0.1 - admin/admin;session=Q9RyqcFgHNCP [17/Jun/2022:09:09:58] "PUT /osm/vnfpkgm/v1/vnf_packages_content/openldap_knf HTTP/1.0" 400 110 "" "curl/7.64.1"
```

Update OpenLDAP KNF descriptor.

```console
curl -v 192.168.64.22/osm/vnfpkgm/v1/vnf_packages_content/4ffdeb67-92e7-46fa-9fa2-331a4d674137 \
  -X PUT \
  -H "Authorization: Bearer ${OSM_TOKEN}" \
  -H 'Accept: application/json' -H 'Content-Type: application/yaml' \
  --data-binary @../osm-pkgs/openldap_knf/openldap_vnfd.yaml
```


List all NS packages in YAML format.

```console
curl -v 192.168.64.22/osm/nsd/v1/ns_descriptors_content \
  -H "Authorization: Bearer ${OSM_TOKEN}"
```

Update OpenLDAP NS descriptor.

```console
curl -v 192.168.64.22/osm/nsd/v1/ns_descriptors_content/6cb736be-8a59-4c60-a979-22328b8094d4 \
  -X PUT \
  -H "Authorization: Bearer ${OSM_TOKEN}" \
  -H 'Accept: application/json' -H 'Content-Type: application/yaml' \
  --data-binary @../osm-pkgs/openldap_ns/openldap_nsd.yaml
```

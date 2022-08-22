HTTP message flow for NS create 2
---------------------------------

### GET NS descriptors

HTTP request

```http
GET /osm/nsd/v1/ns_descriptors HTTP/1.1
Host: localhost
User-Agent: PycURL/7.43.0.6 libcurl/7.58.0 OpenSSL/1.1.1 zlib/1.2.11 libidn2/2.0.4 libpsl/0.19.1 (+libidn2/2.0.4) nghttp2/1.30.0 librtmp/2.3
Accept: application/json
Content-Type: application/yaml
Authorization: Bearer lKGRJmJQdpz9yasXikIUdD5SivCf3tXF
```

HTTP response

```http
HTTP/1.1 200 OK
Server: nginx/1.14.0 (Ubuntu)
Date: Fri, 10 Sep 2021 13:55:58 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 2519
Connection: keep-alive
Set-Cookie: session_id=b4a5a5b5b2519f6550ea3bea5efdbf2f4f9ad517; expires=Fri, 10 Sep 2021 14:55:58 GMT; HttpOnly; Max-Age=3600; Path=/; Secure

[
    {
        "_id": "aba58e40-d65f-4f4e-be0a-e248c14d3e03",
        "id": "openldap_ns",
        "designer": "OSM",
        "version": "1.0",
        "name": "openldap_ns",
        "vnfd-id": [
            "openldap_knf"
        ],
        "virtual-link-desc": [
            {
                "id": "mgmtnet",
                "mgmt-network": true
            }
        ],
        "df": [
            {
                "id": "default-df",
                "vnf-profile": [
                    {
                        "id": "openldap",
                        "virtual-link-connectivity": [
                            {
                                "constituent-cpd-id": [
                                    {
                                        "constituent-base-element-id": "openldap",
                                        "constituent-cpd-id": "mgmt-ext"
                                    }
                                ],
                                "virtual-link-profile-id": "mgmtnet"
                            }
                        ],
                        "vnfd-id": "openldap_knf"
                    }
                ]
            }
        ],
        "description": "NS consisting of a single KNF openldap_knf connected to mgmt network",
        "_admin": {
            "userDefinedData": {},
            "created": 1631268635.96618,
            "modified": 1631268637.8627107,
            "projects_read": [
                "fada443a-905c-4241-8a33-4dcdbdac55e7"
            ],
            "projects_write": [
                "fada443a-905c-4241-8a33-4dcdbdac55e7"
            ],
            "onboardingState": "ONBOARDED",
            "operationalState": "ENABLED",
            "usageState": "NOT_IN_USE",
            "storage": {
                "fs": "mongo",
                "path": "/app/storage/",
                "folder": "aba58e40-d65f-4f4e-be0a-e248c14d3e03",
                "pkg-dir": "openldap_ns",
                "descriptor": "openldap_ns/openldap_nsd.yaml",
                "zipfile": "openldap_ns.tar.gz"
            }
        },
        "nsdOnboardingState": "ONBOARDED",
        "nsdOperationalState": "ENABLED",
        "nsdUsageState": "NOT_IN_USE",
        "_links": {
            "self": {
                "href": "/nsd/v1/ns_descriptors/aba58e40-d65f-4f4e-be0a-e248c14d3e03"
            },
            "nsd_content": {
                "href": "/nsd/v1/ns_descriptors/aba58e40-d65f-4f4e-be0a-e248c14d3e03/nsd_content"
            }
        }
    }
]
```


### GET target VIM account

HTTP request

```http
GET /osm/admin/v1/vim_accounts/4a4425f7-3e72-4d45-a4ec-4241186f3547 HTTP/1.1
Host: localhost
User-Agent: PycURL/7.43.0.6 libcurl/7.58.0 OpenSSL/1.1.1 zlib/1.2.11 libidn2/2.0.4 libpsl/0.19.1 (+libidn2/2.0.4) nghttp2/1.30.0 librtmp/2.3
Accept: application/json
Content-Type: application/yaml
Authorization: Bearer lKGRJmJQdpz9yasXikIUdD5SivCf3tXF
```

HTTP response

```http
HTTP/1.1 200 OK
Server: nginx/1.14.0 (Ubuntu)
Date: Fri, 10 Sep 2021 13:55:58 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 1039
Connection: keep-alive
Set-Cookie: session_id=850ace39d2abbe728dcfd089e7e9407b90a7ec43; expires=Fri, 10 Sep 2021 14:55:58 GMT; HttpOnly; Max-Age=3600; Path=/; Secure

{
    "_id": "4a4425f7-3e72-4d45-a4ec-4241186f3547",
    "name": "mylocation1",
    "vim_type": "dummy",
    "description": null,
    "vim_url": "http://localhost/dummy",
    "vim_user": "u",
    "vim_password": "fNnfmd3KFXvfyVKu3nzItg==",
    "vim_tenant_name": "p",
    "_admin": {
        "created": 1631212983.5388303,
        "modified": 1631212983.5388303,
        "projects_read": [
            "fada443a-905c-4241-8a33-4dcdbdac55e7"
        ],
        "projects_write": [
            "fada443a-905c-4241-8a33-4dcdbdac55e7"
        ],
        "operationalState": "ENABLED",
        "operations": [
            {
                "lcmOperationType": "create",
                "operationState": "COMPLETED",
                "startTime": 1631212983.5930278,
                "statusEnteredTime": 1631212984.0220273,
                "operationParams": null
            }
        ],
        "current_operation": null,
        "detailed-status": ""
    },
    "schema_version": "1.11",
    "admin": {
        "current_operation": 0
    }
}
```

Notice OSM client reissues the same `GET` again after this. This duplication
of HTTP requests might well be a bug...


### POST NS instance content

HTTP request

```http
POST /osm/nslcm/v1/ns_instances_content HTTP/1.1
Host: localhost
User-Agent: PycURL/7.43.0.6 libcurl/7.58.0 OpenSSL/1.1.1 zlib/1.2.11 libidn2/2.0.4 libpsl/0.19.1 (+libidn2/2.0.4) nghttp2/1.30.0 librtmp/2.3
Accept: application/json
Content-Type: application/yaml
Authorization: Bearer lKGRJmJQdpz9yasXikIUdD5SivCf3tXF
Content-Length: 319

{"nsdId": "aba58e40-d65f-4f4e-be0a-e248c14d3e03", "nsName": "ldap2", "nsDescription": "default description", "vimAccountId": "4a4425f7-3e72-4d45-a4ec-4241186f3547", "additionalParamsForVnf": [{"member-vnf-index": "openldap", "additionalParamsForKdu": [{"kdu_name": "ldap", "additionalParams": {"replicaCount": "2"}}]}]}
```

HTTP response

```http
HTTP/1.1 201 Created
Server: nginx/1.14.0 (Ubuntu)
Date: Fri, 10 Sep 2021 13:55:59 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 111
Connection: keep-alive
Location: /osm/nslcm/v1/ns_instances_content/136fcc46-c363-4d74-af14-c115fff7d80a
Set-Cookie: session_id=0d478e27eccc45667ff15c405362764cb6add7cd; expires=Fri, 10 Sep 2021 14:55:58 GMT; HttpOnly; Max-Age=3600; Path=/; Secure

{
    "id": "136fcc46-c363-4d74-af14-c115fff7d80a",
    "nslcmop_id": "b7b6b8a6-f4f2-4862-99bc-f2c73fbf8736"
}
```

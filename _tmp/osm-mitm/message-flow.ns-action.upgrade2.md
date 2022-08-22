HTTP message flow for NS upgrade
--------------------------------

### GET NS instances content

HTTP request

```http
GET /osm/nslcm/v1/ns_instances_content HTTP/1.1
Host: localhost
User-Agent: PycURL/7.43.0.6 libcurl/7.58.0 OpenSSL/1.1.1 zlib/1.2.11 libidn2/2.0.4 libpsl/0.19.1 (+libidn2/2.0.4) nghttp2/1.30.0 librtmp/2.3
Accept: application/json
Content-Type: application/yaml
Authorization: Bearer Rxb9XReQHdW6XmtjKLFLToLs0W0XbD7n
```

HTTP response

```http
HTTP/1.1 200 OK
Server: nginx/1.14.0 (Ubuntu)
Date: Fri, 10 Sep 2021 17:24:59 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 13396
Connection: keep-alive
Set-Cookie: session_id=02efc3019fd72867333ec8223528ed6fbcf022ed; expires=Fri, 10 Sep 2021 18:24:59 GMT; HttpOnly; Max-Age=3600; Path=/; Secure

[
    {
        "_id": "0335c32c-d28c-4d79-9b94-0ffa36326932",
        "name": "ldap",
        "name-ref": "ldap",
        "short-name": "ldap",
        "admin-status": "ENABLED",
        "nsState": "READY",
        "currentOperation": "IDLE",
        "currentOperationID": null,
        "errorDescription": null,
        "errorDetail": null,
        "deploymentStatus": null,
        "configurationStatus": [],
        "vcaStatus": null,
        "nsd": {
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
            }
        },
        "datacenter": "4a4425f7-3e72-4d45-a4ec-4241186f3547",
        "resource-orchestrator": "osmopenmano",
        "description": "default description",
        "constituent-vnfr-ref": [
            "ae63ee09-847f-4108-9a22-852899b6e0ae"
        ],
        "operational-status": "running",
        "config-status": "configured",
        "detailed-status": "Done",
        "orchestration-progress": {},
        "create-time": 1631277626.5666356,
        "nsd-name-ref": "openldap_ns",
        "operational-events": [],
        "nsd-ref": "openldap_ns",
        "nsd-id": "aba58e40-d65f-4f4e-be0a-e248c14d3e03",
        "vnfd-id": [
            "d506d18f-0738-42ab-8b45-cfa98da38e7a"
        ],
        "instantiate_params": {
            "nsdId": "aba58e40-d65f-4f4e-be0a-e248c14d3e03",
            "nsName": "ldap",
            "nsDescription": "default description",
            "vimAccountId": "4a4425f7-3e72-4d45-a4ec-4241186f3547"
        },
        "additionalParamsForNs": null,
        "ns-instance-config-ref": "0335c32c-d28c-4d79-9b94-0ffa36326932",
        "id": "0335c32c-d28c-4d79-9b94-0ffa36326932",
        "ssh-authorized-key": null,
        "flavor": [],
        "image": [],
        "vld": [
            {
                "id": "mgmtnet",
                "mgmt-network": true,
                "name": "mgmtnet",
                "type": null,
                "vim_info": {
                    "vim:4a4425f7-3e72-4d45-a4ec-4241186f3547": {
                        "vim_account_id": "4a4425f7-3e72-4d45-a4ec-4241186f3547",
                        "vim_network_name": null,
                        "vim_details": "{name: mgmtnet, status: ACTIVE}\n",
                        "vim_id": "81a7fb44-b765-4b16-985f-13b481d3b892",
                        "vim_status": "ACTIVE",
                        "vim_name": "mgmtnet"
                    }
                }
            }
        ],
        "_admin": {
            "created": 1631277626.626409,
            "modified": 1631285667.3411994,
            "projects_read": [
                "fada443a-905c-4241-8a33-4dcdbdac55e7"
            ],
            "projects_write": [
                "fada443a-905c-4241-8a33-4dcdbdac55e7"
            ],
            "nsState": "INSTANTIATED",
            "current-operation": null,
            "nslcmop": null,
            "operation-type": null,
            "deployed": {
                "RO": {
                    "vnfd": [],
                    "operational-status": "running"
                },
                "VCA": [],
                "K8s": [
                    {
                        "kdu-instance": "stable-openldap-1-2-3-0098084071",
                        "k8scluster-uuid": "kube-system:b33b0bfd-ce33-47b9-b286-a60c8f04b6d9",
                        "k8scluster-type": "helm-chart-v3",
                        "member-vnf-index": "openldap",
                        "kdu-name": "ldap",
                        "kdu-model": "stable/openldap:1.2.3",
                        "namespace": "fada443a-905c-4241-8a33-4dcdbdac55e7",
                        "kdu-deployment-name": "",
                        "detailed-status": "{'info': {'deleted': '', 'description': 'Rollback \"stable-openldap-1-2-3-0098084071\" failed: cannot patch \"stable-openldap-1-2-3-0098084071\" with kind Service: Service \"stable-openldap-1-2-3-0098084071\" is invalid: spec.clusterIP: Invalid value: \"\": field is immutable', 'first_deployed': '2021-09-10T12:40:56.55575157Z', 'last_deployed': '2021-09-10T14:54:26.378456605Z', 'status': 'failed'}, 'name': 'stable-openldap-1-2-3-0098084071', 'namespace': 'fada443a-905c-4241-8a33-4dcdbdac55e7', 'version': 3}",
                        "operation": "upgrade",
                        "status": "Rollback \"stable-openldap-1-2-3-0098084071\" failed: cannot patch \"stable-openldap-1-2-3-0098084071\" with kind Service: Service \"stab\" is invalid: spec.clusterIP: Invalid value: \"\": field is immutable",
                        "status-time": "1631285667.3301775"
                    }
                ]
            }
        }
    },
    {
        "_id": "136fcc46-c363-4d74-af14-c115fff7d80a",
        "name": "ldap2",
        "name-ref": "ldap2",
        "short-name": "ldap2",
        "admin-status": "ENABLED",
        "nsState": "READY",
        "currentOperation": "IDLE",
        "currentOperationID": null,
        "errorDescription": null,
        "errorDetail": null,
        "deploymentStatus": null,
        "configurationStatus": [],
        "vcaStatus": null,
        "nsd": {
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
            }
        },
        "datacenter": "4a4425f7-3e72-4d45-a4ec-4241186f3547",
        "resource-orchestrator": "osmopenmano",
        "description": "default description",
        "constituent-vnfr-ref": [
            "609ae829-8fbe-44f1-944d-2fba5cd909c2"
        ],
        "operational-status": "running",
        "config-status": "configured",
        "detailed-status": "Done",
        "orchestration-progress": {},
        "create-time": 1631282159.0447648,
        "nsd-name-ref": "openldap_ns",
        "operational-events": [],
        "nsd-ref": "openldap_ns",
        "nsd-id": "aba58e40-d65f-4f4e-be0a-e248c14d3e03",
        "vnfd-id": [
            "d506d18f-0738-42ab-8b45-cfa98da38e7a"
        ],
        "instantiate_params": {
            "nsdId": "aba58e40-d65f-4f4e-be0a-e248c14d3e03",
            "nsName": "ldap2",
            "nsDescription": "default description",
            "vimAccountId": "4a4425f7-3e72-4d45-a4ec-4241186f3547"
        },
        "additionalParamsForNs": null,
        "ns-instance-config-ref": "136fcc46-c363-4d74-af14-c115fff7d80a",
        "id": "136fcc46-c363-4d74-af14-c115fff7d80a",
        "ssh-authorized-key": null,
        "flavor": [],
        "image": [],
        "vld": [
            {
                "id": "mgmtnet",
                "mgmt-network": true,
                "name": "mgmtnet",
                "type": null,
                "vim_info": {
                    "vim:4a4425f7-3e72-4d45-a4ec-4241186f3547": {
                        "vim_account_id": "4a4425f7-3e72-4d45-a4ec-4241186f3547",
                        "vim_network_name": null,
                        "vim_details": "{name: mgmtnet, status: ACTIVE}\n",
                        "vim_id": "81a7fb44-b765-4b16-985f-13b481d3b892",
                        "vim_status": "ACTIVE",
                        "vim_name": "mgmtnet"
                    }
                }
            }
        ],
        "_admin": {
            "created": 1631282159.0555632,
            "modified": 1631285403.5654724,
            "projects_read": [
                "fada443a-905c-4241-8a33-4dcdbdac55e7"
            ],
            "projects_write": [
                "fada443a-905c-4241-8a33-4dcdbdac55e7"
            ],
            "nsState": "INSTANTIATED",
            "current-operation": null,
            "nslcmop": null,
            "operation-type": null,
            "deployed": {
                "RO": {
                    "vnfd": [],
                    "operational-status": "running"
                },
                "VCA": [],
                "K8s": [
                    {
                        "kdu-instance": "stable-openldap-1-2-3-0044064996",
                        "k8scluster-uuid": "kube-system:b33b0bfd-ce33-47b9-b286-a60c8f04b6d9",
                        "k8scluster-type": "helm-chart-v3",
                        "member-vnf-index": "openldap",
                        "kdu-name": "ldap",
                        "kdu-model": "stable/openldap:1.2.3",
                        "namespace": "fada443a-905c-4241-8a33-4dcdbdac55e7",
                        "kdu-deployment-name": "",
                        "detailed-status": "{'config': {'replicaCount': '2'}, 'info': {'deleted': '', 'description': 'Install complete', 'first_deployed': '2021-09-10T13:56:20.089257801Z', 'last_deployed': '2021-09-10T13:56:20.089257801Z', 'status': 'deployed'}, 'name': 'stable-openldap-1-2-3-0044064996', 'namespace': 'fada443a-905c-4241-8a33-4dcdbdac55e7', 'version': 1}",
                        "operation": "install",
                        "status": "Install complete",
                        "status-time": "1631282216.1732676"
                    }
                ]
            }
        }
    }
]
```

Notice OSM client reissues the same `GET` again after this. This duplication
of HTTP requests might well be a bug...


### POST target NS instance action

HTTP request

```http
POST /osm/nslcm/v1/ns_instances/136fcc46-c363-4d74-af14-c115fff7d80a/action HTTP/1.1
Host: localhost
User-Agent: PycURL/7.43.0.6 libcurl/7.58.0 OpenSSL/1.1.1 zlib/1.2.11 libidn2/2.0.4 libpsl/0.19.1 (+libidn2/2.0.4) nghttp2/1.30.0 librtmp/2.3
Accept: application/json
Content-Type: application/yaml
Authorization: Bearer Rxb9XReQHdW6XmtjKLFLToLs0W0XbD7n
Content-Length: 119

{"member_vnf_index": "openldap", "kdu_name": "ldap", "primitive": "upgrade", "primitive_params": {"replicaCount": "3"}}
```

HTTP response

```http
HTTP/1.1 202 Accepted
Server: nginx/1.14.0 (Ubuntu)
Date: Fri, 10 Sep 2021 17:25:00 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 53
Connection: keep-alive
Location: /osm/nslcm/v1/ns_lcm_op_occs/f92f746f-c10a-448e-84e1-3acfd8b684cb
Set-Cookie: session_id=ca59ba9d0d29c0535d2273935498383e9af28a68; expires=Fri, 10 Sep 2021 18:24:59 GMT; HttpOnly; Max-Age=3600; Path=/; Secure

{
    "id": "f92f746f-c10a-448e-84e1-3acfd8b684cb"
}
```

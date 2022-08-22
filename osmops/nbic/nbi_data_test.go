package nbic

// expired on Wed Sep 08 2021 18:52:11 GMT+0000
var expiredNbiTokenPayload = `{
	"issued_at": 1631123531.1251214,
	"expires": 1631127131.1251214,
	"_id": "TuD41hLjDvjlR2cPcAFvWcr6FGvRhIk2",
	"id": "TuD41hLjDvjlR2cPcAFvWcr6FGvRhIk2",
	"project_id": "fada443a-905c-4241-8a33-4dcdbdac55e7",
	"project_name": "admin",
	"username": "admin",
	"user_id": "5c6f2d64-9c23-4718-806a-c74c3fc3c98f",
	"admin": true,
	"roles": [{
		"name": "system_admin",
		"id": "cb545e44-cd2b-4c0b-93aa-7e2cee79afc3"
	}]
}`

// expires on Sat May 17 2053 20:38:51 GMT+0000
var validNbiTokenPayload = `{
	"issued_at": 2631127131.1251214,
	"expires": 2631127131.1251214,
	"_id": "TuD41hLjDvjlR2cPcAFvWcr6FGvRhIk2",
	"id": "TuD41hLjDvjlR2cPcAFvWcr6FGvRhIk2",
	"project_id": "fada443a-905c-4241-8a33-4dcdbdac55e7",
	"project_name": "admin",
	"username": "admin",
	"user_id": "5c6f2d64-9c23-4718-806a-c74c3fc3c98f",
	"admin": true,
	"roles": [{
		"name": "system_admin",
		"id": "cb545e44-cd2b-4c0b-93aa-7e2cee79afc3"
	}]
}`

var vnfDescriptors = `[
    {
        "_id": "4ffdeb67-92e7-46fa-9fa2-331a4d674137",
        "description": "KNF with single KDU using a helm-chart for openldap version 1.2.7",
        "df": [
            {
                "id": "default-df"
            }
        ],
        "ext-cpd": [
            {
                "id": "mgmt-ext",
                "k8s-cluster-net": "mgmtnet"
            }
        ],
        "id": "openldap_knf",
        "k8s-cluster": {
            "nets": [
                {
                    "id": "mgmtnet"
                }
            ]
        },
        "kdu": [
            {
                "name": "ldap",
                "helm-chart": "stable/openldap:1.2.7"
            }
        ],
        "mgmt-cp": "mgmt-ext",
        "product-name": "openldap_knf",
        "provider": "Telefonica",
        "version": "1.0",
        "_admin": {
            "userDefinedData": {},
            "created": 1655475517.840946,
            "modified": 1655478654.0081894,
            "projects_read": [
                "c9e9cf6f-98a4-45f8-b18d-b70d93422d88"
            ],
            "projects_write": [
                "c9e9cf6f-98a4-45f8-b18d-b70d93422d88"
            ],
            "onboardingState": "ONBOARDED",
            "operationalState": "ENABLED",
            "usageState": "NOT_IN_USE",
            "storage": {
                "fs": "mongo",
                "path": "/app/storage/",
                "folder": "4ffdeb67-92e7-46fa-9fa2-331a4d674137",
                "pkg-dir": "openldap_knf",
                "descriptor": "openldap_knf/openldap_vnfd.yaml",
                "zipfile": "openldap_knf.tar.gz"
            }
        },
        "onboardingState": "ONBOARDED",
        "operationalState": "ENABLED",
        "usageState": "NOT_IN_USE",
        "_links": {
            "self": {
                "href": "/vnfpkgm/v1/vnf_packages/4ffdeb67-92e7-46fa-9fa2-331a4d674137"
            },
            "vnfd": {
                "href": "/vnfpkgm/v1/vnf_packages/4ffdeb67-92e7-46fa-9fa2-331a4d674137/vnfd"
            },
            "packageContent": {
                "href": "/vnfpkgm/v1/vnf_packages/4ffdeb67-92e7-46fa-9fa2-331a4d674137/package_content"
            }
        }
    },
    {
        "_id": "5ccfed39-92e7-46fa-9fa2-331a4d674137",
        "description": "Made-up KNF with single KDU using a helm-chart for openldap version 1.2.7",
        "df": [
            {
                "id": "default-df"
            }
        ],
        "ext-cpd": [
            {
                "id": "mgmt-ext",
                "k8s-cluster-net": "mgmtnet"
            }
        ],
        "id": "dummy_knf",
        "k8s-cluster": {
            "nets": [
                {
                    "id": "mgmtnet"
                }
            ]
        },
        "kdu": [
            {
                "name": "ldap",
                "helm-chart": "stable/openldap:1.2.7"
            }
        ],
        "mgmt-cp": "mgmt-ext",
        "product-name": "dummy_knf",
        "provider": "big corp",
        "version": "1.0",
        "_admin": {
            "userDefinedData": {},
            "created": 1655475517.840946,
            "modified": 1655478654.0081894,
            "projects_read": [
                "c9e9cf6f-98a4-45f8-b18d-b70d93422d88"
            ],
            "projects_write": [
                "c9e9cf6f-98a4-45f8-b18d-b70d93422d88"
            ],
            "onboardingState": "ONBOARDED",
            "operationalState": "ENABLED",
            "usageState": "NOT_IN_USE",
            "storage": {
                "fs": "mongo",
                "path": "/app/storage/",
                "folder": "5ccfed39-92e7-46fa-9fa2-331a4d674137",
                "pkg-dir": "dummy_knf",
                "descriptor": "dummy_knf/openldap_vnfd.yaml",
                "zipfile": "dummy_knf.tar.gz"
            }
        },
        "onboardingState": "ONBOARDED",
        "operationalState": "ENABLED",
        "usageState": "NOT_IN_USE",
        "_links": {
            "self": {
                "href": "/vnfpkgm/v1/vnf_packages/5ccfed39-92e7-46fa-9fa2-331a4d674137"
            },
            "vnfd": {
                "href": "/vnfpkgm/v1/vnf_packages/5ccfed39-92e7-46fa-9fa2-331a4d674137/vnfd"
            },
            "packageContent": {
                "href": "/vnfpkgm/v1/vnf_packages/5ccfed39-92e7-46fa-9fa2-331a4d674137/package_content"
            }
        }
    }
]`

var nsDescriptors = `[
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
    },
	{
        "_id": "ddd20a30-d65f-4f4e-be0a-e248c14d3e03",
        "id": "dummy_ns",
        "designer": "OSM",
        "version": "1.0",
        "name": "dummy_ns",
        "vnfd-id": [
            "dummy_knf"
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
                        "id": "dummy",
                        "virtual-link-connectivity": [
                            {
                                "constituent-cpd-id": [
                                    {
                                        "constituent-base-element-id": "dummy",
                                        "constituent-cpd-id": "mgmt-ext"
                                    }
                                ],
                                "virtual-link-profile-id": "mgmtnet"
                            }
                        ],
                        "vnfd-id": "dummy_knf"
                    }
                ]
            }
        ],
        "description": "Made-up NS consisting of a single KNF dummy_knf connected to mgmt network",
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
                "folder": "ddd20a30-d65f-4f4e-be0a-e248c14d3e03",
                "pkg-dir": "dummy_ns",
                "descriptor": "dummy_ns/openldap_nsd.yaml",
                "zipfile": "openldap_ns.tar.gz"
            }
        },
        "nsdOnboardingState": "ONBOARDED",
        "nsdOperationalState": "ENABLED",
        "nsdUsageState": "NOT_IN_USE",
        "_links": {
            "self": {
                "href": "/nsd/v1/ns_descriptors/ddd20a30-d65f-4f4e-be0a-e248c14d3e03"
            },
            "nsd_content": {
                "href": "/nsd/v1/ns_descriptors/ddd20a30-d65f-4f4e-be0a-e248c14d3e03/nsd_content"
            }
        }
    }
]`

var vimAccounts = `[
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
]`

var nsInstancesContent = `[
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
            "modified": 1631285336.7610166,
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
                        "detailed-status": "{'info': {'deleted': '', 'description': 'Install complete', 'first_deployed': '2021-09-10T12:40:56.55575157Z', 'last_deployed': '2021-09-10T12:40:56.55575157Z', 'status': 'deployed'}, 'name': 'stable-openldap-1-2-3-0098084071', 'namespace': 'fada443a-905c-4241-8a33-4dcdbdac55e7', 'version': 1}",
                        "operation": "install",
                        "status": "Install complete",
                        "status-time": "1631277711.4568162"
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
        "nsState": "READ",
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
    },
    {
        "_id": "111fcc46-c363-4d74-af14-c115fff7d80a",
        "name": "dup-name"
    },
    {
        "_id": "222fcc46-c363-4d74-af14-c115fff7d80a",
        "name": "dup-name"
    }
]`

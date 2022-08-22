NS package list example
-----------------------

Example output of a `GET` on `/osm/nsd/v1/ns_descriptors_content`.
Notice this is the same as a `GET` on `/osm/nsd/v1/ns_descriptors`---I
think this is the SOL005 endpoint?

```yaml
-   _admin:
        created: 1655475749.560676
        modified: 1655478812.9101527
        onboardingState: ONBOARDED
        operationalState: ENABLED
        projects_read:
        - c9e9cf6f-98a4-45f8-b18d-b70d93422d88
        projects_write:
        - c9e9cf6f-98a4-45f8-b18d-b70d93422d88
        storage:
            descriptor: openldap_ns/openldap_nsd.yaml
            folder: 6cb736be-8a59-4c60-a979-22328b8094d4
            fs: mongo
            path: /app/storage/
            pkg-dir: openldap_ns
            zipfile: openldap_ns.tar.gz
        usageState: NOT_IN_USE
        userDefinedData: {}
    _id: 6cb736be-8a59-4c60-a979-22328b8094d4
    _links:
        nsd_content:
            href: /nsd/v1/ns_descriptors/6cb736be-8a59-4c60-a979-22328b8094d4/nsd_content
        self:
            href: /nsd/v1/ns_descriptors/6cb736be-8a59-4c60-a979-22328b8094d4
    description: NS consisting of a single KNF openldap_knf connected to mgmt network
    designer: OSM
    df:
    -   id: default-df
        vnf-profile:
        -   id: openldap
            virtual-link-connectivity:
            -   constituent-cpd-id:
                -   constituent-base-element-id: openldap
                    constituent-cpd-id: mgmt-ext
                virtual-link-profile-id: mgmtnet
            vnfd-id: openldap_knf
    id: openldap_ns
    name: openldap_ns
    nsdOnboardingState: ONBOARDED
    nsdOperationalState: ENABLED
    nsdUsageState: NOT_IN_USE
    version: '1.0'
    virtual-link-desc:
    -   id: mgmtnet
        mgmt-network: true
    vnfd-id:
    - openldap_knf
```
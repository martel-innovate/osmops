KNF package list example
------------------------

Example output of a `GET` on `/osm/vnfpkgm/v1/vnf_packages_content`.

```yaml
-   _admin:
        created: 1655475517.840946
        modified: 1655478654.0081894
        onboardingState: ONBOARDED
        operationalState: ENABLED
        projects_read:
        - c9e9cf6f-98a4-45f8-b18d-b70d93422d88
        projects_write:
        - c9e9cf6f-98a4-45f8-b18d-b70d93422d88
        storage:
            descriptor: openldap_knf/openldap_vnfd.yaml
            folder: 4ffdeb67-92e7-46fa-9fa2-331a4d674137
            fs: mongo
            path: /app/storage/
            pkg-dir: openldap_knf
            zipfile: openldap_knf.tar.gz
        usageState: NOT_IN_USE
        userDefinedData: {}
    _id: 4ffdeb67-92e7-46fa-9fa2-331a4d674137
    _links:
        packageContent:
            href: /vnfpkgm/v1/vnf_packages/4ffdeb67-92e7-46fa-9fa2-331a4d674137/package_content
        self:
            href: /vnfpkgm/v1/vnf_packages/4ffdeb67-92e7-46fa-9fa2-331a4d674137
        vnfd:
            href: /vnfpkgm/v1/vnf_packages/4ffdeb67-92e7-46fa-9fa2-331a4d674137/vnfd
    description: KNF with single KDU using a helm-chart for openldap version 1.2.7
    df:
    -   id: default-df
    ext-cpd:
    -   id: mgmt-ext
        k8s-cluster-net: mgmtnet
    id: openldap_knf
    k8s-cluster:
        nets:
        -   id: mgmtnet
    kdu:
    -   helm-chart: stable/openldap:1.2.7
        name: ldap
    mgmt-cp: mgmt-ext
    onboardingState: ONBOARDED
    operationalState: ENABLED
    product-name: openldap_knf
    provider: Telefonica
    usageState: NOT_IN_USE
    version: '1.0'
```
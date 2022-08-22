### Federico's notes about what he'd like to do.

```
helm-repositories/
  helm-repo1.yaml
vnfd-catalog
  fb_magma_knf
   |
   --- fb_magma_knfd.yaml
ns-catalog
  fb_magma_ns
     |
     --- fb_magma_nsd.yaml
ns
  ns.yaml
```

`helm-repo1.yaml`

```yaml
apiVersion: flux.weave.works/v1beta1
kind: HelmRepository
metadata:
  name: magma 
spec:
  id: magma 
  url: https://felipevicens.github.io/fb-magma-helm-chart/
```

^ "check that this may be already covered!!!" (Federico, 9 Mar 2021)

`ns.yaml`

```yaml
apiVersion: flux.weave.works/v1beta1
kind: NSResource
Metadata:
  name: magma_orc8r
spec:
  name: magma_orc8r
  nsd: fb_magma_ns
  vim: <vim_name>
  config:
    -
  parameters:
    - fb_magma_knf:
        kdu_model: "stable/openldap:1.2.2"
```

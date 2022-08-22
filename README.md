OSM Ops
-------
> GitOps for Open Source MANO.

OSM Ops is a cloud-native micro-service to implement GitOps workflows
within [Open Source MANO][osm] (OSM). The basic idea is to describe
the state of an OSM deployment through version-controlled text files
hosted in an online Git repository. Each file declares a desired instantiation
and runtime configuration for some of the services in a specified OSM
cluster. Collectively, the files at a given Git revision describe the
deployment state of the these services at a certain point in time.
OSM Ops monitors the Git repository in order to automatically reconcile
the desired deployment state with the actual live state of the OSM
cluster. OSM Ops is implemented as a [Kubernetes][k8s] operator that
plugs into the [FluxCD][flux] framework in order to leverage the rich
Kubernetes/FluxCD GitOps ecosystem.

This software has been developed by [Martel Innovate][martel] as part
of the [Affordable5G][a5g] EU-funded project. OSM Ops serves the specific
needs of Affordable5G and is not intended as a replacement or alternative
to OSM's own deployment and operations tools, rather as a complement.


### Documentation

The [introduction section][arch.intro] of the [software architecture
document][arch] is the best starting point to learn about OSM Ops:
it is a short read but presents the fundamental ideas clearly with
the help of diagrams. The reader interested in gaining a deeper technical
understanding of OSM Ops is invited to consider the remainder of the
document too. [Hands-on tutorials][demos] demonstrate the core features
and exemplify deployment scenarios.


### Features at a glance

- **Declarative approach**. Edit YAML files to specify which KNFs should
  be in the target OSM cluster and their configuration. OSM Ops determines
  whether to create a new KNF or update an existing one, then issues the
  OSM commands to realise your configuration. OSM Ops can also create or
  update OSM packages.
- **GitOps workflow**. Keep your OSM Ops YAML files in an online Git
  repository. OSM Ops automatically detects new commits and reconciles
  the deployment state declared in the YAML files with the actual live
  state of the OSM cluster.
- **Multi-repo/multi-cluster**. All the OSM Ops files in a Git repository
  target the same OSM cluster. However, you can have OSM Ops monitor multiple
  repositories if you need to manage several distinct OSM clusters at once.
- **Secure handling of OSM credentials**. Use Kubernetes secrets to provide
  the username, password and project for OSM Ops to connect to the target
  OSM cluster.
- **Repo file filters**. Optionally specify filters to match OSM Ops YAML
  files in your repository. Speeds up processing if there are a large number
  of files (e.g. source code, documents, etc.) that OSM Ops should not read.
- **Efficient batch processing**. Up to 6x faster and 89% bandwidth savings
  when processing many KNF create/update operations compared to using the
  `osm` CLI—thanks to caching (NS descriptors, VIM accounts, etc.) and
  smart management of authorisation token lifecycle.


### Project status

- Early days. But the code is solid (modular, close to 100% test coverage)
  and is a good foundation for further development.
- Successfully deployed and run the Malaga Nov 2021 demo; ready for the
  Malaga end-to-end tests in Q3 2022.
- Only create/update KNF available. No rollbacks—delete not implemented.
  But you can still rollback to a previous Git version as long as the set
  of KNFs is the same in both versions.
- OSM packaging functionality partially relies on naming conventions. A
  reasonable choice given the current phase of the project, but it could
  be improved in later iterations. ([Details][pkg].)



[arch]: ./docs/arch/README.md
[arch.intro]: ./docs/arch/intro.md
[a5g]: https://www.affordable5g.eu/
    "Affordable5G"
[demos]: ./docs/demos/README.md
[flux]: https://fluxcd.io/
    "Flux - the GitOps family of projects"
[k8s]: https://en.wikipedia.org/wiki/Kubernetes
    "Kubernetes"
[martel]: https://www.martel-innovate.com/
    "Martel Innovate"
[osm]: https://osm.etsi.org/
    "Open Source MANO"
[pkg]: ./docs/osm-pkgs.md
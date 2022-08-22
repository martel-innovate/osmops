OSM Ops Software Architecture
-----------------------------
> A technical map of the software.

This document describes (only) the **technical** aspects of the OSM
Ops architecture through a set of interlocked architectural viewpoints.
The document is mainly aimed at developers who need to understand
the big picture before modifying the architecture or extending the
code with new functionality.


### Document status

**Work in progress**. Even though this document is a first draft and
many sections need to be written, the included material should be
enough to gain a basic understanding of the OSM Ops architecture.


### Prerequisites

We assume the reader is well versed in distributed systems and cloud
computing. Moreover, the reader should be familiar with the following
technologies: HTTP/REST, Docker, Kubernetes (in particular the Operator
architecture), Go, Kubebuilder, IaC/DevOps/GitOps, FluxCD, Open Source
MANO.


### Table of contents

1. [Introduction][intro]. The basic ideas are summarised here and then
   further developed in later sections.
2. [System requirements][requirements]. An account of functional
   requirements and system quality attributes.
3. [Information model][info-model]. What information the system handles
   and how it is represented and processed.
4. [System decomposition][components]. Components, interfaces and
   modularity.
5. [Interaction mechanics][interaction]. Distributed communication
   protocols and synchronisation, caching.
6. [Implementation][implementation]. Codebase essentials.
7. [Deployment and scalability][deployment].
8. [Quality assurance][qa].




[components]: ./components.md
[deployment]: ./deployment.md
[implementation]: ./implementation.md
[info-model]: ./info-model.md
[interaction]: ./interaction.md
[intro]: ./intro.md
[qa]: ./qa.md
[requirements]: ./requirements.md

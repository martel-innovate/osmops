Introduction
------------
> The why, the what and the how.

This introductory section first touches on project motivation and goals,
then goes on to sketching the architecture conceptual model and how
it has been implemented through a Kubernetes operator plugged into
the FluxCD framework.


### Motivation and goals

The [Affordable5G][a5g] project adopts Open Source MANO (OSM) to
virtualise and orchestrate network functions, simplify infrastructure
operation, and achieve faster service deployment. One of the Affordable5G
objectives is to explore the continuous delivery of services through
GitOps workflows whereby the state of an OSM Kubernetes deployment
is described by version-controlled text files which a tool then interprets
to achieve the desired deployment state in the live OSM cluster.
Although OSM features a sophisticated toolset for the packaging,
deployment and operation of services, GitOps workflows for Kubernetes
network functions (KNFs) are not fully supported yet. Hence the need,
within Affordable5G, of a software to complement OSM’s capabilities
with GitOps.

Automated, version-controlled service delivery has several benefits.
Automation shortens deployment time and ensures reproducibility of
deployment states. In turn, reproducibility dramatically reduces the
time needed to recover from severe production incidents caused by
faulty deployments as the OSM cluster can swiftly be reverted to a
previous, known-to-be-working deployment state stored in the Git
repository. Thus, overall cluster stability and service availability
are enhanced. Moreover, the Git repository stores information about
who modified the OSM cluster state when, thus furnishing an audit
trail that may help to detect security breaches and failure to comply
with regulations such as GDPR.


### Conceptual model

OSM Ops is a cloud-native micro-service to implement GitOps workflows
within OSM. The basic idea is to describe the state of an OSM deployment
through version-controlled text files hosted in an online Git repository.
Each file declares a desired instantiation and runtime configuration
for some of the services in a specified OSM cluster. Collectively,
the files at a given Git revision describe the deployment state of
these services at a certain point in time. OSM Ops monitors the Git
repository in order to automatically reconcile the desired deployment
state with the actual live state of the OSM cluster. OSM Ops is implemented
as a [Kubernetes][k8s] operator that plugs into the [FluxCD][flux]
framework in order to leverage the rich Kubernetes/FluxCD GitOps
ecosystem. The following visual illustrates the context in which
OSM Ops operates and exemplifies the GitOps workflow resulting in
the creation and update of KNFs from version-controlled deployment
declarations.

![Architecture context diagram.][dia.ctx]

From the system administrator's perspective, the GitOps workflow is
as follows. She initially installs, through OSM packages, the deployment
descriptors (typically NSD, VNFD and KDU) for each KNF that she would
like to operate. Each KDU references a Helm chart describing the Kubernetes
resources (service, deployment, etc.) which constitute a KNF. To instantiate
a KNF, OSM has to be able to fetch the corresponding Helm chart. Usually,
Helm charts are maintained in an online repository that the system
administrator takes care of connecting to OSM—e.g. by adding the repository
to the OSM database with the `osm` command line tool. Likewise, since
Helm charts, in turn, reference container images, the system administrator
has to make sure that all the images required for her KNFs can be
downloaded from within the OSM cluster. The usual arrangement here
is that a container registry service provides the needed images to
OSM. Again, the system administrator takes care of this initial setup
step. The diagram exemplifies these initial installation and setup
steps with two sets of OSM deployment descriptors, one for OpenLDAP
and the other for TensorFlow, each referencing their respective Helm
charts in an online Git repository and, in turn, the charts reference
container images in a public Docker registry.

After provisioning KNF descriptors, the system administrator can then
edit text files to declare the desired deployment state of her KNFs.
Soon after she commits these files to the OSM Ops descriptor repository,
a background reconciliation process is set in motion that ultimately
results in NS instances running in the OSM cluster with the desired
deployment configuration. The diagram depicts a scenario where the
system administrator commits a new revision, `v6`, to the OSM Ops
descriptor repository. The `v6` files collectively declare that the
OSM cluster should run an OpenLDAP KNF with three replicas and a
TensorFlow KNF with one replica. As hinted by the diagram, the last
time the reconciliation process ran, it realised the deployment configuration
declared in revision `v5` which demanded an OpenLDAP KNF with two
replicas. Therefore to realise the `v6` configuration, the outcome
of the reconciliation process should be that another replica is added
to the existing OpenLDAP KNF and a brand new TensorFlow KNF is created
with one replica.

We now turn our attention to the reconciliation process that runs
behind the scenes. FluxCD detects any changes to the OSM Ops descriptor
repository and forwards new revisions to OSM Ops for processing. On
receiving a new revision, OSM Ops determines which KNFs to create
and which to update. It then calls the OSM cluster manager to actually
create or update the KNFs declared in that revision. In turn, the OSM
cluster manager orchestrates calls to Helm and Kubernetes to fulfill
the requested create and update operations which usually also involve
fetching Helm charts from a repository and pulling container images.
The diagram illustrates the reconciliation process for revision `v6`.
(Bear in mind, the diagram shows a conceptual, high-level message
flow, the next section provides a more accurate description.)


### Implementation overview

Having defined the abstract ideas, we are now ready to explain how
they have been realised. In a nutshell, OSM Ops is a Kubernetes operator
that gets notified of any changes to an online Git repository monitored
by FluxCD and then uses OSM’s north-bound interface (NBI) to realise
the KNF deployment configurations found in that repository.

These deployment configurations are declared through OSM Ops YAML
files. Each file specifies a desired instantiation and runtime configuration
(e.g. number of replicas) of a KNF previously defined within OSM by
installing suitable OSM descriptor packages, Helm charts, etc. For
example, the following YAML file demands that the live OSM cluster
run a 2-replica NS instance called `ldap` within the VIM identified
by the given VIM account and that the service be configured according
to the definitions found in the named OSM descriptors—the referenced
NSD, VNFD and KDU are actually defined in the OpenLDAP OSM packages
published by Telefonica.

```yaml
kind: NsInstance
name: ldap
description: Demo LDAP NS instance
nsdName: openldap_ns
vnfName: openldap
vimAccountName: mylocation1
kdu:
    name: ldap
    params:
        replicaCount: "2"
```

Source Controller is a FluxCD service that, among other things, manages
interactions with online Git repositories—e.g. repositories hosted
on GitHub or GitLab. OSM Ops depends on it both for monitoring repositories
and for fetching the repository content at a given revision. Source
Controller arranges a Kubernetes custom resource for each repository
that it monitors and then polls each repository to detect new revisions.
As soon as a new revision becomes available, Source Controller updates
the corresponding Git repository custom resource in Kubernetes.

OSM Ops implements the Kubernetes Operator interface to get notified
of any changes to Git repository custom resources. Thus, soon after
Source Controller updates a Git repository custom resource, Kubernetes
dispatches an update event to OSM Ops. This arrangement is akin to
the publish-subscribe pattern often found in messaging systems: Source
Controller, the publisher, sends a message to Kubernetes, the broker,
which results in the broker notifying OSM Ops, the subscriber. The
publisher and the subscriber have no knowledge of each other (no space
coupling) and communication is asynchronous (no time coupling).

At this point, OSM Ops enters the reconcile phase in which it tries
to align the deployment state declared in the OSM Ops YAML files with
that of the live OSM cluster. It fetches the content of the notified
Git revision from Source Controller as a tarball and then uses OSM's
NBI to transition the OSM cluster to the deployment state declared
in the OSM Ops YAML files found in the tarball. For each file, OSM
Ops determines whether to create or update the KNF specified in the
file and then configures it according to the KNF parameters given
in that file.

The UML communication diagram below summarises the typical workflow
through which OSM Ops turns the deployment state declared in a Git
repository into actual NS instances. The workflow begins with a system
administrator pushing a new revision, `v6`, to the online Git repository.
It then continues as just explained, with Source Controller updating
the Git custom resource, Kubernetes notifying OSM Ops and OSM Ops
calling the NBI to achieve the deployment state declared in `v6`.

![Implementation overview.][dia.impl]


### Rationale

What is the rationale behind our design decisions? A few explanatory
words are in order.

**TODO**
- evaluated two leading GitOps solutions: ArgoCD & FluxCD
- similar capabilities but ArgoCD comes with powerful UI
- convergence: the two projects will likely be merged in the
  future—ref merger plans
- FluxCD has better docs about extending it with custom functionality
  which is what in the end tipped the balance in its favour
- Go was a natural PL choice b/c FluxCD and K8s libs are both
  written in Go




[a5g]: https://www.affordable5g.eu/
    "Affordable5G"
[dia.ctx]: ./arch.context.png
[dia.impl]: ./arch.impl-overview.png
[flux]: https://fluxcd.io/
    "Flux - the GitOps family of projects"
[k8s]: https://en.wikipedia.org/wiki/Kubernetes
    "Kubernetes"
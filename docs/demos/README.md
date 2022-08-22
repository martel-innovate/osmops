Demos
-----
> Hands-on journeys through core features and deployment scenarios.

OSM Ops is a proof-of-concept tool to implement IaC workflows with
OSM. The basic idea behind the tool is to describe the state of an
OSM deployment through version-controlled text files which the tool
interprets to achieve the desired deployment state in a given OSM
cluster. If this is Greek to you, go read the architecture [intro][arch.intro]
before venturing in the madness below :-)

We've got a few demos you can run yourself to try out OSM Ops. The
overall goal of each demo is the same: show how an admin could just
edit YAML files in GitHub to create, update and configure KNFs in her
OSM cluster along the lines of the examples in the [intro][arch.intro]—you've
read that, haven't you :-) But each demo happens in a slightly different
deployment setting. So here are the demos:

* [Local clusters][demo.local]. Build a local OSM release 10 cluster
  and a separate Kind cluster to host OSM Ops. Then simulate commits
  to this repository to watch OSM Ops create and update an OpenLDAP
  KNF. All that in the comfort of your laptop.
* [Dev mode][demo.dev]. Same as above, but OSM Ops now runs directly on
  your box, outside the cluster. Comes in handy if you want to easily
  debug OSM Ops to see what it's actually up to.
* [Malaga][demo.malaga]. Demo given at the Affordable5G official review
  in November 2021 with a fully-fledged deployment in Malaga.
* [Packaging][demo.pack]. Build on the local clusters demo to showcase
  how OSM Ops can also create and update OSM packages from sources.




[arch.intro]: ../arch/intro.md
[demo.dev]: ./dev-mode.md
[demo.local]: ./local-clusters.md
[demo.malaga]: ./malaga.md
[demo.pack]: ./pack.md
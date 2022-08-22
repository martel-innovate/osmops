OSM Package Support
-------------------
> Kinda works, but it could be much better!

OSM Ops can create or update OSM packages from package source files
in an OSM Ops repo. For this to work, the operator has to arrange
package files into a directory tree according to some naming conventions.
Nothing is configurable at the moment, so you've got to stick to
the naming conventions down to a tee if you want to make OSM Ops
handle your OSM packages.

Eventually, we could reimplement the packaging functionality properly,
e.g. use a semantic approach (parse, interpret OSM files, etc.) rather
than naming conventions and guesswork.

Anyway, at the moment this functionality is actually stable and sort
of useable. If you'd like to give OSM Ops a shot at managing your
KNF and NS packages, read on.


### TL;DR

To make OSM Ops create or update OSM packages in your repo:

1. Put the files that make up a package in a directory right under
   `<target-dir>/osm-pkgs` where `target-dir` is the deployment
   target directory specified in `osm_ops_config.yaml`.
2. Name the package directory with a `_knf` suffix if it's a KNF
   package or `_ns` if it's an NS.
3. Use the directory name (including suffix) as a package ID in
   the YAML definitions.

If package `p2` depends on `p1`, name their directories in such a
way that `p2`'s directory name comes before `p1`'s in alphabetical
order.


### How it works

#### OSM package tree
OSM Ops expects package source files to be in a directory tree rooted
at `<target-dir>/osm-pkgs`. `target-dir` is the deployment target directory,
within your OSM Ops-managed repo, you specify in the `osm_ops_config.yaml`
config file whereas the `osm-pkgs` bit isn't configurable at the moment.
The source files that make up a package have to be in a directory right
under `osm-pkgs`. How you structure the package directory is up to you
(you could have sub-dirs if you wanted) but the way you name it tells
OSM Ops how to handle the package---more about it later.

Here's an example repo layout with an OSM package tree.

```
my-gitops-repo
 | -- osm_ops_config.yaml
 + -- deployment-target-dir
    + -- kdu
       | -- ldap.ops.yaml
    + -- osm-pkgs
        + -- openldap_knf
           | -- openldap_vnfd.yaml
        + -- openldap_ns
           | -- openldap_nsd.yaml
           | -- README.md
```

`my-gitops-repo` is your repo root dir, e.g. on GitHub it could be
hosted at https://github.com/c0c0n3/my-gitops-repo. `osm_ops_config.yaml`
is the usual OSM Ops config file. In this case it specifies a target
directory of `deployment-target-dir` as in the example below:

```yaml
targetDir: deployment-target-dir
fileExtensions:
  - .ops.yaml
connectionFile: /etc/osmops/nbi-connection.yaml
```

`osm-pkgs` contains the source files of two OSM packages, each in
its own directory. One package `openldap_knf` contains the YAML to
define a KNF for an OpenLDAP service. The other, `openldap_ns`,
contains the YAML to define an NS for the OpenLDAP KNF defined by
`openldap_knf` plus a standard README. Notice you don't need to
add an OSM `checksums.txt` file to each package source directory
since OSM Ops does that for you when uploading the package to OSM
as we'll see later. Finally, there's an `ldap.ops.yaml` file with
some instructions for OSM Ops to manage the deployment of the
OpenLDAP service defined through the above KNF and NS packages.

#### OSM package directory names
At the moment OSM Ops blindly assumes that any sub-directory of
`osm-pkgs` contains either a KNF or NS package. If the directory
name ends with `_knf`, OSM Ops treats the whole directory as a KNF
package. Likewise, if the directory name ends with `_ns`, OSM Ops
treats it as an NS package. (OSM Ops will report an error if the
directory name doesn't have an `_ns` or `_knf` suffix.)

OSM Ops also relies on another naming convention to figure out the
package ID. In fact, it assumes the directory name is also the package
ID declared in the KNF or NS YAML stanza.

So to make OSM Ops manage your package source, you have to:

* name the package directory with a `_knf` suffix if it's a KNF
  package or `_ns` if it's an NS;
* use the directory name (including suffix) as a package ID in
  the YAML definitions.

In our example layout above, `openldap_vnfd.yaml` uses the enclosing
directory name as an ID in the VNFD declaration

```yaml
vnfd:
  id: openldap_knf
# ... rest of the file
```

And as you've guessed already, `openldap_nsd.yaml` does pretty much
the same

```yaml
nsd:
  nsd:
  - id: openldap_ns
    name: openldap_ns
    vnfd-id:
    - openldap_knf
# ... rest of the file
```

#### OSM package dependencies
At the moment OSM Ops does **not explicitly** handle dependencies
among OSM packages. But it does process package directories in the
OSM package tree in alphabetical order. This way, the operator can
name package directories in such a way that if package `p2` depends
on `p1`, `p2`'s name comes before `p1`'s in alphabetical order.

We used this in our example layout. In fact, `openldap_knf` defines
a KNF that's then referenced by `openldap_ns`, so OSM Ops should
process `openldap_knf` before `openldap_ns`. And this is exactly
what happens because `openldap_knf < openldap_ns` in the alphabetical
order.

#### Processing a package tree
So if there's an OSM package tree directory (`osm-pkgs`), OSM Ops
will create or update any OSM packages found in there. To figure
out whether to create or update a package, OSM Ops queries the NBI
upfront to see what packages are there already. If a package source
is in `osm-pkgs` but not in OSM, then OSM Ops creates the package
in OSM from the `osm-pkgs` source, otherwise it's an update. OSM Ops
will skip processing packages if there's no `osm-pkgs` directory or
it has no sub-directories.

OSM Ops blindly assumes that any sub-directory `p` of the OSM package
tree root contains the source files of an OSM package. If `p` has
to be created, OSM Ops reads `p`'s contents to make a gzipped tar
archive in the OSM format (including assembling the `checksums.txt`
file) and then streams it to OSM NBI to create the package in OSM.
On the other hand, if `p` is to be updated, OSM Ops tries locating
the YAML file containing `p`'s VNFD or NSD definition, then uploads
that YAML to OSM.

**NOTE. Package update.**
It's kinda weird the way it works, but most likely I'm missing something.
In fact, our [initial implementation][pr.1] actually uploaded a tarball
to OSM not only for create operations but also for updates. As it turns
out, OSM client does something different when it comes to updating a
package. It tries finding a YAML file in the package dir, blindly assumes
it's a VNFD or NSD and PUTs it in OSM. What if there are other files
in the package? Well, I've got no idea why OSM client does that, but
I've changed our update implementation to be in line with OSM client's.
Have a look at OSM client's [VNFD][osm-client.vnfd] and [NSD][osm-client.nsd]
update implementation.


### How it could work

Surely this is a stopgap solution. Eventually we'll implement proper
(semantic) handling of packages and their dependencies. One obvious
approach would be to:

* parse OSM package definitions;
* interpret the parsed AST to build a dependency graph;
* extract a DAG `d[k]` for each graph component `g[k]`;
* topologically sort `d[k]` to get a sequence of nodes `s[k]`;
* process `s[k]` sequences in parallel.




[osm-client.nsd]: https://osm.etsi.org/gitlab/osm/osmclient/-/blob/master/osmclient/sol005/nsd.py
[osm-client.vnfd]: https://osm.etsi.org/gitlab/osm/osmclient/-/blob/master/osmclient/sol005/vnfd.py
[pr.1]: https://github.com/c0c0n3/source-watcher/pull/1

#!/bin/bash

multipass launch --name osm --cpus 2 --mem 6G --disk 40G 18.04

multipass mount ./ osm:/mnt/osm-install

# multipass exec osm -- cd /mnt/osm-install && ./patched.install_osm.sh 2>&1 | tee install.log
#                                               ^ sudo issue

multipass shell osm

# Base OSM install
#
# $ cd /mnt/osm-install
# $ ./patched.install_osm.sh 2>&1 | tee install.log

# KNF setup for an isolated K8s cluster, copy-pasted from:
# - https://osm.etsi.org/docs/user-guide/05-osm-usage.html#adding-kubernetes-cluster-to-osm
# but changed version to the actual K8s server version returned by `kubectl version`
#
# $ osm vim-create --name mylocation1 --user u --password p --tenant p --account_type dummy --auth_url http://localhost/dummy
# $ osm k8scluster-add cluster --creds .kube/config --vim mylocation1 --k8s-nets '{k8s_net1: null}' --version "v1.15.12" --description="Isolated K8s cluster in mylocation1"

# Some rops where to fetch Helm charts for KNF
#
# $ osm repo-add --type helm-chart --description "Bitnami repo" bitnami https://charts.bitnami.com/bitnami
# $ osm repo-add --type helm-chart --description "Cetic repo" cetic https://cetic.github.io/helm-charts
# $ osm repo-add --type helm-chart --description "Elastic repo" elastic https://helm.elastic.co

# To clean up:
#
# $ multipass stop osm
# $ multipass delete osm
# $ multipass purge

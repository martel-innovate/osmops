#!/bin/bash
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.
#
REPOSITORY_BASE=https://osm-download.etsi.org/repository/osm/debian
RELEASE=ReleaseTEN
REPOSITORY=stable
DOCKER_TAG=10
DEVOPS_PATH=/usr/share/osm-devops

function usage(){
    echo -e "usage: $0 [OPTIONS]"
    echo -e "Install OSM from binaries or source code (by default, from binaries)"
    echo -e "  OPTIONS"
    echo -e "     -h / --help:    print this help"
    echo -e "     -y:             do not prompt for confirmation, assumes yes"
    echo -e "     -r <repo>:      use specified repository name for osm packages"
    echo -e "     -R <release>:   use specified release for osm binaries (deb packages, lxd images, ...)"
    echo -e "     -u <repo base>: use specified repository url for osm packages"
    echo -e "     -k <repo key>:  use specified repository public key url"
    echo -e "     -b <refspec>:   install OSM from source code using a specific branch (master, v2.0, ...) or tag"
    echo -e "                     -b master          (main dev branch)"
    echo -e "                     -b v2.0            (v2.0 branch)"
    echo -e "                     -b tags/v1.1.0     (a specific tag)"
    echo -e "                     ..."
    echo -e "     -c <orchestrator> deploy osm services using container <orchestrator>. Valid values are <k8s> or <swarm>.  If -c is not used then osm will be deployed using default orchestrator. When used with --uninstall, osm services deployed by the orchestrator will be uninstalled"
    echo -e "     -s <stack name> or <namespace>  user defined stack name when installed using swarm or namespace when installed using k8s, default is osm"
    echo -e "     -H <VCA host>   use specific juju host controller IP"
    echo -e "     -S <VCA secret> use VCA/juju secret key"
    echo -e "     -P <VCA pubkey> use VCA/juju public key file"
    echo -e "     -C <VCA cacert> use VCA/juju CA certificate file"
    echo -e "     -A <VCA apiproxy> use VCA/juju API proxy"
    echo -e "     --vimemu:       additionally deploy the VIM emulator as a docker container"
    echo -e "     --elk_stack:    additionally deploy an ELK docker stack for event logging"
    echo -e "     --pla:          install the PLA module for placement support"
    echo -e "     -m <MODULE>:    install OSM but only rebuild the specified docker images (LW-UI, NBI, LCM, RO, MON, POL, KAFKA, MONGO, PROMETHEUS, PROMETHEUS-CADVISOR, KEYSTONE-DB, PLA, NONE)"
    echo -e "     -o <ADDON>:     ONLY (un)installs one of the addons (vimemu, elk_stack, k8s_monitor)"
    echo -e "     -O <openrc file/cloud name>: Install OSM to an OpenStack infrastructure. <openrc file/cloud name> is required. If a <cloud name> is used, the clouds.yaml file should be under ~/.config/openstack/ or /etc/openstack/"
    echo -e "     -N <openstack public network name/ID>: Public network name required to setup OSM to OpenStack"
    echo -e "     -D <devops path> use local devops installation path"
    echo -e "     -w <work dir>   Location to store runtime installation"
    echo -e "     -t <docker tag> specify osm docker tag (default is latest)"
    echo -e "     -l:             LXD cloud yaml file"
    echo -e "     -L:             LXD credentials yaml file"
    echo -e "     -K:             Specifies the name of the controller to use - The controller must be already bootstrapped"
    echo -e "     -d <docker registry URL> use docker registry URL instead of dockerhub"
    echo -e "     -p <docker proxy URL> set docker proxy URL as part of docker CE configuration"
    echo -e "     -T <docker tag> specify docker tag for the modules specified with option -m"
    echo -e "     --nocachelxdimages:  do not cache local lxd images, do not create cronjob for that cache (will save installation time, might affect instantiation time)"
    echo -e "     --nolxd:        do not install and configure LXD, allowing unattended installations (assumes LXD is already installed and confifured)"
    echo -e "     --nodocker:     do not install docker, do not initialize a swarm (assumes docker is already installed and a swarm has been initialized)"
    echo -e "     --nojuju:       do not juju, assumes already installed"
    echo -e "     --nodockerbuild:do not build docker images (use existing locally cached images)"
    echo -e "     --nohostports:  do not expose docker ports to host (useful for creating multiple instances of osm on the same host)"
    echo -e "     --nohostclient: do not install the osmclient"
    echo -e "     --uninstall:    uninstall OSM: remove the containers and delete NAT rules"
    echo -e "     --source:       install OSM from source code using the latest stable tag"
    echo -e "     --develop:      (deprecated, use '-b master') install OSM from source code using the master branch"
    echo -e "     --pullimages:   pull/run osm images from docker.io/opensourcemano"
    echo -e "     --k8s_monitor:  install the OSM kubernetes monitoring with prometheus and grafana"
    echo -e "     --volume:       create a VM volume when installing to OpenStack"
    echo -e "     --showopts:     print chosen options and exit (only for debugging)"
    echo -e "     --charmed:                   Deploy and operate OSM with Charms on k8s"
    echo -e "     [--bundle <bundle path>]:    Specify with which bundle to deploy OSM with charms (--charmed option)"
    echo -e "     [--k8s <kubeconfig path>]:   Specify with which kubernetes to deploy OSM with charms (--charmed option)"
    echo -e "     [--vca <name>]:              Specifies the name of the controller to use - The controller must be already bootstrapped (--charmed option)"
    echo -e "     [--lxd <yaml path>]:         Takes a YAML file as a parameter with the LXD Cloud information (--charmed option)"
    echo -e "     [--lxd-cred <yaml path>]:    Takes a YAML file as a parameter with the LXD Credentials information (--charmed option)"
    echo -e "     [--microstack]:              Installs microstack as a vim. (--charmed option)"
    echo -e "     [--overlay]:                 Add an overlay to override some defaults of the default bundle (--charmed option)"
    echo -e "     [--ha]:                      Installs High Availability bundle. (--charmed option)"
    echo -e "     [--tag]:                     Docker image tag. (--charmed option)"
    echo -e "     [--registry]:                Docker registry with optional credentials as user:pass@hostname:port (--charmed option)"

}

add_repo() {
  REPO_CHECK="^$1"
  grep "${REPO_CHECK/\[arch=amd64\]/\\[arch=amd64\\]}" /etc/apt/sources.list > /dev/null 2>&1
  if [ $? -ne 0 ]
  then
    need_packages_lw="software-properties-common apt-transport-https"
    echo -e "Checking required packages: $need_packages_lw"
    dpkg -l $need_packages_lw &>/dev/null \
      || ! echo -e "One or several required packages are not installed. Updating apt cache requires root privileges." \
      || sudo apt-get -q update \
      || ! echo "failed to run apt-get update" \
      || exit 1
    dpkg -l $need_packages_lw &>/dev/null \
      || ! echo -e "Installing $need_packages_lw requires root privileges." \
      || sudo apt-get install -y $need_packages_lw \
      || ! echo "failed to install $need_packages_lw" \
      || exit 1
    wget -qO - $REPOSITORY_BASE/$RELEASE/OSM%20ETSI%20Release%20Key.gpg | sudo apt-key add -
    sudo DEBIAN_FRONTEND=noninteractive add-apt-repository -y "$1" && sudo DEBIAN_FRONTEND=noninteractive apt-get update
    return 0
  fi

  return 1
}

clean_old_repo() {
dpkg -s 'osm-devops' &> /dev/null
if [ $? -eq 0 ]; then
  # Clean the previous repos that might exist
  sudo sed -i "/osm-download.etsi.org/d" /etc/apt/sources.list
fi
}

while getopts ":b:r:c:n:k:u:R:l:L:K:p:D:o:O:m:N:H:S:s:w:t:U:P:A:d:p:f:F:-: hy" o; do
    case "${o}" in
        D)
            DEVOPS_PATH="${OPTARG}"
            ;;
        r)
            REPOSITORY="${OPTARG}"
            ;;
        R)
            RELEASE="${OPTARG}"
            ;;
        u)
            REPOSITORY_BASE="${OPTARG}"
            ;;
        t)
            DOCKER_TAG="${OPTARG}"
            ;;
        -)
            [ "${OPTARG}" == "help" ] && usage && exit 0
            ;;
        :)
            echo "Option -$OPTARG requires an argument" >&2
            usage && exit 1
            ;;
        \?)
            echo -e "Invalid option: '-$OPTARG'\n" >&2
            usage && exit 1
            ;;
        h)
            usage && exit 0
            ;;
        *)
            ;;
    esac
done

clean_old_repo
add_repo "deb [arch=amd64] $REPOSITORY_BASE/$RELEASE $REPOSITORY devops"
sudo DEBIAN_FRONTEND=noninteractive apt-get -q update
sudo DEBIAN_FRONTEND=noninteractive apt-get install osm-devops
$DEVOPS_PATH/installers/full_install_osm.sh -R $RELEASE -r $REPOSITORY -u $REPOSITORY_BASE -D $DEVOPS_PATH -t $DOCKER_TAG "$@"
```bash

Last login: Sun Sep  5 19:08:04 on ttys015
mactel:source-watcher andrea$ multipass find
Image                       Aliases           Version          Description
snapcraft:core18                              20201111         Snapcraft builder for Core 18
snapcraft:core20                              20201111         Snapcraft builder for Core 20
snapcraft:core                                20210430         Snapcraft builder for Core 16
18.04                       bionic            20210825         Ubuntu 18.04 LTS
20.04                       focal,lts         20210825         Ubuntu 20.04 LTS
mactel:source-watcher andrea$ multipass launch --name osm --cpus 2 --mem 6G --disk 40G 18.04
One quick question before we launch … Would you like to help
the Multipass developers, by sending anonymous usage data?
This includes your operating system, which images you use,
the number of instances, their properties and how long you use them.
We’d also like to measure Multipass’s speed.

Send usage data (yes/no/Later)? no
Launched: osm

###########################################################################################
New Multipass 1.7.0 release
Workflows, auto-bridges and more...

Go here for more information: https://github.com/CanonicalLtd/multipass/releases/tag/v1.7.0
###########################################################################################
mactel:source-watcher andrea$ multipass exec osm -- bash
To run a command as administrator (user "root"), use "sudo <command>".
See "man sudo_root" for details.

ubuntu@osm:~$ lsb_release -a
No LSB modules are available.
Distributor ID:	Ubuntu
Description:	Ubuntu 18.04.5 LTS
Release:	18.04
Codename:	bionic
ubuntu@osm:~$ wget https://osm-download.etsi.org/ftp/osm-10.0-ten/install_osm.sh
--2021-09-06 10:11:29--  https://osm-download.etsi.org/ftp/osm-10.0-ten/install_osm.sh
Resolving osm-download.etsi.org (osm-download.etsi.org)... 195.238.226.47
Connecting to osm-download.etsi.org (osm-download.etsi.org)|195.238.226.47|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 9348 (9.1K) [text/x-sh]
Saving to: ‘install_osm.sh’

install_osm.sh                               100%[============================================================================================>]   9.13K  --.-KB/s    in 0.003s

2021-09-06 10:11:29 (3.12 MB/s) - ‘install_osm.sh’ saved [9348/9348]

ubuntu@osm:~$ chmod +x install_osm.sh
ubuntu@osm:~$ ./install_osm.sh 2>&1 | tee osm_install_log.txt
Checking required packages: software-properties-common apt-transport-https
Warning: apt-key output should not be parsed (stdout is not a terminal)
OK
Hit:1 http://archive.ubuntu.com/ubuntu bionic InRelease
Get:2 http://archive.ubuntu.com/ubuntu bionic-updates InRelease [88.7 kB]
Get:3 http://archive.ubuntu.com/ubuntu bionic-backports InRelease [74.6 kB]
Get:4 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease [4070 B]
Get:5 http://archive.ubuntu.com/ubuntu bionic/universe amd64 Packages [8570 kB]
Get:6 http://archive.ubuntu.com/ubuntu bionic/universe Translation-en [4941 kB]
Get:7 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable/devops amd64 Packages [479 B]
Get:8 http://archive.ubuntu.com/ubuntu bionic/multiverse amd64 Packages [151 kB]
Get:9 http://archive.ubuntu.com/ubuntu bionic/multiverse Translation-en [108 kB]
Get:10 http://archive.ubuntu.com/ubuntu bionic-updates/main amd64 Packages [2192 kB]
Get:11 http://archive.ubuntu.com/ubuntu bionic-updates/main Translation-en [430 kB]
Get:12 http://archive.ubuntu.com/ubuntu bionic-updates/universe amd64 Packages [1748 kB]
Get:13 http://archive.ubuntu.com/ubuntu bionic-updates/universe Translation-en [375 kB]
Get:14 http://archive.ubuntu.com/ubuntu bionic-updates/multiverse amd64 Packages [27.3 kB]
Get:15 http://archive.ubuntu.com/ubuntu bionic-updates/multiverse Translation-en [6808 B]
Get:16 http://security.ubuntu.com/ubuntu bionic-security InRelease [88.7 kB]
Get:17 http://archive.ubuntu.com/ubuntu bionic-backports/main amd64 Packages [10.0 kB]
Get:18 http://archive.ubuntu.com/ubuntu bionic-backports/main Translation-en [4764 B]
Get:19 http://archive.ubuntu.com/ubuntu bionic-backports/universe amd64 Packages [10.3 kB]
Get:20 http://archive.ubuntu.com/ubuntu bionic-backports/universe Translation-en [4588 B]
Get:21 http://security.ubuntu.com/ubuntu bionic-security/main amd64 Packages [1846 kB]
Get:22 http://security.ubuntu.com/ubuntu bionic-security/main Translation-en [338 kB]
Get:23 http://security.ubuntu.com/ubuntu bionic-security/universe amd64 Packages [1137 kB]
Get:24 http://security.ubuntu.com/ubuntu bionic-security/universe Translation-en [259 kB]
Get:25 http://security.ubuntu.com/ubuntu bionic-security/multiverse amd64 Packages [20.9 kB]
Get:26 http://security.ubuntu.com/ubuntu bionic-security/multiverse Translation-en [4732 B]
Fetched 22.4 MB in 5s (4674 kB/s)
Reading package lists...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Hit:1 http://archive.ubuntu.com/ubuntu bionic InRelease
Hit:2 http://archive.ubuntu.com/ubuntu bionic-updates InRelease
Hit:3 http://archive.ubuntu.com/ubuntu bionic-backports InRelease
Hit:4 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease
Hit:5 http://security.ubuntu.com/ubuntu bionic-security InRelease
Reading package lists...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Hit:1 http://archive.ubuntu.com/ubuntu bionic InRelease
Hit:2 http://archive.ubuntu.com/ubuntu bionic-updates InRelease
Hit:3 http://archive.ubuntu.com/ubuntu bionic-backports InRelease
Hit:4 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease
Hit:5 http://security.ubuntu.com/ubuntu bionic-security InRelease
Reading package lists...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Reading package lists...
Building dependency tree...
Reading state information...
The following NEW packages will be installed:
  osm-devops
0 upgraded, 1 newly installed, 0 to remove and 4 not upgraded.
Need to get 824 kB of archives.
After this operation, 9116 kB of additional disk space will be used.
Get:1 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable/devops amd64 osm-devops all 10.0.1-1 [824 kB]
Fetched 824 kB in 0s (2210 kB/s)
                                Selecting previously unselected package osm-devops.
(Reading database ... 60392 files and directories currently installed.)
Preparing to unpack .../osm-devops_10.0.1-1_all.deb ...
Unpacking osm-devops (10.0.1-1) ...
Setting up osm-devops (10.0.1-1) ...
Checking required packages: git wget curl tar
2021-09-06T10:12:14+02:00 INFO Waiting for automatic snapd restart...
jq 1.5+dfsg-1 from Canonical* installed
##  Mon Sep  6 10:12:19 CEST 2021 source: logging sourced
##  Mon Sep  6 10:12:19 CEST 2021 source: config sourced
##  Mon Sep  6 10:12:19 CEST 2021 source: container sourced
##  Mon Sep  6 10:12:19 CEST 2021 source: git_functions sourced
The installation will do the following
        1. Install and configure LXD
        2. Install juju
        3. Install docker CE
        4. Disable swap space
        5. Install and initialize Kubernetes
        as pre-requirements.
        Do you want to proceed (Y/n)? y
Installing lightweight build of OSM
Checking required packages: snapd
* Applying /etc/sysctl.d/10-console-messages.conf ...
kernel.printk = 4 4 1 7
* Applying /etc/sysctl.d/10-ipv6-privacy.conf ...
net.ipv6.conf.all.use_tempaddr = 2
net.ipv6.conf.default.use_tempaddr = 2
* Applying /etc/sysctl.d/10-kernel-hardening.conf ...
kernel.kptr_restrict = 1
* Applying /etc/sysctl.d/10-link-restrictions.conf ...
fs.protected_hardlinks = 1
fs.protected_symlinks = 1
* Applying /etc/sysctl.d/10-lxd-inotify.conf ...
fs.inotify.max_user_instances = 1024
* Applying /etc/sysctl.d/10-magic-sysrq.conf ...
kernel.sysrq = 176
* Applying /etc/sysctl.d/10-network-security.conf ...
net.ipv4.conf.default.rp_filter = 1
net.ipv4.conf.all.rp_filter = 1
net.ipv4.tcp_syncookies = 1
* Applying /etc/sysctl.d/10-ptrace.conf ...
kernel.yama.ptrace_scope = 1
* Applying /etc/sysctl.d/10-zeropage.conf ...
vm.mmap_min_addr = 65536
* Applying /usr/lib/sysctl.d/50-default.conf ...
net.ipv4.conf.all.promote_secondaries = 1
net.core.default_qdisc = fq_codel
* Applying /etc/sysctl.d/60-lxd-production.conf ...
fs.inotify.max_queued_events = 1048576
fs.inotify.max_user_instances = 1048576
fs.inotify.max_user_watches = 1048576
vm.max_map_count = 262144
kernel.dmesg_restrict = 1
net.ipv4.neigh.default.gc_thresh3 = 8192
net.ipv6.neigh.default.gc_thresh3 = 8192
net.core.bpf_jit_limit = 3000000000
kernel.keys.maxkeys = 2000
kernel.keys.maxbytes = 2000000
* Applying /etc/sysctl.d/99-cloudimg-ipv6.conf ...
net.ipv6.conf.all.use_tempaddr = 0
net.ipv6.conf.default.use_tempaddr = 0
* Applying /etc/sysctl.d/99-sysctl.conf ...
* Applying /etc/sysctl.conf ...
Reading package lists...
Building dependency tree...
Reading state information...
The following packages were automatically installed and are no longer required:
  dns-root-data dnsmasq-base ebtables libuv1 uidmap xdelta3
Use 'sudo apt autoremove' to remove them.
The following packages will be REMOVED:
  liblxc-common* liblxc1* lxcfs* lxd* lxd-client*
0 upgraded, 0 newly installed, 5 to remove and 4 not upgraded.
After this operation, 34.1 MB disk space will be freed.
(Reading database ... 61700 files and directories currently installed.)e ...
Removing lxd (3.0.3-0ubuntu1~18.04.1) ...
Removing lxd dnsmasq configuration
Removing lxcfs (3.0.3-0ubuntu1~18.04.2) ...
Removing lxd-client (3.0.3-0ubuntu1~18.04.1) ...
Removing liblxc-common (3.0.3-0ubuntu1~18.04.1) ...
Removing liblxc1 (3.0.3-0ubuntu1~18.04.1) ...
Processing triggers for man-db (2.8.3-2ubuntu0.1) ...
Processing triggers for libc-bin (2.27-3ubuntu1.4) ...
(Reading database ... 61454 files and directories currently installed.)
Purging configuration files for liblxc-common (3.0.3-0ubuntu1~18.04.1) ...
Purging configuration files for lxd (3.0.3-0ubuntu1~18.04.1) ...
Purging configuration files for lxcfs (3.0.3-0ubuntu1~18.04.2) ...
Processing triggers for systemd (237-3ubuntu10.51) ...
Processing triggers for ureadahead (0.100.0-21) ...
lxd 4.17 from Canonical* installed
To start your first instance, try: lxc launch ubuntu:18.04

Installing Docker CE ...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Reading package lists...
Building dependency tree...
Reading state information...
ca-certificates is already the newest version (20210119~18.04.1).
ca-certificates set to manually installed.
software-properties-common is already the newest version (0.96.24.32.14).
software-properties-common set to manually installed.
The following packages were automatically installed and are no longer required:
  dns-root-data dnsmasq-base ebtables libuv1 uidmap xdelta3
Use 'sudo apt autoremove' to remove them.
The following NEW packages will be installed:
  apt-transport-https
0 upgraded, 1 newly installed, 0 to remove and 4 not upgraded.
Need to get 4348 B of archives.
After this operation, 154 kB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu bionic-updates/universe amd64 apt-transport-https all 1.6.14 [4348 B]
Fetched 4348 B in 0s (54.8 kB/s)
                                Selecting previously unselected package apt-transport-https.
(Reading database ... 61437 files and directories currently installed.)
Preparing to unpack .../apt-transport-https_1.6.14_all.deb ...
Unpacking apt-transport-https (1.6.14) ...
Setting up apt-transport-https (1.6.14) ...
Warning: apt-key output should not be parsed (stdout is not a terminal)
OK
Hit:1 http://archive.ubuntu.com/ubuntu bionic InRelease
Get:2 https://download.docker.com/linux/ubuntu bionic InRelease [64.4 kB]
Hit:3 http://archive.ubuntu.com/ubuntu bionic-updates InRelease
Hit:4 http://archive.ubuntu.com/ubuntu bionic-backports InRelease
Hit:5 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease
Hit:6 http://security.ubuntu.com/ubuntu bionic-security InRelease
Get:7 https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages [19.8 kB]
Fetched 84.3 kB in 1s (127 kB/s)
Reading package lists...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Reading package lists...
Building dependency tree...
Reading state information...
The following packages were automatically installed and are no longer required:
  dns-root-data dnsmasq-base ebtables libuv1 uidmap xdelta3
Use 'sudo apt autoremove' to remove them.
The following additional packages will be installed:
  containerd.io docker-ce-cli docker-ce-rootless-extras docker-scan-plugin
  libltdl7 pigz
Suggested packages:
  aufs-tools cgroupfs-mount | cgroup-lite
Recommended packages:
  slirp4netns
The following NEW packages will be installed:
  containerd.io docker-ce docker-ce-cli docker-ce-rootless-extras
  docker-scan-plugin libltdl7 pigz
0 upgraded, 7 newly installed, 0 to remove and 4 not upgraded.
Need to get 96.7 MB of archives.
After this operation, 407 MB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu bionic/universe amd64 pigz amd64 2.4-1 [57.4 kB]
Get:2 https://download.docker.com/linux/ubuntu bionic/stable amd64 containerd.io amd64 1.4.9-1 [24.7 MB]
Get:3 http://archive.ubuntu.com/ubuntu bionic/main amd64 libltdl7 amd64 2.4.6-2 [38.8 kB]
Get:4 https://download.docker.com/linux/ubuntu bionic/stable amd64 docker-ce-cli amd64 5:20.10.8~3-0~ubuntu-bionic [38.8 MB]
Get:5 https://download.docker.com/linux/ubuntu bionic/stable amd64 docker-ce amd64 5:20.10.8~3-0~ubuntu-bionic [21.2 MB]
Get:6 https://download.docker.com/linux/ubuntu bionic/stable amd64 docker-ce-rootless-extras amd64 5:20.10.8~3-0~ubuntu-bionic [7911 kB]
Get:7 https://download.docker.com/linux/ubuntu bionic/stable amd64 docker-scan-plugin amd64 0.8.0~ubuntu-bionic [3888 kB]
Fetched 96.7 MB in 3s (28.7 MB/s)
                                 Selecting previously unselected package pigz.
(Reading database ... 61441 files and directories currently installed.)
Preparing to unpack .../0-pigz_2.4-1_amd64.deb ...
Unpacking pigz (2.4-1) ...
Selecting previously unselected package containerd.io.
Preparing to unpack .../1-containerd.io_1.4.9-1_amd64.deb ...
Unpacking containerd.io (1.4.9-1) ...
Selecting previously unselected package docker-ce-cli.
Preparing to unpack .../2-docker-ce-cli_5%3a20.10.8~3-0~ubuntu-bionic_amd64.deb ...
Unpacking docker-ce-cli (5:20.10.8~3-0~ubuntu-bionic) ...
Selecting previously unselected package docker-ce.
Preparing to unpack .../3-docker-ce_5%3a20.10.8~3-0~ubuntu-bionic_amd64.deb ...
Unpacking docker-ce (5:20.10.8~3-0~ubuntu-bionic) ...
Selecting previously unselected package docker-ce-rootless-extras.
Preparing to unpack .../4-docker-ce-rootless-extras_5%3a20.10.8~3-0~ubuntu-bionic_amd64.deb ...
Unpacking docker-ce-rootless-extras (5:20.10.8~3-0~ubuntu-bionic) ...
Selecting previously unselected package docker-scan-plugin.
Preparing to unpack .../5-docker-scan-plugin_0.8.0~ubuntu-bionic_amd64.deb ...
Unpacking docker-scan-plugin (0.8.0~ubuntu-bionic) ...
Selecting previously unselected package libltdl7:amd64.
Preparing to unpack .../6-libltdl7_2.4.6-2_amd64.deb ...
Unpacking libltdl7:amd64 (2.4.6-2) ...
Setting up containerd.io (1.4.9-1) ...
Created symlink /etc/systemd/system/multi-user.target.wants/containerd.service → /lib/systemd/system/containerd.service.
Setting up docker-ce-rootless-extras (5:20.10.8~3-0~ubuntu-bionic) ...
Setting up docker-scan-plugin (0.8.0~ubuntu-bionic) ...
Setting up libltdl7:amd64 (2.4.6-2) ...
Setting up docker-ce-cli (5:20.10.8~3-0~ubuntu-bionic) ...
Setting up pigz (2.4-1) ...
Setting up docker-ce (5:20.10.8~3-0~ubuntu-bionic) ...
Created symlink /etc/systemd/system/multi-user.target.wants/docker.service → /lib/systemd/system/docker.service.
Created symlink /etc/systemd/system/sockets.target.wants/docker.socket → /lib/systemd/system/docker.socket.
Processing triggers for libc-bin (2.27-3ubuntu1.4) ...
Processing triggers for systemd (237-3ubuntu10.51) ...
Processing triggers for man-db (2.8.3-2ubuntu0.1) ...
Processing triggers for ureadahead (0.100.0-21) ...
Adding user to group 'docker'
... restarted Docker service
Client: Docker Engine - Community
 Version:           20.10.8
 API version:       1.41
 Go version:        go1.16.6
 Git commit:        3967b7d
 Built:             Fri Jul 30 19:54:08 2021
 OS/Arch:           linux/amd64
 Context:           default
 Experimental:      true

Server: Docker Engine - Community
 Engine:
  Version:          20.10.8
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.16.6
  Git commit:       75249d8
  Built:            Fri Jul 30 19:52:16 2021
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.4.9
  GitCommit:        e25210fe30a0a703442421b0f60afac609f950a3
 runc:
  Version:          1.0.1
  GitCommit:        v1.0.1-0-g4144b63
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0
... Docker CE installation done
Creating folders for installation
Hit:1 https://download.docker.com/linux/ubuntu bionic InRelease
Hit:2 http://archive.ubuntu.com/ubuntu bionic InRelease
Hit:3 http://archive.ubuntu.com/ubuntu bionic-updates InRelease
Hit:4 http://archive.ubuntu.com/ubuntu bionic-backports InRelease
Hit:5 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease
Hit:6 http://security.ubuntu.com/ubuntu bionic-security InRelease
Reading package lists...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Reading package lists...
Building dependency tree...
Reading state information...
apt-transport-https is already the newest version (1.6.14).
The following packages were automatically installed and are no longer required:
  dns-root-data dnsmasq-base ebtables libuv1 uidmap xdelta3
Use 'sudo apt autoremove' to remove them.
0 upgraded, 0 newly installed, 0 to remove and 4 not upgraded.
Warning: apt-key output should not be parsed (stdout is not a terminal)
OK
Hit:1 http://archive.ubuntu.com/ubuntu bionic InRelease
Hit:2 https://download.docker.com/linux/ubuntu bionic InRelease
Hit:3 http://archive.ubuntu.com/ubuntu bionic-updates InRelease
Hit:4 http://archive.ubuntu.com/ubuntu bionic-backports InRelease
Hit:5 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease
Hit:7 http://security.ubuntu.com/ubuntu bionic-security InRelease
Get:6 https://packages.cloud.google.com/apt kubernetes-xenial InRelease [9383 B]
Get:8 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages [49.4 kB]
Fetched 58.8 kB in 1s (56.0 kB/s)
Reading package lists...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Hit:1 https://download.docker.com/linux/ubuntu bionic InRelease
Hit:2 http://archive.ubuntu.com/ubuntu bionic InRelease
Hit:3 http://archive.ubuntu.com/ubuntu bionic-updates InRelease
Hit:4 http://archive.ubuntu.com/ubuntu bionic-backports InRelease
Hit:5 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease
Hit:7 http://security.ubuntu.com/ubuntu bionic-security InRelease
Hit:6 https://packages.cloud.google.com/apt kubernetes-xenial InRelease
Reading package lists...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Installing Kubernetes Packages ...
Reading package lists...
Building dependency tree...
Reading state information...
The following packages were automatically installed and are no longer required:
  dns-root-data dnsmasq-base libuv1 uidmap xdelta3
Use 'sudo apt autoremove' to remove them.
The following additional packages will be installed:
  conntrack cri-tools kubernetes-cni socat
The following NEW packages will be installed:
  conntrack cri-tools kubeadm kubectl kubelet kubernetes-cni socat
0 upgraded, 7 newly installed, 0 to remove and 4 not upgraded.
Need to get 71.4 MB of archives.
After this operation, 302 MB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu bionic/main amd64 conntrack amd64 1:1.4.4+snapshot20161117-6ubuntu2 [30.6 kB]
Get:2 http://archive.ubuntu.com/ubuntu bionic/main amd64 socat amd64 1.7.3.2-2ubuntu2 [342 kB]
Get:3 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 cri-tools amd64 1.13.0-01 [8775 kB]
Get:4 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubernetes-cni amd64 0.8.7-00 [25.0 MB]
Get:5 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubelet amd64 1.15.0-00 [20.2 MB]
Get:6 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubectl amd64 1.15.0-00 [8763 kB]
Get:7 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubeadm amd64 1.15.0-00 [8246 kB]
Fetched 71.4 MB in 2s (32.2 MB/s)
                                 Selecting previously unselected package conntrack.
(Reading database ... 61694 files and directories currently installed.)
Preparing to unpack .../0-conntrack_1%3a1.4.4+snapshot20161117-6ubuntu2_amd64.deb ...
Unpacking conntrack (1:1.4.4+snapshot20161117-6ubuntu2) ...
Selecting previously unselected package cri-tools.
Preparing to unpack .../1-cri-tools_1.13.0-01_amd64.deb ...
Unpacking cri-tools (1.13.0-01) ...
Selecting previously unselected package kubernetes-cni.
Preparing to unpack .../2-kubernetes-cni_0.8.7-00_amd64.deb ...
Unpacking kubernetes-cni (0.8.7-00) ...
Selecting previously unselected package socat.
Preparing to unpack .../3-socat_1.7.3.2-2ubuntu2_amd64.deb ...
Unpacking socat (1.7.3.2-2ubuntu2) ...
Selecting previously unselected package kubelet.
Preparing to unpack .../4-kubelet_1.15.0-00_amd64.deb ...
Unpacking kubelet (1.15.0-00) ...
Selecting previously unselected package kubectl.
Preparing to unpack .../5-kubectl_1.15.0-00_amd64.deb ...
Unpacking kubectl (1.15.0-00) ...
Selecting previously unselected package kubeadm.
Preparing to unpack .../6-kubeadm_1.15.0-00_amd64.deb ...
Unpacking kubeadm (1.15.0-00) ...
Setting up conntrack (1:1.4.4+snapshot20161117-6ubuntu2) ...
Setting up kubernetes-cni (0.8.7-00) ...
Setting up cri-tools (1.13.0-01) ...
Setting up socat (1.7.3.2-2ubuntu2) ...
Setting up kubelet (1.15.0-00) ...
Created symlink /etc/systemd/system/multi-user.target.wants/kubelet.service → /lib/systemd/system/kubelet.service.
Setting up kubectl (1.15.0-00) ...
Setting up kubeadm (1.15.0-00) ...
Processing triggers for man-db (2.8.3-2ubuntu0.1) ...
kubelet set on hold.
kubeadm set on hold.
kubectl set on hold.
I0906 10:14:00.311428   10501 version.go:248] remote version is much newer: v1.22.1; falling back to: stable-1.15
[init] Using Kubernetes version: v1.15.12
[preflight] Running pre-flight checks
	[WARNING IsDockerSystemdCheck]: detected "cgroupfs" as the Docker cgroup driver. The recommended driver is "systemd". Please follow the guide at https://kubernetes.io/docs/setup/cri/
	[WARNING SystemVerification]: this Docker version is not on the list of validated versions: 20.10.8. Latest validated version: 18.09
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Activating the kubelet service
[certs] Using certificateDir folder "/etc/kubernetes/pki"
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [osm localhost] and IPs [192.168.64.19 127.0.0.1 ::1]
[certs] Generating "apiserver-etcd-client" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [osm localhost] and IPs [192.168.64.19 127.0.0.1 ::1]
[certs] Generating "etcd/healthcheck-client" certificate and key
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [osm kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 192.168.64.19]
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "sa" key and public key
[kubeconfig] Using kubeconfig folder "/etc/kubernetes"
[kubeconfig] Writing "admin.conf" kubeconfig file
[kubeconfig] Writing "kubelet.conf" kubeconfig file
[kubeconfig] Writing "controller-manager.conf" kubeconfig file
[kubeconfig] Writing "scheduler.conf" kubeconfig file
[control-plane] Using manifest folder "/etc/kubernetes/manifests"
[control-plane] Creating static Pod manifest for "kube-apiserver"
[control-plane] Creating static Pod manifest for "kube-controller-manager"
[control-plane] Creating static Pod manifest for "kube-scheduler"
[etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[apiclient] All control plane components are healthy after 25.511910 seconds
[upload-config] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config-1.15" in namespace kube-system with the configuration for the kubelets in the cluster
[upload-certs] Skipping phase. Please see --upload-certs
[mark-control-plane] Marking the node osm as control-plane by adding the label "node-role.kubernetes.io/master=''"
[mark-control-plane] Marking the node osm as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
[bootstrap-token] Using token: 1bjm75.9ghzdclhrx6enqgb
[bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstrap-token] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.64.19:6443 --token 1bjm75.9ghzdclhrx6enqgb \
    --discovery-token-ca-cert-hash sha256:439d23d440f5fe042d93485a9c94342d6eb934e1c051ff5f196842c3e5135688
Error from server (NotFound): namespaces "osm" not found
podsecuritypolicy.policy/psp.flannel.unprivileged created
clusterrole.rbac.authorization.k8s.io/flannel created
clusterrolebinding.rbac.authorization.k8s.io/flannel created
serviceaccount/flannel created
configmap/kube-flannel-cfg created
daemonset.apps/kube-flannel-ds created
node/osm untainted
error: error reading [/tmp/openebs.eW4VwQ]: recognized file extensions are [.json .yaml .yml]
Waiting for storageclass

### Mon Sep  6 10:22:04 CEST 2021 install_k8s_storageclass: FATAL error: Storageclass not ready after 400 seconds. Cannot install openebs
BACKTRACE:
### FATAL /usr/share/osm-devops/common/logging 39
### install_k8s_storageclass /usr/share/osm-devops/installers/full_install_osm.sh 848
### install_lightweight /usr/share/osm-devops/installers/full_install_osm.sh 1211
### main /usr/share/osm-devops/installers/full_install_osm.sh 1876
-------
ubuntu@osm:~$ exit
exit
mactel:source-watcher andrea$ mutlipass stop osm
-bash: mutlipass: command not found
mactel:source-watcher andrea$ multipass stop osm
mactel:source-watcher andrea$ multipass delete osm
mactel:source-watcher andrea$

```
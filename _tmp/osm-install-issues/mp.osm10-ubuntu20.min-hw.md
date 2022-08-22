```bash

Last login: Fri Sep  3 20:09:34 on ttys011
mactel:source-watcher andrea$ multipass launch --name osm --cpus 2 --mem 6G --disk 40G
Launched: osm
mactel:source-watcher andrea$ multipass exec osm -- bash
To run a command as administrator (user "root"), use "sudo <command>".
See "man sudo_root" for details.

ubuntu@osm:~$ wget https://osm-download.etsi.org/ftp/osm-10.0-ten/install_osm.sh
--2021-09-03 20:20:02--  https://osm-download.etsi.org/ftp/osm-10.0-ten/install_osm.sh
Resolving osm-download.etsi.org (osm-download.etsi.org)... 195.238.226.47
Connecting to osm-download.etsi.org (osm-download.etsi.org)|195.238.226.47|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 9348 (9.1K) [text/x-sh]
Saving to: ‘install_osm.sh’

install_osm.sh                               100%[============================================================================================>]   9.13K  --.-KB/s    in 0.003s

2021-09-03 20:20:02 (3.20 MB/s) - ‘install_osm.sh’ saved [9348/9348]

ubuntu@osm:~$ chmod +x install_osm.sh
ubuntu@osm:~$ ./install_osm.sh 2>&1 | tee osm_install_log.txt
Checking required packages: software-properties-common apt-transport-https
Warning: apt-key output should not be parsed (stdout is not a terminal)
OK
Hit:1 http://archive.ubuntu.com/ubuntu focal InRelease
Get:2 http://archive.ubuntu.com/ubuntu focal-updates InRelease [114 kB]
Get:3 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease [4070 B]
Get:4 http://security.ubuntu.com/ubuntu focal-security InRelease [114 kB]
Get:5 http://archive.ubuntu.com/ubuntu focal-backports InRelease [101 kB]
Get:6 http://archive.ubuntu.com/ubuntu focal/universe amd64 Packages [8628 kB]
Get:7 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable/devops amd64 Packages [479 B]
Get:8 http://archive.ubuntu.com/ubuntu focal/universe Translation-en [5124 kB]
Get:9 http://archive.ubuntu.com/ubuntu focal/universe amd64 c-n-f Metadata [265 kB]
Get:10 http://archive.ubuntu.com/ubuntu focal/multiverse amd64 Packages [144 kB]
Get:11 http://archive.ubuntu.com/ubuntu focal/multiverse Translation-en [104 kB]
Get:12 http://archive.ubuntu.com/ubuntu focal/multiverse amd64 c-n-f Metadata [9136 B]
Get:13 http://archive.ubuntu.com/ubuntu focal-updates/main amd64 Packages [1175 kB]
Get:14 http://archive.ubuntu.com/ubuntu focal-updates/main Translation-en [254 kB]
Get:15 http://archive.ubuntu.com/ubuntu focal-updates/main amd64 c-n-f Metadata [14.1 kB]
Get:16 http://archive.ubuntu.com/ubuntu focal-updates/universe amd64 Packages [853 kB]
Get:17 http://archive.ubuntu.com/ubuntu focal-updates/universe Translation-en [181 kB]
Get:18 http://archive.ubuntu.com/ubuntu focal-updates/universe amd64 c-n-f Metadata [18.8 kB]
Get:19 http://archive.ubuntu.com/ubuntu focal-updates/multiverse amd64 Packages [24.6 kB]
Get:20 http://archive.ubuntu.com/ubuntu focal-updates/multiverse Translation-en [6776 B]
Get:21 http://archive.ubuntu.com/ubuntu focal-updates/multiverse amd64 c-n-f Metadata [620 B]
Get:22 http://archive.ubuntu.com/ubuntu focal-backports/main amd64 Packages [2568 B]
Get:23 http://archive.ubuntu.com/ubuntu focal-backports/main Translation-en [1120 B]
Get:24 http://archive.ubuntu.com/ubuntu focal-backports/main amd64 c-n-f Metadata [400 B]
Get:25 http://archive.ubuntu.com/ubuntu focal-backports/restricted amd64 c-n-f Metadata [116 B]
Get:26 http://archive.ubuntu.com/ubuntu focal-backports/universe amd64 Packages [5812 B]
Get:27 http://archive.ubuntu.com/ubuntu focal-backports/universe Translation-en [2068 B]
Get:28 http://archive.ubuntu.com/ubuntu focal-backports/universe amd64 c-n-f Metadata [288 B]
Get:29 http://archive.ubuntu.com/ubuntu focal-backports/multiverse amd64 c-n-f Metadata [116 B]
Get:30 http://security.ubuntu.com/ubuntu focal-security/main amd64 Packages [830 kB]
Get:31 http://security.ubuntu.com/ubuntu focal-security/main Translation-en [162 kB]
Get:32 http://security.ubuntu.com/ubuntu focal-security/main amd64 c-n-f Metadata [8604 B]
Get:33 http://security.ubuntu.com/ubuntu focal-security/restricted amd64 Packages [374 kB]
Get:34 http://security.ubuntu.com/ubuntu focal-security/restricted Translation-en [53.7 kB]
Get:35 http://security.ubuntu.com/ubuntu focal-security/universe amd64 Packages [638 kB]
Get:36 http://security.ubuntu.com/ubuntu focal-security/universe Translation-en [101 kB]
Get:37 http://security.ubuntu.com/ubuntu focal-security/universe amd64 c-n-f Metadata [12.3 kB]
Get:38 http://security.ubuntu.com/ubuntu focal-security/multiverse amd64 Packages [21.9 kB]
Get:39 http://security.ubuntu.com/ubuntu focal-security/multiverse Translation-en [4948 B]
Get:40 http://security.ubuntu.com/ubuntu focal-security/multiverse amd64 c-n-f Metadata [540 B]
Fetched 19.4 MB in 4s (5011 kB/s)
Reading package lists...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Hit:1 http://archive.ubuntu.com/ubuntu focal InRelease
Hit:2 http://archive.ubuntu.com/ubuntu focal-updates InRelease
Hit:3 http://archive.ubuntu.com/ubuntu focal-backports InRelease
Hit:4 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease
Hit:5 http://security.ubuntu.com/ubuntu focal-security InRelease
Reading package lists...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Hit:1 http://archive.ubuntu.com/ubuntu focal InRelease
Hit:2 http://archive.ubuntu.com/ubuntu focal-updates InRelease
Hit:3 http://archive.ubuntu.com/ubuntu focal-backports InRelease
Hit:4 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease
Hit:5 http://security.ubuntu.com/ubuntu focal-security InRelease
Reading package lists...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Reading package lists...
Building dependency tree...
Reading state information...
The following NEW packages will be installed:
  osm-devops
0 upgraded, 1 newly installed, 0 to remove and 5 not upgraded.
Need to get 824 kB of archives.
After this operation, 9116 kB of additional disk space will be used.
Get:1 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable/devops amd64 osm-devops all 10.0.1-1 [824 kB]
Fetched 824 kB in 0s (2128 kB/s)
                                Selecting previously unselected package osm-devops.
(Reading database ... 63510 files and directories currently installed.)
Preparing to unpack .../osm-devops_10.0.1-1_all.deb ...
Unpacking osm-devops (10.0.1-1) ...
Setting up osm-devops (10.0.1-1) ...
Checking required packages: git wget curl tar
jq 1.5+dfsg-1 from Canonical* installed
##  Fri Sep  3 20:20:51 CEST 2021 source: logging sourced
##  Fri Sep  3 20:20:51 CEST 2021 source: config sourced
##  Fri Sep  3 20:20:51 CEST 2021 source: container sourced
##  Fri Sep  3 20:20:51 CEST 2021 source: git_functions sourced
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
* Applying /etc/sysctl.d/10-magic-sysrq.conf ...
kernel.sysrq = 176
* Applying /etc/sysctl.d/10-network-security.conf ...
net.ipv4.conf.default.rp_filter = 2
net.ipv4.conf.all.rp_filter = 2
* Applying /etc/sysctl.d/10-ptrace.conf ...
kernel.yama.ptrace_scope = 1
* Applying /etc/sysctl.d/10-zeropage.conf ...
vm.mmap_min_addr = 65536
* Applying /usr/lib/sysctl.d/50-default.conf ...
net.ipv4.conf.default.promote_secondaries = 1
sysctl: setting key "net.ipv4.conf.all.promote_secondaries": Invalid argument
net.ipv4.ping_group_range = 0 2147483647
net.core.default_qdisc = fq_codel
fs.protected_regular = 1
fs.protected_fifos = 1
* Applying /usr/lib/sysctl.d/50-pid-max.conf ...
kernel.pid_max = 4194304
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
* Applying /usr/lib/sysctl.d/protect-links.conf ...
fs.protected_fifos = 1
fs.protected_hardlinks = 1
fs.protected_regular = 2
fs.protected_symlinks = 1
* Applying /etc/sysctl.conf ...
Reading package lists...
Building dependency tree...
Reading state information...
Package 'lxcfs' is not installed, so not removed
Package 'lxd' is not installed, so not removed
Package 'lxd-client' is not installed, so not removed
Package 'liblxc1' is not installed, so not removed
0 upgraded, 0 newly installed, 0 to remove and 5 not upgraded.
snap "lxd" is already installed, see 'snap help refresh'
To start your first instance, try: lxc launch ubuntu:18.04

Installing Docker CE ...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Reading package lists...
Building dependency tree...
Reading state information...
ca-certificates is already the newest version (20210119~20.04.1).
ca-certificates set to manually installed.
software-properties-common is already the newest version (0.98.9.5).
software-properties-common set to manually installed.
The following NEW packages will be installed:
  apt-transport-https
0 upgraded, 1 newly installed, 0 to remove and 5 not upgraded.
Need to get 4680 B of archives.
After this operation, 162 kB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu focal-updates/universe amd64 apt-transport-https all 2.0.6 [4680 B]
Fetched 4680 B in 0s (72.8 kB/s)
                                Selecting previously unselected package apt-transport-https.
(Reading database ... 64818 files and directories currently installed.)
Preparing to unpack .../apt-transport-https_2.0.6_all.deb ...
Unpacking apt-transport-https (2.0.6) ...
Setting up apt-transport-https (2.0.6) ...
Warning: apt-key output should not be parsed (stdout is not a terminal)
OK
Hit:1 http://archive.ubuntu.com/ubuntu focal InRelease
Get:2 https://download.docker.com/linux/ubuntu focal InRelease [52.1 kB]
Hit:3 http://archive.ubuntu.com/ubuntu focal-updates InRelease
Hit:4 http://archive.ubuntu.com/ubuntu focal-backports InRelease
Hit:5 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease
Hit:6 http://security.ubuntu.com/ubuntu focal-security InRelease
Get:7 https://download.docker.com/linux/ubuntu focal/stable amd64 Packages [10.7 kB]
Fetched 62.9 kB in 1s (82.5 kB/s)
Reading package lists...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Reading package lists...
Building dependency tree...
Reading state information...
The following additional packages will be installed:
  containerd.io docker-ce-cli docker-ce-rootless-extras docker-scan-plugin
  pigz slirp4netns
Suggested packages:
  aufs-tools cgroupfs-mount | cgroup-lite
The following NEW packages will be installed:
  containerd.io docker-ce docker-ce-cli docker-ce-rootless-extras
  docker-scan-plugin pigz slirp4netns
0 upgraded, 7 newly installed, 0 to remove and 5 not upgraded.
Need to get 96.7 MB of archives.
After this operation, 406 MB of additional disk space will be used.
Get:1 https://download.docker.com/linux/ubuntu focal/stable amd64 containerd.io amd64 1.4.9-1 [24.7 MB]
Get:2 http://archive.ubuntu.com/ubuntu focal/universe amd64 pigz amd64 2.4-1 [57.4 kB]
Get:3 http://archive.ubuntu.com/ubuntu focal/universe amd64 slirp4netns amd64 0.4.3-1 [74.3 kB]
Get:4 https://download.docker.com/linux/ubuntu focal/stable amd64 docker-ce-cli amd64 5:20.10.8~3-0~ubuntu-focal [38.8 MB]
Get:5 https://download.docker.com/linux/ubuntu focal/stable amd64 docker-ce amd64 5:20.10.8~3-0~ubuntu-focal [21.2 MB]
Get:6 https://download.docker.com/linux/ubuntu focal/stable amd64 docker-ce-rootless-extras amd64 5:20.10.8~3-0~ubuntu-focal [7917 kB]
Get:7 https://download.docker.com/linux/ubuntu focal/stable amd64 docker-scan-plugin amd64 0.8.0~ubuntu-focal [3889 kB]
Fetched 96.7 MB in 2s (46.9 MB/s)
                                 Selecting previously unselected package pigz.
(Reading database ... 64822 files and directories currently installed.)
Preparing to unpack .../0-pigz_2.4-1_amd64.deb ...
Unpacking pigz (2.4-1) ...
Selecting previously unselected package containerd.io.
Preparing to unpack .../1-containerd.io_1.4.9-1_amd64.deb ...
Unpacking containerd.io (1.4.9-1) ...
Selecting previously unselected package docker-ce-cli.
Preparing to unpack .../2-docker-ce-cli_5%3a20.10.8~3-0~ubuntu-focal_amd64.deb ...
Unpacking docker-ce-cli (5:20.10.8~3-0~ubuntu-focal) ...
Selecting previously unselected package docker-ce.
Preparing to unpack .../3-docker-ce_5%3a20.10.8~3-0~ubuntu-focal_amd64.deb ...
Unpacking docker-ce (5:20.10.8~3-0~ubuntu-focal) ...
Selecting previously unselected package docker-ce-rootless-extras.
Preparing to unpack .../4-docker-ce-rootless-extras_5%3a20.10.8~3-0~ubuntu-focal_amd64.deb ...
Unpacking docker-ce-rootless-extras (5:20.10.8~3-0~ubuntu-focal) ...
Selecting previously unselected package docker-scan-plugin.
Preparing to unpack .../5-docker-scan-plugin_0.8.0~ubuntu-focal_amd64.deb ...
Unpacking docker-scan-plugin (0.8.0~ubuntu-focal) ...
Selecting previously unselected package slirp4netns.
Preparing to unpack .../6-slirp4netns_0.4.3-1_amd64.deb ...
Unpacking slirp4netns (0.4.3-1) ...
Setting up slirp4netns (0.4.3-1) ...
Setting up docker-scan-plugin (0.8.0~ubuntu-focal) ...
Setting up containerd.io (1.4.9-1) ...
Created symlink /etc/systemd/system/multi-user.target.wants/containerd.service → /lib/systemd/system/containerd.service.
Setting up docker-ce-cli (5:20.10.8~3-0~ubuntu-focal) ...
Setting up pigz (2.4-1) ...
Setting up docker-ce-rootless-extras (5:20.10.8~3-0~ubuntu-focal) ...
Setting up docker-ce (5:20.10.8~3-0~ubuntu-focal) ...
Created symlink /etc/systemd/system/multi-user.target.wants/docker.service → /lib/systemd/system/docker.service.
Created symlink /etc/systemd/system/sockets.target.wants/docker.socket → /lib/systemd/system/docker.socket.
Processing triggers for man-db (2.9.1-1) ...
Processing triggers for systemd (245.4-4ubuntu3.11) ...
Adding user to group 'docker'
... restarted Docker service
Client: Docker Engine - Community
 Version:           20.10.8
 API version:       1.41
 Go version:        go1.16.6
 Git commit:        3967b7d
 Built:             Fri Jul 30 19:54:27 2021
 OS/Arch:           linux/amd64
 Context:           default
 Experimental:      true

Server: Docker Engine - Community
 Engine:
  Version:          20.10.8
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.16.6
  Git commit:       75249d8
  Built:            Fri Jul 30 19:52:33 2021
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
Hit:1 https://download.docker.com/linux/ubuntu focal InRelease
Hit:2 http://archive.ubuntu.com/ubuntu focal InRelease
Hit:3 http://archive.ubuntu.com/ubuntu focal-updates InRelease
Hit:4 http://archive.ubuntu.com/ubuntu focal-backports InRelease
Hit:5 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease
Hit:6 http://security.ubuntu.com/ubuntu focal-security InRelease
Reading package lists...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Reading package lists...
Building dependency tree...
Reading state information...
apt-transport-https is already the newest version (2.0.6).
0 upgraded, 0 newly installed, 0 to remove and 5 not upgraded.
Warning: apt-key output should not be parsed (stdout is not a terminal)
OK
Hit:1 https://download.docker.com/linux/ubuntu focal InRelease
Hit:2 http://archive.ubuntu.com/ubuntu focal InRelease
Hit:3 http://archive.ubuntu.com/ubuntu focal-updates InRelease
Hit:4 http://archive.ubuntu.com/ubuntu focal-backports InRelease
Hit:5 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease
Hit:7 http://security.ubuntu.com/ubuntu focal-security InRelease
Get:6 https://packages.cloud.google.com/apt kubernetes-xenial InRelease [9383 B]
Get:8 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages [49.4 kB]
Fetched 58.8 kB in 1s (55.8 kB/s)
Reading package lists...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Hit:1 https://download.docker.com/linux/ubuntu focal InRelease
Hit:2 http://archive.ubuntu.com/ubuntu focal InRelease
Hit:3 http://archive.ubuntu.com/ubuntu focal-updates InRelease
Hit:4 http://archive.ubuntu.com/ubuntu focal-backports InRelease
Hit:5 https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease
Hit:7 http://security.ubuntu.com/ubuntu focal-security InRelease
Get:6 https://packages.cloud.google.com/apt kubernetes-xenial InRelease [9383 B]
Fetched 9383 B in 1s (10.7 kB/s)
Reading package lists...
W: Conflicting distribution: https://osm-download.etsi.org/repository/osm/debian/ReleaseTEN stable InRelease (expected stable but got )
Installing Kubernetes Packages ...
Reading package lists...
Building dependency tree...
Reading state information...
The following additional packages will be installed:
  conntrack cri-tools ebtables kubernetes-cni socat
Suggested packages:
  nftables
The following NEW packages will be installed:
  conntrack cri-tools ebtables kubeadm kubectl kubelet kubernetes-cni socat
0 upgraded, 8 newly installed, 0 to remove and 5 not upgraded.
Need to get 71.5 MB of archives.
After this operation, 303 MB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu focal/main amd64 conntrack amd64 1:1.4.5-2 [30.3 kB]
Get:2 http://archive.ubuntu.com/ubuntu focal/main amd64 ebtables amd64 2.0.11-3build1 [80.3 kB]
Get:3 http://archive.ubuntu.com/ubuntu focal/main amd64 socat amd64 1.7.3.3-2 [323 kB]
Get:4 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 cri-tools amd64 1.13.0-01 [8775 kB]
Get:5 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubernetes-cni amd64 0.8.7-00 [25.0 MB]
Get:6 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubelet amd64 1.15.0-00 [20.2 MB]
Get:7 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubectl amd64 1.15.0-00 [8763 kB]
Get:8 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubeadm amd64 1.15.0-00 [8246 kB]
Fetched 71.5 MB in 2s (32.8 MB/s)
                                 Selecting previously unselected package conntrack.
(Reading database ... 65073 files and directories currently installed.)
Preparing to unpack .../0-conntrack_1%3a1.4.5-2_amd64.deb ...
Unpacking conntrack (1:1.4.5-2) ...
Selecting previously unselected package cri-tools.
Preparing to unpack .../1-cri-tools_1.13.0-01_amd64.deb ...
Unpacking cri-tools (1.13.0-01) ...
Selecting previously unselected package ebtables.
Preparing to unpack .../2-ebtables_2.0.11-3build1_amd64.deb ...
Unpacking ebtables (2.0.11-3build1) ...
Selecting previously unselected package kubernetes-cni.
Preparing to unpack .../3-kubernetes-cni_0.8.7-00_amd64.deb ...
Unpacking kubernetes-cni (0.8.7-00) ...
Selecting previously unselected package socat.
Preparing to unpack .../4-socat_1.7.3.3-2_amd64.deb ...
Unpacking socat (1.7.3.3-2) ...
Selecting previously unselected package kubelet.
Preparing to unpack .../5-kubelet_1.15.0-00_amd64.deb ...
Unpacking kubelet (1.15.0-00) ...
Selecting previously unselected package kubectl.
Preparing to unpack .../6-kubectl_1.15.0-00_amd64.deb ...
Unpacking kubectl (1.15.0-00) ...
Selecting previously unselected package kubeadm.
Preparing to unpack .../7-kubeadm_1.15.0-00_amd64.deb ...
Unpacking kubeadm (1.15.0-00) ...
Setting up conntrack (1:1.4.5-2) ...
Setting up kubectl (1.15.0-00) ...
Setting up ebtables (2.0.11-3build1) ...
Setting up socat (1.7.3.3-2) ...
Setting up cri-tools (1.13.0-01) ...
Setting up kubernetes-cni (0.8.7-00) ...
Setting up kubelet (1.15.0-00) ...
Created symlink /etc/systemd/system/multi-user.target.wants/kubelet.service → /lib/systemd/system/kubelet.service.
Setting up kubeadm (1.15.0-00) ...
Processing triggers for man-db (2.9.1-1) ...
kubelet set on hold.
kubeadm set on hold.
kubectl set on hold.
I0903 20:22:12.541425    9240 version.go:248] remote version is much newer: v1.22.1; falling back to: stable-1.15
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
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [osm localhost] and IPs [192.168.64.19 127.0.0.1 ::1]
[certs] Generating "etcd/healthcheck-client" certificate and key
[certs] Generating "apiserver-etcd-client" certificate and key
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [osm localhost] and IPs [192.168.64.19 127.0.0.1 ::1]
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [osm kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 192.168.64.19]
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
[apiclient] All control plane components are healthy after 26.006114 seconds
[upload-config] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config-1.15" in namespace kube-system with the configuration for the kubelets in the cluster
[upload-certs] Skipping phase. Please see --upload-certs
[mark-control-plane] Marking the node osm as control-plane by adding the label "node-role.kubernetes.io/master=''"
[mark-control-plane] Marking the node osm as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
[bootstrap-token] Using token: 4tmh7l.ishe157blg1ilg1q
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

kubeadm join 192.168.64.19:6443 --token 4tmh7l.ishe157blg1ilg1q \
    --discovery-token-ca-cert-hash sha256:3bf3f4d5d3483c4a9eb1d0281caf1d9353b515b163b2c38afaa046294a24b01b
Error from server (NotFound): namespaces "osm" not found
podsecuritypolicy.policy/psp.flannel.unprivileged created
clusterrole.rbac.authorization.k8s.io/flannel created
clusterrolebinding.rbac.authorization.k8s.io/flannel created
serviceaccount/flannel created
configmap/kube-flannel-cfg created
daemonset.apps/kube-flannel-ds created
node/osm untainted
error: error reading [/tmp/openebs.CeXhWG]: recognized file extensions are [.json .yaml .yml]
Waiting for storageclass

### Fri Sep  3 20:30:21 CEST 2021 install_k8s_storageclass: FATAL error: Storageclass not ready after 400 seconds. Cannot install openebs
BACKTRACE:
### FATAL /usr/share/osm-devops/common/logging 39
### install_k8s_storageclass /usr/share/osm-devops/installers/full_install_osm.sh 848
### install_lightweight /usr/share/osm-devops/installers/full_install_osm.sh 1211
### main /usr/share/osm-devops/installers/full_install_osm.sh 1876
-------
ubuntu@osm:~$ exit
exit
mactel:source-watcher andrea$ multipass stop osm
mactel:source-watcher andrea$ multipass delete osm
mactel:source-watcher andrea$ multipass purge
mactel:source-watcher andrea$ multipass purge
mactel:source-watcher andrea$

```
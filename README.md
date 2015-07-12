# dockerfileview - a public Dockerfile viewer

dockerfileview is a dead simple command line tool to view public Dockerfile. `dockerfileview` command will display Docerfile of specified base image by parsing `FROM` keyword and fetching Dockerfile again until `FROM scratch` statement is found, to see all the instructions which constructs specified base image.

![rubima wars pic](http://k.swd.cc/dockerfileview/resource/screenshot/example-usage.gif)

## Installation

To install dockerfileview, please use `go get`.

```
$ go get github.com/mitchellh/gox
```

If you have not installed go on your system, precompiled executables available at [release page](https://github.com/remore/dockerfileview/releases) is for you.

## Usage

dockerfileview can take a local file path for Dockerfile or public docker image name.

```
$ dockerfileview /path/to/Dockerfile
...
$ dockerfileview centos
...
$ dockerfileview nginx:1.9.2
...
$ dockerfileview jwilder/nginx-proxy
```

In case of redirecting standard output to to other command such as `less` or `tail` via pipe, `--text` option is recommended to use.

```bash
$ dockerfileview --text centos | less

#
# centos
# https://raw.githubusercontent.com/CentOS/sig-cloud-instance-images/0a6a7fa816e834b29222fce2df0b858ab1b97a87/docker/Dockerfile
#

FROM scratch
MAINTAINER The CentOS Project <cloud-ops@centos.org> - ami_creator
ADD centos-7-20150616_1752-docker.tar.xz /
# Volumes for systemd
# VOLUME ["/run", "/tmp"]

# Environment for systemd
# ENV container=docker

# For systemd usage this changes to /usr/sbin/init
# Keeping it as /bin/bash for compatability with previous


CMD ["/bin/bash"]

#
# (Direct Input from CLI)
#

FROM centos
```

## Limitations(there are certain types of unsupported image name)

Please be noted that regarding with docker image name, official images and latest image registered at public docker registry are supported. Meanwhile, here is the list of example image names we doesn't support.

- public_user/repo:<tag name>
- localhost:50111/foobar

This is because internally dockerfileview command will only fetch public information by connection to either `registry.hub.docker.com` or `raw.githubusercontent.com` to retrieve public Dockerfile.

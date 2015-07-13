# dockerfileview - a Dockerfile viewer to trace ancestry of the base image

dockerfileview is a dead simple command line tool enable you to trace ancestry of the base image. Here is an example.

```
$ dockerfileview nginx:1.9.2

#
# 🐳  debian:jessie
# https://raw.githubusercontent.com/tianon/docker-brew-debian/e9bafb113f432c48c7e86c616424cb4b2f2c7a51/jessie/Dockerfile
#

FROM scratch
ADD rootfs.tar.xz /
CMD ["/bin/bash"]

#
# 🐳  nginx:1.9.2
# https://raw.githubusercontent.com/nginxinc/docker-nginx/1eea9f7d082dff426e7923a90138de804038266d/Dockerfile
#

FROM debian:jessie

MAINTAINER NGINX Docker Maintainers "docker-maint@nginx.com"

 .
 .
 .
```

There are tons of useful public docker images such as official images(e.g. `ubuntu:14.04` or `centos:latest`) and personal images hosted at registry.hub.docker.com(e.g. `jwilder/nginx-proxy`), but to download Dockerfiles of those public images are boring and tiresome work. To make matters worse, you sometimes need to trace ancestry to investigate "base image of base image of current image" kind of thing manually. This is where dockerfileview comes in.

![command line example](http://k.swd.cc/dockerfileview/resource/screenshot/example-usage.gif)

## Installation

To install dockerfileview, please use `go get`.

```
$ go get github.com/remore/dockerfileview
```

If you have not installed go on your system, precompiled executables are available at [release page](https://github.com/remore/dockerfileview/releases) is for you. Or, simply type `docker run` command such as:

```
$ docker run remore/dockerfileview dockerfileview ubuntu:14.04
```

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

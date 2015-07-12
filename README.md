# dockerfileview - a public Dockerfile viewer

`dockerfileview` is a dead simple command line tool to view public Dockerfile. `dockerfileview` will keep showing public Dockerfiles by parsing `FROM` keyword of each Dockerfile recursively until `FROM scratch` statement is found.

![rubima wars pic](http://k.swd.cc/burn/resource/screenshot/rubima-wars.png)

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

## About Input Parameter and Limitations(Unsupported Image Name)

dockerfileview takes either local file path for Dockerfile(such as `./Dockerfile`) or public docker image name(e.g. `centos`, `nginx:1.9.2`, ` jwilder/nginx-proxy`) as an input parameter. Regarding with docker image name, official images and latest image registered at public docker registry are supported.

Meanwhile, here is the list of example image names we doesn't support.

- public_user/repo:<tag name>
- localhost:50111/foobar

Internally `dockerfileview` command will fetch either `registry.hub.docker.com` or `raw.githubusercontent.com` to retrieve public Dockerfile.

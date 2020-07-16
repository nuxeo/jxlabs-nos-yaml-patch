FROM centos:7

RUN yum install -y git

ENTRYPOINT ["yp"]

COPY build/linux/yaml-patch /usr/local/bin

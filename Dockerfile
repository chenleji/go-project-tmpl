FROM centos:7
MAINTAINER chenleji@gmail.com

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY ./go-project-tmpl /

ENTRYPOINT /go-project-tmpl


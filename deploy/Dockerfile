FROM centos:7

WORKDIR /opt

COPY ../build/xdoubletest /opt

COPY ../example/prod.json /opt

RUN xdoubletest -c prod.json

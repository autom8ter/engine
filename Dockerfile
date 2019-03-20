FROM golang:1.11

COPY plugins /.plugins

RUN go get github.com/autom8ter/engine/enginectl

ENTRYPOINT [ "enginectl", "init" ]

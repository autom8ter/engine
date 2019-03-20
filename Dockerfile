FROM golang:1.11

WORKDIR app
ENV HOME=/app
RUN git clone https://github.com/autom8ter/plugins.git && mv plugins .plugins
RUN go get github.com/autom8ter/engine/enginectl
ENTRYPOINT [ "enginectl" ]
FROM golang:alpine

RUN apk update \
  && apk add git bash

WORKDIR app
ENV HOME=/app
RUN git clone https://github.com/autom8ter/plugins.git
RUN mv plugins .plugins
RUN go get github.com/autom8ter/engine/enginectl
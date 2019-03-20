FROM alpine
RUN apk update \
  && apk add git wget curl
RUN git clone https://github.com/autom8ter/plugins.git && mv plugins .plugins
RUN curl -L -o enginectl https://github.com/autom8ter/engine/releases/download/v1.0/enginectl
RUN chmod +x enginectl && mv enginectl /usr/local/bin
RUN enginectl
ENTRYPOINT [ "enginectl" ]
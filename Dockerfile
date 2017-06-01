FROM alpine:3.5

ADD . /go-soil
RUN \
  apk add --update git go make gcc musl-dev linux-headers && \
  (cd go-soil && make gsoil)                           && \
  cp go-soil/build/bin/gsoil /gsoil                     && \
  apk del git go make gcc musl-dev linux-headers          && \
  rm -rf /go-soil && rm -rf /var/cache/apk/*

EXPOSE 39421
EXPOSE 30403
EXPOSE 30403/udp

ENTRYPOINT ["/gsoil"]

FROM alpine:3.5

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.5 https://github.com/ethereum/go-soil && \
  (cd go-soil && make gsoil) && \
  cp go-soil/build/bin/gsoil /gsoil && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /go-soil && rm -rf /var/cache/apk/*

EXPOSE 39421
EXPOSE 30403

ENTRYPOINT ["/gsoil"]

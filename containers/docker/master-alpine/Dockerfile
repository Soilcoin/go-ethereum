FROM alpine:3.5

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.6 https://github.com/Soilcoin/go-soil && \
  (cd go-soil && make gsoil) && \
  cp go-ethereum/build/bin/gsoil /gsoil && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /go-ethereum && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30403

ENTRYPOINT ["/geth"]

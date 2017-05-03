FROM ubuntu:xenial

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang git make gcc libc-dev ca-certificates && \
  git clone --depth 1 https://github.com/ethereum/go-soil && \
  (cd go-soil && make gsoil) && \
  cp go-soil/build/bin/gsoil /gsoil && \
  apt-get remove -y golang git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /go-soil

EXPOSE 39421
EXPOSE 30403

ENTRYPOINT ["/gsoil"]

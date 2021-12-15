FROM alpine
ARG TARBALL="v5.1.3.tar.gz"
WORKDIR /src
RUN apk add --no-cache build-base curl perl && \
  curl -sL https://github.com/kdlucas/byte-unixbench/archive/refs/tags/${TARBALL} | tar --strip-components=1 -zxf - && \
  cd UnixBench && make

WORKDIR /src/UnixBench
ENTRYPOINT ["./Run"]
CMD []

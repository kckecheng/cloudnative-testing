FROM alpine
ARG TARBALL="vdbench50407.zip"
WORKDIR /src
COPY ${TARBALL} .
RUN apk add --no-cache bash openjdk17 && \
  unzip ${TARBALL} -d vdbench && rm -f ${TARBALL}
ENTRYPOINT ["sh", "vdbench/vdbench"]
CMD []

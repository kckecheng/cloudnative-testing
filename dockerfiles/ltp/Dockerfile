FROM centos
ARG VERSION=20210927
WORKDIR /src
RUN yum install -y gcc make pkgconf autoconf automake bison flex m4 kernel-headers glibc-headers && \
  curl -sL https://github.com/linux-test-project/ltp/releases/download/${VERSION}/ltp-full-${VERSION}.tar.xz | tar -Jxf - && \
  cd ltp-full-${VERSION} && make autotools && ./configure && make && make install && \
  cd .. && rm -rf ltp-full-${VERSION} && yum clean all -y
WORKDIR /opt/ltp
ENTRYPOINT ["./runltp"]
CMD ["--help"]


FROM alpine
ARG TARBALL="bonnie++_1.04.tgz"
WORKDIR /src
RUN apk add --no-cache build-base g++ make curl && \
  curl -sL  https://www.coker.com.au/bonnie++/${TARBALL} | tar -zxf - && \
  src_dir=$(find . -type d -name "bonnie++*") && cd ${src_dir} && \
  ./configure && make && make install && cd .. && \
  rm -rf ${src_dir}
ENTRYPOINT ["bonnie++"]
CMD ["sh", "-c"]

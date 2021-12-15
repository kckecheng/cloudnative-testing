FROM alpine
RUN apk add --no-cache iperf3
EXPOSE 5201/tcp
EXPOSE 5201/udp
ENTRYPOINT ["/usr/bin/iperf3"]
CMD ["--help"]

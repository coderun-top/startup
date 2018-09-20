FROM alpine:3.4

ADD dist/startup /bin/
ENTRYPOINT ["/bin/startup"]

FROM daocloud.io/library/golang:1.9

COPY startup /go/bin

CMD ["startup"]

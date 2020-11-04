FROM alpine:latest as builder

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/tianxinbaiyun/mws
COPY ./mws $GOPATH/src/github.com/tianxinbaiyun/mws/mws

EXPOSE 8000
ENTRYPOINT ["./mws"]

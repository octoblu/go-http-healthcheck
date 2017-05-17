FROM golang:1.8
MAINTAINER Octoblu, Inc. <docker@octoblu.com>
WORKDIR /go/src/github.com/octoblu/shorter-url

COPY . /go/src/github.com/octoblu/go-http-healthcheck

RUN env CGO_ENABLED=0 go build -o go-http-healthcheck -a -ldflags '-s' .

CMD ["./go-http-healthcheck"]

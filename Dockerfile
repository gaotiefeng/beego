FROM  golang:1.12-alpine as bulider

MAINTAINER  Gaotiefeng "1096392101@qq.com"

ENV GOPROXY https://goproxy.io
ENV GO111MODULE=on

WORKDIR $GOPATH/src/beego-api

ADD . .
ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . $GOPATH/src/beego-api

EXPOSE  8080

ENTRYPOINT go run main.go
FROM golang:1.16.3-alpine
RUN apk add --update &&  apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app

RUN go get -u gorm.io/gorm && \ 
    go get -u gorm.io/driver/mysql && \ 
    go get github.com/labstack/echo/v4 && \ 
    go get github.com/labstack/echo/v4/middleware
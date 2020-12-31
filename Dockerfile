FROM golang:alpine

MAINTAINER Maintainer

ENV GIN_MODE=release
ENV PORT=3004

WORKDIR /go/src/crawler-docker
COPY main.go .

RUN apk update && apk add --no-cache git
RUN go get github.com/ryulitaro/crawler
RUN go get github.com/gin-gonic/gin

RUN go build main.go

EXPOSE $PORT

ENTRYPOINT ["./main"]
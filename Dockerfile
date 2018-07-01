FROM golang:1.10.3-alpine

LABEL maintainer "aimof <aimof.aimof@mgail.com">

RUN apk update --no-cache

COPY . /go/src/github.com/aimof/catalpha/
WORKDIR /go/src/github.com/aimof/catalpha/cmd/catalpha/

RUN go build

CMD ./catalpha

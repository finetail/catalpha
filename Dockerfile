FROM golang:1.10.3-alpine

LABEL maintainer "aimof <aimof.aimof@gmail.com>"

RUN apk update --no-cache && \
    apk add git

COPY . /go/src/github.com/aimof/catalpha/
WORKDIR /go/src/github.com/aimof/catalpha/cmd/catalpha/

RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure && \
    go build

CMD ./catalpha

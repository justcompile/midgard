FROM golang:1.13-alpine as deps

ENV GOLINT_VERSION 1.21.0

ENV GO111MODULE on

RUN mkdir -p $GOPATH/src/github.com/justcompile/midgard && \ 
    mkdir -p /build && \
    apk --no-cache add curl git bash gcc libc-dev ca-certificates && \
    update-ca-certificates && \
    curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v$GOLINT_VERSION

FROM deps as common

WORKDIR $GOPATH/src/github.com/justcompile/midgard/common

COPY ./scripts $GOPATH/src/github.com/justcompile/midgard/scripts

ADD ./common/go.mod $GOPATH/src/github.com/justcompile/midgard/common/go.mod
ADD ./common/go.sum $GOPATH/src/github.com/justcompile/midgard/common/go.sum

RUN go mod download

COPY ./common $GOPATH/src/github.com/justcompile/midgard/common

FROM common as web

WORKDIR $GOPATH/src/github.com/justcompile/midgard/web

ADD ./web/go.mod $GOPATH/src/github.com/justcompile/midgard/web/go.mod
ADD ./web/go.sum $GOPATH/src/github.com/justcompile/midgard/web/go.sum

RUN go mod download

COPY ./web $GOPATH/src/github.com/justcompile/midgard/web

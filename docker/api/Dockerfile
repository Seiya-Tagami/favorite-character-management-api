ARG GO_VERSION=1.20.7
ARG ALPINE_VERSION=3.18

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}

WORKDIR /app

RUN apk update \
&& apk add --no-cache git \
&& apk add bash 

RUN go install github.com/cosmtrek/air@latest

COPY app/go.mod .
COPY app/go.sum .

RUN go mod download
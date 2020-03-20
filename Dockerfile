FROM golang:1.14.0-alpine3.11

RUN apk add --no-cache git

EXPOSE 8080
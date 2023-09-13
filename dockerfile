FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

WORKDIR /lagoserv

COPY usermgmt usermgmt
COPY usermgmt_server usermgmt_server

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o lagoserv usermgmt_server/user_server.go

EXPOSE 50051

CMD ./lagoserv
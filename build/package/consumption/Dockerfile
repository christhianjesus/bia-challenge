# syntax=docker/dockerfile:1
FROM golang:1.19-bullseye

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/

RUN go build -o /echo-server cmd/consumption/main.go

RUN ["chmod", "+x", "/echo-server"]

EXPOSE 8080

ENTRYPOINT ["/echo-server"]
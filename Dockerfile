FROM golang:1.12

ENV GOPATH=/go
ENV GO111MODULE=on

ENV APP=/fance
ENV PROJECT_ROOT=$APP

WORKDIR $APP
COPY . $APP

RUN go get ./...
RUN go get github.com/githubnemo/CompileDaemon

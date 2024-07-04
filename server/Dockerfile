FROM golang:alpine

COPY . /server

WORKDIR /server
RUN go build

ENTRYPOINT ./po-key-mon
FROM golang:1.13.3

# go mod proj path
RUN mkdir /go_mod
WORKDIR /go_mod

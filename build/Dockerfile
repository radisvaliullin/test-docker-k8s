FROM golang:1.12.7

COPY ./bin /go/src/github.com/radisvaliullin/test-docker-k8s/bin

WORKDIR /go/src/github.com/radisvaliullin/test-docker-k8s/bin

EXPOSE 7374

ENTRYPOINT /go/src/github.com/radisvaliullin/test-docker-k8s/bin/srv

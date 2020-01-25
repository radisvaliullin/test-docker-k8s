FROM golang:1.13.3

# Install componets
RUN apt-get update && apt-get -y install unzip

# Install protoc
ENV PROTOC_ZIP=protoc-3.11.2-linux-x86_64.zip
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.2/${PROTOC_ZIP}
RUN unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
RUN unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
RUN rm -f $PROTOC_ZIP

# Install golang protobuf dependencies (for protoc gen)
RUN GO111MODULE=on go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.12.2
RUN GO111MODULE=on go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v1.12.2
RUN GO111MODULE=on go get github.com/golang/protobuf/protoc-gen-go@v1.3.2

# go mod projects path
RUN mkdir /go_mod
WORKDIR /go_mod

PROJ_REPO := github.com/radisvaliullin/test-docker-k8s

# protoc imports
PROTOC_IMPORTS := $()
# PROTOC_IMPORTS := $(PROTOC_IMPORTS)-I $(PROTOC_PATH)/include
PROTOC_IMPORTS := $(PROTOC_IMPORTS) -I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

# gen go code from proto
gen:
	@protoc -I $(GOPATH)/src/$(PROJ_REPO)/proto \
	$(PROTOC_IMPORTS) \
	--go_out=plugins=grpc:$(GOPATH)/src \
	--grpc-gateway_out=logtostderr=true:$(GOPATH)/src \
	$(GOPATH)/src/$(PROJ_REPO)/proto/pb/v1/tsrv/test.proto

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o bin/srv cmd/srv/*

image:
	@docker build -t test-docker-k8s -f ./build/Dockerfile .

start:
	@docker stop test-docker-k8s || true
	@docker rm test-docker-k8s || true
	@docker run -d -p 7374:7374 --name test-docker-k8s --rm test-docker-k8s

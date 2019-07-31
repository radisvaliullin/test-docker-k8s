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
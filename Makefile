PROJ_REPO := github.com/radisvaliullin/test-docker-k8s
PROTOC_PATH := ~/bin

# gen go code from proto
gen:
	@protoc -I $(GOPATH)/src/$(PROJ_REPO)/proto \
	-I $(PROTOC_PATH)/include \
	-I $(GOPATH)/src \
	-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:$(GOPATH)/src \
	--grpc-gateway_out=logtostderr=true:$(GOPATH)/src \
	$(GOPATH)/src/$(PROJ_REPO)/proto/pb/v1/tsrv/test.proto
PROJ_REPO := github.com/radisvaliullin/test-docker-k8s

# gen go code from proto
gen:
	@protoc -I $(GOPATH)/src/$(PROJ_REPO)/proto/ \
	--go_out=plugins=grpc:$(GOPATH)/src \
	$(GOPATH)/src/$(PROJ_REPO)/proto/pb/v1/tsrv/test.proto
# VAR
# Iages Name
DEV_PROTOC_GET_IMAGE=$(PROJ_NAME)-dev-protoc-gen
DEV_ENV_IMAGE=$(PROJ_NAME)-dev-env
# Containers Name
DEV_PROTOC_GEN_CONT=$(PROJ_NAME)-dev-protoc-gen
DEV_ENV_CONT=$(PROJ_NAME)-dev-env
DEV_TEST_CONT=$(PROJ_NAME)-dev-test

# dev protobuf gen image
dev-protoc-gen-image:
	@docker build -t $(DEV_PROTOC_GET_IMAGE) -f ./build/dev-env/gen.Dockerfile .

# dev env image
dev-env-image:
	@docker build -t $(DEV_ENV_IMAGE) -f ./build/dev-env/dev.Dockerfile .

# generate go code from proto
protoc-gen:
	@docker run -ti --rm --name $(DEV_PROTOC_GEN_CONT) \
	-v $(CURDIR):/go_mod/$(PROJ_REPO) \
	-w /go_mod/$(PROJ_REPO) \
	$(DEV_PROTOC_GET_IMAGE) \
	protoc \
	-I /go_mod/$(PROJ_REPO)/proto \
	-I /go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.12.2/third_party/googleapis \
	--go_out=plugins=grpc:/go_mod \
	--grpc-gateway_out=logtostderr=true:/go_mod \
	/go_mod/$(PROJ_REPO)/proto/pb/v1/tsrv/test.proto

# dev test
dev-test:
	@docker run -ti --rm --name $(DEV_TEST_CONT) \
	-v $(CURDIR):/go_mod/$(PROJ_REPO) \
	$(DEV_ENV_IMAGE) \
	/bin/bash

#
dev-env:
	@docker run -ti --name $(DEV_ENV_CONT) \
	-v $(CURDIR):/go_mod/$(PROJ_NAME) \
	-w /go_mod/$(PROJ_NAME) \
	$(DEV_ENV_IMAGE) \
	/bin/bash \
	|| \
	docker start -i $(DEV_ENV_CONT)

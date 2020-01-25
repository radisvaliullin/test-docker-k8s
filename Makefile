# imports
include scripts/makefiles/dev.make

# project 
PROJ_NAME := test-docker-k8s
PROJ_REPO := github.com/radisvaliullin/$(PROJ_NAME)

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o bin/srv cmd/srv/*

image:
	@docker build -t test-docker-k8s -f ./build/Dockerfile .

start:
	@docker stop test-docker-k8s || true
	@docker rm test-docker-k8s || true
	@docker run -d -p 7374:7374 --name test-docker-k8s --rm test-docker-k8s

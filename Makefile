.DEFAULT_GOAL	:= example-rds-server

#------------------------------------------------------------------------------
# Variables
#------------------------------------------------------------------------------

SHELL 	:= /bin/bash
PKG 	:= github.com/envoyproxy/meta-protocol-control-plane-api
OUT?=./out
DOCKER_TMP?=$(OUT)/docker_temp/
DOCKER_TAG?=ghcr.io/aeraki-mesh/example-rds-server:latest
BINARY_NAME?=$(OUT)/example-rds-server

.PHONY: build
build:
	GOOS=linux go build ./meta_protocol_proxy/...

#--------------------------------------
#-- example xDS control plane server
#--------------------------------------
.PHONY: $(OUT)/example-rds-server example-rds-server

$(OUT)/example-rds-server:
	GOOS=linux go build  -o $@ example/main/main.go

example-rds-server: $(OUT)/example-rds-server

docker-build: example-rds-server
	rm -rf $(DOCKER_TMP)
	mkdir $(DOCKER_TMP)
	cp example/docker/Dockerfile $(DOCKER_TMP)
	cp $(BINARY_NAME) $(DOCKER_TMP)
	docker build -t $(DOCKER_TAG) $(DOCKER_TMP)
	rm -rf $(DOCKER_TMP)
docker-push: docker-build
	docker push $(DOCKER_TAG)


.PHONY: proto

DIR=$(shell pwd)
GOPATH=$(shell go env -json | jq -r '.GOPATH')
PROTO_FILES=$(shell find ./api -name "*.proto" | sed 's/\.\/api\///g' )
PBDIR="pkg/pbgen"
PROTOC_IMAGE=ghcr.io/krishnaiyer/protoc-go
PROTOC_IMAGE_VERSION=v0.0.1-val

init:
	@echo "Initialize repository"
	go get golang.org/x/tools/cmd/goimports

deps:
	go mod tidy

proto:
	@echo "Generate proto files"
	@rm -rf ${PBDIR}
	@mkdir -p ${PBDIR}
	$(foreach file,$(PROTO_FILES), docker run --rm -v ${DIR}/api:/proto -v ${GOPATH}/src:/go-out  ${PROTOC_IMAGE}:${PROTOC_IMAGE_VERSION} $(file);)
	@goimports -w ${PBDIR}

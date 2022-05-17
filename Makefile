GO111MODULE:=on
GOBIN:=$(GOPATH)/bin
PROTOTOOL:=$(GOBIN)/prototool
PATH:=$(GOBIN):$(PATH)
PROJECT_DIR=$(shell pwd)
GEN_DIR=$(PROJECT_DIR)/internal/pkg/gen

clean:
	rm -rf vendor

install: clean
	go get google.golang.org/grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install github.com/alta/protopatch/cmd/protoc-gen-go-patch@latest

lint:
	$(PROTOTOOL) lint

vendor:
	go mod download
	go mod vendor

generate_proto: clean
	rm -rf $(GEN_DIR)
	buf generate
	go mod vendor


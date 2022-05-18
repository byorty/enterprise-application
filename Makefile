GO111MODULE:=on
GOBIN:=$(GOPATH)/bin
PATH:=$(GOBIN):$(PATH)
PROJECT_NAME:=enterprise_application
PROJECT_DIR:=$(shell pwd)
GEN_DIR:=$(PROJECT_DIR)/pkg/common/gen
BUILD_DIR:=$(PROJECT_DIR)/build
VENDOR_DIR:=$(PROJECT_DIR)/vendor
DOCKER_COMPOSE:=$(PROJECT_DIR)/deployments/docker-compose.yml

export

define build_app
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(1)/main.go $(PROJECT_DIR)/cmd/$(1)/main.go
endef

clean:
	rm -rf $(VENDOR_DIR)

install: clean
	go get google.golang.org/grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install github.com/alta/protopatch/cmd/protoc-gen-go-patch@latest

vendor:
	go mod vendor

buf_update:
	buf mod update

buf_generate: clean
	rm -rf $(GEN_DIR)
	buf generate
	make vendor

build_apps:
	rm -rf $(BUILD_DIR)
	@$(foreach BUILD_APP, $(shell ls $(PROJECT_DIR)/cmd), CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BUILD_APP) $(PROJECT_DIR)/cmd/$(BUILD_APP)/main.go;)

up: build_apps
	docker-compose -p $(PROJECT_NAME) -f $(DOCKER_COMPOSE) up -d --build --force-recreate

down:
	docker-compose -p $(PROJECT_NAME) -f $(DOCKER_COMPOSE) down --remove-orphans
	docker image prune -f
	docker volume prune -f

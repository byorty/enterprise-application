GO111MODULE:=on
GOBIN:=$(GOPATH)/bin
PROTOTOOL:=$(GOBIN)/prototool
PATH:=$(GOBIN):$(PATH)

clean:
	rm -rf vendor

install: clean
	go get google.golang.org/grpc
	go get github.com/gogo/protobuf/proto
	go get github.com/gogo/protobuf/jsonpb
	go get github.com/gogo/protobuf/gogoproto
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install github.com/gogo/protobuf/protoc-gen-gogo
	go install github.com/gogo/protobuf/protoc-gen-gogoslick
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

lint:
	$(PROTOTOOL) lint

vendor:
	go mod download
	go mod vendor

generate_proto: clean
	buf generate
	go mod vendor


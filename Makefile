GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)

.PHONY: init
# init env
init:
	go get -d -u  github.com/tkeel-io/tkeel-interface/openapi
	go get -d -u  github.com/tkeel-io/kit
	go get -d -u  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.7.0

	go install github.com/ocavue/protoc-gen-typescript@latest
	go install  github.com/tkeel-io/tkeel-interface/tool/cmd/artisan@latest
	go install  google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install  google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install  github.com/tkeel-io/tkeel-interface/protoc-gen-go-http@latest
	go install  github.com/tkeel-io/tkeel-interface/protoc-gen-go-errors@latest
	go install  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.7.0

.PHONY: proto
proto:
	API_PROTO_FILES=( $$(find {api,mindspore_serving} -name '*.proto') ); \
	for f in "$${API_PROTO_FILES[@]}"; do \
  		echo 'Processing '+$$f;\
		protoc \
        		-I ./ \
        		--gogo_out=plugins=grpc,:.\
        		--govalidators_out=gogoimport=true,:.\
		$$f;\
	done


.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go generate ./...


.PHONY: all
# generate all
all:
	make api;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

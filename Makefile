PROTOC_VERSION=3.17.3

.PHONY: dev
dev:
	go run cmd/cuneum-api/main.go

.PHONY: docker
docker:
	apt update
	apt install unzip -y
	$(eval TEMP_DIR := $(shell mktemp -d))
	wget https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-linux-x86_64.zip -O $(TEMP_DIR)/protoc.zip
	unzip -o $(TEMP_DIR)/protoc.zip -d /usr/local

	make install

.PHONY: install
install:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

.PHONY: build
build:
	go build -ldflags "-s -w" -o bin/cuneum-api cmd/cuneum-api/main.go

.PHONT: gen
gen:
	go generate ./...
	protoc \
		-I. \
		-I/usr/local/include \
		-I/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		-I/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.5.0 \
		--experimental_allow_proto3_optional \
		--go_out=. --go_opt=module=github.com/cuneum/cuneum-api \
		--go-grpc_out=. --go-grpc_opt=module=github.com/cuneum/cuneum-api \
		--grpc-gateway_out=. --grpc-gateway_opt=module=github.com/cuneum/cuneum-api \
		--openapiv2_out=./docs --openapiv2_opt=use_go_templates=true --openapiv2_opt=allow_merge=true \
		proto/*.proto

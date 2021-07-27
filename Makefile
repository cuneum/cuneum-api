.PHONY: dev
dev:
	go run cmd/cuneum-api/main.go

.PHONY: install
install:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

.PHONY: build
build:
	go build -ldflags "-s -w" -o bin/cuneum-api cmd/cuneum-api/main.go

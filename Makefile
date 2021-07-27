.PHONY: dev
dev:
	go run cmd/cuneum-api/main.go

.PHONY: install
install:
	exit 0

.PHONY: build
build:
	go build -ldflags "-s -w" -o bin/cuneum-api cmd/cuneum-api/main.go

.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

test: vet
	go test ./... -v
.PHONY:test

build: test
	go build cpfgen.go cnpjgen.go
.PHONY:build

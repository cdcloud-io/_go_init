# cdcloud-io Makefile for GO.
# @: only show the output of the command, not the command itself

# build variables
BIN_DIR := bin
MODULE_NAME := 
URL_PATH := 
.DEFAULT_GOAL := run

# .PHONY as targets do not represent files.
.PHONY: all initapi build test test-with-cover generate-mocks clean run deps mod prod asm lint

all:
	@echo "**********************************************************"
	@echo "**          cdcloud-io GO build tool                    **"
	@echo "**********************************************************"

$(BIN_DIR):
	@mkdir -p $@

lint:
	@golangci-lint run --enable-all

build: | $(BIN_DIR)
	@go build -v -o ${BIN_DIR}/$(MODULE_NAME) cmd/${MODULE_NAME}/main.go

run:
	@go run cmd/${MODULE_NAME}/main.go

test:
	@go test -v $(shell go list ./... | grep -v /test/)

test-with-cover:
	@go test -v -coverprofile=cover.out $(shell go list ./... | grep -v /test/)
	@go tool cover -html=cover.out -o cover.html

generate-mocks:
	@mockery --all --with-expecter --keeptree

clean:
	@go clean
	@rm -rf ${BIN_DIR}/${MODULE_NAME}
	@rm -rf vendor

deps:
	@go get ./...

mod: deps
	@go mod download
	@go mod tidy
	@go mod vendor

prod-build:
	@mkdir -p ${BIN_DIR}/${MODULE_NAME}
	@go build -mod=vendor -ldflags="-s -w" -o ${BIN_DIR}/${MODULE_NAME}/${MODULE_NAME} ./cmd/${MODULE_NAME} || (echo "Build failed with exit code $$?"; exit 1)

asm:
	@go tool compile -S cmd/${MODULE_NAME}/main.go > ${MODULE_NAME}.asm

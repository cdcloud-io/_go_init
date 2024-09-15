# cdcloud-io Makefile for GO.
# @: only show the output of the command, not the command itself

# build variables
BIN_DIR := bin
MODULE_NAME := 
URL_PATH := 

.DEFAULT_GOAL := all

# .PHONY as targets do not represent files.
.PHONY: all initapi build test test-with-cover generate-mocks clean run deps mod prod-build asm lint

# Print header information
all:
	@echo "**********************************************************"
	@echo "**          cdcloud-io GO build tool                    **"
	@echo "**********************************************************"

# Ensure BIN_DIR exists
$(BIN_DIR):
	@mkdir -p $@

# Lint the code
lint:
	@golangci-lint run --enable-all

# Build the application
build: | $(BIN_DIR)
	@go build -v -o ${BIN_DIR}/$(MODULE_NAME) cmd/${MODULE_NAME}/main.go

build-proto:
	@protoc --proto_path=protobuf protobuf/*.proto --go_out=/internal/ --go-grpc_out=.

# Run the application
run:
	@go run cmd/${MODULE_NAME}/main.go

# Run tests
test:
	@go test -v $(shell go list ./... | grep -v /test/)

# Run tests with coverage
test-with-cover:
	@go test -v -coverprofile=cover.out $(shell go list ./... | grep -v /test/)
	@go tool cover -html=cover.out -o cover.html

# Generate mocks
generate-mocks:
	@mockery --all --with-expecter --keeptree

# Clean the build artifacts
clean:
	@go clean
	@rm -rf ${BIN_DIR}/${MODULE_NAME}
	@rm -rf vendor

# Get dependencies
deps:
	@go get ./...

# Manage Go modules
mod: 
	@go mod tidy
	@go mod download
	@go mod vendor

# Production build
prod-build: | $(BIN_DIR)
	@go build -mod=vendor -ldflags="-s -w" -o ${BIN_DIR}/${MODULE_NAME}/${MODULE_NAME} ./cmd/${MODULE_NAME} || (echo "Build failed with exit code $$?"; exit 1)

# Generate assembly code
asm:
	@go tool compile -S cmd/${MODULE_NAME}/main.go > ${MODULE_NAME}.asm

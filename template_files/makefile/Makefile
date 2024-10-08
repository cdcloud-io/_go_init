# cdcloud-io Makefile for GO.

# INFO
# @         | only show the output of the command, not the command itself
# -         | Ignores errors from the command. ie @-<command> means silent no stderr out on command
# ./...     | instructs go tool to apply the command to all files in local and subdirectories
# printf -- | instructs printf to interpret -- as string and not options

# load variables from a .env file.
-include .env

# path variables
GOBASE ?= $(CURDIR)

# build variables
VERSION := $(shell git describe --tags 2>/dev/null || echo "v0.0.1")
BUILD := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

PROJECTNAME := $(shell basename "$(PWD)")

BIN_DIR := bin
MODULE_NAME := <<__MODULE_NAME__>>
URL_PATH := <<__URL_PATH__>>

.DEFAULT_GOAL := help

# .PHONY as targets do not represent files.
.PHONY: help fmt vet lint run build build-prod build-proto test test-with-cover generate-mocks clean deps tidy asm

# Print header information and available targets
help:
	@printf "\n"
	@printf "##########################################################\n"
	@printf "#          📘 cdcloud-io GO build tool 📘                #\n"
	@printf "##########################################################\n"
	@printf "\n"
	@printf "📄 Makefile variables 📄\n"
	@printf -- "------------------------\n"
	@printf "VERSION: %s\n" "$(VERSION)"
	@printf "BUILD: %s\n" "$(BUILD)"
	@printf "BIN_DIR: %s\n" "$(BIN_DIR)"
	@printf "MODULE_NAME: %s\n" "$(MODULE_NAME)"
	@printf "URL_PATH: %s\n" "$(URL_PATH)"
	@printf "GOBASE: %s\n" "$(GOBASE)"
	@printf "\n"
	@printf "🎯 Available targets 🎯\n"
	@printf -- "------------------------\n"
	@printf "help              - Show this help information\n"
	@printf "fmt               - Format the code using gofmt\n"
	@printf "vet               - Run go vet to check for issues\n"
	@printf "lint              - Run golangci-lint to analyze the code\n"
	@printf "build             - Build the project\n"
	@printf "build-prod        - Build the project for production with stripped symbols\n"
	@printf "build-proto       - Build protobuf definitions\n"
	@printf "run               - Run the application\n"
	@printf "test              - Run the tests\n"
	@printf "test-with-cover   - Run tests with coverage\n"
	@printf "generate-mocks    - Generate mocks using mockery\n"
	@printf "clean             - Clean the project and build artifacts\n"
	@printf "deps              - Get dependencies\n"
	@printf "tidy              - Tidy and download Go modules\n"
	@printf "asm               - Generate assembly code\n"
	@printf "\n"

# Ensure BIN_DIR exists
$(BIN_DIR):
	@mkdir -p $@

# Formats the code to the official GO style guideline
fmt:
	@go fmt ./...

# Analyzes GO code for potential errors or suspicious constructs, but does not modify the code.
vet:
	@go vet ./...

# Lint analyzes source code to identify potential issues.
lint:
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "golangci-lint is not installed. Please install it first."; exit 1; \
	fi
	@golangci-lint run --enable-all

# Run the application
run:
	@go run cmd/$(MODULE_NAME)/main.go

# Build the application
build: | $(BIN_DIR)
	@echo "  >  Building binary..."
	@go build -v -o $(BIN_DIR)/$(MODULE_NAME) cmd/$(MODULE_NAME)/main.go

# Build the application for production delivery (strip symbols)
build-prod: | $(BIN_DIR)
	@echo "  >  Building Production binary..."
	@go build -mod=vendor -ldflags="-s -w" -o $(BIN_DIR)/$(MODULE_NAME)/$(MODULE_NAME) ./cmd/$(MODULE_NAME) || (echo "Build failed with exit code $$?"; exit 1)

# Build proto-buf code
build-proto:
	@-protoc --proto_path=protobuf protobuf/*.proto --go_out=/internal/ --go-grpc_out=.

# Run tests
test:
	@go test -v $(shell go list ./... | grep -v /test/)

# Run tests with coverage
test-with-cover:
	@go test -v -coverprofile=cover.out $(shell go list ./... | grep -v /test/)
	@go tool cover -html=cover.out -o cover.html

# Generate mocks
generate-mocks:
	@-mockery --all --with-expecter --keeptree

# Clean the build artifacts
clean:
	@go clean
	@-rm -rf $(BIN_DIR)/$(MODULE_NAME)
	@-rm -rf vendor
	@-rm -f cover.out cover.html

# Get dependencies
deps:
	@go get ./...

# Fix/Tidy GO modules
tidy: deps
	@go mod tidy
	@go mod download
	@go mod vendor

# Generate assembly code
asm:
	@go tool compile -S cmd/$(MODULE_NAME)/main.go > $(MODULE_NAME).asm

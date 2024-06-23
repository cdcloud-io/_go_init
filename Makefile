# cdcloud-io Makefile for GO.
# @: only show the output of the command, not the command itself

# build variables
BIN_DIR := bin
ARTIFACT_NAME := go_proj1
URL_PATH := 
.DEFAULT_GOAL := run

# .PHONY as targets do not represent files.
.PHONY: all init build test test-with-cover generate-mocks clean run deps mod prod asm lint

all: test build

initapi:
	@if [ ! -f go.mod ]; then \
		if [ -d "cmd" ]; then \
			echo "Directory ./cmd already exists. Initialization aborted to prevent overwriting existing code."; \
		else \
			if [ -z "${URL_PATH}" ]; then \
				echo "Initializing Go module..."; \
				go mod init ${ARTIFACT_NAME}; \
			else \
				echo "Initializing Go module with URL path..."; \
				go mod init ${URL_PATH}; \
			fi; \
			mkdir -p api;
			mkdir -p cmd/${ARTIFACT_NAME}; \
			mkdir -p bin; \
			mkdir -p config; \
			touch config/config.go; \
			mkdir -p docs; \
			mkdir -p examples; \
			mkdir -p internal/app; \
			touch internal/app/app.go; \
			mkdir -p internal/router; \
			touch internal/router/router.go; \
			mkdir -p internal/endpoint1; \
			touch internal/endpoint1/handler.go; \
			touch internal/endpoint1/endpoint1.go; \
			mkdir -p internal/endpoint2; \
			touch internal/endpoint2/handler.go; \
			touch internal/endpoint2/endpoint2.go; \
			mkdir -p internal/middleware; \
			touch internal/middleware/middleware.go; \
			touch internal/middleware/logging.go; \
			touch internal/middleware/auth.go; \
			touch internal/middleware/logic.go; \
			mkdir -p pkg; \
			mkdir -p test; \
			printf "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"%s\")\n}" "${ARTIFACT_NAME}" > cmd/${ARTIFACT_NAME}/main.go; \
			touch README.md; \
			printf "# %s" "${ARTIFACT_NAME}" > README.md; \
		fi; \
	else \
		echo "Go module already initialized."; \
	fi

$(BIN_DIR):
    mkdir $@

build: | $(BIN_DIR)
	@go build -v ${BIN_DIR}/$(ARTIFACT}/${ARTIFACT} cmd/${ARTIFACT}/main.go

test:
	@go test -v $(shell go list ./... | grep -v /test/)

test-with-cover:
	@go test -v -coverprofile=cover.out $(shell go list ./... | grep -v /test/)
	@go tool cover -html=cover.out -o cover.html

generate-mocks:
	@mockery --all --with-expecter --keeptree

clean:
	@go clean
	@rm -rf ${BIN_DIR}/${ARTIFACT_NAME}
	@rm -rf vendor

run:
	@go run cmd/${ARTIFACT_NAME}/main.go

deps:
	@go get ./...

mod: deps
	@go mod download
	@go mod tidy
	@go mod vendor

prod:
	@mkdir -p ${BIN_DIR}/${ARTIFACT_NAME}
	@go build -mod=vendor -ldflags="-s -w" -o ${BIN_DIR}/${ARTIFACT_NAME}/${ARTIFACT_NAME} ./cmd/${ARTIFACT_NAME} || (echo "Build failed with exit code $$?"; exit 1)

asm:
	@go tool compile -S cmd/${ARTIFACT_NAME}/main.go > ${ARTIFACT_NAME}.asm

lint:
	@golangci-lint run --enable-all

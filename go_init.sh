#!/usr/bin/env bash

# ==============================================================================
# Title          : Go Init Script
# Description    : This script sets up a Go project environment and should be sourced.
# Company        : cdcloud-io
# Author         : cd-stephen
# References     : [URLs or other references]
# Last Modified  : 7/2/2024
# Version        : 1.0
# Usage          : source go_init.sh
# Notes          : Ensure this script is sourced to run the function in the current shell.
# ==============================================================================

clear
echo '🟨 sourcing go_init.sh'
sleep 1

# Set GO BIN
export PATH=$PATH:$(go env GOPATH):$(go env GOPATH)/bin

# Function to set up a Go project
function go_init() {
    # Check if the script is being run in a subdirectory of the user's home directory
    if [ "$PWD" == "$HOME" ] || [[ "$PWD" != "$HOME/"* ]]; then
        echo ''
        echo "🟥 Error: Script must be run in a subdirectory of ${HOME}. Exiting..."
        return 1
    fi

    # Check if the directory contains only .git/ and README.md
    for file in * .*; do
        if [ "$file" != "." ] && [ "$file" != ".." ] && [ "$file" != "*" ] && [ "$file" != ".*" ] && [ "$file" != ".git" ] && [ "$file" != "README.md" ]; then
            echo ''
            echo "🟥 Error: Not an empty project. Only .git/ and README.md are allowed. Exiting..."
            return 1
        fi
    done

    # Scoped variable
    URL_PATH_VALUE=false

    # Download the Makefile and .gitignore
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/api_templates_files/Makefile -O Makefile
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/api_templates_files/.gitignore -O .gitignore
    
    # Extract the MODULE_NAME from the current directory name
    MODULE_NAME=$(basename "$(pwd)")
    
    # Extract the URL_PATH from the Git configuration if a .git directory exists
    if [ -d ".git" ]; then
        GIT_URL=$(git config --get remote.origin.url)
        echo "GIT_URL: $GIT_URL"
        URL_PATH=$(echo "$GIT_URL" | sed -E "s|git@([^:]+):([^/]+/[^/]+)\.git$|\\1/\\2|")
        echo "GO MOD: $URL_PATH"
    else
        URL_PATH=""
    fi
    
    # Use sed to replace the placeholders in the Makefile
    sed -i "s|^MODULE_NAME :=.*|MODULE_NAME := $MODULE_NAME|" Makefile
    sed -i "s|^URL_PATH :=.*|URL_PATH := $URL_PATH|" Makefile

    echo ''
    echo "🟩 Makefile has been set up with MODULE_NAME: $MODULE_NAME and URL_PATH: $URL_PATH"
    echo ''

    if [ -z "${URL_PATH}" ]; then
        echo "Initializing Go module..."
        go mod init "${MODULE_NAME}"
    else
        echo "Initializing Go module with URL path..."
        go mod init "${URL_PATH}"
        URL_PATH_VALUE=true
    fi

    # Create necessary directories
    mkdir -p api > /dev/null 2>&1                            ## openapi spec
    mkdir -p bin > /dev/null 2>&1                            ## compilation bin destination
    mkdir -p build/{docker,k8s/kustomize} > /dev/null 2>&1   ## scripts for build, run, deploy
    mkdir -p cmd/${MODULE_NAME} > /dev/null 2>&1             ## application entry point. main.go
    mkdir -p config > /dev/null 2>&1                         ## config.yaml used by internal/config/config.go
    mkdir -p docs/img > /dev/null 2>&1                       ## module/app documentation and images
    mkdir -p example > /dev/null 2>&1                        ## optional use for app/code usage examples
    mkdir -p internal/{app,config,${MODULE_NAME}_proto,endpoint/user,middleware/auth,middleware/logging,middleware/trace,server} > /dev/null 2>&1 ## module/app internal packages
    mkdir -p protobuf
    mkdir -p test > /dev/null 2>&1                           ## unit/integration tests

    # Download template files
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/api_templates_files/main.go.tmpl -O ./cmd/${MODULE_NAME}/main.go > /dev/null 2>&1
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/api_templates_files/app.go.tmpl -O ./internal/app/app.go > /dev/null 2>&1
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/api_templates_files/config.go.tmpl -O ./internal/config/config.go > /dev/null 2>&1
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/api_templates_files/config.yaml -O ./config/config.yaml > /dev/null 2>&1
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/api_templates_files/handler.go.tmpl -O ./internal/endpoint/user/handler.go > /dev/null 2>&1
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/api_templates_files/user.go.tmpl -O ./internal/endpoint/user/user.go > /dev/null 2>&1
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/api_templates_files/server.go.tmpl -O ./internal/server/server.go > /dev/null 2>&1
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/api_templates_files/build-dockerfile.sh -O ./build/build-dockerfile.sh > /dev/null 2>&1
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/api_templates_files/build-run-with-env.sh -O ./build/build-run-with-env.sh > /dev/null 2>&1
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/api_templates_files/Dockerfile -O ./build/docker/Dockerfile > /dev/null 2>&1

    # Replace template placeholders in the downloaded files
    if [ "$URL_PATH_VALUE" = true ]; then
        sed -i "s|__MODULE_NAME__|$URL_PATH|g" ./cmd/${MODULE_NAME}/main.go
        sed -i "s|__MODULE_NAME__|$URL_PATH|g" ./internal/app/app.go
        sed -i "s|__MODULE_NAME__|$URL_PATH|g" ./internal/config/config.go
        sed -i "s|__MODULE_NAME__|$MODULE_NAME|g" ./config/config.yaml
        sed -i "s|__MODULE_NAME__|$URL_PATH|g" ./internal/endpoint/user/handler.go
        sed -i "s|__MODULE_NAME__|$URL_PATH|g" ./internal/endpoint/user/user.go
        sed -i "s|__MODULE_NAME__|$URL_PATH|g" ./internal/server/server.go
        sed -i "s|__MODULE_NAME__|$MODULE_NAME|g" ./build/build-dockerfile.sh
        sed -i "s|__MODULE_NAME__|$MODULE_NAME|g" ./build/build-run-with-env.sh
        sed -i "s|__MODULE_NAME__|$MODULE_NAME|g" ./build/docker/Dockerfile
    else
        sed -i "s|__MODULE_NAME__|$MODULE_NAME|g" ./cmd/${MODULE_NAME}/main.go
        sed -i "s|__MODULE_NAME__|$MODULE_NAME|g" ./internal/app/app.go
        sed -i "s|__MODULE_NAME__|$MODULE_NAME|g" ./internal/config/config.go
        sed -i "s|__MODULE_NAME__|$MODULE_NAME|g" ./config/config.yaml
        sed -i "s|__MODULE_NAME__|$MODULE_NAME|g" ./internal/endpoint/user/handler.go
        sed -i "s|__MODULE_NAME__|$MODULE_NAME|g" ./internal/endpoint/user/user.go
        sed -i "s|__MODULE_NAME__|$MODULE_NAME|g" ./internal/server/server.go
        sed -i "s|__MODULE_NAME__|$MODULE_NAME|g" ./build/build-dockerfile.sh
        sed -i "s|__MODULE_NAME__|$MODULE_NAME|g" ./build/build-run-with-env.sh
        sed -i "s|__MODULE_NAME__|$MODULE_NAME|g" ./build/docker/Dockerfile
    fi

    echo '⤵️ Updating go.mod'
    go get -u ./... > /dev/null 2>&1 
    go mod tidy > /dev/null 2>&1
    go mod download > /dev/null 2>&1
    go mod vendor > /dev/null 2>&1

    # Create README.md if it does not exist
    if [ ! -f README.md ]; then
        printf "# %s" "${MODULE_NAME}" > README.md
        echo ''
        echo '🟩 INFO: Go module has been initialized'
        echo ''
    else
        rm -f README.md
        printf "# %s" "${MODULE_NAME}" > README.md
        echo ''
        echo '🟨 WARN: README.md has been modified'
        echo '🟩 INFO: Go module has been initialized'
        echo ''
    fi
}

clear
echo '🟩 sourcing go_init.sh'

# Make sure to source this file in .bashrc
# source /path/to/this/file/go_init.sh or source $HOME/go_init.sh

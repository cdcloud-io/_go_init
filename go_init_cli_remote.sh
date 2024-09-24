#!/usr/bin/env bash

# ==============================================================================
# Title          : Go Init Script
# Description    : This script sets up a GO project environment adhering to hexagonal architecture.
# Company        : cdcloud-io
# Author         : cd-stephen
# Last Modified  : 09-23-2024
# Version        : 1.5
# Usage          : source go_init.sh
# Notes          : Ensure this script is sourced to run the function in the current shell.
# ==============================================================================

# Function to set up a Go project
function go_init() {
    clear
    echo '🟨 Setup Started: GO CLI project'
    sleep 1

    # Check if the script is being run in a subdirectory of the user's home directory
    if [[ "$PWD" == "$HOME" || "${PWD##$HOME/}" == "$PWD" ]]; then
        echo "🟥 Error: Script must be run in a subdirectory of ${HOME}. Exiting..."
        echo ''
        exit 1
    fi

    # Check if the directory contains only .git/ and README.md
    shopt -s dotglob nullglob
    files=(*)
    for file in "${files[@]}"; do
        if [[ "$file" != ".git" && "$file" != "README.md" ]]; then
            echo "🟥 Error: Not an empty project. Only .git/ and README.md are allowed. Exiting..."
            echo ''
            exit 1
        fi
    done
    shopt -u dotglob nullglob

    # Scoped variable
    URL_PATH_VALUE=""

    # Download the Makefile and .gitignore
    wget -q https://raw.githubusercontent.com/cdcloud-io/go-init/refs/heads/develop/template_files/makefile/Makefile -O Makefile
    wget -q https://raw.githubusercontent.com/cdcloud-io/go-init/refs/heads/develop/template_files/git/.gitignore -O .gitignore

    # Extract the MODULE_NAME from the current directory name
    MODULE_NAME=$(basename "$(pwd)")

    # Extract the URL_PATH from the Git configuration if a .git directory exists
    if [ -d ".git" ]; then
        GIT_URL=$(git config --get remote.origin.url)
        echo "GIT_URL: $GIT_URL"
        URL_PATH=$(echo "$GIT_URL" | sed -E "s|git@([^:]+):([^/]+/[^/]+)\.git$|\\1/\\2|")
        URL_PATH_MSG=$URL_PATH
        echo "GO MOD: $URL_PATH"
    else
        URL_PATH=""
        URL_PATH_MSG="<local project>"
    fi

    # Use sed to replace the placeholders in the Makefile
    sed -i "s|^MODULE_NAME :=.*|MODULE_NAME := $MODULE_NAME|" Makefile
    sed -i "s|^URL_PATH :=.*|URL_PATH := $URL_PATH|" Makefile

    echo '│'
    echo "└── Makefile has been set up with MODULE_NAME: $MODULE_NAME and URL_PATH: $URL_PATH_MSG"

    if [ -z "${URL_PATH}" ]; then
        echo "└── Initializing Go module..."
        go mod init "${MODULE_NAME}"
    else
        echo '└── Initializing Go module with URL path...'
        go mod init "${URL_PATH}"
        URL_PATH_VALUE=true
    fi

    # Create necessary directories
    mkdir -p api >/dev/null 2>&1                                         ## openapi spec
    mkdir -p bin >/dev/null 2>&1                                         ## compilation bin destination
    mkdir -p build/{docker,k8s/kustomize} >/dev/null 2>&1                ## scripts for build, run, deploy
    mkdir -p cmd/${MODULE_NAME} >/dev/null 2>&1                          ## application entry point. main.go
    mkdir -p config >/dev/null 2>&1                                      ## config.yaml used by internal/config/config.go
    mkdir -p docs/img >/dev/null 2>&1                                    ## module/app documentation and images
    mkdir -p example >/dev/null 2>&1                                     ## optional use for app/code usage examples
    mkdir -p internal/{app,config,domain,repository} >/dev/null 2>&1     ## create internal package structure
}

go_init
echo '│'
echo '🟩 Setup Complete: GO CLI Project'
echo ''

#!/usr/bin/env bash

# ==============================================================================
# Title          : Go Init Script
# Description    : This script sets up a GO project environment adhering to hexagonal architecture.
# Company        : cdcloud-io
# Author         : cd-stephen
# Last Modified  : 09-23-2024
# Version        : 1.5
# Usage          : source go_init.sh
# Notes          : written and tested against bash in linux
# ==============================================================================

clear
echo '🟨 sourcing go_init.sh'
sleep 1

# Set GO BIN
export PATH=$PATH:$(go env GOPATH):$(go env GOPATH)/bin

# Function to set up a Go project
function go_init() {
    # Check if curl is installed
    if ! command -v curl &>/dev/null; then
        echo "Error: curl is not installed. Please install curl. # sudo apt install curl"
        exit 1
    fi

    curl -s https://raw.githubusercontent.com/cdcloud-io/go-init/refs/heads/develop/go_init_cli_remote.sh | bash

    if [ $? -eq 0 ]; then
        echo "GO template created successfully."
        exit 0
    else
        echo "Error: GO template creation FAILED!"
        exit 1
    fi
}

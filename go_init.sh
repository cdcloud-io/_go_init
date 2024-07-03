# Source File
# Add to .bashrc

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

echo 'sourcing go_init.sh'

# Function to set up a Go project
function go_init() {
    if [ ! -f go.mod ]; then
    # Check if the "cmd" directory exists
    if [ -d "cmd" ]; then
        echo "Directory ./cmd already exists. Initialization aborted to prevent overwriting existing code."
    else
        # Initialize Go module
        if [ -z "${URL_PATH}" ]; then
            echo "Initializing GO module..."
            go mod init "${MODULE_NAME}"
        else
            echo "Initializing GO module with URL path..."
            go mod init "${URL_PATH}"
        fi

    # Extract the MODULE_NAME from the current directory name
    MODULE_NAME=$(basename "$(pwd)")

    # Extract the URL_PATH from the .git config if a .git directory exists
    if [ -d ".git" ]; then
        GIT_URL=$(git config --get remote.origin.url)
        URL_PATH=$(echo "$GIT_URL" | sed -E "s|.*[:/]([^/]+/[^/]+)\.git$|\\1|")
    else
        URL_PATH=""
    fi

    # Use sed to replace the placeholders in the Makefile
    sed -i "s|^MODULE_NAME :=.*|MODULE_NAME := $MODULE_NAME|" Makefile
    sed -i "s|^URL_PATH :=.*|URL_PATH := $URL_PATH|" Makefile

    echo "Makefile has been set up with MODULE_NAME: $MODULE_NAME and URL_PATH: $URL_PATH"
    echo ''
    echo 'Initializing Module File and Directory Structure'
}


#!/usr/bin/env bash

# add this file to a path directory, or use the bash alias

# Function to set up a Go project
function setup_goproj() {
    # Download the Makefile and .gitignore
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/Makefile -O Makefile
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/.gitignore -O .gitignore

    # Extract the ARTIFACT_NAME from the current directory name
    ARTIFACT_NAME=$(basename "$(pwd)")

    # Extract the URL_PATH from the Git configuration if a .git directory exists
    if [ -d ".git" ]; then
        GIT_URL=$(git config --get remote.origin.url)
        URL_PATH=$(echo "$GIT_URL" | sed -E "s|.*[:/]([^/]+/[^/]+)\.git$|\\1|")
    else
        URL_PATH=""
    fi

    # Use sed to replace the placeholders in the Makefile
    sed -i "s|^ARTIFACT_NAME :=.*|ARTIFACT_NAME := $ARTIFACT_NAME|" Makefile
    sed -i "s|^URL_PATH :=.*|URL_PATH := $URL_PATH|" Makefile

    echo "Makefile has been set up with ARTIFACT_NAME: $ARTIFACT_NAME and URL_PATH: $URL_PATH"
}

setup_goproj

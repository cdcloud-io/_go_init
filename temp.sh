#!/bin/bash

# Loop through all items in the current directory
for item in * .[^.]*; do
    # Check if the item is not README.md or .git
    if [ "$item" != "README.md" ] && [ "$item" != ".git" ]; then
        echo "other files found"
    fi
done

# If no other files were found
echo "ok"

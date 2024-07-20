#!/bin/bash

# Check if the directory exists
DIR="."

# Iterate over all files in the directory
for file in "$DIR"/*.txt; do
  # Check if it's a file (and not a directory)
  if [ -f "$file" ]; then
    # Sort the file and save the output to a temporary file
    sort "$file" -o "$file"
    echo "Sorted lines in file: $file"
  fi
done

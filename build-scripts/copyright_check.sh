#!/bin/bash

current_year=$(date +"%Y")
missing_copyright=false

for file in $(find . -name "*.go"); do
    first_line=$(head -n 1 "$file")
    if [[ ! "$first_line" =~ ^\/\/.*avijit\ bhattacharjee.*$current_year ]]; then
        echo "File $file is missing copyright comment"
        missing_copyright=true
    fi
done

if [ "$missing_copyright" = true ]; then
    exit 1
fi
#!/usr/bin/env bash

# example
# static_web.sh input.html output.go

FILE="$1"
FILE_NOPREFIX="${1/*web/}"
FILE_SCRUBBED="${FILE//[^0-9a-zA-Z]/_}"
{
    # Scrub file
    sed 's/`/\\&/g' "$1"
} | {
    # Convert string to golang
    echo "package web_build"
    echo 'import (
        "net/http"
        "fmt"
    )'
    echo -n 'const' "$FILE_SCRUBBED" '= `'
    cat
    echo '`'
    echo "func init() {
        http.HandleFunc(\"$FILE_NOPREFIX\", func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintf(w, $FILE_SCRUBBED)
        })
    }"
} | tee "$2"
# Write to file and stdout

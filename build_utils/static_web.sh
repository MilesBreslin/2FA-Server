#!/bin/bash

# example
# static_web.sh input.html output.go

FILE="$1"
FILE_NOPREFIX="${1/*web/}"
FILE_SCRUBBED="${FILE//[^0-9a-zA-Z]/_}"
{
    sed 's/`/\\&/g;s/%/\\&/g' "$1"
} | {
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
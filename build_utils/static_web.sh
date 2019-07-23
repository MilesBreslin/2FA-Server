#!/bin/bash

# example
# static_web.sh input.html output.go

{
    cat "$1"
} | {
    echo "package web_built"
    echo 'import (
        "net/http"
        "fmt"
    )'
    echo -n 'const static = `'
    cat
    echo '`'
    echo "func init() {
        http.HandleFunc(\"${1/*web/}\", func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintf(w, static)
        })
    }"
} | tee "$2"
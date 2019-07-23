#!/bin/bash

# example
# static_web.sh input.html output.go

{
    cat "$1"
} | {
    echo "package web_built"
    echo -n 'const file = `'
    cat
    echo '`'
    echo 'import http'
    echo 'import fmt'
    echo "func init() {
        http.HandleFunc(\"${1/*web/}\", func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintf(w, static)
        })
    }"
} | tee "$2"
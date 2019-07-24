#!/bin/bash

mkdir web_build 2>/dev/null

rm -rf web_build/* 2>/dev/null

for file in web/* ; do
    build_utils/static_web.sh "$file" "web_build/${file/web\//}.go"
done

mkdir build 2>/dev/null
go build -o build/server cmd/server/main.go
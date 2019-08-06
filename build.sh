#!/usr/bin/env bash

testDir() {
    init_dir="$PWD"
    cd "$1"
    go test .
    [ "$?" == "0" ] || exit "1"
    cd "$init_dir"
}

PROJECT_NAME="github.com/opensource2fa/server"
PROJECT_PATH="$GOPATH/src/$PROJECT_NAME"
if [ "$PWD" != "$PROJECT_PATH" ]; then
    if ! [ -e "$PROJECT_PATH" ]; then
        mkdir -p "$GOPATH/src/github.com/opensource2fa"
        rm -rf "$PROJECT_PATH" 2> /dev/null
        ln -s "PWD" "$PROJECT_PATH"
    fi
fi

if [ "$1" == "test" ]; then
    testDir pkg/keys/keychain
    testDir pkg/totp
    testDir pkg/realtime_api
    exit 0
elif [ "$1" == "deps" ]; then
    go get github.com/gorilla/websocket
elif [ "$1" == "fromscratch" ]; then
    ./build.sh deps || exit 1
    ./build.sh test || exit 1
    ./build.sh || exit 1
elif ! [ -z "$1" ]; then
    echo "./build.sh test"
    echo $'\t' "Runs all tests on the code"
    echo "./build.sh deps"
    echo $'\t' "Download all dependencies"
    echo "./build.sh"
    echo $'\t' "Builds all of the code"
    exit 0
fi

mkdir web_build 2>/dev/null

rm -rf web_build/* 2>/dev/null

for file in web/* ; do
    build_utils/static_web.sh "$file" "web_build/${file/web\//}.go"
done

mkdir build 2>/dev/null
go build -o build/server cmd/server/main.go

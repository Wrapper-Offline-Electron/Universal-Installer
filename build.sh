#!/bin/sh
# Copyright (c) 2022 Wrapper-Offline-Electron. MIT license.
# Please use this only for Unix-like operating systems! Use 'build.ps1' for Windows!

windows="amd64"
darwin="amd64"
linux="amd64"

runplatform() {
    out="WOE-Universal-Installer-$1-$2"
    if [ $1 = "windows" ]; then
        out="$out.exe"
    fi
    GOOS=$1 GOARCH=$2 go build -o $out
}

# gc (golang) exists in this system
if type "go" > /dev/null; then
    echo "Starting to compile"
    for platform in $windows; do
        runplatform "windows" $platform
    done
    echo "Windows compiled"
    for platform in $darwin; do
        runplatform "darwin" $platform
    done
    echo "Darwin compiled"
    for platform in $linux; do
        runplatform "linux" $platform
    done
    echo "Linux compiled"
    echo "WOE-Universal-Installer was built successfully into the root directory"
else
    echo "Error: gc (golang) doesn't exist in this system, please install download Go (https://golang.org/dl/)"
    exit 1
fi

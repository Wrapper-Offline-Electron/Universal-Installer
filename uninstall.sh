#!/bin/bash

abort() {
    printf "%s\n" "$@"
    exit 1
}

# Make sure BASH Is installed
if [ -z "${BASH_VERSION:-}" ]
then
    abort "Bash is required to execute this script."
fi

# Check OS
OS="$(uname)"
MACHINE="$(uname -m)"
if [[ "${OS}" != "Darwin" && "${MACHINE}" != "arm64" && "${OS}" != "Linux" && "${MACHINE}" != "amd64" ]]
then
    abort "Wrapper Offline Electron Universal Installer only supports Windows, Linux, macOS 64-bit."
fi

rm -rf ~/Desktop/WOE-Universal-Installer

echo Successfully uninstalled.

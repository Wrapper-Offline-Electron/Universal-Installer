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
if [[ "${OS}" != "Darwin" && "${MACHINE}" != "amd64" && "${MACHINE}" != "arm64" ]]
then
    abort "For now, Wrapper Offline Electron Universal Installer Install File is only supported for MacOS (Darwin)."
fi

# Uninstall
rm -rf ~/Desktop/WOE-Universal-Installer

echo Successfully uninstalled.
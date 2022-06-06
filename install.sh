#!/bin/bash

abort() {
    printf "%s\n" "$@"
    exit 1
}

lower() {
    echo "$1" | tr '[:upper:]' '[:lower:]'
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

# Download
rm -rf ~/Desktop/WOE-Universal-Installer # Make sure universal installer is removed before installing it with curl
curl -L https://github.com/Wrapper-Offline-Electron/Universal-Installer/releases/download/v0.2.1-alpha+windows/WOE-Universal-Installer-$(lower $OS)-amd64 -o ~/Desktop/WOE-Universal-Installer && echo "Successfully installed." || echo "Failed to install with curl."
chmod +x ~/Desktop/WOE-Universal-Installer # make script be able to execute

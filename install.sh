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

# Download
rm -rf ~/Desktop/WOE-Universal-Installer # Make sure calcead is removed before installing it with curl
curl -f -L https://github.com/Wrapper-Offline-Electron/Universal-Installer/releases/download/v0.1.1/WOE-Universal-Installer-darwin-amd64 -o ~/Desktop/WOE-Universal-Installer && echo "Successfully installed." || echo "Failed to install with curl."
chmod +x ~/Desktop/WOE-Universal-Installer # make executable be able to execute

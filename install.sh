#!/bin/bash

function abort {
  local s
  s="$1"
  "echo" "-e" "$s"
  exit 1
  
}
OS="$(uname)"
MACHINE="$(uname -m)"


if [ "$MACHINE" != "amd64" ]; then
  if [ "$MACHINE" != "arm64" ]; then
    "abort" "Wrapper Offline Electron is only supported for 64-bit architecture for Windows, MacOS, and Linux."  
fi
fi
rm -rf ~/Desktop/WOE-Universal-Installer

curl -f -L https://github.com/Wrapper-Offline-Electron/Universal-Installer/releases/download/v0.1.1/WOE-Universal-Installer-$OS-amd64 -o ~/Desktop/WOE-Universal-Installer && echo "Successfully installed." || echo "Failed to install with curl."
chmod +x ~/Desktop/WOE-Universal-Installer


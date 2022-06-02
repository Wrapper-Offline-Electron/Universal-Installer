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

"echo" "-e" "Successfully uninstalled."
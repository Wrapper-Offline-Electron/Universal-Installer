# Universal-Installer
Wrapper Offline Electron Universal Installer

## Binaries

**Two options:**

**Option 1 - Script to automatically install into Desktop:**

These scripts install latest version.

Windows:
To install: go to https://raw.githubusercontent.com/Wrapper-Offline-Electron/Universal-Installer/main/install.bat and download it by `Ctrl+S` or `File > Save Page As`,

then run the bat file (double click it).


To uninstall: go to https://raw.githubusercontent.com/Wrapper-Offline-Electron/Universal-Installer/main/uninstall.bat and download it by `Ctrl+S` or `File > Save Page As`,

then run the bat file (double click it).

MacOS, Linux:
To install, go into the terminal, and run:
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Wrapper-Offline-Electron/Universal-Installer/main/install.sh)"
```
To uninstall, go into the terminal, and run:
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Wrapper-Offline-Electron/Universal-Installer/main/uninstall.sh)"
```

**Option 2 - Check releases and install:**
https://youtu.be/kSB5gtzq8gw

## Building from source

**It is recommended that you clone exactly this git repository and build from root directory.**

**Requirements:**

* Go compiler 1.17

**How to build:**

Assuming you have cloned this git repository...

If you are on Windows, choose either options to build:

1. Run cmd in root directory and run:
```powershell
powershell -executionpolicy bypass -File .\build.ps1
```
2. Launch Windows Powershell, `cd` to root directory, and run:
```powershell
.\build.ps1
```

If you are on a Unix-like OS:

1. Go to terminal and `cd` to root directory. Then, run:
```sh
chmod +x ./build.sh
```
(to give execute permission to the script), and run:
```sh
./build.sh
```

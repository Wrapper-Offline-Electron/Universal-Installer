# Universal-Installer
Wrapper Offline Electron Universal Installer

## Binaries

### Automated script download (into the Desktop)
*Note: This installs the latest version.*

**MacOS, Linux:**
Go to the Terminal.
Install:
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Wrapper-Offline-Electron/Universal-Installer/main/install.sh)"
```

Uninstall:
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Wrapper-Offline-Electron/Universal-Installer/main/uninstall.sh)"
```

**Windows:**
Script not implemented yet for Windows.

### Straight to downloading the binary (don't do this if you used the automated script)

Check releases.

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

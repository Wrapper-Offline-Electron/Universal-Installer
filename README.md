# Universal-Installer
Wrapper Offline Electron Universal Installer

## Reason for Archive

Wrapper Offline Electron Universal Installer was made specifically for bug testing and non-production ready testing. I made this to make it easier to test Wrapper Offline Electron without having to fully package it and then publish it. With this, I did not have to package the application, which makes it easier for OTHERS to test if the application works.

This universal installer was supposed to be easy to use, but it was hard to some of the testers, and there were always bugs and issues popping up, which made development even harder for me. Simple things just kept having issues and bugs, and it just got to the point where I couldn't fix some of them (due to like connection errors for example).

It was awesome to develop this Universal Installer, as it made me think new ways and let me learn so many new things. Sadly, due to too many issues with me and the testers with this, I've decided to stop maintenance of this.

Thank you for helping us, Universal Installer.

Regards,
JackProgramsJP

## Binaries

Go to [releases](https://github.com/Wrapper-Offline-Electron/Universal-Installer/releases)!

Download the latest version!

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

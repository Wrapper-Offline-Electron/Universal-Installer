#!/usr/bin/env pwsh
# Only use for Windows. MacOS, Linux is fine, but you would have to install Powershell yourself, so it would be better to use the shell files.

if ([System.Version]$PSVersionTable.PSVersion -lt [System.Version]"3.0.0.4080") {
    Write-Output "Powershell Version 3.0 (3.0.0.4080) is needed in order for this script to work."
    exit 1
}

if (-not [Environment]::Is64BitProcess) {
    Write-Output "Wrapper Offline Electron Universal Installer only works for 64-bit Windows"
    exit 1
}

# Download

$DesktopPath = [Environment]::GetFolderPath("Desktop")
$InstallerPath = $DesktopPath + "\WOE-Universal-Installer.exe"

if (Test-Path $InstallerPath) {
    Remove-Item $InstallerPath
}

Invoke-WebRequest -Uri "https://github.com/Wrapper-Offline-Electron/Universal-Installer/releases/download/v0.3.1-beta%2B1/WOE-Universal-Installer-windows-amd64.exe" -OutFile $InstallerPath 
Write-Output "Successfully installed."

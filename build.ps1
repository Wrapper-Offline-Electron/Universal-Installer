#!/usr/bin/env pwsh
# Copyright (c) 2022 Wrapper-Offline-Electron. MIT license.
# Please use this only for Windows! Use 'build.sh' for Unix-like operating systems instead!

$Platforms = @{
    windows = "386", "amd64"
    darwin = "amd64"
    linux = "amd64"
}

# gc (golang) exists in this system
if (Get-Command "go.exe" -ErrorAction SilentlyContinue) {
    Write-Output "Starting to compile"
    foreach ($Platform in $Platforms.GetEnumerator()) {
        for ($i = 0; $i -lt $Platform.Value.Count; $i++) {
            $Env:GOOS = $Platform.Name; $Env:GOARCH = $Platform.Value[$i]; go build -o ("WOE-Universal-Installer-$($Platform.Name)-$($Platform.Value[$i])")
        }
        Write-Output "$($Platform.Name) compiled"
    }
    Write-Output "WOE-Universal-Installer was built successfully in the root directory"
} else {
    Write-Error "Error: gc (golang) doesn't exist in this system, please install download Go (https://golang.org/dl/)"
    exit 1
}

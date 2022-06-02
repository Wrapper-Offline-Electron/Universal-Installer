@echo off
setlocal EnableDelayedExpansion
setlocal EnableExtensions

rem !/bin/bash
@echo OFF


set OS=Windows
reg Query "HKLM\Hardware\Description\System\CentralProcessor\0" | find /i "x86" > NUL && set MACHINE=i386 || set MACHINE=amd64
if !MACHINE! NEQ amd64 (
  set _0=Wrapper Offline Electron is only supported for 64-bit architecture for Windows^, MacOS^, and Linux.
  call :abort _1 0 _0
  echo | set /p ^=!_1!
)

del "%USERPROFILE%\Desktop\WOE-Universal-Installer.exe"


powershell -Command "(New-Object Net.WebClient).DownloadFile('https://github.com/Wrapper-Offline-Electron/Universal-Installer/releases/download/v0.1.1/WOE-Universal-Installer-windows-amd64', \"$HOME\Desktop\WOE-Universal-Installer.exe\")"
echo Successfully installed.

goto :EOF
:abort
set s_%~2=!%~3!
echo !s_%~2!

exit /b 1
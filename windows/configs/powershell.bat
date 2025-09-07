@echo off
echo Powershell Setup Start !!

set "BASE_DIR=%~1"
set "PS1_FILE=Microsoft.PowerShell_profile.ps1"
set "PS1_SRC=%BASE_DIR%\powershell\%PS1_FILE%"

:: Get PowerShell profile path
for /f "delims=" %%i in ('powershell -NoProfile -Command "$PROFILE"') do set "PS1_DES=%%i"

:: Ensure source file exists
if not exist "%PS1_SRC%" (
  echo ERROR: Source file not found: "%PS1_SRC%"
  pause
  exit /b 1
)

:: Ensure target directory exists
if not exist "%PS1_DES%" (
  echo Powershell profile target directory is not exist
  echo Target file will be created at: "%PS1_DES%"
  echo Press Ctrl+C to cancel
  pause >nul
  type nul > "%PS1_DES%"
)

:: Remove existing file
echo Removing existing profile.ps1 if they exist
if exist "%PS1_DES%" del "%PS1_DES%"

echo Creating symlinks
mklink "%PS1_DES%" "%PS1_SRC%"

echo Powershell Setup Done !!

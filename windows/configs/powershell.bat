@echo off
echo Powershell Setup Start !!

set "BASE_DIR=%~1"

set "PS1_FILE=Microsoft.PowerShell_profile.ps1"
set "PS1_SRC=%BASE_DIR%\powershell\%PS1_FILE%"

set "OMP_THEME_FILE=oh-my-posh-theme.omp.json"
set "OMP_THEME_SRC=%BASE_DIR%\powershell\%OMP_THEME_FILE%"

:: Get PowerShell profile path
for /f "delims=" %%i in ('powershell -NoProfile -Command "$PROFILE"') do (
  set "PS1_DES_DIR=%%~dpi"
)

set "PS1_DES=%PS1_DES_DIR%%PS1_FILE%"
set "OMP_THEME_DES=%PS1_DES_DIR%%OMP_THEME_FILE%"

:: Ensure source file exists
if not exist "%PS1_SRC%" (
  echo ERROR: Source file not found: "%PS1_SRC%"
  pause
  exit /b 1
)

if not exist "%OMP_THEME_SRC%" (
  echo ERROR: Source theme not found: "%OMP_THEME_SRC%"
  pause
  exit /b 1
)

:: Ensure target directory exists
if not exist "%PS1_DES_DIR%" (
  echo Powershell profile target directory is not exist
  echo Target folder will be created at: "%PS1_DES_DIR%"
  echo Press Ctrl+C to cancel
  pause >nul
  mkdir "%PS1_DES_DIR%"
)

:: Remove existing file
echo Removing existing profile.ps1 if they exist
if exist "%PS1_DES%" del "%PS1_DES%"

echo Removing existing omp.theme if they exist
if exist "%OMP_THEME_DES%" del "%OMP_THEME_DES%"

:: Create symlinks
echo Creating symlinks profile.ps1
mklink "%PS1_DES%" "%PS1_SRC%"

echo Creating symlinks omp.theme
mklink "%OMP_THEME_DES%" "%OMP_THEME_SRC%"

echo Powershell Setup Done !!

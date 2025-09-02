@echo off
echo Starting Powershell setup !!

:: Get PowerShell profile path
for /f "delims=" %%i in ('powershell -NoProfile -Command "$PROFILE"') do set "PS_PROFILE=%%i"
echo %PS_PROFILE%

echo Removing existing profile.ps1 if they exist
if exist "%PS_PROFILE%" del "%PS_PROFILE%"

echo Creating symlinks
mklink "%PS_PROFILE%" "%~dp0Microsoft.PowerShell_profile.ps1"

echo Powershell setup done !!

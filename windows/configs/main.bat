@echo off
echo.
echo Start Windows Configs Setup !!
echo.

:: resolve repo root (2 levels up)
for %%I in ("%~dp0..\..") do set "REPO_ROOT=%%~fI"

echo Repo root detected: %REPO_ROOT%

echo.

call "%~dp0\nvim.bat" "%REPO_ROOT%"

echo.

call "%~dp0\vscode.bat" "%REPO_ROOT%"

echo.

call "%~dp0\powershell.bat" "%REPO_ROOT%"

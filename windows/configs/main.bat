@echo off
echo.
echo Start Windows Configs Setup !!
echo.

:: resolve repo root (2 levels up)
for %%I in ("%~dp0..\..") do set "REPO_ROOT=%%~fI"

echo Repo root detected: %REPO_ROOT%

:: No Powershell check since it is Windows
call "%~dp0\powershell.bat" "%REPO_ROOT%"

echo.

where nvim >nul 2>nul
if %ERRORLEVEL%==0 (
  call "%~dp0\nvim.bat" "%REPO_ROOT%"
)

echo.

where code >nul 2>nul
if %ERRORLEVEL%==0 (
  call "%~dp0\vscode.bat" "%REPO_ROOT%"
)

echo.

echo We done!!!!!

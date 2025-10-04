@echo off
echo VSCode Setup Start !!

set "BASE_DIR=%~1%"
set "VSCODE_DES=%USERPROFILE%\AppData\Roaming\Code\User"
set "VSCODE_SRC=%BASE_DIR%\vscode"
set "VSCODE_EXTENSIONS=%VSCODE_SRC%\extensions.txt"

:: Ensure source file exists
if not exist "%VSCODE_SRC%" (
  echo ERROR: Source file not found: "%VSCODE_SRC%"
  pause
  exit /b 1
)

:: Ensure target directory exists
if not exist "%VSCODE_DES%" (
  echo VSCODE target directory is not exist
  echo Target directory will be created at: "%VSCODE_DES%"
  echo Press Ctrl+C to cancel
  pause >nul
  mkdir "%VSCODE_DES%"
)

:: Remove existing init file
if exist "%VSCODE_DES%\settings.json" del "%VSCODE_DES%\settings.json"
if exist "%VSCODE_DES%\keybindings.json" del "%VSCODE_DES%\keybindings.json"

:: Create symlinks
echo Creating symlinks
mklink "%VSCODE_DES%\settings.json" "%VSCODE_SRC%\settings.json"
mklink "%VSCODE_DES%\keybindings.json" "%VSCODE_SRC%\keybindings.json"

:: Install VSCode Extensions
echo Install extensions

:: Ensure extensions file exists
if not exist "%VSCODE_EXTENSIONS%" (
  echo ERROR: Extensions file not found: "%VSCODE_EXTENSIONS%"
  pause
  exit /b 1
)

:: Loop through each extension 
for /f "usebackq delims=" %%e in ("%VSCODE_EXTENSIONS%") do ( 
  call code --install-extension %%e --extensions-dir "%USERPROFILE%\.vscode\extensions"
)

echo VSCode Setup Done !!

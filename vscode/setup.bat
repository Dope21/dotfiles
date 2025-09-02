@echo off
echo Starting VSCode setup !!

set "VSCODE_USER=%USERPROFILE%\AppData\Roaming\Code\User"

echo Removing existing settings.json and keybindings.json if they exist
if exist "%VSCODE_USER%\settings.json" del "%VSCODE_USER%\settings.json"
if exist "%VSCODE_USER%\keybindings.json" del "%VSCODE_USER%\keybindings.json"

echo Creating symlinks
mklink "%VSCODE_USER%\settings.json" "%~dp0settings.json"
mklink "%VSCODE_USER%\keybindings.json" "%~dp0keybindings.json"

echo VSCode setup done !!
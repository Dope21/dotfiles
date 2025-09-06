@echo off
echo Starting NeoVim setup !!

set "INIT_FILE=init.vim"
set "VIMRC=%LOCALAPPDATA%\nvim\%INIT_FILE%"

echo Removing existing init if they exist
if exist "%VIMRC%" del "%VIMRC%"

echo Creating symlinks
mklink "%VIMRC%" "%~dp0%INIT_FILE%"

echo NeoVim setup done !!

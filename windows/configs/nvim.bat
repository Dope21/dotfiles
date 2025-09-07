@echo off
echo Neovim Setup Start !!

set "BASE_DIR=%~1"
set "INIT_FILE=init.vim"
set "VIMRC_DES=%LOCALAPPDATA%\nvim\%INIT_FILE%"
set "VIMRC_SRC=%BASE_DIR%\nvim\%INIT_FILE%"

:: Ensure source file exists
if not exist "%VIMRC_SRC%" (
  echo ERROR: Source file not found: "%VIMRC_SRC%"
  pause
  exit /b 1
)

:: Ensure target directory exists
if not exist "%LOCALAPPDATA%\nvim" (
  echo Nvim target directory is not exist
  echo Target directory will be created at: "%LOCALAPPDATA%\nvim"
  echo Press Ctrl+C to cancel
  pause >nul
  mkdir "%LOCALAPPDATA%\nvim"
)

:: Remove existing init file
if exist "%VIMRC_DES%" del "%VIMRC_DES%"

:: Create symlinks
echo Creating symlinks
mklink "%VIMRC_DES%" "%VIMRC_SRC%"
if errorlevel 1 (
  echo ERROR: Failed to create symlink
  pause
  exit /b 1
)

echo NeoVim Setup Done !!
echo.

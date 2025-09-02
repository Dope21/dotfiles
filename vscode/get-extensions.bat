@echo off
setlocal enabledelayedexpansion

set "OUTPUT_FILE=%~dp0extensions.json"

echo {> "%OUTPUT_FILE%"
echo   "recommendations": [>> "%OUTPUT_FILE%"

for /f "delims=" %%e in ('code --list-extensions') do (
    set "ext=%%e"
    echo     "!ext!",>> "%OUTPUT_FILE%"
)

echo   ]>> "%OUTPUT_FILE%"
echo }>> "%OUTPUT_FILE%"

echo ✅ Generated %OUTPUT_FILE% with recommended extensions.

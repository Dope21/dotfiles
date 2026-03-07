for %%I in ("%~dp0") do set "VSCODE_SRC=%%~fI"
set "VSCODE_EXTENSIONS=%VSCODE_SRC%\extensions.txt"

:: Loop through each extension 
for /f "usebackq delims=" %%e in ("%VSCODE_EXTENSIONS%") do ( 
  call code --install-extension %%e --extensions-dir "%USERPROFILE%\.vscode\extensions"
)
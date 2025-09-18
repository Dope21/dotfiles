Clear-Host
Write-Host "Well cum!!, $env:USERNAME!"

Set-Alias g git
Set-Alias vim nvim

Import-Module -Name Terminal-Icons

Set-PSReadLineKeyHandler -Chord "Ctrl+f" -Function ForwardWord
Set-PSReadLineKeyHandler -Chord "Ctrl-l" -Function AcceptSuggestion

$theme = Join-Path (Split-Path $PROFILE) "oh-my-posh-theme.omp.json"
oh-my-posh init pwsh --config $theme | Invoke-Expression

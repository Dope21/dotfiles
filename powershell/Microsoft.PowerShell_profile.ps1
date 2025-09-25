Clear-Host

if (Get-Command git -ErrorAction SilentlyContinue) {
  Set-Alias g git
}

if (Get-Command nvim -ErrorAction SilentlyContinue) {
  Set-Alias vim nvim
}

if (Get-Module -ListAvailable -Name Terminal-Icons) {
  Import-Module -Name Terminal-Icons
}

if (Get-Module -ListAvailable -Name PSReadLine) {
  Set-PSReadLineKeyHandler -Chord "Ctrl+f" -Function ForwardWord
  Set-PSReadLineKeyHandler -Chord "Ctrl-l" -Function AcceptSuggestion
}

if (Get-Command oh-my-posh -ErrorAction SilentlyContinue) {
  $theme = Join-Path (Split-Path $PROFILE) "oh-my-posh-theme.omp.json"
  oh-my-posh init pwsh --config $theme | Invoke-Expression
}

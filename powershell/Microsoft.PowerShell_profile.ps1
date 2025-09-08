
Set-Alias g git
Set-Alias vim nvim

$theme = Join-Path (Split-Path $PROFILE) "oh-my-posh-theme.omp.json"
oh-my-posh init pwsh --config $theme | Invoke-Expression


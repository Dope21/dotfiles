set fish_greeting ""

set -gx PATH /opt/homebrew/bin $PATH
set -gx PATH /usr/local/bin $PATH
set -gx PATH /home/linuxbrew/.linuxbrew/bin $PATH

# alias
command -qv nvim && alias vim nvim
alias g git
alias cls clear

set fish_greeting ""

set -g theme_color_scheme terminal-dark
set -g fish_prompt_pwd_dir_length 1
set -g theme_display_user yes
set -g theme_hide_hostname no
set -g theme_hostname always

eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"

alias vim "nvim"
alias python "python3"
alias py "python3"

if type -q exa
  alias ll "exa -l -g --icons"
  alias lla "ll -a"
end

if status is-interactive
    # Commands to run in interactive sessions can go here
end

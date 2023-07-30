set fish_greeting ""

if type -q exa
  alias ll "exa -l -g --icons"
  alias lla "ll -a"
end

alias nv "nvim"
alias py "python3"

if status is-interactive
    # Commands to run in interactive sessions can go here
end

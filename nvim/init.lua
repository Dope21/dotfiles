require('highlights')
require('plugins')
require('base')
require('mapping')

-- copy to clip board for WSL
vim.cmd [[
  augroup Yank
  autocmd!
  autocmd TextYankPost * :call system('/mnt/c/windows/system32/clip.exe ',@")
  augroup END
]]

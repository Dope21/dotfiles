local status, packer = pcall(require, "packer")
if (not status) then
  print("packer is not installed")
  return
end

vim.cmd [[packadd packer.nvim]]

packer.startup(function(use)
  use 'wbthomason/packer.nvim'
  use {
    'svrana/neosolarized.nvim',
    requires = { 'tjdevries/colorbuddy.nvim' }
  }
  use {
    'nvim-lualine/lualine.nvim',
    requires = { 'nvim-tree/nvim-web-devicons', opt = true }
  }
  use 'neovim/nvim-lspconfig' -- LSP server
  use 'onsails/lspkind-nvim'  --
  use 'hrsh7th/cmp-buffer'    --
  use 'hrsh7th/cmp-nvim-lsp'  --
  use 'hrsh8th/nvim-cmp'      --
  use 'L3MON4D3/LuaSnip'      -- Snip engine
end)

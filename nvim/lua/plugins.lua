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
    'nvim-lualine/lualine.nvim', -- status line in nvim
    requires = { 'nvim-tree/nvim-web-devicons', opt = true }
  }
  use {
    'nvim-treesitter/nvim-treesitter', -- syntax highlight
    run = ':TSUpdate'
  }
  use 'windwp/nvim-autopairs'        -- open-close auto complete
  use 'windwp/nvim-ts-autotag'       -- auto close tag html
  use 'kyazdani42/nvim-web-devicons' -- icons
  use 'akinsho/nvim-bufferline.lua'  -- tabs buffer

  -- LSP server and auto complete plugins
  use 'neovim/nvim-lspconfig'
  use 'onsails/lspkind-nvim'
  use 'hrsh7th/cmp-buffer'
  use 'hrsh7th/cmp-nvim-lsp'
  use 'hrsh7th/nvim-cmp'
  use 'L3MON4D3/LuaSnip'

  -- Telescope
  use 'nvim-telescope/telescope.nvim'
  use 'nvim-telescope/telescope-file-browser.nvim'
  use 'nvim-lua/plenary.nvim'

  use 'norcalli/nvim-colorizer.lua' -- color highlight
  use 'lewis6991/gitsigns.nvim'     -- git marker
  use 'dinhhuy258/git.nvim'
end)

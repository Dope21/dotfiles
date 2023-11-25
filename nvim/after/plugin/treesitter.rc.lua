local status, ts = pcall(require, 'nvim-treesitter.configs')
if (not status) then return end

ts.setup {
  highlight = {
    enable = true,
    disable = {},
  },
  indent = {
    enable = true,
    disable = {},
  },
  ensure_installed = {
    'javascript',
    'html',
    'tsx',
    'lua',
    'json',
    'css',
    'python'
  },
  autotag = {
    enable = true,
  }
}

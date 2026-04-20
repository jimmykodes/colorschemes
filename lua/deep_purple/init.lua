local M = {}
local theme = require("deep_purple.theme")

function M.setup()
	vim.o.background = 'dark'
	if vim.fn.exists('syntax_on') then
		vim.cmd('syntax reset')
	end

	vim.o.termguicolors = true
	vim.g.colors_name = 'deep_purple'

	theme.setup()
end

return M

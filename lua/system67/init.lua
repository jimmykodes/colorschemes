local M = {}
local theme = require("system67.theme")

function M.setup()
	vim.o.background = 'light'
	if vim.fn.exists('syntax_on') then
		vim.cmd('syntax reset')
	end

	vim.o.termguicolors = true
	vim.g.colors_name = 'system67'

	theme.setup()
end

return M

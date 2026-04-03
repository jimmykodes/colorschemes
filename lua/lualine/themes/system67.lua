local colors = {
	normal  = '#148ba3',
	insert  = '#5e9366',
	visual  = '#956d90',
	replace = '#e95a21',
	command = '#9e744d',

	fg      = '#222222',
	bg      = '#eeeeee',
	bg_alt  = '#d5d5d5',
}

return {
	normal = {
		a = { fg = colors.bg, bg = colors.normal },
		b = { fg = colors.normal, bg = colors.bg },
		c = { fg = colors.fg, bg = colors.bg_alt },
	},
	insert = { a = { fg = colors.bg, bg = colors.insert } },
	visual = { a = { fg = colors.bg, bg = colors.visual } },
	command = { a = { fg = colors.bg, bg = colors.command } },
	replace = { a = { fg = colors.bg, bg = colors.replace } },

	inactive = {
		a = { bg = colors.bg, fg = colors.normal },
		b = { bg = colors.bg, fg = colors.bg_alt },
		c = { bg = colors.bg, fg = colors.bg_alt },
	},
}

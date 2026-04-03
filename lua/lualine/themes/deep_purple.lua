local colors = {
	normal  = '#90b0ff',
	insert  = '#49a078',
	visual  = '#bf79bf',
	replace = '#9d5060',
	command = '#d99058',

	fg      = '#bec4de',
	bg      = '#030027',
	bg_alt  = '#00001d',
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

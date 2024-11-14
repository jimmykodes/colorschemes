local colors = {
	normal  = '#80a2eb',
	insert  = '#80ddd0',
	visual  = '#b680e5',
	replace = '#dd7383',
	command = '#e5b680',

	fg      = '#b8d7db',
	bg      = '#1e2530',
	bg_alt  = '#141b26',
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

local colors = {
	normal  = '#5ca1c1',
	insert  = '#8fbc95',
	visual  = '#7c76c0',
	replace = '#BF616A',
	command = '#D08770',

	fg      = '#E5E9F0',
	bg      = '#222834',
	bg_alt  = '#181e2a',
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

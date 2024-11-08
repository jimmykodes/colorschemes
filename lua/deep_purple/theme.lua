local M = {}

local c = require("deep_purple.colors")
local hl = vim.api.nvim_set_hl

function M.setup()
	hl(0, "Normal", { fg = c.fg, bg = c.bg,  })
	-- alpha
	hl(0, "DashboardCenter", { fg = c.purple,  })
	hl(0, "DashboardFooter", { fg = c.cyan,  })
	hl(0, "DashboardHeader", { fg = c.light_blue,  })

	-- cursor
	hl(0, "Cursor", { fg = c.cursor_fg, bg = c.cursor_bg,  })
	hl(0, "CursorLine", { bg = c.dark,  })

	-- diff
	hl(0, "DiffAdd", { fg = c.alt_bg, bg = c.sign_add,  })
	hl(0, "DiffChange", { fg = c.alt_bg, bg = c.sign_change, underline = true,  })
	hl(0, "DiffDelete", { fg = c.alt_bg, bg = c.sign_delete,  })

	-- git
	hl(0, "DiffViewNormal", { fg = c.gray, bg = c.alt_bg,  })
	hl(0, "DiffviewFilePanelDeletion", { fg = c.sign_delete,  })
	hl(0, "DiffviewFilePanelInsertion", { fg = c.sign_add,  })
	hl(0, "DiffviewStatusAdded", { fg = c.sign_add,  })
	hl(0, "DiffviewStatusDeleted", { fg = c.sign_delete,  })
	hl(0, "DiffviewStatusModified", { fg = c.sign_change,  })
	hl(0, "DiffviewStatusRenamed", { fg = c.sign_change,  })
	hl(0, "DiffviewVertSplit", { bg = c.bg,  })
	hl(0, "GitSignsAdd", { fg = c.sign_add,  })
	hl(0, "GitSignsChange", { fg = c.sign_change,  })
	hl(0, "GitSignsDelete", { fg = c.sign_delete,  })
	hl(0, "SignAdd", { fg = c.sign_add,  })
	hl(0, "SignChange", { fg = c.sign_change,  })
	hl(0, "SignDelete", { fg = c.sign_delete,  })
	hl(0, "diffAdded", { fg = c.sign_add,  })
	hl(0, "diffFile", { fg = c.alt_bg,  })
	hl(0, "diffFileId", { fg = c.blue,  })
	hl(0, "diffNewFile", { fg = c.green,  })
	hl(0, "diffOldFile", { fg = c.red,  })
	hl(0, "diffRemoved", { fg = c.sign_delete,  })

	-- highlights
	hl(0, "Bold", { bold = true,  })
	hl(0, "ColorColumn", { bg = c.dark,  })
	hl(0, "Comment", { fg = c.comment, italic = true,  })
	hl(0, "Conceal", { fg = c.accent,  })
	hl(0, "Debug", { fg = c.red,  })
	hl(0, "Delimiter", { fg = c.gray,  })
	hl(0, "Directory", { fg = c.blue,  })
	hl(0, "EndOfBuffer", { fg = c.bg,  })
	hl(0, "Error", { fg = c.error_red, bg = c.bg, bold = true,  })
	hl(0, "ErrorMsg", { fg = c.error_red, bg = c.bg, bold = true,  })
	hl(0, "FloatBorder", { fg = c.gray, bg = c.alt_bg,  })
	hl(0, "FoldColumn", { fg = c.accent, bg = c.alt_bg,  })
	hl(0, "Folded", { fg = c.accent, bg = c.alt_bg,  })
	hl(0, "Ignore", { fg = c.cyan, bg = c.bg, bold = true,  })
	hl(0, "IncSearch", { fg = c.light_gray, bg = c.search_blue,  })
	hl(0, "Italic", { italic = true,  })
	hl(0, "LineNr", { fg = c.gray,  })
	hl(0, "MatchParen", { fg = c.hint_blue, bg = c.bg, underline = true,  })
	hl(0, "MatchParenCur", { underline = true,  })
	hl(0, "MatchWord", { underline = true,  })
	hl(0, "MatchWordCur", { underline = true,  })
	hl(0, "ModeMsg", { fg = c.fg, bg = c.bg,  })
	hl(0, "MoreMsg", { fg = c.orange,  })
	hl(0, "MsgArea", { fg = c.fg, bg = c.bg,  })
	hl(0, "MsgSeparator", { fg = c.fg, bg = c.bg,  })
	hl(0, "NonText", { fg = c.bg,  })
	hl(0, "NormalFloat", { bg = c.dark,  })
	hl(0, "NormalNC", { fg = c.fg, bg = c.bg,  })
	hl(0, "Pmenu", { fg = c.light_gray,  })
	hl(0, "PmenuSbar", { bg = c.dark,  })
	hl(0, "PmenuSel", { fg = c.alt_bg, bg = c.blue,  })
	hl(0, "PmenuThumb", { bg = c.gray,  })
	hl(0, "Question", { fg = c.orange,  })
	hl(0, "QuickFixLine", { bg = c.black,  })
	hl(0, "Search", { fg = c.light_gray, bg = c.search_blue,  })
	hl(0, "SignColumn", { bg = c.bg,  })
	hl(0, "SpecialComment", { fg = c.comment,  })
	hl(0, "SpecialKey", { fg = c.blue, bold = true,  })
	hl(0, "SpellBad", { fg = c.error_red, underline = true,  })
	hl(0, "SpellCap", { fg = c.yellow, underline = true,  })
	hl(0, "SpellLocal", { fg = c.green, underline = true,  })
	hl(0, "SpellRare", { fg = c.purple, underline = true,  })
	hl(0, "Substitute", { fg = c.light_gray, bg = c.search_orange,  })
	hl(0, "TabLine", { fg = c.light_gray, bg = c.alt_bg,  })
	hl(0, "TabLineFill", { fg = c.white, bg = c.alt_bg,  })
	hl(0, "TabLineSel", { fg = c.white, bg = c.alt_bg,  })
	hl(0, "Title", { fg = c.blue, bold = true,  })
	hl(0, "Todo", { fg = c.purple, bg = c.bg, bold = true,  })
	hl(0, "Underlined", { underline = true,  })
	hl(0, "VertSplit", { fg = c.fg, bg = c.bg,  })
	hl(0, "Visual", { bg = c.ui_blue,  })
	hl(0, "VisualNOS", { bg = c.alt_bg,  })
	hl(0, "WarningMsg", { fg = c.error_red, bg = c.bg,  })
	hl(0, "Whitespace", { fg = c.bg,  })
	hl(0, "WildMenu", { fg = c.alt_bg, bg = c.blue,  })

	-- illuminate
	hl(0, "IlluminatedWordRead", { bg = c.highlight_alt_bg,  })
	hl(0, "IlluminatedWordText", { bg = c.highlight_alt_bg,  })
	hl(0, "IlluminatedWordWrite", { bg = c.highlight_alt_bg,  })

	-- lsp
	hl(0, "LspDiagnosticsDefaultError", { fg = c.error_red,  })
	hl(0, "LspDiagnosticsDefaultHint", { fg = c.hint_blue,  })
	hl(0, "LspDiagnosticsDefaultInformation", { fg = c.info_yellow,  })
	hl(0, "LspDiagnosticsDefaultWarning", { fg = c.warning_orange,  })
	hl(0, "LspDiagnosticsError", { fg = c.error_red,  })
	hl(0, "LspDiagnosticsFloatingError", { fg = c.error_red,  })
	hl(0, "LspDiagnosticsFloatingHint", { fg = c.hint_blue,  })
	hl(0, "LspDiagnosticsFloatingInformation", { fg = c.info_yellow,  })
	hl(0, "LspDiagnosticsFloatingWarning", { fg = c.warning_orange,  })
	hl(0, "LspDiagnosticsHint", { fg = c.hint_blue,  })
	hl(0, "LspDiagnosticsInformation", { fg = c.info_yellow,  })
	hl(0, "LspDiagnosticsSignError", { fg = c.error_red,  })
	hl(0, "LspDiagnosticsSignHint", { fg = c.hint_blue,  })
	hl(0, "LspDiagnosticsSignInformation", { fg = c.info_yellow,  })
	hl(0, "LspDiagnosticsSignWarning", { fg = c.warning_orange,  })
	hl(0, "LspDiagnosticsUnderlineError", { underline = true,  })
	hl(0, "LspDiagnosticsUnderlineHint", { underline = true,  })
	hl(0, "LspDiagnosticsUnderlineInformation", { underline = true,  })
	hl(0, "LspDiagnosticsUnderlineWarning", { underline = true,  })
	hl(0, "LspDiagnosticsVirtualTextError", { fg = c.error_red,  })
	hl(0, "LspDiagnosticsVirtualTextHint", { fg = c.hint_blue,  })
	hl(0, "LspDiagnosticsVirtualTextInformation", { fg = c.info_yellow,  })
	hl(0, "LspDiagnosticsVirtualTextWarning", { fg = c.warning_orange,  })
	hl(0, "LspDiagnosticsWarning", { fg = c.warning_orange,  })
	hl(0, "QuickScopePrimary", { fg = c.magenta, underline = true,  })
	hl(0, "QuickScopeSecondary", { fg = c.hint_blue, underline = true,  })

	-- lualine
	hl(0, "LuaLineDiffAdd", { fg = c.sign_add,  })
	hl(0, "LuaLineDiffChange", { fg = c.sign_change,  })
	hl(0, "LuaLineDiffDelete", { fg = c.sign_delete,  })

	-- nvim_tree
	hl(0, "NvimTreeExecFile", { fg = c.green,  })
	hl(0, "NvimTreeFolderIcon", { fg = c.blue,  })
	hl(0, "NvimTreeFolderName", { fg = c.blue,  })
	hl(0, "NvimTreeGitDeleted", { fg = c.sign_delete,  })
	hl(0, "NvimTreeGitDirty", { fg = c.sign_add,  })
	hl(0, "NvimTreeGitMerge", { fg = c.sign_change,  })
	hl(0, "NvimTreeGitNew", { fg = c.sign_add,  })
	hl(0, "NvimTreeGitRenamed", { fg = c.sign_change,  })
	hl(0, "NvimTreeGitStaged", { fg = c.sign_add,  })
	hl(0, "NvimTreeImageFile", { fg = c.purple,  })
	hl(0, "NvimTreeIndentMarker", { fg = c.gray,  })
	hl(0, "NvimTreeNormal", { fg = c.light_gray, bg = c.alt_bg,  })
	hl(0, "NvimTreeOpenedFolderName", { fg = c.cyan, italic = true,  })
	hl(0, "NvimTreeRootFolder", { fg = c.fg, bold = true,  })
	hl(0, "NvimTreeSpecialFile", { fg = c.orange,  })
	hl(0, "NvimTreeSymlink", { fg = c.cyan,  })
	hl(0, "NvimTreeVertSplit", { fg = c.alt_bg, bg = c.alt_bg,  })

	-- telescope
	hl(0, "TelescopeMatching", { fg = c.info_yellow, bold = true,  })
	hl(0, "TelescopeSelection", { fg = c.hint_blue,  })

	-- treesitter
	hl(0, "TSAnnotation", { fg = c.yellow,  })
	hl(0, "TSAttribute", { fg = c.orange,  })
	hl(0, "TSBoolean", { fg = c.blue,  })
	hl(0, "TSCharacter", { fg = c.orange,  })
	hl(0, "TSComment", { fg = c.comment,  })
	hl(0, "TSConditional", { fg = c.blue,  })
	hl(0, "TSConstBuiltin", { fg = c.blue,  })
	hl(0, "TSConstMacro", { fg = c.orange,  })
	hl(0, "TSConstant", { fg = c.yellow,  })
	hl(0, "TSConstructor", { fg = c.orange,  })
	hl(0, "TSEmphasis", { italic = true,  })
	hl(0, "TSError", { fg = c.error_red,  })
	hl(0, "TSException", { fg = c.purple,  })
	hl(0, "TSField", { fg = c.light_blue,  })
	hl(0, "TSFloat", { fg = c.green,  })
	hl(0, "TSFuncBuiltin", { fg = c.yellow,  })
	hl(0, "TSFuncMacro", { fg = c.yellow,  })
	hl(0, "TSFunction", { fg = c.yellow,  })
	hl(0, "TSInclude", { fg = c.purple,  })
	hl(0, "TSInstalled", { fg = c.green, bg = c.black,  })
	hl(0, "TSKeyword", { fg = c.blue,  })
	hl(0, "TSKeywordFunction", { fg = c.blue,  })
	hl(0, "TSKeywordOperator", { fg = c.blue,  })
	hl(0, "TSKeywordReturn", { fg = c.purple,  })
	hl(0, "TSLabel", { fg = c.light_blue,  })
	hl(0, "TSLiteral", { fg = c.yellow_orange,  })
	hl(0, "TSMethod", { fg = c.yellow,  })
	hl(0, "TSMissing", { fg = c.error_red, bg = c.black,  })
	hl(0, "TSNamespace", { fg = c.orange,  })
	hl(0, "TSNumber", { fg = c.green,  })
	hl(0, "TSOperator", { fg = c.fg,  })
	hl(0, "TSParameter", { fg = c.light_blue,  })
	hl(0, "TSParameterReference", { fg = c.light_blue,  })
	hl(0, "TSProperty", { fg = c.light_blue,  })
	hl(0, "TSPunctBracket", { fg = c.fg,  })
	hl(0, "TSPunctDelimiter", { fg = c.fg,  })
	hl(0, "TSPunctSpecial", { fg = c.fg,  })
	hl(0, "TSQueryLinterError", { fg = c.warning_orange,  })
	hl(0, "TSRepeat", { fg = c.purple,  })
	hl(0, "TSString", { fg = c.cyan,  })
	hl(0, "TSStringEscape", { fg = c.green,  })
	hl(0, "TSStringRegex", { fg = c.cyan,  })
	hl(0, "TSStrong", { fg = c.yellow_orange,  })
	hl(0, "TSStructure", { fg = c.light_blue,  })
	hl(0, "TSSymbol", { fg = c.light_blue,  })
	hl(0, "TSTag", { fg = c.blue,  })
	hl(0, "TSTagDelimiter", { fg = c.gray,  })
	hl(0, "TSText", { fg = c.fg,  })
	hl(0, "TSTitle", { fg = c.blue, bold = true,  })
	hl(0, "TSType", { fg = c.orange,  })
	hl(0, "TSTypeBuiltin", { fg = c.blue,  })
	hl(0, "TSURI", { fg = c.yellow_orange, underline = true,  })
	hl(0, "TSUnderline", { underline = true,  })
	hl(0, "TSVariable", { fg = c.light_blue,  })
	hl(0, "TSVariableBuiltin", { fg = c.light_blue,  })

	-- types
	hl(0, "Boolean", { fg = c.blue,  })
	hl(0, "Character", { fg = c.orange,  })
	hl(0, "Conditional", { fg = c.blue,  })
	hl(0, "Constant", { fg = c.blue,  })
	hl(0, "Define", { fg = c.purple,  })
	hl(0, "Exception", { fg = c.purple,  })
	hl(0, "Float", { fg = c.green,  })
	hl(0, "Function", { fg = c.yellow,  })
	hl(0, "Identifier", { fg = c.light_blue,  })
	hl(0, "Include", { fg = c.purple,  })
	hl(0, "Keyword", { fg = c.purple,  })
	hl(0, "Label", { fg = c.purple,  })
	hl(0, "Macro", { fg = c.purple,  })
	hl(0, "Number", { fg = c.green,  })
	hl(0, "Operator", { fg = c.fg,  })
	hl(0, "PreCondit", { fg = c.purple,  })
	hl(0, "PreProc", { fg = c.purple,  })
	hl(0, "Repeat", { fg = c.purple,  })
	hl(0, "Special", { fg = c.orange,  })
	hl(0, "SpecialChar", { fg = c.white,  })
	hl(0, "Statement", { fg = c.purple,  })
	hl(0, "StorageClass", { fg = c.blue,  })
	hl(0, "String", { fg = c.cyan,  })
	hl(0, "Structure", { fg = c.blue,  })
	hl(0, "Tag", { fg = c.blue,  })
	hl(0, "Type", { fg = c.blue,  })
	hl(0, "Typedef", { fg = c.blue,  })
	hl(0, "Variable", { fg = c.light_blue,  })

end

return M

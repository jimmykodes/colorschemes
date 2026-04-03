package arctic

import "github.com/jimmykodes/colorschemes"

var (
	fg        = colorschemes.ColorMustFromHex("#E5E9F0")
	bg        = colorschemes.ColorMustFromHex("#222834")
	red       = colorschemes.ColorMustFromHex("#BF616A")
	orange    = colorschemes.ColorMustFromHex("#D08770")
	yellow    = colorschemes.ColorMustFromHex("#EBCB8B")
	green     = colorschemes.ColorMustFromHex("#8fbc95")
	cyan      = colorschemes.ColorMustFromHex("#88e0d0")
	blue      = colorschemes.ColorMustFromHex("#5ca1c1")
	purple    = colorschemes.ColorMustFromHex("#7c76c0")
	magenta   = colorschemes.ColorMustFromHex("#B48EAD")
	white     = colorschemes.ColorMustFromHex("#ceeff4")
	gray      = colorschemes.ColorMustFromHex("#6b7486")
	dark_gray = colorschemes.ColorMustFromHex("#3B4252")
	black     = bg.Darken(0.1)
)

var Arctic = &colorschemes.Colorscheme{
	Metadata: colorschemes.Metadata{
		Name:  "arctic",
		Theme: colorschemes.Dark,
	},

	Fg:             fg,
	Bg:             bg,
	BgDark:         bg.Darken(.1),
	BgLight:        bg.Lighten(.1),
	HighlightAltBg: bg.Lighten(.2),

	Red:      red,
	Orange:   orange,
	Yellow:   yellow,
	Green:    green,
	Cyan:     cyan,
	Blue:     blue,
	Purple:   purple,
	Magenta:  magenta,
	White:    white,
	Gray:     gray,
	DarkGray: dark_gray,

	LightBlack:  dark_gray,
	LightRed:    red.Lighten(.20),
	LightGreen:  green.Lighten(.20),
	LightYellow: yellow.Lighten(.20),
	LightBlue:   blue.Lighten(.20),
	LightPurple: purple.Lighten(.20),
	LightCyan:   cyan.Lighten(.20),
	LightWhite:  white.Lighten(.20),
	LightGray:   gray.Lighten(.20),

	Hint:    magenta,
	Info:    blue,
	Warning: yellow,
	Error:   red,
	Accent:  dark_gray,
	Search:  magenta,
	Replace: orange,
	Comment: gray,
	Context: gray,

	CursorFg:   bg,
	CursorBg:   fg,
	CursorLine: bg.Darken(.05),

	SignDiff:         yellow,
	SignDiffBg:       yellow.Darken(.25),
	SignAdd:          green,
	SignAddBg:        green.Darken(.25),
	SignStagedAdd:    green.Lerp(bg, .25),
	SignChange:       blue,
	SignChangeBg:     blue.Darken(.25),
	SignStagedChange: blue.Lerp(bg, .25),
	SignDelete:       red,
	SignDeleteBg:     red.Darken(.25),
	SignStagedDelete: red.Lerp(bg, 0.25),
}

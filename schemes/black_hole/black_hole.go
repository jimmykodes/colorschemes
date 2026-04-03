package black_hole

import "github.com/jimmykodes/colorschemes"

var (
	fg        = colorschemes.ColorMustFromHex("#fde9c3")
	bg        = colorschemes.ColorMustFromHex("#080028")
	red       = colorschemes.ColorMustFromHex("#fe5683")
	orange    = colorschemes.ColorMustFromHex("#fd7e00")
	yellow    = colorschemes.ColorMustFromHex("#fdca33")
	green     = colorschemes.ColorMustFromHex("#ffb102")
	cyan      = colorschemes.ColorMustFromHex("#a9bdff")
	blue      = colorschemes.ColorMustFromHex("#5d6aff")
	purple    = colorschemes.ColorMustFromHex("#b203ea")
	magenta   = colorschemes.ColorMustFromHex("#ab1a95")
	white     = colorschemes.ColorMustFromHex("#ffffff")
	gray      = colorschemes.ColorMustFromHex("#817ee0")
	dark_gray = gray.Darken(.50)
	black     = bg.Darken(0.1)
)

var BlackHole = &colorschemes.Colorscheme{
	Metadata: colorschemes.Metadata{
		Name:  "black_hole",
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
	LightGray:   gray.Lighten(.50),

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
	SignDelete:       orange,
	SignDeleteBg:     orange.Darken(.25),
	SignStagedDelete: orange.Lerp(bg, 0.25),
}

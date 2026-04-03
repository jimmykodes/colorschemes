package deep_purple

import "github.com/jimmykodes/colorschemes"

var (
	fg      = colorschemes.ColorMustFromHex("#bec4de")
	bg      = colorschemes.ColorMustFromHex("#030027")
	red     = colorschemes.ColorMustFromHex("#9d5060")
	green   = colorschemes.ColorMustFromHex("#49a078")
	yellow  = colorschemes.ColorMustFromHex("#d7ba89")
	blue    = colorschemes.ColorMustFromHex("#90b0ff")
	purple  = colorschemes.ColorMustFromHex("#bf79bf")
	cyan    = colorschemes.ColorMustFromHex("#a1cca5")
	orange  = colorschemes.ColorMustFromHex("#d99058")
	magenta = colorschemes.ColorMustFromHex("#bf40bf")
	white   = colorschemes.ColorMustFromHex("#e5dada")
	gray    = colorschemes.ColorMustFromHex("#8c92ac")

	dark_gray = colorschemes.ColorMustFromHex("#36454f").Lighten(.50)
	black     = bg.Darken(0.1)
)

var DeepPurple = &colorschemes.Colorscheme{
	Metadata: colorschemes.Metadata{
		Name:  "deep_purple",
		Theme: colorschemes.Dark,
	},

	Fg:             fg,
	Bg:             bg,
	BgDark:         bg.Darken(.1),
	BgLight:        bg.Lighten(.1),
	HighlightAltBg: bg.Lighten(.1).Lighten(.2),

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
	Black:    black,

	LightBlack:  dark_gray,
	LightRed:    red.Lighten(.20),
	LightGreen:  green.Lighten(.20),
	LightYellow: yellow.Lighten(.20),
	LightBlue:   blue.Lighten(.20),
	LightPurple: colorschemes.ColorMustFromHex("#736b92"),
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
	SignDiffBg:       yellow.Darken(.50),
	SignAdd:          green,
	SignAddBg:        green.Darken(.50),
	SignStagedAdd:    green.Lerp(bg, .25),
	SignChange:       blue,
	SignChangeBg:     blue.Darken(.50),
	SignStagedChange: blue.Lerp(bg, .25),
	SignDelete:       orange,
	SignDeleteBg:     orange.Darken(.50),
	SignStagedDelete: orange.Lerp(bg, 0.25),
}

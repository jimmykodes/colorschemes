package system76

import "github.com/jimmykodes/colorschemes"

var (
	fg       = colorschemes.ColorMustFromHex("#cccccc")
	bg       = colorschemes.ColorMustFromHex("#222222")
	red      = colorschemes.ColorMustFromHex("#f15d22")
	green    = colorschemes.ColorMustFromHex("#6eaa76")
	yellow   = colorschemes.ColorMustFromHex("#eece61")
	blue     = colorschemes.ColorMustFromHex("#18a9c7")
	purple   = colorschemes.ColorMustFromHex("#ad7fa8")
	cyan     = colorschemes.ColorMustFromHex("#7ec589")
	orange   = colorschemes.ColorMustFromHex("#d19a66")
	magenta  = colorschemes.ColorMustFromHex("#D16D9E")
	white    = colorschemes.ColorMustFromHex("#D8DEE9")
	gray     = colorschemes.ColorMustFromHex("#98918d")
	darkGray = colorschemes.ColorMustFromHex("#3e3e3e")
	darkBlue = colorschemes.ColorMustFromHex("#223E55")

	signAdd    = colorschemes.ColorMustFromHex("#587c0c")
	signChange = colorschemes.ColorMustFromHex("#0c7d9d")
	signDelete = colorschemes.ColorMustFromHex("#b81a22")
)

var System76 = &colorschemes.Colorscheme{
	Metadata: colorschemes.Metadata{
		Name:  "system76",
		Theme: colorschemes.Dark,
	},

	Fg:             fg,
	Bg:             bg,
	BgDark:         bg.Darken(.1),
	BgLight:        bg.Lighten(.05),
	HighlightAltBg: bg.Lighten(.1),

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
	DarkGray: darkGray,
	Black:    bg.Darken(.1),

	LightBlack:  darkGray,
	LightRed:    colorschemes.ColorMustFromHex("#D16969"),
	LightGreen:  colorschemes.ColorMustFromHex("#B5CEA8"),
	LightYellow: yellow.Lighten(.20),
	LightBlue:   colorschemes.ColorMustFromHex("#81c8d6"),
	LightPurple: purple.Lighten(.20),
	LightCyan:   cyan.Lighten(.20),
	LightWhite:  white.Lighten(.20),
	LightGray:   colorschemes.ColorMustFromHex("#c8c9c1"),

	Hint:    colorschemes.ColorMustFromHex("#4FC1FF"),
	Info:    colorschemes.ColorMustFromHex("#264F78"),
	Warning: colorschemes.ColorMustFromHex("#ff8800"),
	Error:   colorschemes.ColorMustFromHex("#F44747"),

	Accent:  colorschemes.ColorMustFromHex("#BBBBBB"),
	Search:  darkBlue,
	Replace: colorschemes.ColorMustFromHex("#613214"),
	Comment: colorschemes.ColorMustFromHex("#787878"),
	Context: colorschemes.ColorMustFromHex("#606060"),

	CursorFg:   bg,
	CursorBg:   cyan,
	CursorLine: bg.Darken(.05),

	SignDiff:         colorschemes.ColorMustFromHex("#D7BA7D"),
	SignDiffBg:       colorschemes.ColorMustFromHex("#583a00"),
	SignAdd:          signAdd,
	SignAddBg:        colorschemes.ColorMustFromHex("#17310a"),
	SignStagedAdd:    signAdd.Lerp(bg, .5),
	SignChange:       signChange,
	SignChangeBg:     colorschemes.ColorMustFromHex("#0c3058"),
	SignStagedChange: signChange.Lerp(bg, .5),
	SignDelete:       signDelete,
	SignDeleteBg:     colorschemes.ColorMustFromHex("#461a22"),
	SignStagedDelete: signDelete.Lerp(bg, .5),
}

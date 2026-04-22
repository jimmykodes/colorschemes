package system67

import "github.com/jimmykodes/colorschemes"

var (
	fg       = colorschemes.ColorMustFromHex("#222222")
	bg       = colorschemes.ColorMustFromHex("#eeeeee")
	red      = colorschemes.ColorMustFromHex("#e95a21")
	green    = colorschemes.ColorMustFromHex("#5e9366")
	yellow   = colorschemes.ColorMustFromHex("#907d3b")
	blue     = colorschemes.ColorMustFromHex("#148ba3")
	purple   = colorschemes.ColorMustFromHex("#956d90")
	cyan     = colorschemes.ColorMustFromHex("#5d9266")
	orange   = colorschemes.ColorMustFromHex("#9e744d")
	magenta  = colorschemes.ColorMustFromHex("#ad5a83")
	white    = colorschemes.ColorMustFromHex("#7b7e84")
	gray     = colorschemes.ColorMustFromHex("#52524f")
	darkGray = colorschemes.ColorMustFromHex("#302f2f")

	signAdd    = colorschemes.ColorMustFromHex("#587c0c")
	signChange = colorschemes.ColorMustFromHex("#0c7d9d")
	signDelete = colorschemes.ColorMustFromHex("#b81a22")
)

var System67 = &colorschemes.Colorscheme{
	Metadata: colorschemes.Metadata{
		Name:  "system67",
		Theme: colorschemes.Light,
	},

	Fg:             fg,
	Bg:             bg,
	BgDark:         bg.Darken(.1),
	BgLight:        bg.Lighten(.1),
	HighlightAltBg: bg.Lighten(.1).Lighten(.1),

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

	LightBlack:  darkGray,
	LightRed:    colorschemes.ColorMustFromHex("#b25a5a"),
	LightGreen:  colorschemes.ColorMustFromHex("#6e7e66"),
	LightYellow: yellow.Lighten(.20),
	LightBlue:   colorschemes.ColorMustFromHex("#5a8b95"),
	LightPurple: purple.Lighten(.20),
	LightCyan:   cyan.Lighten(.20),
	LightWhite:  white.Lighten(.20),
	LightGray:   colorschemes.ColorMustFromHex("#7f7f7a"),

	Hint:    colorschemes.ColorMustFromHex("#3b91bf"),
	Info:    colorschemes.ColorMustFromHex("#264F78"),
	Success: signAdd,
	Warning: colorschemes.ColorMustFromHex("#c66900"),
	Error:   colorschemes.ColorMustFromHex("#F44747"),

	Accent:  colorschemes.ColorMustFromHex("#7d7d7d"),
	Search:  colorschemes.ColorMustFromHex("#5e81ac"),
	Replace: colorschemes.ColorMustFromHex("#613214"),
	Comment: colorschemes.ColorMustFromHex("#787878"),
	Context: colorschemes.ColorMustFromHex("#606060"),

	CursorFg:   bg,
	CursorBg:   cyan,
	CursorLine: bg.Darken(.05),

	SignDiff:         colorschemes.ColorMustFromHex("#87754e"),
	SignDiffBg:       colorschemes.ColorMustFromHex("#583a00"),
	SignAdd:          signAdd,
	SignAddBg:        colorschemes.ColorMustFromHex("#17310a"),
	SignStagedAdd:    signAdd.Lerp(bg, .25),
	SignChange:       signChange,
	SignChangeBg:     colorschemes.ColorMustFromHex("#0c3058"),
	SignStagedChange: signChange.Lerp(bg, .25),
	SignDelete:       signDelete,
	SignDeleteBg:     colorschemes.ColorMustFromHex("#461a22"),
	SignStagedDelete: signDelete.Lerp(bg, .25),
}

package pokemon

import "github.com/jimmykodes/colorschemes"

var (
	fg = colorschemes.ColorMustFromHex("#d1e5d9")
	bg = colorschemes.ColorMustFromHex("#232d36")

	_normal           = colorschemes.ColorMustFromHex("#9fa19f")
	_flying           = colorschemes.ColorMustFromHex("#81b9ef")
	_poison           = colorschemes.ColorMustFromHex("#9141cb")
	_ground           = colorschemes.ColorMustFromHex("#915121")
	_rock             = colorschemes.ColorMustFromHex("#afa981")
	_steel            = colorschemes.ColorMustFromHex("#60a1b8")
	_fire             = colorschemes.ColorMustFromHex("#e62829")
	_water            = colorschemes.ColorMustFromHex("#2980ef")
	_grass            = colorschemes.ColorMustFromHex("#3fa129")
	_electric         = colorschemes.ColorMustFromHex("#fac000")
	_ice              = colorschemes.ColorMustFromHex("#3dcef3")
	_dragon           = colorschemes.ColorMustFromHex("#5060e1")
	_fairy            = colorschemes.ColorMustFromHex("#ef70ef")
	_bug              = colorschemes.ColorMustFromHex("#91a119")
	_ghost            = colorschemes.ColorMustFromHex("#704170")
	_dark             = colorschemes.ColorMustFromHex("#624d4e")
	_fighting         = colorschemes.ColorMustFromHex("#ff8000")
	_psychic          = colorschemes.ColorMustFromHex("#ef4179")
	_hp               = colorschemes.ColorMustFromHex("#69dc12")
	_light_hp         = colorschemes.ColorMustFromHex("#9ee865")
	_attack           = colorschemes.ColorMustFromHex("#efcc18")
	_light_attack     = colorschemes.ColorMustFromHex("#f5de69")
	_defense          = colorschemes.ColorMustFromHex("#e86412")
	_light_defense    = colorschemes.ColorMustFromHex("#f09a65")
	_sp_attack        = colorschemes.ColorMustFromHex("#14c3f1")
	_light_sp_attack  = colorschemes.ColorMustFromHex("#66d8f6")
	_sp_defense       = colorschemes.ColorMustFromHex("#4a6adf")
	_light_sp_defense = colorschemes.ColorMustFromHex("#899eea")
	_speed            = colorschemes.ColorMustFromHex("#d51dad")
	_light_speed      = colorschemes.ColorMustFromHex("#e46cca")

	red       = _fire.Lighten(0.40)
	green     = _light_hp
	yellow    = _light_attack
	blue      = _light_sp_defense
	purple    = _poison.Lighten(0.40)
	cyan      = _light_sp_attack
	orange    = _light_defense
	magenta   = _light_speed
	white     = colorschemes.ColorMustFromHex("#e5f1ec")
	gray      = _normal
	dark_gray = _dark
	black     = bg.Darken(0.10)

	light_black  = dark_gray
	light_red    = red.Lighten(0.20)
	light_green  = green.Lighten(0.20)
	light_yellow = yellow.Lighten(0.20)
	light_blue   = blue.Lighten(0.20)
	light_purple = purple.Lighten(0.20)
	light_cyan   = cyan.Lighten(0.20)
	light_white  = white.Lighten(0.20)
	light_gray   = gray.Lighten(0.50)
)

var Pokemon = &colorschemes.Colorscheme{
	Metadata: colorschemes.Metadata{
		Name:  "pokemon",
		Theme: colorschemes.Dark,
	},

	Fg:             fg,
	Bg:             bg,
	BgDark:         black,
	BgLight:        bg.Lighten(0.10),
	HighlightAltBg: bg.Lighten(0.10).Lighten(0.10),

	Red:      red,
	Green:    green,
	Yellow:   yellow,
	Blue:     blue,
	Purple:   purple,
	Cyan:     cyan,
	Orange:   orange,
	Magenta:  magenta,
	White:    white,
	Gray:     gray,
	DarkGray: dark_gray,
	Black:    black,

	LightBlack:  light_black,
	LightRed:    light_red,
	LightGreen:  light_green,
	LightYellow: light_yellow,
	LightBlue:   light_blue,
	LightPurple: light_purple,
	LightCyan:   light_cyan,
	LightWhite:  light_white,
	LightGray:   light_gray,

	Hint:    _dragon,
	Info:    _electric,
	Warning: orange,
	Error:   red,

	Accent:  _steel,
	Search:  _dragon,
	Replace: yellow,

	Comment: _steel,
	Context: _steel,

	CursorFg:   bg,
	CursorBg:   fg,
	CursorLine: black,

	SignAdd:    green,
	SignChange: blue,
	SignDelete: orange,
	SignDiff:   yellow,

	SignStagedAdd:    green.Darken(0.25),
	SignStagedChange: blue.Darken(0.25),
	SignStagedDelete: orange.Darken(0.25),

	SignAddBg:    green.Darken(0.50),
	SignChangeBg: blue.Darken(0.50),
	SignDeleteBg: orange.Darken(0.50),
	SignDiffBg:   yellow.Darken(0.50),
}

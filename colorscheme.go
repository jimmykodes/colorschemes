package colorschemes

import (
	"encoding/json"
)

type Theme string

const (
	Dark  Theme = "dark"
	Light Theme = "light"
)

type Metadata struct {
	Name  string
	Theme Theme
}

type Colorscheme struct {
	Metadata Metadata `json:"-"`
	Fg       Color    `json:"fg"`
	Bg       Color    `json:"bg"`
	Red      Color    `json:"red"`
	Orange   Color    `json:"orange"`
	Yellow   Color    `json:"yellow"`
	Green    Color    `json:"green"`
	Cyan     Color    `json:"cyan"`
	Blue     Color    `json:"blue"`
	Purple   Color    `json:"purple"`
	Magenta  Color    `json:"magenta"`
	White    Color    `json:"white"`
	Gray     Color    `json:"gray"`
	DarkGray Color    `json:"dark_gray"`
	Black    Color    `json:"black"`

	LightBlack  Color `json:"light_black"`
	LightRed    Color `json:"light_red"`
	LightGreen  Color `json:"light_green"`
	LightYellow Color `json:"light_yellow"`
	LightBlue   Color `json:"light_blue"`
	LightPurple Color `json:"light_purple"`
	LightCyan   Color `json:"light_cyan"`
	LightWhite  Color `json:"light_white"`
	LightGray   Color `json:"light_gray"`

	Hint    Color `json:"hint"`
	Info    Color `json:"info"`
	Warning Color `json:"warning"`
	Error   Color `json:"error"`

	Accent  Color `json:"accent"`
	Search  Color `json:"search"`
	Replace Color `json:"replace"`

	BgDark         Color `json:"bg_dark"`
	BgLight        Color `json:"bg_light"`
	HighlightAltBg Color `json:"highlight_alt_bg"`

	Comment Color `json:"comment"`
	Context Color `json:"context"`

	CursorFg   Color `json:"cursor_fg"`
	CursorBg   Color `json:"cursor_bg"`
	CursorLine Color `json:"cursor_line"`

	SignDiff   Color `json:"sign_diff"`
	SignDiffBg Color `json:"sign_diff_bg"`

	SignAdd       Color `json:"sign_add"`
	SignStagedAdd Color `json:"sign_staged_add"`
	SignAddBg     Color `json:"sign_add_bg"`

	SignChange       Color `json:"sign_change"`
	SignChangeBg     Color `json:"sign_change_bg"`
	SignStagedChange Color `json:"sign_staged_change"`

	SignDelete       Color `json:"sign_delete"`
	SignDeleteBg     Color `json:"sign_delete_bg"`
	SignStagedDelete Color `json:"sign_staged_delete"`
}

func (c Colorscheme) Map() (map[string]string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	var out map[string]string
	err = json.Unmarshal(data, &out)
	return out, err
}

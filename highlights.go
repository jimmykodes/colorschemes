package colorschemes

import (
	"bytes"
	"fmt"
)

type Highlight struct {
	FG string
	BG string

	Bold      bool
	Underline bool
	Italic    bool

	Link string
}

func (h *Highlight) UnmarshalText(text []byte) error {
	parts := bytes.Split(text, []byte{' '})
	switch len(parts) {
	case 0:
		return fmt.Errorf("invalid data - must provide at least one value")
	case 1:
		h.FG = string(parts[0])
	case 2:
		switch p := string(parts[0]); p {
		case "link":
			h.Link = string(parts[1])
		default:
			h.FG = string(parts[0])
			h.BG = string(parts[1])
		}
	default:
		h.FG = string(parts[0])
		h.BG = string(parts[1])
		for _, part := range parts[2:] {
			switch string(part) {
			case "u", "ul", "underline":
				h.Underline = true
			case "b", "bold":
				h.Bold = true
			case "i", "italic":
				h.Italic = true
			}
		}
	}
	return nil
}

type LualineHL struct {
	Normal  string `yaml:"normal"`
	Insert  string `yaml:"insert"`
	Visual  string `yaml:"visual"`
	Replace string `yaml:"replace"`
	Command string `yaml:"command"`
	Fg      string `yaml:"fg"`
	Bg      string `yaml:"bg"`
	BgAlt   string `yaml:"bg_alt"`
}

type TerminalHL struct {
	Foreground   string `yaml:"fg"`
	Background   string `yaml:"bg"`
	CursorFG     string `yaml:"cursor_fg"`
	CursorBG     string `yaml:"cursor_bg"`
	SelectionBG  string `yaml:"selection_bg"`
	SelectionFG  string `yaml:"selection_fg"`
	Black        string `yaml:"black"`
	Red          string `yaml:"red"`
	Green        string `yaml:"green"`
	Yellow       string `yaml:"yellow"`
	Blue         string `yaml:"blue"`
	Purple       string `yaml:"purple"`
	Cyan         string `yaml:"cyan"`
	White        string `yaml:"white"`
	BrightBlack  string `yaml:"bright_black"`
	BrightRed    string `yaml:"bright_red"`
	BrightGreen  string `yaml:"bright_green"`
	BrightYellow string `yaml:"bright_yellow"`
	BrightBlue   string `yaml:"bright_blue"`
	BrightPurple string `yaml:"bright_purple"`
	BrightCyan   string `yaml:"bright_cyan"`
	BrightWhite  string `yaml:"bright_white"`
}

type K9sHL struct {
	FG           string `yaml:"fg"`
	BG           string `yaml:"bg"`
	Black        string `yaml:"black"`
	Blue         string `yaml:"blue"`
	Green        string `yaml:"green"`
	Gray         string `yaml:"gray"`
	Orange       string `yaml:"orange"`
	Purple       string `yaml:"purple"`
	Red          string `yaml:"red"`
	Yellow       string `yaml:"yellow"`
	BrightYellow string `yaml:"brightYellow"`
}

type Highlights struct {
	Lualine  *LualineHL  `yaml:"lualine"`
	Terminal *TerminalHL `yaml:"terminal"`
	K9s      *K9sHL      `yaml:"k9s"`
	Normal   Highlight   `yaml:"Normal"`

	Groups map[string]map[string]Highlight `yaml:"groups"`
}

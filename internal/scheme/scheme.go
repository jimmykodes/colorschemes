package scheme

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Scheme struct {
	Metadata struct {
		Name  string `yaml:"name"`
		Theme string `yaml:"theme"`
	} `yaml:"metadata"`
	Lualine struct {
		Normal  string `yaml:"normal"`
		Insert  string `yaml:"insert"`
		Visual  string `yaml:"visual"`
		Replace string `yaml:"replace"`
		Command string `yaml:"command"`
		Fg      string `yaml:"fg"`
		Bg      string `yaml:"bg"`
		BgAlt   string `yaml:"bg_alt"`
	} `yaml:"lualine"`
	Normal Highlight                       `yaml:"Normal"`
	HTML   []string                        `yaml:"html"`
	Colors map[string]*Color               `yaml:"colors"`
	Groups map[string]map[string]Highlight `yaml:"groups"`
}

type Highlight struct {
	FG string
	BG string

	Bold      bool
	Underline bool
	Italic    bool
}

func (h *Highlight) UnmarshalText(text []byte) error {
	parts := bytes.Split(text, []byte{' '})
	switch len(parts) {
	case 0:
		return fmt.Errorf("invalid data - must provide at least one value")
	case 1:
		h.FG = string(parts[0])
	case 2:
		h.FG = string(parts[0])
		h.BG = string(parts[1])
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

type Color struct {
	value string
	f     string
	args  []string
}

func (c Color) String() string {
	return c.value
}

func (c Color) R(colors map[string]*Color) (int64, error) {
	val, err := c.Value(colors)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(val[1:3], 16, 64)
}

func (c Color) G(colors map[string]*Color) (int64, error) {
	val, err := c.Value(colors)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(val[3:5], 16, 64)
}

func (c Color) B(colors map[string]*Color) (int64, error) {
	val, err := c.Value(colors)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(val[5:], 16, 64)
}

func (c *Color) UnmarshalText(text []byte) error {
	if text[0] == '#' {
		c.value = string(text)
		return nil
	}
	f, argbytes, found := bytes.Cut(text, []byte{' '})
	if !found {
		return fmt.Errorf("invalid data - must provide a hex color or a function and at least one arg")
	}
	c.f = string(f)
	c.args = strings.Split(string(argbytes), " ")
	return nil
}

func clamp(val, lower, upper int64) int64 {
	return max(min(val, upper), lower)
}

func rgbClamp(val int64) int64 {
	return clamp(val, 0, 255)
}

type _funcs struct{}

var funcs = _funcs{}

func (f _funcs) darken(colors map[string]*Color, c string, amount int64) (string, error) {
	col, ok := colors[c]
	if !ok {
		return "", fmt.Errorf("invalid color %s", c)
	}

	r, err := col.R(colors)
	if err != nil {
		return "", err
	}
	g, err := col.G(colors)
	if err != nil {
		return "", err
	}
	b, err := col.B(colors)
	if err != nil {
		return "", err
	}

	r = rgbClamp(r - amount)
	g = rgbClamp(g - amount)
	b = rgbClamp(b - amount)

	return fmt.Sprintf("#%02x%02x%02x", r, g, b), nil
}

func (f _funcs) lighten(colors map[string]*Color, c string, amount int64) (string, error) {
	return f.darken(colors, c, -amount)
}

func (f _funcs) ref(colors map[string]*Color, c string) (string, error) {
	col, ok := colors[c]
	if !ok {
		return "", fmt.Errorf("invalid color %s", c)
	}
	return col.Value(colors)
}

func (c *Color) Value(colors map[string]*Color) (string, error) {
	if c.value != "" {
		return c.value, nil
	}
	switch c.f {
	case "darken":
		amount, err := strconv.ParseInt(c.args[1], 10, 64)
		if err != nil {
			return "", err
		}
		val, err := funcs.darken(colors, c.args[0], amount)
		if err != nil {
			return "", err
		}
		c.value = val
		return val, nil
	case "lighten":
		amount, err := strconv.ParseInt(c.args[1], 10, 64)
		if err != nil {
			return "", err
		}
		val, err := funcs.lighten(colors, c.args[0], amount)
		if err != nil {
			return "", err
		}
		c.value = val
		return val, nil
	case "ref":
		val, err := funcs.ref(colors, c.args[0])
		if err != nil {
			return "", err
		}
		c.value = val
		return val, nil
	default:
		return "", fmt.Errorf("invalid function %s", c.f)
	}
}

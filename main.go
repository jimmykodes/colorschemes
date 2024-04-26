package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

func main() {
	scheme := flag.String("scheme", "", "generate a specific colorscheme")
	flag.Parse()

	if err := run(*scheme); err != nil {
		os.Exit(1)
	}
}

func run(scheme string) error {
	schemes, err := loadSchemes(scheme)
	if err != nil {
		return err
	}

	for _, scheme := range schemes {
		fmt.Println("Processing", scheme.Name)
		if schemeErr := genScheme(scheme); schemeErr != nil {
			fmt.Println("Error processing", scheme.Name, "-", schemeErr)
			err = errors.Join(err, schemeErr)
		}
	}
	return err
}

type SchemeTemplate struct {
	Name string
	Ext  string
}

func loadSchemes(scheme string) ([]SchemeTemplate, error) {
	entries, err := os.ReadDir("templates")
	if err != nil {
		return nil, err
	}
	out := make([]SchemeTemplate, 0, len(entries))
	for _, entry := range entries {
		name := entry.Name()
		ext := filepath.Ext(name)
		if ext != ".yaml" && ext != ".yml" {
			continue
		}
		name = strings.TrimSuffix(name, ext)
		if scheme == "" {
			out = append(out, SchemeTemplate{Name: name, Ext: ext})
			continue
		}

		if name == scheme {
			return []SchemeTemplate{{Name: name, Ext: ext}}, nil
		}
	}
	return out, nil
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
	Colors map[string]*Color               `yaml:"colors"`
	Groups map[string]map[string]Highlight `yaml:"groups"`
}

func genScheme(schemeTemplate SchemeTemplate) error {
	fp := filepath.Join("templates", schemeTemplate.Name+schemeTemplate.Ext)
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	var scheme Scheme
	if err := yaml.NewDecoder(f).Decode(&scheme); err != nil {
		return err
	}

	for _, color := range scheme.Colors {
		_, err := color.Value(scheme.Colors)
		if err != nil {
			return err
		}
	}

	colorDir := "colors"
	if err := os.MkdirAll(colorDir, 0766); err != nil && !os.IsExist(err) {
		return err
	}
	if err := createVim(colorDir, scheme); err != nil {
		return err
	}

	lualineDir := filepath.Join("lua", "lualine", "themes")
	if err := createLualine(lualineDir, scheme); err != nil {
		return err
	}

	luaDir := filepath.Join("lua", scheme.Metadata.Name)
	if err := os.MkdirAll(luaDir, 0766); err != nil && !os.IsExist(err) {
		return err
	}
	if err := createInit(luaDir, scheme); err != nil {
		return err
	}
	if err := createColors(luaDir, scheme); err != nil {
		return err
	}
	if err := createTheme(luaDir, scheme); err != nil {
		return err
	}
	return nil
}

func createLualine(themesDir string, scheme Scheme) error {
	f, err := os.Create(filepath.Join(themesDir, scheme.Metadata.Name+".lua"))
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl := template.Must(template.New("lualine").Parse(lualineTmpl))
	return tmpl.Execute(f, scheme)
}

func createVim(colorDir string, scheme Scheme) error {
	f, err := os.Create(filepath.Join(colorDir, scheme.Metadata.Name+".vim"))
	if err != nil {
		return err
	}
	defer f.Close()
	tmpl := template.Must(template.New("scheme").Parse(vimTmpl))
	if err := tmpl.Execute(f, scheme); err != nil {
		return err
	}
	return nil
}

func createInit(luaDir string, scheme Scheme) error {
	f, err := os.Create(filepath.Join(luaDir, "init.lua"))
	if err != nil {
		return err
	}
	defer f.Close()
	tmpl := template.Must(template.New("scheme").Parse(initTmpl))
	if err := tmpl.Execute(f, scheme); err != nil {
		return err
	}
	return nil
}

func createColors(luaDir string, scheme Scheme) error {
	f, err := os.Create(filepath.Join(luaDir, "colors.lua"))
	if err != nil {
		return err
	}
	defer f.Close()
	tmpl := template.Must(template.New("scheme").Parse(colorsTmpl))
	if err := tmpl.Execute(f, scheme); err != nil {
		return err
	}
	return nil
}

func createTheme(luaDir string, scheme Scheme) error {
	f, err := os.Create(filepath.Join(luaDir, "theme.lua"))
	if err != nil {
		return err
	}
	defer f.Close()
	tmpl := template.Must(template.New("scheme").Parse(themeTmpl))
	if err := tmpl.Execute(f, scheme); err != nil {
		return err
	}
	return nil
}

var (
	themeTmpl = `
{{- define "style" -}}
{ {{ if ne .FG "-" }}fg = c.{{ .FG }}, {{ end }}
{{- if .BG }}{{if ne .BG "-" }}bg = c.{{ .BG }}, {{end}}{{ end -}}
{{- if .Bold }}bold = true, {{ end -}}
{{- if .Underline }}underline = true, {{ end -}}
{{- if .Italic }}italic = true, {{ end }} }
{{- end -}}

local M = {}

local c = require("{{ .Metadata.Name }}.colors")
local hl = vim.api.nvim_set_hl

function M.setup()
	hl(0, "Normal", {{ template "style" .Normal }})
{{- range $group, $gmap := .Groups }}
	-- {{ $group }}
	{{- range $name, $hl := $gmap }}
	hl(0, "{{ $name }}", {{ template "style" $hl }})
	{{- end }}
{{ end }}
end

return M
`

	initTmpl = `local M = {}
local theme = require("{{ .Metadata.Name }}.theme")

function M.setup()
	vim.cmd('hi clear')

	vim.o.background = '{{ .Metadata.Theme }}'
	if vim.fn.exists('syntax_on') then
		vim.cmd('syntax reset')
	end

	vim.o.termguicolors = true
	vim.g.colors_name = '{{ .Metadata.Name }}'

	theme.setup()
end

return M
`

	colorsTmpl = `local M = {}
{{- $colors := .Colors }}
{{ range $name, $color := .Colors }}M.{{$name}} = '{{ $color.Value $colors }}'
{{end}}
return M
`

	vimTmpl = `lua << EOF
local {{.Metadata.Name }} = require("{{ .Metadata.Name }}")
{{ .Metadata.Name }}.setup()
EOF
`

	lualineTmpl = `local colors = {
	normal  = '{{ .Lualine.Normal | index .Colors }}',
	insert  = '{{ .Lualine.Insert | index .Colors }}',
	visual  = '{{ .Lualine.Visual | index .Colors }}',
	replace = '{{ .Lualine.Replace | index .Colors }}',
	command = '{{ .Lualine.Command | index .Colors }}',

	fg      = '{{ .Lualine.Fg | index .Colors }}',
	bg      = '{{ .Lualine.Bg | index .Colors }}',
	bg_alt  = '{{ .Lualine.BgAlt | index .Colors }}',
}

return {
	normal = {
		a = { fg = colors.bg, bg = colors.normal },
		b = { fg = colors.normal, bg = colors.bg },
		c = { fg = colors.fg, bg = colors.bg_alt },
	},
	insert = { a = { fg = colors.bg, bg = colors.insert } },
	visual = { a = { fg = colors.bg, bg = colors.visual } },
	command = { a = { fg = colors.bg, bg = colors.command } },
	replace = { a = { fg = colors.bg, bg = colors.replace } },

	inactive = {
		a = { bg = colors.bg, fg = colors.normal },
		b = { bg = colors.bg, fg = colors.bg_alt },
		c = { bg = colors.bg, fg = colors.bg_alt },
	},
}
`
)

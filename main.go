package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
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

type Scheme struct {
	Metadata struct {
		Name  string `yaml:"name"`
		Theme string `yaml:"theme"`
	} `yaml:"metadata"`
	Colors map[string]string               `yaml:"colors"`
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
	colorDir := "colors"
	if err := os.MkdirAll(colorDir, 0766); err != nil && !os.IsExist(err) {
		return err
	}
	if err := createVim(colorDir, scheme); err != nil {
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

var themeTmpl = `
{{- define "style" -}}
{ fg = {{ if ne .FG "-" }}c.{{ .FG }}{{else}}"fg"{{ end }}
{{- if .BG }}, bg = {{if ne .BG "-" }}{{ .BG }}{{else}}"bg"{{end}}{{ end -}}
{{- if .Bold }}, bold = true{{ end -}}
{{- if .Underline }}, underline = true{{ end -}}
{{- if .Italic }}, italic = true{{ end }} }
{{- end -}}

local M = {}

local c = require("{{ .Metadata.Name }}.colors")
local hl = vim.api.nvim_set_hl

function M.setup()
{{- range $group, $gmap := .Groups }}
	-- {{ $group }}
	{{- range $name, $hl := $gmap }}
	hl(0, "{{ $name }}", {{ template "style" $hl }})
	{{- end }}
{{ end }}
end

return M
`

var initTmpl = `local M = {}
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

var colorsTmpl = `local M = {}
{{ range $name, $color := .Colors }}M.{{$name}} = '{{ $color }}'
{{end}}
return M
`

var vimTmpl = `lua << EOF
local {{.Metadata.Name }} = require("{{ .Metadata.Name }}")
{{ .Metadata.Name }}.setup()
EOF
`

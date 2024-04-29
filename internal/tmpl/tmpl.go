package tmpl

import (
	"embed"
	"os"
	"path/filepath"
	"text/template"

	"github.com/jimmykodes/colorschemes/internal/scheme"
)

//go:embed templates
var TemplatesDir embed.FS

type Tmpl struct {
	tmpl *template.Template
}

func New() (*Tmpl, error) {
	tmpl, err := template.New("colors").ParseFS(TemplatesDir, "templates/*.gotmpl")
	if err != nil {
		return nil, err
	}
	return &Tmpl{tmpl: tmpl}, nil
}

func (t *Tmpl) Lualine(themesDir string, scheme scheme.Scheme) error {
	f, err := os.Create(filepath.Join(themesDir, scheme.Metadata.Name+".lua"))
	if err != nil {
		return err
	}
	defer f.Close()

	return t.tmpl.ExecuteTemplate(f, "lualine.gotmpl", scheme)
}

func (t *Tmpl) HTML(htmlDir string, scheme scheme.Scheme) error {
	f, err := os.Create(filepath.Join(htmlDir, scheme.Metadata.Name+".html"))
	if err != nil {
		return err
	}
	defer f.Close()
	return t.tmpl.ExecuteTemplate(f, "html.gotmpl", scheme)
}

func (t *Tmpl) Vim(colorDir string, scheme scheme.Scheme) error {
	f, err := os.Create(filepath.Join(colorDir, scheme.Metadata.Name+".vim"))
	if err != nil {
		return err
	}
	defer f.Close()
	return t.tmpl.ExecuteTemplate(f, "vim.gotmpl", scheme)
}

func (t *Tmpl) Init(luaDir string, scheme scheme.Scheme) error {
	f, err := os.Create(filepath.Join(luaDir, "init.lua"))
	if err != nil {
		return err
	}
	defer f.Close()
	return t.tmpl.ExecuteTemplate(f, "init.gotmpl", scheme)
}

func (t *Tmpl) Colors(luaDir string, scheme scheme.Scheme) error {
	f, err := os.Create(filepath.Join(luaDir, "colors.lua"))
	if err != nil {
		return err
	}
	defer f.Close()
	return t.tmpl.ExecuteTemplate(f, "colors.gotmpl", scheme)
}

func (t *Tmpl) Theme(luaDir string, scheme scheme.Scheme) error {
	f, err := os.Create(filepath.Join(luaDir, "theme.lua"))
	if err != nil {
		return err
	}
	defer f.Close()
	return t.tmpl.ExecuteTemplate(f, "theme.gotmpl", scheme)
}

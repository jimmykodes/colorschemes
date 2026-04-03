package tmpl

import (
	"embed"
	"os"
	"path/filepath"
	"text/template"

	"github.com/jimmykodes/colorschemes"
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

// func (t *Tmpl) Lualine(themesDir string, scheme scheme.Scheme, name string) error {
// 	f, err := os.Create(filepath.Join(themesDir, name+".lua"))
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
//
// 	return t.tmpl.ExecuteTemplate(f, "lualine.gotmpl", scheme)
// }

// func (t *Tmpl) HTML(htmlDir string, scheme scheme.Scheme, name string) error {
// 	f, err := os.Create(filepath.Join(htmlDir, name+".html"))
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	return t.tmpl.ExecuteTemplate(f, "html.gotmpl", scheme)
// }

// func (t *Tmpl) Ghostty(dir string, scheme scheme.Scheme, name string) error {
// 	f, err := os.Create(filepath.Join(dir, name))
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	return t.tmpl.ExecuteTemplate(f, "ghostty.gotmpl", scheme)
// }
//
// func (t *Tmpl) WezTerm(weztermDir string, scheme scheme.Scheme, name string) error {
// 	f, err := os.Create(filepath.Join(weztermDir, name+".toml"))
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	return t.tmpl.ExecuteTemplate(f, "wezterm.gotmpl", scheme)
// }

// func (t *Tmpl) K9s(dir string, scheme scheme.Scheme, name string) error {
// 	f, err := os.Create(filepath.Join(dir, name+".yaml"))
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	return t.tmpl.ExecuteTemplate(f, "k9s.gotmpl", scheme)
// }

type TmplContext struct {
	Metadata colorschemes.Metadata
	HL       *colorschemes.Highlights
	Colors   map[string]string
}

func (t *Tmpl) Lua(colorDir string, data *TmplContext) error {
	f, err := os.Create(filepath.Join(colorDir, data.Metadata.Name+".lua"))
	if err != nil {
		return err
	}
	defer f.Close()
	return t.tmpl.ExecuteTemplate(f, "lua.gotmpl", data)
}

func (t *Tmpl) Init(luaDir string, data *TmplContext) error {
	f, err := os.Create(filepath.Join(luaDir, "init.lua"))
	if err != nil {
		return err
	}
	defer f.Close()
	return t.tmpl.ExecuteTemplate(f, "init.gotmpl", data)
}

func (t *Tmpl) Colors(luaDir string, data *TmplContext) error {
	f, err := os.Create(filepath.Join(luaDir, "colors.lua"))
	if err != nil {
		return err
	}
	defer f.Close()
	return t.tmpl.ExecuteTemplate(f, "colors.gotmpl", data)
}

func (t *Tmpl) Theme(luaDir string, data *TmplContext) error {
	f, err := os.Create(filepath.Join(luaDir, "theme.lua"))
	if err != nil {
		return err
	}
	defer f.Close()
	return t.tmpl.ExecuteTemplate(f, "theme.gotmpl", data)
}

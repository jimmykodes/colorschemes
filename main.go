package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jimmykodes/colorschemes/internal/scheme"
	"github.com/jimmykodes/colorschemes/internal/tmpl"
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

func genScheme(schemeTemplate SchemeTemplate) error {
	fp := filepath.Join("templates", schemeTemplate.Name+schemeTemplate.Ext)
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	var scheme scheme.Scheme
	if err := yaml.NewDecoder(f).Decode(&scheme); err != nil {
		return err
	}

	// MARK: parsing
	for _, color := range scheme.Colors {
		_, err := color.Value(scheme.Colors)
		if err != nil {
			return err
		}
	}

	for _, group := range scheme.Groups {
		for name, hl := range group {
			if _, ok := scheme.Colors[hl.FG]; hl.FG != "" && hl.FG != "-" && !ok {
				fmt.Printf("%s: invalid fg color '%s' for name %s\n", schemeTemplate.Name, hl.FG, name)
			}
			if _, ok := scheme.Colors[hl.BG]; hl.BG != "" && hl.BG != "-" && !ok {
				fmt.Printf("%s: invalid bg color '%s' for name %s\n", schemeTemplate.Name, hl.BG, name)
			}
		}
	}
	templates, err := tmpl.New()
	if err != nil {
		return err
	}

	weztermDir := filepath.Join("configs", "wezterm")
	if err := os.MkdirAll(weztermDir, 0766); err != nil && !os.IsExist(err) {
		return err
	}
	if err := templates.WezTerm(weztermDir, scheme); err != nil {
		return err
	}

	k9sDir := filepath.Join("configs", "k9s")
	if err := os.MkdirAll(k9sDir, 0766); err != nil && !os.IsExist(err) {
		return err
	}
	if err := templates.K9s(k9sDir, scheme); err != nil {
		return err
	}

	htmlDir := filepath.Join("examples", "html")
	if err := os.MkdirAll(htmlDir, 0766); err != nil && !os.IsExist(err) {
		return err
	}
	if err := templates.HTML(htmlDir, scheme); err != nil {
		return err
	}

	colorDir := "colors"
	if err := os.MkdirAll(colorDir, 0766); err != nil && !os.IsExist(err) {
		return err
	}
	if err := templates.Vim(colorDir, scheme); err != nil {
		return err
	}

	lualineDir := filepath.Join("lua", "lualine", "themes")
	if err := templates.Lualine(lualineDir, scheme); err != nil {
		return err
	}

	luaDir := filepath.Join("lua", scheme.Metadata.Name)
	if err := os.MkdirAll(luaDir, 0766); err != nil && !os.IsExist(err) {
		return err
	}
	if err := templates.Init(luaDir, scheme); err != nil {
		return err
	}
	if err := templates.Colors(luaDir, scheme); err != nil {
		return err
	}
	if err := templates.Theme(luaDir, scheme); err != nil {
		return err
	}
	return nil
}

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jimmykodes/colorschemes/internal/scheme"
	"github.com/jimmykodes/colorschemes/internal/tmpl"
	"gopkg.in/yaml.v3"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	scheme, err := decodeScheme()
	if err != nil {
		fmt.Println("Error decoding scheme", "-", err)
		return err
	}

	fmt.Println("processing", scheme.Metadata.Name)
	if err := genScheme(scheme); err != nil {
		fmt.Println("Error processing", scheme.Metadata.Name, "-", err)
		return err
	}

	return nil
}

func decodeScheme() (scheme.Scheme, error) {
	var scheme scheme.Scheme
	err := yaml.NewDecoder(os.Stdin).Decode(&scheme)
	return scheme, err
}

func genScheme(scheme scheme.Scheme) error {
	for _, color := range scheme.Colors {
		_, err := color.Value(scheme.Colors)
		if err != nil {
			return err
		}
	}

	for _, group := range scheme.Groups {
		for name, hl := range group {
			if _, ok := scheme.Colors[hl.FG]; hl.FG != "" && hl.FG != "-" && !ok {
				fmt.Printf("%s: invalid fg color '%s' for name %s\n", scheme.Metadata.Name, hl.FG, name)
			}
			if _, ok := scheme.Colors[hl.BG]; hl.BG != "" && hl.BG != "-" && !ok {
				fmt.Printf("%s: invalid bg color '%s' for name %s\n", scheme.Metadata.Name, hl.BG, name)
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

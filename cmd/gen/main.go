package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jimmykodes/colorschemes"
	"github.com/jimmykodes/colorschemes/internal/tmpl"
	"github.com/jimmykodes/colorschemes/schemes"
	"github.com/jimmykodes/gommand"
)

func main() {
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

var root = &gommand.Command{
	Name: "gen",
	Run: func(ctx *gommand.Context) error {
		templates, err := tmpl.New()
		if err != nil {
			return err
		}
		for _, cs := range schemes.All {
			colors, err := cs.Map()
			if err != nil {
				return err
			}
			data := &tmpl.TmplContext{HL: colorschemes.BaseHighlights, Colors: colors, Metadata: cs.Metadata}
			if err := errors.Join(
				nvim(templates, data),
				k9s(templates, data),
				ghostty(templates, data),
			); err != nil {
				return err
			}
		}
		return nil
	},
}

func nvim(templates *tmpl.Tmpl, data *tmpl.TmplContext) error {
	nvimDir := "./"

	colorDir := filepath.Join(nvimDir, "colors")
	if err := os.MkdirAll(colorDir, 0o766); err != nil && !os.IsExist(err) {
		return fmt.Errorf("make colors dir: %w", err)
	}
	if err := templates.Lua(colorDir, data); err != nil {
		return fmt.Errorf("generate colors lua: %w", err)
	}

	luaDir := filepath.Join(nvimDir, "lua", data.Metadata.Name)
	if err := os.MkdirAll(luaDir, 0o766); err != nil && !os.IsExist(err) {
		return fmt.Errorf("make lua dir: %w", err)
	}

	if err := templates.Init(luaDir, data); err != nil {
		return fmt.Errorf("generate lua init: %w", err)
	}

	if err := templates.Colors(luaDir, data); err != nil {
		return fmt.Errorf("generate lua colors: %w", err)
	}

	if err := templates.Theme(luaDir, data); err != nil {
		return fmt.Errorf("generate themes: %w", err)
	}

	lualineDir := filepath.Join(nvimDir, "lua", "lualine", "themes")
	if err := os.MkdirAll(lualineDir, 0o766); err != nil && !os.IsExist(err) {
		return fmt.Errorf("make lualine dir: %w", err)
	}
	if err := templates.Lualine(lualineDir, data); err != nil {
		return fmt.Errorf("generate lualine: %w", err)
	}

	return nil
}

func ghostty(templates *tmpl.Tmpl, data *tmpl.TmplContext) error {
	dir := "./extras/k9s/"

	if err := os.MkdirAll(dir, 0o766); err != nil && !os.IsExist(err) {
		return fmt.Errorf("make ghostty dir: %w", err)
	}
	if err := templates.Ghostty(dir, data); err != nil {
		return fmt.Errorf("generate ghostty: %w", err)
	}
	return nil
}

func k9s(templates *tmpl.Tmpl, data *tmpl.TmplContext) error {
	dir := "./extras/k9s/"

	if err := os.MkdirAll(dir, 0o766); err != nil && !os.IsExist(err) {
		return fmt.Errorf("make k9s dir: %w", err)
	}
	if err := templates.K9s(dir, data); err != nil {
		return fmt.Errorf("generate k9s: %w", err)
	}
	return nil
}

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
	"github.com/jimmykodes/gommand/flags"
)

func main() {
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

var root = &gommand.Command{
	Name: "gen",
	FlagSet: flags.NewFlagSet().AddFlags(
		flags.StringFlag("nvim-dir", "", "Directory for nvim colors"),
		flags.StringFlag("k9s-dir", "", "Directory for k9s colors"),
		flags.StringFlag("ghostty-dir", "", "Directory for k9s colors"),
	),
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
				nvim(ctx, templates, data),
				// k9s(ctx, templates, s),
				// ghostty(ctx, templates, s),
			); err != nil {
				return err
			}
		}
		return nil
	},
}

func nvim(ctx *gommand.Context, templates *tmpl.Tmpl, data *tmpl.TmplContext) error {
	nvimDir := ctx.Flags().String("nvim-dir")

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
	//
	// lualineDir := filepath.Join(nvimDir, "lua", "lualine", "themes")
	// if err := os.MkdirAll(lualineDir, 0o766); err != nil && !os.IsExist(err) {
	// 	return fmt.Errorf("make lualine dir: %w", err)
	// }
	// if err := templates.Lualine(lualineDir, s, name); err != nil {
	// 	return fmt.Errorf("generate lualine: %w", err)
	// }

	return nil
}

// func ghostty(ctx *gommand.Context, templates *tmpl.Tmpl, s scheme.Scheme, name string) error {
// 	dir := ctx.Flags().String("ghostty-dir")
// 	if err := os.MkdirAll(dir, 0o766); err != nil && !os.IsExist(err) {
// 		return fmt.Errorf("make ghostty dir: %w", err)
// 	}
// 	if err := templates.Ghostty(dir, s, name); err != nil {
// 		return fmt.Errorf("generate ghostty: %w", err)
// 	}
// 	return nil
// }
//
// func k9s(ctx *gommand.Context, templates *tmpl.Tmpl, s scheme.Scheme, name string) error {
// 	k9sDir := ctx.Flags().String("k9s-dir")
// 	if err := os.MkdirAll(k9sDir, 0o766); err != nil && !os.IsExist(err) {
// 		return fmt.Errorf("make k9s dir: %w", err)
// 	}
// 	if err := templates.K9s(k9sDir, s, name); err != nil {
// 		return fmt.Errorf("generate k9s: %w", err)
// 	}
// 	return nil
// }

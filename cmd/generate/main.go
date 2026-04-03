package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jimmykodes/colorschemes/internal/parser"
	"github.com/jimmykodes/colorschemes/internal/scheme"
	"github.com/jimmykodes/colorschemes/internal/tmpl"
	"github.com/jimmykodes/gommand"
	"github.com/jimmykodes/gommand/flags"
)

func main() {
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

var root = &gommand.Command{
	Name: "generate",
	FlagSet: flags.NewFlagSet().AddFlags(
		flags.StringFlag("nvim-dir", "", "Directory for nvim colors"),
		flags.StringFlag("k9s-dir", "", "Directory for k9s colors"),
		flags.StringFlag("ghostty-dir", "", "Directory for k9s colors"),
	),
	Run: parser.WithParsedSchema(func(ctx *gommand.Context, s scheme.Scheme) error {
		templates, err := tmpl.New()
		if err != nil {
			return err
		}
		return errors.Join(
			k9s(ctx, templates, s),
			nvim(ctx, templates, s),
			ghostty(ctx, templates, s),
		)
	}),
}

func ghostty(ctx *gommand.Context, templates *tmpl.Tmpl, s scheme.Scheme) error {
	dir := ctx.Flags().String("ghostty-dir")
	if err := os.MkdirAll(dir, 0o766); err != nil && !os.IsExist(err) {
		return fmt.Errorf("make ghostty dir: %w", err)
	}
	if err := templates.Ghostty(dir, s); err != nil {
		return fmt.Errorf("generate ghostty: %w", err)
	}
	return nil
}

func k9s(ctx *gommand.Context, templates *tmpl.Tmpl, s scheme.Scheme) error {
	k9sDir := ctx.Flags().String("k9s-dir")
	if err := os.MkdirAll(k9sDir, 0o766); err != nil && !os.IsExist(err) {
		return fmt.Errorf("make k9s dir: %w", err)
	}
	if err := templates.K9s(k9sDir, s); err != nil {
		return fmt.Errorf("generate k9s: %w", err)
	}
	return nil
}

func nvim(ctx *gommand.Context, templates *tmpl.Tmpl, s scheme.Scheme) error {
	nvimDir := ctx.Flags().String("nvim-dir")

	colorDir := filepath.Join(nvimDir, "colors")
	if err := os.MkdirAll(colorDir, 0o766); err != nil && !os.IsExist(err) {
		return fmt.Errorf("make colors dir: %w", err)
	}
	if err := templates.Vim(colorDir, s); err != nil {
		return fmt.Errorf("generate colors: %w", err)
	}

	luaDir := filepath.Join(nvimDir, "lua", s.Metadata.Name)
	if err := os.MkdirAll(luaDir, 0o766); err != nil && !os.IsExist(err) {
		return fmt.Errorf("make lua dir: %w", err)
	}
	if err := templates.Init(luaDir, s); err != nil {
		return fmt.Errorf("generate lua: %w", err)
	}
	if err := templates.Colors(luaDir, s); err != nil {
		return fmt.Errorf("generate colors: %w", err)
	}
	if err := templates.Theme(luaDir, s); err != nil {
		return fmt.Errorf("generate themes: %w", err)
	}

	lualineDir := filepath.Join(nvimDir, "lua", "lualine", "themes")
	if err := os.MkdirAll(lualineDir, 0o766); err != nil && !os.IsExist(err) {
		return fmt.Errorf("make lualine dir: %w", err)
	}
	if err := templates.Lualine(lualineDir, s); err != nil {
		return fmt.Errorf("generate lualine: %w", err)
	}

	return nil
}

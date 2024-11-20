package nvim

import (
	"os"
	"path/filepath"

	"github.com/jimmykodes/colorschemes/internal/parser"
	"github.com/jimmykodes/colorschemes/internal/scheme"
	"github.com/jimmykodes/colorschemes/internal/tmpl"
	"github.com/jimmykodes/gommand"
	"github.com/jimmykodes/gommand/flags"
)

func Cmd() *gommand.Command {
	cmd := &gommand.Command{
		Name: "nvim --nvim-dir <dir>",
		FlagSet: flags.NewFlagSet().AddFlag(
			flags.StringFlag("nvim-dir", "", "Directory for nvim colors"),
		),
		Run: parser.WithParsedSchema(func(ctx *gommand.Context, s scheme.Scheme) error {
			templates, err := tmpl.New()
			if err != nil {
				return err
			}
			nvimDir := ctx.Flags().String("nvim-dir")

			colorDir := filepath.Join(nvimDir, "colors")
			if err := os.MkdirAll(colorDir, 0766); err != nil && !os.IsExist(err) {
				return err
			}
			if err := templates.Vim(colorDir, s); err != nil {
				return err
			}

			luaDir := filepath.Join(nvimDir, "lua", s.Metadata.Name)
			if err := os.MkdirAll(luaDir, 0766); err != nil && !os.IsExist(err) {
				return err
			}
			if err := templates.Init(luaDir, s); err != nil {
				return err
			}
			if err := templates.Colors(luaDir, s); err != nil {
				return err
			}
			if err := templates.Theme(luaDir, s); err != nil {
				return err
			}

			lualineDir := filepath.Join(nvimDir, "lua", "lualine", "themes")
			if err := os.MkdirAll(lualineDir, 0766); err != nil && !os.IsExist(err) {
				return err
			}
			if err := templates.Lualine(lualineDir, s); err != nil {
				return err
			}

			return nil
		}),
	}
	return cmd
}

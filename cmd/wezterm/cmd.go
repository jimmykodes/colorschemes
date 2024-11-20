package wezterm

import (
	"os"

	"github.com/jimmykodes/colorschemes/internal/parser"
	"github.com/jimmykodes/colorschemes/internal/scheme"
	"github.com/jimmykodes/colorschemes/internal/tmpl"
	"github.com/jimmykodes/gommand"
	"github.com/jimmykodes/gommand/flags"
)

func Cmd() *gommand.Command {
	cmd := &gommand.Command{
		Name: "wezterm --wezterm-dir <dir>",
		FlagSet: flags.NewFlagSet().AddFlag(
			flags.StringFlag("wezterm-dir", "", "Directory for wezterm colors"),
		),
		Run: parser.WithParsedSchema(func(ctx *gommand.Context, s scheme.Scheme) error {
			templates, err := tmpl.New()
			if err != nil {
				return err
			}

			weztermDir := ctx.Flags().String("wezterm-dir")
			if err := os.MkdirAll(weztermDir, 0766); err != nil && !os.IsExist(err) {
				return err
			}
			if err := templates.WezTerm(weztermDir, s); err != nil {
				return err
			}
			return nil
		}),
	}
	return cmd
}

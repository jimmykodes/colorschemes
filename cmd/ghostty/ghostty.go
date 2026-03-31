package ghostty

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
		Name: "ghostty",
		FlagSet: flags.NewFlagSet().AddFlag(
			flags.StringFlag("dir", "", "Directory for ghostty colors"),
		),
		Run: parser.WithParsedSchema(func(ctx *gommand.Context, s scheme.Scheme) error {
			templates, err := tmpl.New()
			if err != nil {
				return err
			}

			dir := ctx.Flags().String("dir")
			if err := os.MkdirAll(dir, 0o766); err != nil && !os.IsExist(err) {
				return err
			}
			if err := templates.Ghostty(dir, s); err != nil {
				return err
			}
			return nil
		}),
	}
	return cmd
}

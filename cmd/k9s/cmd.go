package k9s

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
		Name: "k9s --k9s-dir <dir>",
		FlagSet: flags.NewFlagSet().AddFlag(
			flags.StringFlag("k9s-dir", "", "Directory for k9s colors"),
		),
		Run: parser.WithParsedSchema(func(ctx *gommand.Context, s scheme.Scheme) error {
			templates, err := tmpl.New()
			if err != nil {
				return err
			}

			k9sDir := ctx.Flags().String("k9s-dir")
			if err := os.MkdirAll(k9sDir, 0766); err != nil && !os.IsExist(err) {
				return err
			}
			if err := templates.K9s(k9sDir, s); err != nil {
				return err
			}
			return nil
		}),
	}
	return cmd
}

package images

import (
	"os"
	"path/filepath"

	"github.com/jimmykodes/colorschemes/internal/parser"
	"github.com/jimmykodes/colorschemes/internal/scheme"
	"github.com/jimmykodes/colorschemes/internal/tmpl"
	"github.com/jimmykodes/gommand"
)

func Cmd() *gommand.Command {
	cmd := &gommand.Command{
		Name: "images",
		Run: parser.WithParsedSchema(func(ctx *gommand.Context, s scheme.Scheme) error {
			templates, err := tmpl.New()
			if err != nil {
				return err
			}

			htmlDir := filepath.Join("examples", "html")
			if err := os.MkdirAll(htmlDir, 0766); err != nil && !os.IsExist(err) {
				return err
			}
			if err := templates.HTML(htmlDir, s); err != nil {
				return err
			}

			return nil
		}),
	}
	return cmd
}

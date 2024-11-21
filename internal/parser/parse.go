package parser

import (
	"errors"
	"fmt"
	"os"

	"github.com/jimmykodes/colorschemes/internal/scheme"
	"github.com/jimmykodes/gommand"
	"gopkg.in/yaml.v3"
)

func WithParsedSchema(f func(*gommand.Context, scheme.Scheme) error) func(ctx *gommand.Context) error {
	return func(ctx *gommand.Context) error {
		s, err := decodeScheme()
		if err != nil {
			fmt.Println("Error decoding scheme", "-", err)
			return err
		}
		for _, color := range s.Colors {
			_, err := color.Value(s.Colors)
			if err != nil {
				return err
			}
		}
		var invalidColors error
		for _, group := range s.Groups {
			for name, hl := range group {
				if _, ok := s.Colors[hl.FG]; hl.FG != "" && hl.FG != "-" && !ok {
					invalidColors = errors.Join(
						invalidColors,
						fmt.Errorf("%s: invalid fg color '%s' for name %s\n", s.Metadata.Name, hl.FG, name),
					)
				}
				if _, ok := s.Colors[hl.BG]; hl.BG != "" && hl.BG != "-" && !ok {
					invalidColors = errors.Join(
						invalidColors,
						fmt.Errorf("%s: invalid bg color '%s' for name %s\n", s.Metadata.Name, hl.BG, name),
					)
				}
			}
		}
		if invalidColors != nil {
			return invalidColors
		}
		return f(ctx, s)
	}
}

func decodeScheme() (scheme.Scheme, error) {
	var scheme scheme.Scheme
	err := yaml.NewDecoder(os.Stdin).Decode(&scheme)
	return scheme, err
}

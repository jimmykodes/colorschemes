package cmd

import (
	"github.com/jimmykodes/colorschemes/cmd/k9s"
	"github.com/jimmykodes/colorschemes/cmd/nvim"
	"github.com/jimmykodes/colorschemes/cmd/wezterm"
	"github.com/jimmykodes/gommand"
)

func Cmd() *gommand.Command {
	cmd := &gommand.Command{
		Name: "colorschemes",
	}
	cmd.SubCommand(
		wezterm.Cmd(),
		k9s.Cmd(),
		nvim.Cmd(),
	)
	return cmd
}

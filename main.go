package main

import (
	"os"

	"github.com/jimmykodes/colorschemes/cmd"
)

func main() {
	if err := cmd.Cmd().Execute(); err != nil {
		os.Exit(1)
	}
}

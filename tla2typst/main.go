package main

import (
	"os"

	"github.com/randyttruong/tla2typst/cli/maincommand"
)

func main() {
	c := maincommand.Command()

	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}

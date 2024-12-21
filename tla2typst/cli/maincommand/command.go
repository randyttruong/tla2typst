package maincommand

import (
	"os"

	"github.com/randyttruong/tla2typst/cli/flags"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	c := &cobra.Command{
		SilenceUsage: true,
		Use:          os.Args[0],
	}

	flags.AddOutputFormatFlags(c)

	// c.AddCommand()

	return c
}

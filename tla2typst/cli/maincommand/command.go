package maincommand

import (
	"os"

	"github.com/randyttruong/tla2typst/cli/flags"

	"github.com/spf13/cobra"
)

var (
	Filename = ""
)

func Command() *cobra.Command {
	c := &cobra.Command{
		SilenceUsage: true,
		Use:          os.Args[0],
		Short:        "Read and process a TLA specification.",
		Args:         cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			Filename = args[0]
		},
	}

	flags.AddOutputFormatFlags(c)

	c.AddCommand()

	return c
}

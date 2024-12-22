package maincommand

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/randyttruong/tla2typst/cli/flags"
	"github.com/randyttruong/tla2typst/scanner"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			Filename = args[0]

			err := scanner.LoadDocument(Filename) // TODO: Does this belong here?

			if err != nil {
				return errors.Wrapf(err, "Unable to load document %v, got %v\n", Filename, err)
			}

			fmt.Printf("Beginning parsing\n")

			loader := scanner.GetLoader()

			err = scanner.InitScanner(loader)

			if err != nil {
				return errors.Wrapf(err, "Unable to start scanner, got %v", err)
			}

			return nil
		},
	}

	flags.AddOutputFormatFlags(c)
	flags.AddOutputFileFlags(c)
	flags.AddCommentShadingFlags(c)
	flags.AddConfigFlags(c)

	c.AddCommand()

	return c
}

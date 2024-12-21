package flags

import (
	"github.com/spf13/cobra"
)

var (
	outputFilenameDefault = "" // TODO: Make this the .tla input filename
	outputFilename        string

	alignmentOutputFilenameDefault = "" // TODO: Make this the .tla input filename
	alignmentOutputFilename        string

	tlaOutputFilenameDefault = "" // TODO: Make this the .tla input filename
	tlaOutputFilename        string

	stylePackageFilenameDefault = "" // NOTE: Make this none by default
	stylePackageFilename        string
)

var (
	outputFilenameFlagName          = "out"
	alignmentOutputFilenameFlagName = "alignOut"
	tlaOutputFilenameFlagName       = "tlaOut"
	stylePackageFilenameFlagName    = "style"
)

func AddOutputFileFlags(c *cobra.Command) {
	addOutputFilenameFlag(c)
	addAlignmentOutputFilenameFlag(c)
	addTlaOutputFilenameFlag(c)
	addStylePackageFilenameFlag(c)
}

func OutputFilename() string {
	return outputFilename
}

func AlignmentOutputFilename() string {
	return alignmentOutputFilename
}

func TlaOutputFilename() string {
	return tlaOutputFilename
}

func StylePackageFilename() string {
	return stylePackageFilename
}

// TODO: Add descriptors
func addOutputFilenameFlag(c *cobra.Command) {
	c.PersistentFlags().StringVarP(
		&outputFilename,
		outputFilenameFlagName,
		"",
		outputFilenameDefault,
		"",
	)
}

func addAlignmentOutputFilenameFlag(c *cobra.Command) {
	c.PersistentFlags().StringVarP(
		&alignmentOutputFilename,
		alignmentOutputFilenameFlagName,
		"",
		alignmentOutputFilenameDefault,
		"",
	)
}

func addTlaOutputFilenameFlag(c *cobra.Command) {
	c.PersistentFlags().StringVarP(
		&tlaOutputFilename,
		tlaOutputFilenameFlagName,
		"",
		tlaOutputFilenameDefault,
		"",
	)
}

func addStylePackageFilenameFlag(c *cobra.Command) {
	c.PersistentFlags().StringVarP(
		&stylePackageFilename,
		stylePackageFilenameFlagName,
		"",
		stylePackageFilenameDefault,
		"",
	)
}

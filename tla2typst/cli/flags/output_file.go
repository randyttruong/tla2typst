package flags

import (
	"github.com/spf13/cobra"
)

var (
	outputFilenameDefault = "" // TODO: Make this the .tla input filename
	outputFilename        string
	outputFilenameSet     = false

	alignmentOutputFilenameDefault = "" // TODO: Make this the .tla input filename
	alignmentOutputFilename        string
	alignmentOutputFilenameSet     = false

	tlaOutputFilenameDefault = "" // TODO: Make this the .tla input filename
	tlaOutputFilename        string
	tlaOutputFilenameSet     = false

	stylePackageFilenameDefault = "" // NOTE: Make this none by default
	stylePackageFilename        string
	stylePackageFilenameSet     = false
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

	outputFilenameSet = c.PersistentFlags().Lookup(outputFilenameFlagName).Changed
}

func addAlignmentOutputFilenameFlag(c *cobra.Command) {
	c.PersistentFlags().StringVarP(
		&alignmentOutputFilename,
		alignmentOutputFilenameFlagName,
		"",
		alignmentOutputFilenameDefault,
		"",
	)

	alignmentOutputFilenameSet = c.PersistentFlags().Lookup(alignmentOutputFilenameFlagName).Changed
}

func addTlaOutputFilenameFlag(c *cobra.Command) {
	c.PersistentFlags().StringVarP(
		&tlaOutputFilename,
		tlaOutputFilenameFlagName,
		"",
		tlaOutputFilenameDefault,
		"",
	)

	tlaOutputFilenameSet = c.PersistentFlags().Lookup(tlaOutputFilenameFlagName).Changed
}

func addStylePackageFilenameFlag(c *cobra.Command) {
	c.PersistentFlags().StringVarP(
		&stylePackageFilename,
		stylePackageFilenameFlagName,
		"",
		stylePackageFilenameDefault,
		"",
	)

	stylePackageFilenameSet = c.PersistentFlags().Lookup(stylePackageFilenameFlagName).Changed
}

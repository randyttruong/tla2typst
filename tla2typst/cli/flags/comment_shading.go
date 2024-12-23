package flags

import (
	"github.com/spf13/cobra"
)

var (
	useShadingDefault = false
	useShading        bool
	useShadingSet     = false

	usePcalShadingDefault = true
	usePcalShading        bool
	usePcalShadingSet     = false

	commentGrayLevelDefault float32 = 0.85
	commentGrayLevel        float32
	commentGrayLevelSet     = false

	// created with -shade by default, otherwise no
	createPostScriptDefault = false
	createPostScript        bool
	createPostScriptSet     = false
)

var (
	useShadingFlagName       = "shade"
	usePcalShadingFlagName   = "pcalShade" // TODO: Should this be changed?
	commentGrayLevelFlagName = "grayLevel"
	createPostScriptFlagName = "createPs"
)

func AddCommentShadingFlags(c *cobra.Command) {
	addUseShadingFlag(c)
	addUsePcalShadingFlag(c)
	addCommentGrayLevelFlag(c)
	addCreatePostScriptFlag(c)
}

func UseShading() bool {
	return useShading
}

func UsePcalShading() bool {
	return usePcalShading
}

func CommentGrayLevel() float32 {
	return commentGrayLevel
}

func CreatePostScript() bool {
	return createPostScript
}

func addUseShadingFlag(c *cobra.Command) {
	c.PersistentFlags().BoolVarP(
		&useShading,
		useShadingFlagName,
		"",
		useShadingDefault,
		"",
	)

	useShadingSet = c.PersistentFlags().Lookup(useShadingFlagName).Changed
}

func addUsePcalShadingFlag(c *cobra.Command) {
	c.PersistentFlags().BoolVarP(
		&usePcalShading,
		usePcalShadingFlagName,
		"",
		usePcalShadingDefault,
		"",
	)

	usePcalShadingSet = c.PersistentFlags().Lookup(usePcalShadingFlagName).Changed
}

func addCommentGrayLevelFlag(c *cobra.Command) {
	c.PersistentFlags().Float32VarP(
		&commentGrayLevel,
		commentGrayLevelFlagName,
		"",
		commentGrayLevelDefault,
		"",
	)

	commentGrayLevelSet = c.PersistentFlags().Lookup(commentGrayLevelFlagName).Changed
}

func addCreatePostScriptFlag(c *cobra.Command) {
	c.PersistentFlags().BoolVarP(
		&createPostScript,
		createPostScriptFlagName,
		"",
		createPostScriptDefault,
		"",
	)

	createPostScriptSet = c.PersistentFlags().Lookup(createPostScriptFlagName).Changed
}

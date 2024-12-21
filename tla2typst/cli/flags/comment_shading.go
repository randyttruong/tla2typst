package flags

import (
	"github.com/spf13/cobra"
)

var (
	useShadingDefault = false
	useShading        bool

	usePcalShadingDefault = true
	usePcalShading        bool

	commentGrayLevelDefault float32 = 0.85
	commentGrayLevel        float32

	// created with -shade by default, otherwise no
	createPostScriptDefault = false
	createPostScript        bool
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
}

func addUsePcalShadingFlag(c *cobra.Command) {
	c.PersistentFlags().BoolVarP(
		&usePcalShading,
		usePcalShadingFlagName,
		"",
		usePcalShadingDefault,
		"",
	)
}

func addCommentGrayLevelFlag(c *cobra.Command) {
	c.PersistentFlags().Float32VarP(
		&commentGrayLevel,
		commentGrayLevelFlagName,
		"",
		commentGrayLevelDefault,
		"",
	)
}

func addCreatePostScriptFlag(c *cobra.Command) {
	c.PersistentFlags().BoolVarP(
		&createPostScript,
		createPostScriptFlagName,
		"",
		createPostScriptDefault,
		"",
	)
}

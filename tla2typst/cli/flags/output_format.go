package flags

import (
	"github.com/spf13/cobra"
)

var (
	pointSizeDefault int8 = 10
	pointSize        int8
	pointSizeChanged = false

	textWidthDefault int32 = 255
	textWidth        int32
	textWidthChanged = false

	textHeightDefault int32 = 255
	textHeight        int32
	textHeightChanged = false

	horizontalOffsetDefault int32 = 255
	horizontalOffset        int32
	horizontalOffsetChanged = false

	verticalOffsetDefault int32 = 255
	verticalOffset        int32
	verticalOffsetChanged = false
)

var (
	pointSizeFlagName        = "ptSize"
	textWidthFlagName        = "textwidth"
	textHeightFlagName       = "textheight"
	horizontalOffsetFlagName = "hoffset"
	verticalOffsetFlagName   = "voffset"
)

func AddOutputFormatFlags(c *cobra.Command) {
	addPointSize(c)
	addTextWidth(c)
	addTextHeight(c)
	addHorizontalOffset(c)
	addVerticalOffset(c)
}

func addPointSize(c *cobra.Command) {
	c.PersistentFlags().Int8Var(
		&pointSize,
		pointSizeFlagName,
		pointSizeDefault,
		"Specifies the size of the font.  Legal values are 10, 11, or 12, which cause the specification to be typeset in a 10-, 11-, or 12-point font. The default value is 10.\n",
	)
}

func addTextWidth(c *cobra.Command) {
	c.PersistentFlags().Int32Var(
		&textWidth,
		textWidthFlagName,
		textWidthDefault,
		"Specifies the width of the typeset output, in points.  A point is 1/72 of an inch, or about 1/3 mm.",
	)
}

func addTextHeight(c *cobra.Command) {
	c.PersistentFlags().Int32Var(
		&textHeight,
		textHeightFlagName,
		textHeightDefault,
		"Specifies the height of the typeset output, in points.  A point is 1/72 of an inch, or about 1/3 mm.",
	)
}

func addHorizontalOffset(c *cobra.Command) {
	c.PersistentFlags().Int32Var(
		&horizontalOffset,
		horizontalOffsetFlagName,
		horizontalOffsetDefault,
		"Specifies distances, in points, by which the text should be moved horizontally or vertically on the page.  Exactly where on a page the text appears depends on the printer or screen-display program.  You may have to adjust this value to get the output to appear centered on the printed page, or for the entire output to be visible when viewed on the screen.",
	)
}

func addVerticalOffset(c *cobra.Command) {
	c.PersistentFlags().Int32Var(
		&verticalOffset,
		verticalOffsetFlagName,
		verticalOffsetDefault,
		"Specifies distances, in points, by which the text should be moved horizontally or vertically on the page.  Exactly where on a page the text appears depends on the printer or screen-display program.  You may have to adjust this value to get the output to appear centered on the printed page, or for the entire output to be visible when viewed on the screen.",
	)
}

func PointSize() int8 {
	return pointSize
}

func TextWidth() int32 {
	return textWidth
}

func TextHeight() int32 {
	return textHeight
}

func HorizontalOffset() int32 {
	return horizontalOffset
}

func VerticalOffset() int32 {
	return verticalOffset
}
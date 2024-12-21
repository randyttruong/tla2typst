package flags

import (
	"github.com/spf13/cobra"
)

type TlaConfig struct {
	Name string `json:"name"` // identifier

	// Output Format Options
	PointSize        int8 `json:"pointSize,omitempty"`
	TextWidth        int8 `json:"textWidth,omitempty"`
	TextHeight       int8 `json:"textHeight,omitempty"`
	HorizontalOffset int8 `json:"horizontalOffset,omitempty"`
	VerticalOffset   int8 `json:"verticalOffset,omitempty"`

	// Comment Shading Options
	UseShading       bool    `json:"shading,omitempty"`
	UsePcalShading   bool    `json:"pcalShading,omitempty"`
	CommentGrayLevel float32 `json:"commentGrayLevel,omitempty"`
	CreatePostScript float32 `json:"postscript,omitempty"`

	// Output File Options
	OutputFilename          string `json:"output,omitempty"`
	AlignmentOutputFilename string `json:"alignmentOutput,omitempty"`
	TlaOutputFilename       string `json:"tlaOutput,omitempty"`
	StylePackageFilename    string `json:"styling,omitempty"`
}

type Config struct {
	TlaConfigs map[string]*TlaConfig
}

var (
	configFilenameDefault = ""
	configFilename        string
	configSet             bool = false

	config *Config
)

var (
	ConfigFlagName = "config"
)

func AddConfigFlags(c *cobra.Command) {
	addConfig(c)
}

func addConfig(c *cobra.Command) {
	c.PersistentFlags().StringVarP(
		&configFilename,
		"config",
		"",
		configFilenameDefault,
		"",
	)
}

// TODO:
func readConfig() {}

// TODO:
func LoadConfig() {}

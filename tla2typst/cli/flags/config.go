package flags

import (
	"os"

	"github.com/randyttruong/tla2typst/pkg/util"
	"github.com/spf13/cobra"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

type TlaConfig struct {
	Name string `json:"name"` // identifier

	// Output Format Options
	PointSize        int8  `json:"pointSize,omitempty"`
	TextWidth        int32 `json:"textWidth,omitempty"`
	TextHeight       int32 `json:"textHeight,omitempty"`
	HorizontalOffset int32 `json:"horizontalOffset,omitempty"`
	VerticalOffset   int32 `json:"verticalOffset,omitempty"`

	// Comment Shading Options
	UseShading       bool    `json:"shading,omitempty"`
	UsePcalShading   bool    `json:"pcalShading,omitempty"`
	CommentGrayLevel float32 `json:"commentGrayLevel,omitempty"`
	CreatePostScript bool    `json:"postscript,omitempty"`

	// Output File Options
	OutputFilename          string `json:"output,omitempty"`
	AlignmentOutputFilename string `json:"alignmentOutput,omitempty"`
	TlaOutputFilename       string `json:"tlaOutput,omitempty"`
	StylePackageFilename    string `json:"styling,omitempty"`
}

type Config struct {
	TlaConfigs TlaConfig
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

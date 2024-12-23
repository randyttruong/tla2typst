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
	configFilenameSet     bool = false

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

func loadFlags() {

	if pointSizeSet {
		config.TlaConfigs.PointSize = PointSize()
	}
	if textWidthSet {
		config.TlaConfigs.TextWidth = TextWidth()
	}
	if textHeightSet {
		config.TlaConfigs.TextHeight = TextHeight()
	}
	if horizontalOffsetSet {
		config.TlaConfigs.HorizontalOffset = HorizontalOffset()
	}
	if verticalOffsetSet {
		config.TlaConfigs.VerticalOffset = VerticalOffset()
	}

	if useShadingSet {
		config.TlaConfigs.UseShading = UseShading()
	}
	if usePcalShadingSet {
		config.TlaConfigs.UsePcalShading = UsePcalShading()
	}
	if commentGrayLevelSet {
		config.TlaConfigs.CommentGrayLevel = CommentGrayLevel()
	}
	if createPostScriptSet {
		config.TlaConfigs.CreatePostScript = CreatePostScript()
	}

	if outputFilenameSet {
		config.TlaConfigs.OutputFilename = OutputFilename()
	}
	if alignmentOutputFilenameSet {
		config.TlaConfigs.AlignmentOutputFilename = AlignmentOutputFilename()
	}
	if tlaOutputFilenameSet {
		config.TlaConfigs.TlaOutputFilename = TlaOutputFilename()
	}
	if stylePackageFilenameSet {
		config.TlaConfigs.StylePackageFilename = StylePackageFilename()
	}
}

func readConfig(filepath string) error {

	err := util.CheckFilePermissionsAndOwnership(filepath)

	if err != nil {
		errors.Wrapf(err, "Failed to load config, instead got: %v", err)
	}

	bytes, err := os.ReadFile(filepath)

	if err != nil {
		if os.IsNotExist(err) {
			return errors.Wrapf(err, "File at %v does not exist.", filepath)
		}

		return errors.Wrapf(err, "Failed to read file, got %v", err)
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return errors.Wrapf(err, "failed to retrieve config from file %q", filepath)
	}

	return nil
}

func LoadConfig(cmd *cobra.Command, args []string) error {

	if configFilename == "" || !configFilenameSet {
		return nil
	}

	err := readConfig(configFilename)

	if err != nil {
		err = errors.Errorf("Error reading instance config file: %v", err)
		return err
	}

	// populate config with flags to overwrite config options
	loadFlags()

	return nil
}

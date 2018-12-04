package commands

import (
	"log"

	"github.com/starkandwayne/pivy/om"
	"github.com/starkandwayne/pivy/pivnet"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type ConfigTemplateCommand struct {
	global       *global
	logger       *log.Logger
	tile         pivnet.TileFile
	placeholders bool
}

const configTemplateName = "generate-config-template"

func (cmd *ConfigTemplateCommand) register(app *kingpin.Application) {
	c := app.Command(configTemplateName, "Generate om cli compatible config template").Action(cmd.run)
	registerTileFlags(c, &cmd.tile)
	app.Flag(
		"include-placeholders",
		"replace obscured credentials with interpolatable placeholders",
	).Default("false").BoolVar(&cmd.placeholders)
}

func (cmd *ConfigTemplateCommand) run(c *kingpin.ParseContext) error {
	piv, err := cmd.global.NewPivnetSdk(cmd.logger)
	if err != nil {
		return err
	}

	link, err := piv.TileDownloadLink(cmd.tile)
	if err != nil {
		return err
	}

	metaData, err := piv.DownloadTileMetaDataFromLink(link)
	if err != nil {
		return err
	}

	configTemplate, err := om.TileMetaDataToConfigTemplate(metaData, cmd.placeholders)
	if err != nil {
		return err
	}

	cmd.logger.Println(string(configTemplate))

	return nil
}

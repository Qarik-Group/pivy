package commands

import (
	"log"

	"github.com/pivotalservices/tile-config-generator/generator"
	"github.com/starkandwayne/pivy/pivnet"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type TileConfigCommand struct {
	global                     *global
	logger                     *log.Logger
	tile                       pivnet.TileFile
	baseDirectory              string
	doNotIncludeProductVersion bool
	includeErrands             bool
}

const tileConfigName = "generate-tile-config"

func (cmd *TileConfigCommand) register(app *kingpin.Application) {
	c := app.Command(tileConfigName, "Generate tile-config-generator compatible tile config directory").Action(cmd.run)
	registerTileFlags(c, &cmd.tile)
	app.Flag(
		"base-directory",
		"base directory to place generated config templates",
	).Required().StringVar(&cmd.baseDirectory)
	app.Flag(
		"do-not-include-product-version",
		"flag to use a flat output folder",
	).Default("false").BoolVar(&cmd.doNotIncludeProductVersion)
	app.Flag(
		"include-errands",
		"feature flag to include errands",
	).Default("false").BoolVar(&cmd.includeErrands)
}

func (cmd *TileConfigCommand) run(c *kingpin.ParseContext) error {
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

	return generator.NewExecutor(metaData, cmd.baseDirectory, cmd.doNotIncludeProductVersion, cmd.includeErrands).Generate()
}

package commands

import (
	"log"

	"github.com/starkandwayne/pivy/pivnet"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type DownloadCommand struct {
	global *global
	logger *log.Logger
	tile   pivnet.TileFile
}

const downloadName = "download-product-template"

func (cmd *DownloadCommand) register(app *kingpin.Application) {
	c := app.Command(downloadName, "Download raw tile metadata").Action(cmd.run)
	registerTileFlags(c, &cmd.tile)
}

func (cmd *DownloadCommand) run(c *kingpin.ParseContext) error {
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

	cmd.logger.Println(string(metaData))

	return nil
}

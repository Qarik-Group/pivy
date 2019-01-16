package commands

import (
	"fmt"
	"log"

	"github.com/starkandwayne/pivy/metadata"
	"github.com/starkandwayne/pivy/pivnet"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type GolangTemplateCommand struct {
	global *global
	logger *log.Logger
	tile   pivnet.TileFile
}

const golangTemplateName = "generate-golang-template"

func (cmd *GolangTemplateCommand) register(app *kingpin.Application) {
	c := app.Command(golangTemplateName, "Generate golang template").Action(cmd.run)
	registerTileFlags(c, &cmd.tile)
}

func (cmd *GolangTemplateCommand) run(c *kingpin.ParseContext) error {
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

	parser, err := metadata.NewParser(metaData)
	if err != nil {
		return fmt.Errorf("creating parser failed: %s", err)
	}
	f, err := parser.RenderToGolang()
	if err != nil {
		return fmt.Errorf("redendering golang failed: %s", err)
	}
	fmt.Printf("%#v", f)
	return nil
}

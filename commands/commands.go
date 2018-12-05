package commands

import (
	"log"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/starkandwayne/pivy/pivnet"
)

type register interface {
	register(app *kingpin.Application)
}

type global struct {
	acceptEULA  bool
	pivnetToken string
}

// Configure sets up the kingpin commands for the omg-cli.
func Configure(logger *log.Logger, app *kingpin.Application) {
	var global global
	app.Flag(
		"accept-eula",
		"Automatically accept EULA if necessary (Available to select users only)",
	).Default("false").OverrideDefaultFromEnvar("PIVNET_ACCEPT_EULA").BoolVar(&global.acceptEULA)

	app.Flag(
		"pivnet-api-token",
		"API token for network.pivotal.io (see: https://network.pivotal.io/users/dashboard/edit-profile)",
	).Required().OverrideDefaultFromEnvar("PIVNET_API_TOKEN").Short('t').StringVar(&global.pivnetToken)

	cmds := []register{
		&DownloadCommand{logger: logger, global: &global},
		&ConfigTemplateCommand{logger: logger, global: &global},
		&GolangTemplateCommand{logger: logger, global: &global},
	}

	for _, c := range cmds {
		c.register(app)
	}
}

func registerTileFlags(c *kingpin.CmdClause, tile *pivnet.TileFile) {
	c.Flag("product-slug", "Product slug e.g. p-mysql").Required().Short('p').StringVar(&tile.ProductSlug)
	c.Flag("release-version", "Release version e.g. 0.1.2-rc1").Required().Short('r').StringVar(&tile.ReleaseVersion)

	c.Flag("glob",
		"Glob to match product name e.g. *aws* should include *.pivotal",
	).Short('g').Default("*.pivotal").StringVar(&tile.Glob)
}

func (g global) NewPivnetSdk(logger *log.Logger) (*pivnet.Sdk, error) {
	return pivnet.NewSdk(g.pivnetToken, g.acceptEULA, logger)
}

package om

import (
	"bytes"
	"log"

	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/commands/fakes"
	"github.com/pivotal-cf/om/extractor"
)

type TileMetaData []byte
type TileConfigTemplate []byte

func TileMetaDataToConfigTemplate(data TileMetaData, placeholders bool) (TileConfigTemplate, error) {
	metadataExtractor := &fakes.MetadataExtractor{}
	b := &bytes.Buffer{}
	logger := log.New(b, "", 0)
	command := commands.NewConfigTemplate(metadataExtractor, logger)
	metadataExtractor.ExtractMetadataReturns(extractor.Metadata{Raw: data}, nil)
	args := []string{
		"--product", "/not/used/but/required.pivotal",
	}
	if placeholders {
		args = append(args, "--include-placeholders")
	}
	err := command.Execute(args)
	if err != nil {
		return []byte{}, err
	}
	return b.Bytes(), nil
}

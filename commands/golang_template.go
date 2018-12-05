package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/pivotal-cf/kiln/proofing"
	"github.com/starkandwayne/pivy/pivnet"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	yaml "gopkg.in/yaml.v2"
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

	var template proofing.ProductTemplate
	err = yaml.Unmarshal(metaData, &template)
	if err != nil {
		return fmt.Errorf("could not parse metadata: %s", err)
	}

	var fields []Code
	for _, property := range template.AllPropertyBlueprints() {
		if !property.Configurable {
			continue
		}
		tag := jsonTag(property.Property, !property.Required || property.Default != nil)
		label, _ := lookUpLabel(template, property.Property)
		if label != "" {
			fields = append(fields, Comment(label))
		}
		if property.Default != nil {
			d, _ := json.Marshal(property.Default)
			fields = append(fields, Commentf("default: %s", string(d)))
		}

		fields = append(fields, Id(propertyToId(property)).
			Struct(propertyToValueStruct(property)).Tag(tag), Line())

	}

	f := NewFile("tiles")
	f.Type().Id("properties").Struct(fields...)

	fmt.Printf("%#v", f)
	return nil
}

func propertyToId(property proofing.NormalizedPropertyBlueprint) string {
	parts := strings.Split(property.Property, ".")
	return strcase.ToCamel(parts[len(parts)-1])
}

func propertyToValueStruct(property proofing.NormalizedPropertyBlueprint) Code {
	switch property.Type {
	case "boolean":
		return Id("Value").Bool().Tag(jsonTag("value", false))
	case "integer", "port":
		return Id("Value").Int().Tag(jsonTag("value", false))
	case "secret":
		return Id("Value").Struct(
			Id("Secret").String().Tag(jsonTag("secret", false)),
		).Tag(jsonTag("value", false))
	case "simple_credentials":
		return Id("Value").Struct(
			Id("Identity").String().Tag(jsonTag("identity", false)),
			Id("Password").String().Tag(jsonTag("password", false)),
		).Tag(jsonTag("value", false))
	case "rsa_cert_credentials":
		return Id("Value").Struct(
			Id("CertPem").String().Tag(jsonTag("cert_pem", false)),
			Id("PrivateKeyPem").String().Tag(jsonTag("private_key_pem", false)),
		).Tag(jsonTag("value", false))
	case "rsa_pkey_credentials":
		return Id("Value").Struct(
			Id("PublicKeyPem").String().Tag(jsonTag("public_key_pem", false)),
			Id("PrivateKeyPem").String().Tag(jsonTag("private_key_pem", false)),
		).Tag(jsonTag("value", false))
	case "salted_credentials":
		return Id("Value").Struct(
			Id("Identity").String().Tag(jsonTag("identity", false)),
			Id("Password").String().Tag(jsonTag("password", false)),
			Id("Salt").String().Tag(jsonTag("salt", false)),
		).Tag(jsonTag("value", false))
	default:
		return Id("Value").String().Tag(jsonTag("value", false))
	}

}

func jsonTag(str string, omitEmpty bool) map[string]string {
	if omitEmpty {
		return map[string]string{"json": str + ",omitempty"}
	} else {
		return map[string]string{"json": str}
	}
}

func lookUpLabel(template proofing.ProductTemplate, reference string) (string, bool) {
	for _, form := range template.FormTypes {
		for _, prop := range form.PropertyInputs {
			if prop, ok := prop.(proofing.SimplePropertyInput); ok {
				if prop.Reference == reference {
					return prop.Label, true
				}
			}
			if prop, ok := prop.(proofing.CollectionPropertyInput); ok {
				if prop.Reference == reference {
					return prop.Label, true
				}
			}
			if prop, ok := prop.(proofing.CollectionSubfieldPropertyInput); ok {
				if prop.Reference == reference {
					return prop.Label, true
				}
			}
			if prop, ok := prop.(proofing.SelectorPropertyInput); ok {
				if prop.Reference == reference {
					return prop.Label, true
				}
			}
			if prop, ok := prop.(proofing.SelectorOptionPropertyInput); ok {
				if prop.Reference == reference {
					return prop.Label, true
				}
			}
		}
	}
	return "", false
}

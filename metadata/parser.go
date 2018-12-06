package metadata

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/ghodss/yaml"
	"github.com/pivotal-cf/kiln/proofing"
)

type metaDataParser struct {
	parsedTemplate proofing.ProductTemplate
	rawTemplate    *gabs.Container
}

func NewParser(metaData []byte) (*metaDataParser, error) {
	var p metaDataParser
	var err error
	r := bytes.NewReader(metaData)
	p.parsedTemplate, err = proofing.Parse(r)
	if err != nil {
		return &metaDataParser{},
			fmt.Errorf("could not parse metadata: %s", err)
	}

	templateJSON, err := yaml.YAMLToJSON(metaData)
	if err != nil {
		return &metaDataParser{},
			fmt.Errorf("could not convert metadata to json: %s", err)
	}

	p.rawTemplate, err = gabs.ParseJSON(templateJSON)
	if err != nil {
		return &metaDataParser{},
			fmt.Errorf("second parse metadata failed: %s", err)
	}

	return &p, err

}

func (p *metaDataParser) AllPropertyBlueprints() []proofing.NormalizedPropertyBlueprint {
	return p.parsedTemplate.AllPropertyBlueprints()
}

func (p *metaDataParser) GetPropertyLabel(property proofing.NormalizedPropertyBlueprint) (string, bool) {
	return p.lookupPropertyFieldString(property, "label")
}

func (p *metaDataParser) GetPropertyDescription(property proofing.NormalizedPropertyBlueprint) (string, bool) {
	return p.lookupPropertyFieldString(property, "description")
}

func (p *metaDataParser) lookupPropertyFieldString(property proofing.NormalizedPropertyBlueprint, field string) (string, bool) {
	rp, ok := p.lookupRawProperty(property)
	if !ok {
		return "", false
	}
	if !rp.Exists(field) {
		return "", false
	}
	return rp.Path(field).Data().(string), true
}

func (p *metaDataParser) lookupRawProperty(property proofing.NormalizedPropertyBlueprint) (*gabs.Container, bool) {
	parts := strings.Split(property.Property, ".")
	name := parts[len(parts)-1]
	properties, _ := p.rawTemplate.S("property_blueprints").Children()
	for _, property := range properties {
		if property.Path("name").Data().(string) == name {
			return property, true
		}
	}
	return &gabs.Container{}, false
}

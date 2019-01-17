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
	collection     bool
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
	p.collection = false
	return &p, err

}

func (p *metaDataParser) AllPropertyBlueprints() []proofing.NormalizedPropertyBlueprint {
	if p.collection {
		var propertyBlueprints []proofing.NormalizedPropertyBlueprint

		propertyBlueprints = make([]proofing.NormalizedPropertyBlueprint, 0, len(p.parsedTemplate.PropertyBlueprints))

		for _, pb := range p.parsedTemplate.PropertyBlueprints {
			propertyBlueprints = append(propertyBlueprints, pb.Normalize("")...)
		}
		return propertyBlueprints
	}
	return p.parsedTemplate.AllPropertyBlueprints()
}

func (p *metaDataParser) GetCollectionParser(property proofing.NormalizedPropertyBlueprint) *metaDataParser {
	pb, ok := p.lookupRawPropertyField(property, "property_blueprints")
	if !ok {
		return &metaDataParser{}
	}
	o := gabs.New()
	o.Array("property_blueprints")
	children, err := pb.Children()
	if err != nil {
		return &metaDataParser{}
	}
	for _, child := range children {
		o.ArrayAppend(child.Data(), "property_blueprints")
	}
	parser, err := NewParser(o.Bytes())
	if err != nil {
		return &metaDataParser{}
	}
	parser.collection = true
	return parser
}

func (p *metaDataParser) GetPropertyLabel(property proofing.NormalizedPropertyBlueprint) (string, bool) {
	return p.lookupPropertyFieldString(property, "label")
}

func (p *metaDataParser) GetPropertyDescription(property proofing.NormalizedPropertyBlueprint) (string, bool) {
	return p.lookupPropertyFieldString(property, "description")
}

func (p *metaDataParser) lookupRawPropertyField(property proofing.NormalizedPropertyBlueprint, field string) (*gabs.Container, bool) {
	rp, ok := p.lookupRawProperty(property)
	if !ok {
		return &gabs.Container{}, false
	}
	if !rp.Exists(field) {
		return &gabs.Container{}, false
	}
	return rp.Path(field), true
}

func (p *metaDataParser) lookupPropertyFieldString(property proofing.NormalizedPropertyBlueprint, field string) (string, bool) {
	rp, ok := p.lookupRawPropertyField(property, field)
	if !ok {
		return "", false
	}
	return rp.Data().(string), true
}

func (p *metaDataParser) lookupRawProperty(property proofing.NormalizedPropertyBlueprint) (*gabs.Container, bool) {
	parts := strings.Split(property.Property, ".")
	name := parts[len(parts)-1]
	properties, _ := p.rawTemplate.S("property_blueprints").Children()
	for _, property := range properties {
		n, ok := property.Path("name").Data().(string)
		if ok && n == name {
			return property, true
		}
	}
	return &gabs.Container{}, false
}

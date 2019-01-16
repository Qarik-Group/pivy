package metadata

import (
	"encoding/json"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/pivotal-cf/kiln/proofing"
)

type structRenderer struct {
	fields *[]Code
	parser *metaDataParser
}

func newStructRenderer(p *metaDataParser) *structRenderer {
	var fields []Code
	r := structRenderer{
		fields: &fields,
		parser: p,
	}
	return &r
}

func (p *metaDataParser) RenderToGolang() (string, error) {
	r := newStructRenderer(p)
	r.addFieldsForProperties(p.AllPropertyBlueprints())
	f := NewFile("tiles")
	f.Type().Id("properties").Struct(*r.fields...)
	return f.GoString(), nil
}

func propertyToId(property proofing.NormalizedPropertyBlueprint) string {
	parts := strings.Split(property.Property, ".")
	return strcase.ToCamel(parts[len(parts)-1])
}

func (r *structRenderer) addFieldsForProperties(properties []proofing.NormalizedPropertyBlueprint) {
	for _, property := range properties {
		if !property.Configurable {
			continue
		}
		label, ok := r.parser.GetPropertyLabel(property)
		if ok {
			*r.fields = append(*r.fields, Comment(label))
		}
		description, ok := r.parser.GetPropertyDescription(property)
		if ok {
			*r.fields = append(*r.fields, Comment(description))
		}
		if property.Default != nil {
			d, _ := json.Marshal(property.Default)
			*r.fields = append(*r.fields, Commentf("default: %s", string(d)))
		}
		tag := jsonTag(property.Property, !property.Required || property.Default != nil)

		switch property.Type {
		case "collection":
			cp := r.parser.GetCollectionParser(property)
			cr := newStructRenderer(cp)
			cr.addFieldsForProperties(cp.AllPropertyBlueprints())
			*r.fields = append(*r.fields, Id(propertyToId(property)).
				Index().Struct(*cr.fields...).Tag(tag), Line())
		default:
			*r.fields = append(*r.fields, Id(propertyToId(property)).
				Struct(propertyToValueStruct(property)).Tag(tag), Line())
		}
	}
}

func propertyToValueStruct(property proofing.NormalizedPropertyBlueprint) Code {
	return Id("Value").Add(propertyToStruct(property)).Tag(jsonTag("value", false))
}

func propertyToStruct(property proofing.NormalizedPropertyBlueprint) Code {
	switch property.Type {
	case "boolean":
		return Bool()
	case "integer", "port":
		return Int()
	case "secret":
		return Struct(Id("Secret").String().Tag(jsonTag("secret", false)))
	case "simple_credentials":
		return Struct(
			Id("Identity").String().Tag(jsonTag("identity", false)),
			Id("Password").String().Tag(jsonTag("password", false)),
		)
	case "rsa_cert_credentials":
		return Struct(
			Id("CertPem").String().Tag(jsonTag("cert_pem", false)),
			Id("PrivateKeyPem").String().Tag(jsonTag("private_key_pem", false)),
		)
	case "rsa_pkey_credentials":
		return Struct(
			Id("PublicKeyPem").String().Tag(jsonTag("public_key_pem", false)),
			Id("PrivateKeyPem").String().Tag(jsonTag("private_key_pem", false)),
		)
	case "salted_credentials":
		return Struct(
			Id("Identity").String().Tag(jsonTag("identity", false)),
			Id("Password").String().Tag(jsonTag("password", false)),
			Id("Salt").String().Tag(jsonTag("salt", false)),
		)
	default:
		return String()
	}

}

func jsonTag(str string, omitEmpty bool) map[string]string {
	if omitEmpty {
		return map[string]string{"json": str + ",omitempty"}
	} else {
		return map[string]string{"json": str}
	}
}

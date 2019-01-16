package metadata

import (
	"encoding/json"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/pivotal-cf/kiln/proofing"
)

func (p *metaDataParser) RenderToGolang() (*File, error) {

	var fields []Code
	for _, property := range p.AllPropertyBlueprints() {
		if !property.Configurable {
			continue
		}
		tag := jsonTag(property.Property, !property.Required || property.Default != nil)
		label, ok := p.GetPropertyLabel(property)
		if ok {
			fields = append(fields, Comment(label))
		}
		description, ok := p.GetPropertyDescription(property)
		if ok {
			fields = append(fields, Comment(description))
		}
		if property.Default != nil {
			d, _ := json.Marshal(property.Default)
			fields = append(fields, Commentf("default: %s", string(d)))
		}

		switch property.Type {
		case "collection":
			fields = append(fields, Id(propertyToId(property)).
				Struct(collectionPropertyToValueStruct(p.GetCollectionProperties(property))).Tag(tag), Line())
		default:
			fields = append(fields, Id(propertyToId(property)).
				Struct(propertyToValueStruct(property)).Tag(tag), Line())
		}

	}

	f := NewFile("tiles")
	f.Type().Id("properties").Struct(fields...)

	return f, nil
}

func propertyToId(property proofing.NormalizedPropertyBlueprint) string {
	parts := strings.Split(property.Property, ".")
	return strcase.ToCamel(parts[len(parts)-1])
}

func collectionPropertyToValueStruct(properties []proofing.NormalizedPropertyBlueprint) Code {
	var props []Code
	for _, prop := range properties {
		props = append(props, Id(propertyToId(prop)).Add(propertyToStruct(prop)))
	}
	return Id("Value").Struct(props...).Tag(jsonTag("value", false))
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

// func propertyToStruct(property proofing.NormalizedPropertyBlueprint) func(*Statement) {
// 	return func(s *Statement) {
// 		switch property.Type {
// 		case "boolean":
// 			s.Bool()
// 		case "integer", "port":
// 			s.Int()
// 		case "secret":
// 			s.Struct(Id("Secret").String().Tag(jsonTag("secret", false)))
// 		case "simple_credentials":
// 			s.Struct(
// 				Id("Identity").String().Tag(jsonTag("identity", false)),
// 				Id("Password").String().Tag(jsonTag("password", false)),
// 			)
// 		case "rsa_cert_credentials":
// 			s.Struct(
// 				Id("CertPem").String().Tag(jsonTag("cert_pem", false)),
// 				Id("PrivateKeyPem").String().Tag(jsonTag("private_key_pem", false)),
// 			)
// 		case "rsa_pkey_credentials":
// 			s.Struct(
// 				Id("PublicKeyPem").String().Tag(jsonTag("public_key_pem", false)),
// 				Id("PrivateKeyPem").String().Tag(jsonTag("private_key_pem", false)),
// 			)
// 		case "salted_credentials":
// 			s.Struct(
// 				Id("Identity").String().Tag(jsonTag("identity", false)),
// 				Id("Password").String().Tag(jsonTag("password", false)),
// 				Id("Salt").String().Tag(jsonTag("salt", false)),
// 			)
// 		default:
// 			s.String()
// 		}
// 	}
// }

func jsonTag(str string, omitEmpty bool) map[string]string {
	if omitEmpty {
		return map[string]string{"json": str + ",omitempty"}
	} else {
		return map[string]string{"json": str}
	}
}

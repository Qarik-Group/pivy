package metadata

import (
	"encoding/json"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/pivotal-cf/kiln/proofing"
)

type structRenderer struct {
	parser *metaDataParser
}

func newStructRenderer(p *metaDataParser) *structRenderer {
	r := structRenderer{parser: p}
	return &r
}

func (p *metaDataParser) RenderToGolang() (string, error) {
	r := newStructRenderer(p)
	label := strcase.ToCamel(p.parsedTemplate.Label)
	name := p.parsedTemplate.Name
	f := NewFile("tiles")
	f.Type().Id(label).Struct(
		Id("ProductName").String().Tag(jsonTag("product-name", false)),
		Id("ProductProperties").Struct(
			r.fieldsForProperties(p.AllPropertyBlueprints())...,
		).Tag(jsonTag("product-properties", false)),
		Id("ResourceConfig").Struct(
			r.fieldsForResources(p.parsedTemplate.JobTypes)...,
		).Tag(jsonTag("resource-config", false)),
		Id("NetworkProperties").Struct(
			r.fieldsForNetworks(p.parsedTemplate.ServiceBroker)...,
		).Tag(jsonTag("network-properties", false)),
	)
	f.Func().Params(Id("pc").Op("*").Id(label)).Id("ToJson").Params().
		Params(Index().Byte(), Id("error")).Block(
		Id("pc").Dot("ProductName").Op("=").Lit(name),
		Return().Qual("encoding/json", "Marshal").Call(Id("pc")),
	)
	return f.GoString(), nil
}

func propertyToName(property proofing.NormalizedPropertyBlueprint) string {
	parts := strings.Split(property.Property, ".")
	return parts[len(parts)-1]
}

func propertyToId(property proofing.NormalizedPropertyBlueprint) string {
	return strcase.ToCamel(propertyToName(property))
}

func (r *structRenderer) fieldsForProperties(properties []proofing.NormalizedPropertyBlueprint) []Code {
	var fields []Code
	for _, property := range properties {
		if !property.Configurable {
			continue
		}
		label, ok := r.parser.GetPropertyLabel(property)
		if ok {
			fields = append(fields, Comment(label))
		}
		description, ok := r.parser.GetPropertyDescription(property)
		if ok {
			fields = append(fields, Comment(description))
		}

		var dc Code
		if property.Default != nil {
			d, _ := json.Marshal(property.Default)
			dc = Commentf("default: %s", string(d))
		}

		tagName := property.Property
		if r.parser.collection {
			tagName = propertyToName(property)
		}

		omitEmpty := !property.Required || property.Default != nil
		tag := jsonTag(tagName, omitEmpty)
		pointer := Null()
		if omitEmpty {
			pointer = Op("*")
		}
		prop := Id(propertyToId(property)).Add(pointer)

		if property.Type == "collection" {
			cp := r.parser.GetCollectionParser(property)
			cr := newStructRenderer(cp)
			field := prop.Struct(
				Id("Value").Index().Struct(
					cr.fieldsForProperties(
						cp.AllPropertyBlueprints())...).
					Tag(jsonTag("value", false)).Add(dc),
			).Tag(tag)
			fields = append(fields, field)
			continue
		}
		if r.parser.collection {
			fields = append(fields, prop.Add(propertyToStruct(property)).Tag(tag).Add(dc))
			continue
		}
		fields = append(fields, prop.Struct(propertyToValueStruct(property)).Tag(tag), Line())

	}
	return fields
}

func (r *structRenderer) fieldsForResources(jobs []proofing.JobType) []Code {
	var fields []Code
	for _, job := range jobs {
		instance := job.InstanceDefinition
		if !instance.Configurable {
			continue
		}
		omitEmpty := (instance.ZeroIf != proofing.ZeroIfBinding{} || instance.Default != 0)
		tag := jsonTag(job.Name, omitEmpty)
		fields = append(fields, Id(strcase.ToCamel(job.Name)).
			Struct(
				Id("Instances").String().Tag(jsonTag("instances", omitEmpty)).
					Commentf("default: %v", instance.Default),
				Id("InstanceType").Struct(
					Id("ID").String().Tag(jsonTag("id", false)),
				).Tag(jsonTag("instance_type", omitEmpty)),
				Id("PersistentDisk").Struct(
					Id("SizeMB").String().Tag(jsonTag("size_mb", false)),
				).Tag(jsonTag("persistent_disk", omitEmpty)),
				Id("InternetConnected").Bool().Tag(jsonTag("internet_connected", omitEmpty)),
				Id("ELBNames").Index().String().Tag(jsonTag("elb_names", omitEmpty)),
			).Tag(tag), Line())

	}
	return fields
}

func (r *structRenderer) fieldsForNetworks(serviceBroker bool) []Code {
	fields := []Code{
		Id("SingletonAvailabilityZone").Struct(
			Id("Name").String().Tag(jsonTag("name", false)),
		).Tag(jsonTag("singleton_availability_zone", false)),
		Id("OtherAvailabilityZones").Index().Struct(
			Id("Name").String().Tag(jsonTag("name", false)),
		).Tag(jsonTag("other_availability_zones", false)),
		Id("Network").Struct(
			Id("Name").String().Tag(jsonTag("name", false)),
		).Tag(jsonTag("network", false)),
	}
	if serviceBroker {
		fields = append(fields, Id("ServiceNetwork").Struct(
			Id("Name").String().Tag(jsonTag("name", false)),
		).Tag(jsonTag("service_network", false)))
	}
	return fields
}

func propertyToValueStruct(property proofing.NormalizedPropertyBlueprint) Code {
	var dc Code
	if property.Default != nil {
		d, _ := json.Marshal(property.Default)
		dc = Commentf("default: %s", string(d))
	}
	return Id("Value").Add(propertyToStruct(property)).
		Tag(jsonTag("value", false)).Add(dc)
}

func propertyToStruct(property proofing.NormalizedPropertyBlueprint) Code {
	switch property.Type {
	case "boolean":
		return Bool()
	case "integer", "port":
		return Int()
	case "secret":
		return Struct(
			Id("Secret").String().Tag(jsonTag("secret", false)),
		)
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

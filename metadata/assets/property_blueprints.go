package tiles

import "encoding/json"

type FooBar struct {
	ProductName       string `json:"product-name"`
	ProductProperties struct {
		Org *struct {
			Value string `json:"value"` // default: "system"
		} `json:".properties.org,omitempty"`

		Users *struct {
			Value []struct {
				Name        string  `json:"name"`
				DisplayName *string `json:"display_name,omitempty"` // default: "anonymous"
				Password    struct {
					Secret string `json:"secret"`
				} `json:"password"`
			} `json:"value"`
		} `json:".properties.users,omitempty"`
		Db *struct {
			Value string `json:"value"` // default: "internal"
		} `json:".properties.db,omitempty"`

		DbExternalHostname struct {
			Value string `json:"value"`
		} `json:".properties.db.external.hostname"`
	} `json:"product-properties"`
	ResourceConfig    struct{} `json:"resource-config"`
	NetworkProperties struct {
		SingletonAvailabilityZone struct {
			Name string `json:"name"`
		} `json:"singleton_availability_zone"`
		OtherAvailabilityZones []struct {
			Name string `json:"name"`
		} `json:"other_availability_zones"`
		Network struct {
			Name string `json:"name"`
		} `json:"network"`
	} `json:"network-properties"`
}

func (pc *FooBar) ToJson() ([]byte, error) {
	pc.ProductName = "foo-bar"
	return json.Marshal(pc)
}

package tiles

import "encoding/json"

type ExampleServiceTile struct {
	ProductName       string   `json:"product-name"`
	ProductProperties struct{} `json:"product-properties"`
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
		ServiceNetwork struct {
			Name string `json:"name"`
		} `json:"service_network"`
	} `json:"network-properties"`
}

func (pc *ExampleServiceTile) ToJson() ([]byte, error) {
	pc.ProductName = "example"
	return json.Marshal(pc)
}

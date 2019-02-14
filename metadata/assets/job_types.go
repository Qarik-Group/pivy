package tiles

import "encoding/json"

type ExampleTile struct {
	ProductName       string   `json:"product-name"`
	ProductProperties struct{} `json:"product-properties"`
	ResourceConfig    struct {
		Mysql struct {
			Instances    string `json:"instances,omitempty"` // default: 1
			InstanceType struct {
				ID string `json:"id"`
			} `json:"instance_type,omitempty"`
			PersistentDisk struct {
				SizeMB string `json:"size_mb"`
			} `json:"persistent_disk,omitempty"`
			InternetConnected bool     `json:"internet_connected,omitempty"`
			ELBNames          []string `json:"elb_names,omitempty"`
		} `json:"mysql,omitempty"`
	} `json:"resource-config"`
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

func (pc *ExampleTile) ToJson() ([]byte, error) {
	pc.ProductName = "example"
	return json.Marshal(pc)
}

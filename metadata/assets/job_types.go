package example

type ProductConfig struct {
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
}

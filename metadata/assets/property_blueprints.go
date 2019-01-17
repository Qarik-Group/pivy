package foo_bar

type ProductConfig struct {
	ProductProperties struct {
		Org struct {
			Value string `json:"value"` // default: "system"
		} `json:".properties.org,omitempty"`

		Users struct {
			Value []struct {
				Name        string `json:"name"`
				DisplayName string `json:"display_name,omitempty"` // default: "anonymous"
				Password    struct {
					Secret string `json:"secret"`
				} `json:"password"`
			} `json:"value"`
		} `json:".properties.users,omitempty"`
	} `json:"product-properties"`
	ResourceConfig struct{} `json:"resource-config"`
}

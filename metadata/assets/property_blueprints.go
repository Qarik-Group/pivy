package tiles

type properties struct {
	// org label
	// org description
	// default: "system"
	Org struct {
		Value string `json:"value"`
	} `json:".properties.org,omitempty"`

	BigtableCustomPlans []struct {
		Name struct {
			Value string `json:"value"`
		} `json:".properties.name"`

		// Display Name
		// Name of the plan to be displayed to users.
		DisplayName struct {
			Value string `json:"value"`
		} `json:".properties.display_name"`
	} `json:".properties.bigtable_custom_plans,omitempty"`
}

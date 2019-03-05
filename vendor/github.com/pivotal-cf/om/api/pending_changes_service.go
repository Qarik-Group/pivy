package api

import (
	"encoding/json"
	"fmt"
)

const pendingChangesEndpoint = "/api/v0/staged/pending_changes"

type PendingChangesOutput struct {
	ChangeList []ProductChange `json:"product_changes"`
}

type CompletenessChecks struct {
	ConfigurationComplete       bool `json:"configuration_complete"`
	StemcellPresent             bool `json:"stemcell_present"`
	ConfigurablePropertiesValid bool `json:"configurable_properties_valid"`
}

type ProductChange struct {
	GUID               string              `json:"guid"`
	Action             string              `json:"action"`
	Errands            []Errand            `json:"errands"`
	CompletenessChecks *CompletenessChecks `json:"completeness_checks,omitempty"`
}

func (a Api) ListStagedPendingChanges() (PendingChangesOutput, error) {
	resp, err := a.sendAPIRequest("GET", pendingChangesEndpoint, nil)
	if err != nil {
		return PendingChangesOutput{}, fmt.Errorf("failed to submit request: %s", err)
	}
	defer resp.Body.Close()

	var pendingChanges PendingChangesOutput
	if err := json.NewDecoder(resp.Body).Decode(&pendingChanges); err != nil {
		return PendingChangesOutput{}, fmt.Errorf("could not unmarshal pending_changes response: %s", err)
	}

	return pendingChanges, nil
}

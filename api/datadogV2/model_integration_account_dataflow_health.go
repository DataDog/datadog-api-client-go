// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAccountDataflowHealth Collection health of a single dataflow.
type IntegrationAccountDataflowHealth string

// List of IntegrationAccountDataflowHealth.
const (
	INTEGRATIONACCOUNTDATAFLOWHEALTH_OK      IntegrationAccountDataflowHealth = "DATAFLOW_HEALTH_OK"
	INTEGRATIONACCOUNTDATAFLOWHEALTH_BROKEN  IntegrationAccountDataflowHealth = "DATAFLOW_HEALTH_BROKEN"
	INTEGRATIONACCOUNTDATAFLOWHEALTH_UNKNOWN IntegrationAccountDataflowHealth = "DATAFLOW_HEALTH_UNKNOWN"
)

var allowedIntegrationAccountDataflowHealthEnumValues = []IntegrationAccountDataflowHealth{
	INTEGRATIONACCOUNTDATAFLOWHEALTH_OK,
	INTEGRATIONACCOUNTDATAFLOWHEALTH_BROKEN,
	INTEGRATIONACCOUNTDATAFLOWHEALTH_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IntegrationAccountDataflowHealth) GetAllowedValues() []IntegrationAccountDataflowHealth {
	return allowedIntegrationAccountDataflowHealthEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IntegrationAccountDataflowHealth) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IntegrationAccountDataflowHealth(value)
	return nil
}

// NewIntegrationAccountDataflowHealthFromValue returns a pointer to a valid IntegrationAccountDataflowHealth
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIntegrationAccountDataflowHealthFromValue(v string) (*IntegrationAccountDataflowHealth, error) {
	ev := IntegrationAccountDataflowHealth(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IntegrationAccountDataflowHealth: valid values are %v", v, allowedIntegrationAccountDataflowHealthEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IntegrationAccountDataflowHealth) IsValid() bool {
	for _, existing := range allowedIntegrationAccountDataflowHealthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntegrationAccountDataflowHealth value.
func (v IntegrationAccountDataflowHealth) Ptr() *IntegrationAccountDataflowHealth {
	return &v
}

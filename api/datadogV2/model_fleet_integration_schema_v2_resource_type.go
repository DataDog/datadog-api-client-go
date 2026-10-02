// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetIntegrationSchemaV2ResourceType The type of the integration schema resource.
type FleetIntegrationSchemaV2ResourceType string

// List of FleetIntegrationSchemaV2ResourceType.
const (
	FLEETINTEGRATIONSCHEMAV2RESOURCETYPE_INTEGRATION_SCHEMA FleetIntegrationSchemaV2ResourceType = "integration_schema"
)

var allowedFleetIntegrationSchemaV2ResourceTypeEnumValues = []FleetIntegrationSchemaV2ResourceType{
	FLEETINTEGRATIONSCHEMAV2RESOURCETYPE_INTEGRATION_SCHEMA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FleetIntegrationSchemaV2ResourceType) GetAllowedValues() []FleetIntegrationSchemaV2ResourceType {
	return allowedFleetIntegrationSchemaV2ResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FleetIntegrationSchemaV2ResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FleetIntegrationSchemaV2ResourceType(value)
	return nil
}

// NewFleetIntegrationSchemaV2ResourceTypeFromValue returns a pointer to a valid FleetIntegrationSchemaV2ResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFleetIntegrationSchemaV2ResourceTypeFromValue(v string) (*FleetIntegrationSchemaV2ResourceType, error) {
	ev := FleetIntegrationSchemaV2ResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FleetIntegrationSchemaV2ResourceType: valid values are %v", v, allowedFleetIntegrationSchemaV2ResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FleetIntegrationSchemaV2ResourceType) IsValid() bool {
	for _, existing := range allowedFleetIntegrationSchemaV2ResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FleetIntegrationSchemaV2ResourceType value.
func (v FleetIntegrationSchemaV2ResourceType) Ptr() *FleetIntegrationSchemaV2ResourceType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAgentV2ResourceType The type of the agent resource.
type FleetAgentV2ResourceType string

// List of FleetAgentV2ResourceType.
const (
	FLEETAGENTV2RESOURCETYPE_AGENT FleetAgentV2ResourceType = "agent"
)

var allowedFleetAgentV2ResourceTypeEnumValues = []FleetAgentV2ResourceType{
	FLEETAGENTV2RESOURCETYPE_AGENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FleetAgentV2ResourceType) GetAllowedValues() []FleetAgentV2ResourceType {
	return allowedFleetAgentV2ResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FleetAgentV2ResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FleetAgentV2ResourceType(value)
	return nil
}

// NewFleetAgentV2ResourceTypeFromValue returns a pointer to a valid FleetAgentV2ResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFleetAgentV2ResourceTypeFromValue(v string) (*FleetAgentV2ResourceType, error) {
	ev := FleetAgentV2ResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FleetAgentV2ResourceType: valid values are %v", v, allowedFleetAgentV2ResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FleetAgentV2ResourceType) IsValid() bool {
	for _, existing := range allowedFleetAgentV2ResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FleetAgentV2ResourceType value.
func (v FleetAgentV2ResourceType) Ptr() *FleetAgentV2ResourceType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAccessScopeResourceType The JSON:API resource type for a fleet access scope.
type FleetAccessScopeResourceType string

// List of FleetAccessScopeResourceType.
const (
	FLEETACCESSSCOPERESOURCETYPE_ACCESS_SCOPE FleetAccessScopeResourceType = "access_scope"
)

var allowedFleetAccessScopeResourceTypeEnumValues = []FleetAccessScopeResourceType{
	FLEETACCESSSCOPERESOURCETYPE_ACCESS_SCOPE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FleetAccessScopeResourceType) GetAllowedValues() []FleetAccessScopeResourceType {
	return allowedFleetAccessScopeResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FleetAccessScopeResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FleetAccessScopeResourceType(value)
	return nil
}

// NewFleetAccessScopeResourceTypeFromValue returns a pointer to a valid FleetAccessScopeResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFleetAccessScopeResourceTypeFromValue(v string) (*FleetAccessScopeResourceType, error) {
	ev := FleetAccessScopeResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FleetAccessScopeResourceType: valid values are %v", v, allowedFleetAccessScopeResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FleetAccessScopeResourceType) IsValid() bool {
	for _, existing := range allowedFleetAccessScopeResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FleetAccessScopeResourceType value.
func (v FleetAccessScopeResourceType) Ptr() *FleetAccessScopeResourceType {
	return &v
}

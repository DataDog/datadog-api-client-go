// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapLayoutType Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapLayoutType string

// List of CustomTopologyMapLayoutType.
const (
	CUSTOMTOPOLOGYMAPLAYOUTTYPE_HIERARCHICAL CustomTopologyMapLayoutType = "hierarchical"
)

var allowedCustomTopologyMapLayoutTypeEnumValues = []CustomTopologyMapLayoutType{
	CUSTOMTOPOLOGYMAPLAYOUTTYPE_HIERARCHICAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomTopologyMapLayoutType) GetAllowedValues() []CustomTopologyMapLayoutType {
	return allowedCustomTopologyMapLayoutTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomTopologyMapLayoutType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomTopologyMapLayoutType(value)
	return nil
}

// NewCustomTopologyMapLayoutTypeFromValue returns a pointer to a valid CustomTopologyMapLayoutType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomTopologyMapLayoutTypeFromValue(v string) (*CustomTopologyMapLayoutType, error) {
	ev := CustomTopologyMapLayoutType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomTopologyMapLayoutType: valid values are %v", v, allowedCustomTopologyMapLayoutTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomTopologyMapLayoutType) IsValid() bool {
	for _, existing := range allowedCustomTopologyMapLayoutTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomTopologyMapLayoutType value.
func (v CustomTopologyMapLayoutType) Ptr() *CustomTopologyMapLayoutType {
	return &v
}

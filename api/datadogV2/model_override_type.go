// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OverrideType The definition of `OverrideType` object.
type OverrideType string

// List of OverrideType.
const (
	OVERRIDETYPE_OVERRIDES OverrideType = "overrides"
)

var allowedOverrideTypeEnumValues = []OverrideType{
	OVERRIDETYPE_OVERRIDES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OverrideType) GetAllowedValues() []OverrideType {
	return allowedOverrideTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OverrideType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OverrideType(value)
	return nil
}

// NewOverrideTypeFromValue returns a pointer to a valid OverrideType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOverrideTypeFromValue(v string) (*OverrideType, error) {
	ev := OverrideType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OverrideType: valid values are %v", v, allowedOverrideTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OverrideType) IsValid() bool {
	for _, existing := range allowedOverrideTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OverrideType value.
func (v OverrideType) Ptr() *OverrideType {
	return &v
}

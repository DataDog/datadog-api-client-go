// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OverrideDataType Indicates that the resource is of type 'overrides'.
type OverrideDataType string

// List of OverrideDataType.
const (
	OVERRIDEDATATYPE_OVERRIDES OverrideDataType = "overrides"
)

var allowedOverrideDataTypeEnumValues = []OverrideDataType{
	OVERRIDEDATATYPE_OVERRIDES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OverrideDataType) GetAllowedValues() []OverrideDataType {
	return allowedOverrideDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OverrideDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OverrideDataType(value)
	return nil
}

// NewOverrideDataTypeFromValue returns a pointer to a valid OverrideDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOverrideDataTypeFromValue(v string) (*OverrideDataType, error) {
	ev := OverrideDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OverrideDataType: valid values are %v", v, allowedOverrideDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OverrideDataType) IsValid() bool {
	for _, existing := range allowedOverrideDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OverrideDataType value.
func (v OverrideDataType) Ptr() *OverrideDataType {
	return &v
}

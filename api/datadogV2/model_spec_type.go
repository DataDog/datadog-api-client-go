// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpecType Type of the spec resource.
type SpecType string

// List of SpecType.
const (
	SPECTYPE_SPEC SpecType = "spec"
)

var allowedSpecTypeEnumValues = []SpecType{
	SPECTYPE_SPEC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpecType) GetAllowedValues() []SpecType {
	return allowedSpecTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpecType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpecType(value)
	return nil
}

// NewSpecTypeFromValue returns a pointer to a valid SpecType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpecTypeFromValue(v string) (*SpecType, error) {
	ev := SpecType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpecType: valid values are %v", v, allowedSpecTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpecType) IsValid() bool {
	for _, existing := range allowedSpecTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpecType value.
func (v SpecType) Ptr() *SpecType {
	return &v
}

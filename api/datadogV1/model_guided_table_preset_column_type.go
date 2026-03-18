// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTablePresetColumnType
type GuidedTablePresetColumnType string

// List of GuidedTablePresetColumnType.
const (
	GUIDEDTABLEPRESETCOLUMNTYPE_PRESET GuidedTablePresetColumnType = "preset"
)

var allowedGuidedTablePresetColumnTypeEnumValues = []GuidedTablePresetColumnType{
	GUIDEDTABLEPRESETCOLUMNTYPE_PRESET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GuidedTablePresetColumnType) GetAllowedValues() []GuidedTablePresetColumnType {
	return allowedGuidedTablePresetColumnTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GuidedTablePresetColumnType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GuidedTablePresetColumnType(value)
	return nil
}

// NewGuidedTablePresetColumnTypeFromValue returns a pointer to a valid GuidedTablePresetColumnType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGuidedTablePresetColumnTypeFromValue(v string) (*GuidedTablePresetColumnType, error) {
	ev := GuidedTablePresetColumnType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GuidedTablePresetColumnType: valid values are %v", v, allowedGuidedTablePresetColumnTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GuidedTablePresetColumnType) IsValid() bool {
	for _, existing := range allowedGuidedTablePresetColumnTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuidedTablePresetColumnType value.
func (v GuidedTablePresetColumnType) Ptr() *GuidedTablePresetColumnType {
	return &v
}

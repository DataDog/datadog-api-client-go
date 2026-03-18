// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTablePresetColumnPresetType
type GuidedTablePresetColumnPresetType string

// List of GuidedTablePresetColumnPresetType.
const (
	GUIDEDTABLEPRESETCOLUMNPRESETTYPE_PREVIOUS_PERIOD GuidedTablePresetColumnPresetType = "previous_period"
)

var allowedGuidedTablePresetColumnPresetTypeEnumValues = []GuidedTablePresetColumnPresetType{
	GUIDEDTABLEPRESETCOLUMNPRESETTYPE_PREVIOUS_PERIOD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GuidedTablePresetColumnPresetType) GetAllowedValues() []GuidedTablePresetColumnPresetType {
	return allowedGuidedTablePresetColumnPresetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GuidedTablePresetColumnPresetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GuidedTablePresetColumnPresetType(value)
	return nil
}

// NewGuidedTablePresetColumnPresetTypeFromValue returns a pointer to a valid GuidedTablePresetColumnPresetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGuidedTablePresetColumnPresetTypeFromValue(v string) (*GuidedTablePresetColumnPresetType, error) {
	ev := GuidedTablePresetColumnPresetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GuidedTablePresetColumnPresetType: valid values are %v", v, allowedGuidedTablePresetColumnPresetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GuidedTablePresetColumnPresetType) IsValid() bool {
	for _, existing := range allowedGuidedTablePresetColumnPresetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuidedTablePresetColumnPresetType value.
func (v GuidedTablePresetColumnPresetType) Ptr() *GuidedTablePresetColumnPresetType {
	return &v
}

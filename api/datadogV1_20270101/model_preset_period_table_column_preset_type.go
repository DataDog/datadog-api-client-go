// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PresetPeriodTableColumnPresetType Test-only generated schema for the dashboard versioning proof of concept.
type PresetPeriodTableColumnPresetType string

// List of PresetPeriodTableColumnPresetType.
const (
	PRESETPERIODTABLECOLUMNPRESETTYPE_PREVIOUS_PERIOD PresetPeriodTableColumnPresetType = "previous_period"
)

var allowedPresetPeriodTableColumnPresetTypeEnumValues = []PresetPeriodTableColumnPresetType{
	PRESETPERIODTABLECOLUMNPRESETTYPE_PREVIOUS_PERIOD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PresetPeriodTableColumnPresetType) GetAllowedValues() []PresetPeriodTableColumnPresetType {
	return allowedPresetPeriodTableColumnPresetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PresetPeriodTableColumnPresetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PresetPeriodTableColumnPresetType(value)
	return nil
}

// NewPresetPeriodTableColumnPresetTypeFromValue returns a pointer to a valid PresetPeriodTableColumnPresetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPresetPeriodTableColumnPresetTypeFromValue(v string) (*PresetPeriodTableColumnPresetType, error) {
	ev := PresetPeriodTableColumnPresetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PresetPeriodTableColumnPresetType: valid values are %v", v, allowedPresetPeriodTableColumnPresetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PresetPeriodTableColumnPresetType) IsValid() bool {
	for _, existing := range allowedPresetPeriodTableColumnPresetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PresetPeriodTableColumnPresetType value.
func (v PresetPeriodTableColumnPresetType) Ptr() *PresetPeriodTableColumnPresetType {
	return &v
}

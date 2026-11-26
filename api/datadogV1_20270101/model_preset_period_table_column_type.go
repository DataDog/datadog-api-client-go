// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PresetPeriodTableColumnType Test-only generated schema for the dashboard versioning proof of concept.
type PresetPeriodTableColumnType string

// List of PresetPeriodTableColumnType.
const (
	PRESETPERIODTABLECOLUMNTYPE_PRESET PresetPeriodTableColumnType = "preset"
)

var allowedPresetPeriodTableColumnTypeEnumValues = []PresetPeriodTableColumnType{
	PRESETPERIODTABLECOLUMNTYPE_PRESET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PresetPeriodTableColumnType) GetAllowedValues() []PresetPeriodTableColumnType {
	return allowedPresetPeriodTableColumnTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PresetPeriodTableColumnType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PresetPeriodTableColumnType(value)
	return nil
}

// NewPresetPeriodTableColumnTypeFromValue returns a pointer to a valid PresetPeriodTableColumnType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPresetPeriodTableColumnTypeFromValue(v string) (*PresetPeriodTableColumnType, error) {
	ev := PresetPeriodTableColumnType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PresetPeriodTableColumnType: valid values are %v", v, allowedPresetPeriodTableColumnTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PresetPeriodTableColumnType) IsValid() bool {
	for _, existing := range allowedPresetPeriodTableColumnTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PresetPeriodTableColumnType value.
func (v PresetPeriodTableColumnType) Ptr() *PresetPeriodTableColumnType {
	return &v
}

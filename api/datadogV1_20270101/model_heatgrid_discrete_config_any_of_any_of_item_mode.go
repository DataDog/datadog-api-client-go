// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridDiscreteConfigAnyOfAnyOfItemMode Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridDiscreteConfigAnyOfAnyOfItemMode string

// List of HeatgridDiscreteConfigAnyOfAnyOfItemMode.
const (
	HEATGRIDDISCRETECONFIGANYOFANYOFITEMMODE_DISCRETE HeatgridDiscreteConfigAnyOfAnyOfItemMode = "discrete"
)

var allowedHeatgridDiscreteConfigAnyOfAnyOfItemModeEnumValues = []HeatgridDiscreteConfigAnyOfAnyOfItemMode{
	HEATGRIDDISCRETECONFIGANYOFANYOFITEMMODE_DISCRETE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridDiscreteConfigAnyOfAnyOfItemMode) GetAllowedValues() []HeatgridDiscreteConfigAnyOfAnyOfItemMode {
	return allowedHeatgridDiscreteConfigAnyOfAnyOfItemModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridDiscreteConfigAnyOfAnyOfItemMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridDiscreteConfigAnyOfAnyOfItemMode(value)
	return nil
}

// NewHeatgridDiscreteConfigAnyOfAnyOfItemModeFromValue returns a pointer to a valid HeatgridDiscreteConfigAnyOfAnyOfItemMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridDiscreteConfigAnyOfAnyOfItemModeFromValue(v string) (*HeatgridDiscreteConfigAnyOfAnyOfItemMode, error) {
	ev := HeatgridDiscreteConfigAnyOfAnyOfItemMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridDiscreteConfigAnyOfAnyOfItemMode: valid values are %v", v, allowedHeatgridDiscreteConfigAnyOfAnyOfItemModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridDiscreteConfigAnyOfAnyOfItemMode) IsValid() bool {
	for _, existing := range allowedHeatgridDiscreteConfigAnyOfAnyOfItemModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridDiscreteConfigAnyOfAnyOfItemMode value.
func (v HeatgridDiscreteConfigAnyOfAnyOfItemMode) Ptr() *HeatgridDiscreteConfigAnyOfAnyOfItemMode {
	return &v
}

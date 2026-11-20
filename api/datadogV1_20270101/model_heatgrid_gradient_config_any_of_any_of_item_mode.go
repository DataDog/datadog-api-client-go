// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridGradientConfigAnyOfAnyOfItemMode Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridGradientConfigAnyOfAnyOfItemMode string

// List of HeatgridGradientConfigAnyOfAnyOfItemMode.
const (
	HEATGRIDGRADIENTCONFIGANYOFANYOFITEMMODE_GRADIENT HeatgridGradientConfigAnyOfAnyOfItemMode = "gradient"
)

var allowedHeatgridGradientConfigAnyOfAnyOfItemModeEnumValues = []HeatgridGradientConfigAnyOfAnyOfItemMode{
	HEATGRIDGRADIENTCONFIGANYOFANYOFITEMMODE_GRADIENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridGradientConfigAnyOfAnyOfItemMode) GetAllowedValues() []HeatgridGradientConfigAnyOfAnyOfItemMode {
	return allowedHeatgridGradientConfigAnyOfAnyOfItemModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridGradientConfigAnyOfAnyOfItemMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridGradientConfigAnyOfAnyOfItemMode(value)
	return nil
}

// NewHeatgridGradientConfigAnyOfAnyOfItemModeFromValue returns a pointer to a valid HeatgridGradientConfigAnyOfAnyOfItemMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridGradientConfigAnyOfAnyOfItemModeFromValue(v string) (*HeatgridGradientConfigAnyOfAnyOfItemMode, error) {
	ev := HeatgridGradientConfigAnyOfAnyOfItemMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridGradientConfigAnyOfAnyOfItemMode: valid values are %v", v, allowedHeatgridGradientConfigAnyOfAnyOfItemModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridGradientConfigAnyOfAnyOfItemMode) IsValid() bool {
	for _, existing := range allowedHeatgridGradientConfigAnyOfAnyOfItemModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridGradientConfigAnyOfAnyOfItemMode value.
func (v HeatgridGradientConfigAnyOfAnyOfItemMode) Ptr() *HeatgridGradientConfigAnyOfAnyOfItemMode {
	return &v
}

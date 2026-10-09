// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridPresetColorSource Use a preset color palette.
type HeatgridPresetColorSource string

// List of HeatgridPresetColorSource.
const (
	HEATGRIDPRESETCOLORSOURCE_PRESET HeatgridPresetColorSource = "preset"
)

var allowedHeatgridPresetColorSourceEnumValues = []HeatgridPresetColorSource{
	HEATGRIDPRESETCOLORSOURCE_PRESET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridPresetColorSource) GetAllowedValues() []HeatgridPresetColorSource {
	return allowedHeatgridPresetColorSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridPresetColorSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridPresetColorSource(value)
	return nil
}

// NewHeatgridPresetColorSourceFromValue returns a pointer to a valid HeatgridPresetColorSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridPresetColorSourceFromValue(v string) (*HeatgridPresetColorSource, error) {
	ev := HeatgridPresetColorSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridPresetColorSource: valid values are %v", v, allowedHeatgridPresetColorSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridPresetColorSource) IsValid() bool {
	for _, existing := range allowedHeatgridPresetColorSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridPresetColorSource value.
func (v HeatgridPresetColorSource) Ptr() *HeatgridPresetColorSource {
	return &v
}

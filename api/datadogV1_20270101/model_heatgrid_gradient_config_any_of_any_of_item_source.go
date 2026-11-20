// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridGradientConfigAnyOfAnyOfItemSource Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridGradientConfigAnyOfAnyOfItemSource string

// List of HeatgridGradientConfigAnyOfAnyOfItemSource.
const (
	HEATGRIDGRADIENTCONFIGANYOFANYOFITEMSOURCE_CUSTOM HeatgridGradientConfigAnyOfAnyOfItemSource = "custom"
)

var allowedHeatgridGradientConfigAnyOfAnyOfItemSourceEnumValues = []HeatgridGradientConfigAnyOfAnyOfItemSource{
	HEATGRIDGRADIENTCONFIGANYOFANYOFITEMSOURCE_CUSTOM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridGradientConfigAnyOfAnyOfItemSource) GetAllowedValues() []HeatgridGradientConfigAnyOfAnyOfItemSource {
	return allowedHeatgridGradientConfigAnyOfAnyOfItemSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridGradientConfigAnyOfAnyOfItemSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridGradientConfigAnyOfAnyOfItemSource(value)
	return nil
}

// NewHeatgridGradientConfigAnyOfAnyOfItemSourceFromValue returns a pointer to a valid HeatgridGradientConfigAnyOfAnyOfItemSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridGradientConfigAnyOfAnyOfItemSourceFromValue(v string) (*HeatgridGradientConfigAnyOfAnyOfItemSource, error) {
	ev := HeatgridGradientConfigAnyOfAnyOfItemSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridGradientConfigAnyOfAnyOfItemSource: valid values are %v", v, allowedHeatgridGradientConfigAnyOfAnyOfItemSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridGradientConfigAnyOfAnyOfItemSource) IsValid() bool {
	for _, existing := range allowedHeatgridGradientConfigAnyOfAnyOfItemSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridGradientConfigAnyOfAnyOfItemSource value.
func (v HeatgridGradientConfigAnyOfAnyOfItemSource) Ptr() *HeatgridGradientConfigAnyOfAnyOfItemSource {
	return &v
}

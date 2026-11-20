// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridGradientConfigAnyOfAnyOfItemSource4447622208 Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridGradientConfigAnyOfAnyOfItemSource4447622208 string

// List of HeatgridGradientConfigAnyOfAnyOfItemSource4447622208.
const (
	HEATGRIDGRADIENTCONFIGANYOFANYOFITEMSOURCE4447622208_PRESET HeatgridGradientConfigAnyOfAnyOfItemSource4447622208 = "preset"
)

var allowedHeatgridGradientConfigAnyOfAnyOfItemSource4447622208EnumValues = []HeatgridGradientConfigAnyOfAnyOfItemSource4447622208{
	HEATGRIDGRADIENTCONFIGANYOFANYOFITEMSOURCE4447622208_PRESET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridGradientConfigAnyOfAnyOfItemSource4447622208) GetAllowedValues() []HeatgridGradientConfigAnyOfAnyOfItemSource4447622208 {
	return allowedHeatgridGradientConfigAnyOfAnyOfItemSource4447622208EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridGradientConfigAnyOfAnyOfItemSource4447622208) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridGradientConfigAnyOfAnyOfItemSource4447622208(value)
	return nil
}

// NewHeatgridGradientConfigAnyOfAnyOfItemSource4447622208FromValue returns a pointer to a valid HeatgridGradientConfigAnyOfAnyOfItemSource4447622208
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridGradientConfigAnyOfAnyOfItemSource4447622208FromValue(v string) (*HeatgridGradientConfigAnyOfAnyOfItemSource4447622208, error) {
	ev := HeatgridGradientConfigAnyOfAnyOfItemSource4447622208(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridGradientConfigAnyOfAnyOfItemSource4447622208: valid values are %v", v, allowedHeatgridGradientConfigAnyOfAnyOfItemSource4447622208EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridGradientConfigAnyOfAnyOfItemSource4447622208) IsValid() bool {
	for _, existing := range allowedHeatgridGradientConfigAnyOfAnyOfItemSource4447622208EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridGradientConfigAnyOfAnyOfItemSource4447622208 value.
func (v HeatgridGradientConfigAnyOfAnyOfItemSource4447622208) Ptr() *HeatgridGradientConfigAnyOfAnyOfItemSource4447622208 {
	return &v
}

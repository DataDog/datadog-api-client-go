// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168 Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168 string

// List of HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168.
const (
	HEATGRIDDISCRETECONFIGANYOFANYOFITEMSOURCE4447619168_PRESET HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168 = "preset"
)

var allowedHeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168EnumValues = []HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168{
	HEATGRIDDISCRETECONFIGANYOFANYOFITEMSOURCE4447619168_PRESET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168) GetAllowedValues() []HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168 {
	return allowedHeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168(value)
	return nil
}

// NewHeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168FromValue returns a pointer to a valid HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168FromValue(v string) (*HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168, error) {
	ev := HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168: valid values are %v", v, allowedHeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168) IsValid() bool {
	for _, existing := range allowedHeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168 value.
func (v HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168) Ptr() *HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168 {
	return &v
}

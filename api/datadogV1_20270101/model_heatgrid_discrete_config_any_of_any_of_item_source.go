// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridDiscreteConfigAnyOfAnyOfItemSource Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridDiscreteConfigAnyOfAnyOfItemSource string

// List of HeatgridDiscreteConfigAnyOfAnyOfItemSource.
const (
	HEATGRIDDISCRETECONFIGANYOFANYOFITEMSOURCE_CUSTOM HeatgridDiscreteConfigAnyOfAnyOfItemSource = "custom"
)

var allowedHeatgridDiscreteConfigAnyOfAnyOfItemSourceEnumValues = []HeatgridDiscreteConfigAnyOfAnyOfItemSource{
	HEATGRIDDISCRETECONFIGANYOFANYOFITEMSOURCE_CUSTOM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridDiscreteConfigAnyOfAnyOfItemSource) GetAllowedValues() []HeatgridDiscreteConfigAnyOfAnyOfItemSource {
	return allowedHeatgridDiscreteConfigAnyOfAnyOfItemSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridDiscreteConfigAnyOfAnyOfItemSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridDiscreteConfigAnyOfAnyOfItemSource(value)
	return nil
}

// NewHeatgridDiscreteConfigAnyOfAnyOfItemSourceFromValue returns a pointer to a valid HeatgridDiscreteConfigAnyOfAnyOfItemSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridDiscreteConfigAnyOfAnyOfItemSourceFromValue(v string) (*HeatgridDiscreteConfigAnyOfAnyOfItemSource, error) {
	ev := HeatgridDiscreteConfigAnyOfAnyOfItemSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridDiscreteConfigAnyOfAnyOfItemSource: valid values are %v", v, allowedHeatgridDiscreteConfigAnyOfAnyOfItemSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridDiscreteConfigAnyOfAnyOfItemSource) IsValid() bool {
	for _, existing := range allowedHeatgridDiscreteConfigAnyOfAnyOfItemSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridDiscreteConfigAnyOfAnyOfItemSource value.
func (v HeatgridDiscreteConfigAnyOfAnyOfItemSource) Ptr() *HeatgridDiscreteConfigAnyOfAnyOfItemSource {
	return &v
}

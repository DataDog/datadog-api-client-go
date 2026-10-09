// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridSortByValueProperty Sort by value.
type HeatgridSortByValueProperty string

// List of HeatgridSortByValueProperty.
const (
	HEATGRIDSORTBYVALUEPROPERTY_VALUE HeatgridSortByValueProperty = "value"
)

var allowedHeatgridSortByValuePropertyEnumValues = []HeatgridSortByValueProperty{
	HEATGRIDSORTBYVALUEPROPERTY_VALUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridSortByValueProperty) GetAllowedValues() []HeatgridSortByValueProperty {
	return allowedHeatgridSortByValuePropertyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridSortByValueProperty) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridSortByValueProperty(value)
	return nil
}

// NewHeatgridSortByValuePropertyFromValue returns a pointer to a valid HeatgridSortByValueProperty
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridSortByValuePropertyFromValue(v string) (*HeatgridSortByValueProperty, error) {
	ev := HeatgridSortByValueProperty(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridSortByValueProperty: valid values are %v", v, allowedHeatgridSortByValuePropertyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridSortByValueProperty) IsValid() bool {
	for _, existing := range allowedHeatgridSortByValuePropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridSortByValueProperty value.
func (v HeatgridSortByValueProperty) Ptr() *HeatgridSortByValueProperty {
	return &v
}

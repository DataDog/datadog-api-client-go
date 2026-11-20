// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800 Test-only generated schema for the dashboard versioning proof of concept.
type ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800 string

// List of ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800.
const (
	SCATTERPLOTLEGENDTYPEANYOFANYOFITEMTYPE4451196800_NONE ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800 = "none"
)

var allowedScatterPlotLegendTypeAnyOfAnyOfItemType4451196800EnumValues = []ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800{
	SCATTERPLOTLEGENDTYPEANYOFANYOFITEMTYPE4451196800_NONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800) GetAllowedValues() []ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800 {
	return allowedScatterPlotLegendTypeAnyOfAnyOfItemType4451196800EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800(value)
	return nil
}

// NewScatterPlotLegendTypeAnyOfAnyOfItemType4451196800FromValue returns a pointer to a valid ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScatterPlotLegendTypeAnyOfAnyOfItemType4451196800FromValue(v string) (*ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800, error) {
	ev := ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800: valid values are %v", v, allowedScatterPlotLegendTypeAnyOfAnyOfItemType4451196800EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800) IsValid() bool {
	for _, existing := range allowedScatterPlotLegendTypeAnyOfAnyOfItemType4451196800EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800 value.
func (v ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800) Ptr() *ScatterPlotLegendTypeAnyOfAnyOfItemType4451196800 {
	return &v
}

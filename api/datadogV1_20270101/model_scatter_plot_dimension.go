// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterPlotDimension Test-only generated schema for the dashboard versioning proof of concept.
type ScatterPlotDimension string

// List of ScatterPlotDimension.
const (
	SCATTERPLOTDIMENSION_COLOR  ScatterPlotDimension = "color"
	SCATTERPLOTDIMENSION_X      ScatterPlotDimension = "x"
	SCATTERPLOTDIMENSION_Y      ScatterPlotDimension = "y"
	SCATTERPLOTDIMENSION_RADIUS ScatterPlotDimension = "radius"
)

var allowedScatterPlotDimensionEnumValues = []ScatterPlotDimension{
	SCATTERPLOTDIMENSION_COLOR,
	SCATTERPLOTDIMENSION_X,
	SCATTERPLOTDIMENSION_Y,
	SCATTERPLOTDIMENSION_RADIUS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScatterPlotDimension) GetAllowedValues() []ScatterPlotDimension {
	return allowedScatterPlotDimensionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScatterPlotDimension) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScatterPlotDimension(value)
	return nil
}

// NewScatterPlotDimensionFromValue returns a pointer to a valid ScatterPlotDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScatterPlotDimensionFromValue(v string) (*ScatterPlotDimension, error) {
	ev := ScatterPlotDimension(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScatterPlotDimension: valid values are %v", v, allowedScatterPlotDimensionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScatterPlotDimension) IsValid() bool {
	for _, existing := range allowedScatterPlotDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScatterPlotDimension value.
func (v ScatterPlotDimension) Ptr() *ScatterPlotDimension {
	return &v
}

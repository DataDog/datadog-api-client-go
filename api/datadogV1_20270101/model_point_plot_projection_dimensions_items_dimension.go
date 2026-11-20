// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PointPlotProjectionDimensionsItemsDimension Test-only generated schema for the dashboard versioning proof of concept.
type PointPlotProjectionDimensionsItemsDimension string

// List of PointPlotProjectionDimensionsItemsDimension.
const (
	POINTPLOTPROJECTIONDIMENSIONSITEMSDIMENSION_GROUP  PointPlotProjectionDimensionsItemsDimension = "group"
	POINTPLOTPROJECTIONDIMENSIONSITEMSDIMENSION_Y      PointPlotProjectionDimensionsItemsDimension = "y"
	POINTPLOTPROJECTIONDIMENSIONSITEMSDIMENSION_RADIUS PointPlotProjectionDimensionsItemsDimension = "radius"
	POINTPLOTPROJECTIONDIMENSIONSITEMSDIMENSION_TIME   PointPlotProjectionDimensionsItemsDimension = "time"
)

var allowedPointPlotProjectionDimensionsItemsDimensionEnumValues = []PointPlotProjectionDimensionsItemsDimension{
	POINTPLOTPROJECTIONDIMENSIONSITEMSDIMENSION_GROUP,
	POINTPLOTPROJECTIONDIMENSIONSITEMSDIMENSION_Y,
	POINTPLOTPROJECTIONDIMENSIONSITEMSDIMENSION_RADIUS,
	POINTPLOTPROJECTIONDIMENSIONSITEMSDIMENSION_TIME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PointPlotProjectionDimensionsItemsDimension) GetAllowedValues() []PointPlotProjectionDimensionsItemsDimension {
	return allowedPointPlotProjectionDimensionsItemsDimensionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PointPlotProjectionDimensionsItemsDimension) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PointPlotProjectionDimensionsItemsDimension(value)
	return nil
}

// NewPointPlotProjectionDimensionsItemsDimensionFromValue returns a pointer to a valid PointPlotProjectionDimensionsItemsDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPointPlotProjectionDimensionsItemsDimensionFromValue(v string) (*PointPlotProjectionDimensionsItemsDimension, error) {
	ev := PointPlotProjectionDimensionsItemsDimension(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PointPlotProjectionDimensionsItemsDimension: valid values are %v", v, allowedPointPlotProjectionDimensionsItemsDimensionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PointPlotProjectionDimensionsItemsDimension) IsValid() bool {
	for _, existing := range allowedPointPlotProjectionDimensionsItemsDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PointPlotProjectionDimensionsItemsDimension value.
func (v PointPlotProjectionDimensionsItemsDimension) Ptr() *PointPlotProjectionDimensionsItemsDimension {
	return &v
}

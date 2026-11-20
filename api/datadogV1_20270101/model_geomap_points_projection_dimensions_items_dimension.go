// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapPointsProjectionDimensionsItemsDimension Test-only generated schema for the dashboard versioning proof of concept.
type GeomapPointsProjectionDimensionsItemsDimension string

// List of GeomapPointsProjectionDimensionsItemsDimension.
const (
	GEOMAPPOINTSPROJECTIONDIMENSIONSITEMSDIMENSION_COLOR     GeomapPointsProjectionDimensionsItemsDimension = "color"
	GEOMAPPOINTSPROJECTIONDIMENSIONSITEMSDIMENSION_LATITUDE  GeomapPointsProjectionDimensionsItemsDimension = "latitude"
	GEOMAPPOINTSPROJECTIONDIMENSIONSITEMSDIMENSION_LONGITUDE GeomapPointsProjectionDimensionsItemsDimension = "longitude"
)

var allowedGeomapPointsProjectionDimensionsItemsDimensionEnumValues = []GeomapPointsProjectionDimensionsItemsDimension{
	GEOMAPPOINTSPROJECTIONDIMENSIONSITEMSDIMENSION_COLOR,
	GEOMAPPOINTSPROJECTIONDIMENSIONSITEMSDIMENSION_LATITUDE,
	GEOMAPPOINTSPROJECTIONDIMENSIONSITEMSDIMENSION_LONGITUDE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeomapPointsProjectionDimensionsItemsDimension) GetAllowedValues() []GeomapPointsProjectionDimensionsItemsDimension {
	return allowedGeomapPointsProjectionDimensionsItemsDimensionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeomapPointsProjectionDimensionsItemsDimension) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeomapPointsProjectionDimensionsItemsDimension(value)
	return nil
}

// NewGeomapPointsProjectionDimensionsItemsDimensionFromValue returns a pointer to a valid GeomapPointsProjectionDimensionsItemsDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeomapPointsProjectionDimensionsItemsDimensionFromValue(v string) (*GeomapPointsProjectionDimensionsItemsDimension, error) {
	ev := GeomapPointsProjectionDimensionsItemsDimension(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeomapPointsProjectionDimensionsItemsDimension: valid values are %v", v, allowedGeomapPointsProjectionDimensionsItemsDimensionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeomapPointsProjectionDimensionsItemsDimension) IsValid() bool {
	for _, existing := range allowedGeomapPointsProjectionDimensionsItemsDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeomapPointsProjectionDimensionsItemsDimension value.
func (v GeomapPointsProjectionDimensionsItemsDimension) Ptr() *GeomapPointsProjectionDimensionsItemsDimension {
	return &v
}

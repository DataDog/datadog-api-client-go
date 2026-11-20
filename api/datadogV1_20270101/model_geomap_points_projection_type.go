// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapPointsProjectionType Test-only generated schema for the dashboard versioning proof of concept.
type GeomapPointsProjectionType string

// List of GeomapPointsProjectionType.
const (
	GEOMAPPOINTSPROJECTIONTYPE_GEOMAP_POINTS GeomapPointsProjectionType = "geomap_points"
)

var allowedGeomapPointsProjectionTypeEnumValues = []GeomapPointsProjectionType{
	GEOMAPPOINTSPROJECTIONTYPE_GEOMAP_POINTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeomapPointsProjectionType) GetAllowedValues() []GeomapPointsProjectionType {
	return allowedGeomapPointsProjectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeomapPointsProjectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeomapPointsProjectionType(value)
	return nil
}

// NewGeomapPointsProjectionTypeFromValue returns a pointer to a valid GeomapPointsProjectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeomapPointsProjectionTypeFromValue(v string) (*GeomapPointsProjectionType, error) {
	ev := GeomapPointsProjectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeomapPointsProjectionType: valid values are %v", v, allowedGeomapPointsProjectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeomapPointsProjectionType) IsValid() bool {
	for _, existing := range allowedGeomapPointsProjectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeomapPointsProjectionType value.
func (v GeomapPointsProjectionType) Ptr() *GeomapPointsProjectionType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendDefinitionType string

// List of GeomapBackendDefinitionType.
const (
	GEOMAPBACKENDDEFINITIONTYPE_GEOMAP GeomapBackendDefinitionType = "geomap"
)

var allowedGeomapBackendDefinitionTypeEnumValues = []GeomapBackendDefinitionType{
	GEOMAPBACKENDDEFINITIONTYPE_GEOMAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeomapBackendDefinitionType) GetAllowedValues() []GeomapBackendDefinitionType {
	return allowedGeomapBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeomapBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeomapBackendDefinitionType(value)
	return nil
}

// NewGeomapBackendDefinitionTypeFromValue returns a pointer to a valid GeomapBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeomapBackendDefinitionTypeFromValue(v string) (*GeomapBackendDefinitionType, error) {
	ev := GeomapBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeomapBackendDefinitionType: valid values are %v", v, allowedGeomapBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeomapBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedGeomapBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeomapBackendDefinitionType value.
func (v GeomapBackendDefinitionType) Ptr() *GeomapBackendDefinitionType {
	return &v
}

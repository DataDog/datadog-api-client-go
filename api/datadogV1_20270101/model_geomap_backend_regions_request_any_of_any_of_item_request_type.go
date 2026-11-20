// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType string

// List of GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType.
const (
	GEOMAPBACKENDREGIONSREQUESTANYOFANYOFITEMREQUESTTYPE_DDSQL GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType = "ddsql"
)

var allowedGeomapBackendRegionsRequestAnyOfAnyOfItemRequestTypeEnumValues = []GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType{
	GEOMAPBACKENDREGIONSREQUESTANYOFANYOFITEMREQUESTTYPE_DDSQL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType) GetAllowedValues() []GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType {
	return allowedGeomapBackendRegionsRequestAnyOfAnyOfItemRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType(value)
	return nil
}

// NewGeomapBackendRegionsRequestAnyOfAnyOfItemRequestTypeFromValue returns a pointer to a valid GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeomapBackendRegionsRequestAnyOfAnyOfItemRequestTypeFromValue(v string) (*GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType, error) {
	ev := GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType: valid values are %v", v, allowedGeomapBackendRegionsRequestAnyOfAnyOfItemRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType) IsValid() bool {
	for _, existing := range allowedGeomapBackendRegionsRequestAnyOfAnyOfItemRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType value.
func (v GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType) Ptr() *GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType {
	return &v
}

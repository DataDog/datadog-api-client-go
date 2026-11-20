// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendPointsRequestAnyOfAnyOfItemRequestType Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendPointsRequestAnyOfAnyOfItemRequestType string

// List of GeomapBackendPointsRequestAnyOfAnyOfItemRequestType.
const (
	GEOMAPBACKENDPOINTSREQUESTANYOFANYOFITEMREQUESTTYPE_NETWORK_DEVICES_GEOMAP GeomapBackendPointsRequestAnyOfAnyOfItemRequestType = "network_devices_geomap"
)

var allowedGeomapBackendPointsRequestAnyOfAnyOfItemRequestTypeEnumValues = []GeomapBackendPointsRequestAnyOfAnyOfItemRequestType{
	GEOMAPBACKENDPOINTSREQUESTANYOFANYOFITEMREQUESTTYPE_NETWORK_DEVICES_GEOMAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeomapBackendPointsRequestAnyOfAnyOfItemRequestType) GetAllowedValues() []GeomapBackendPointsRequestAnyOfAnyOfItemRequestType {
	return allowedGeomapBackendPointsRequestAnyOfAnyOfItemRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeomapBackendPointsRequestAnyOfAnyOfItemRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeomapBackendPointsRequestAnyOfAnyOfItemRequestType(value)
	return nil
}

// NewGeomapBackendPointsRequestAnyOfAnyOfItemRequestTypeFromValue returns a pointer to a valid GeomapBackendPointsRequestAnyOfAnyOfItemRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeomapBackendPointsRequestAnyOfAnyOfItemRequestTypeFromValue(v string) (*GeomapBackendPointsRequestAnyOfAnyOfItemRequestType, error) {
	ev := GeomapBackendPointsRequestAnyOfAnyOfItemRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeomapBackendPointsRequestAnyOfAnyOfItemRequestType: valid values are %v", v, allowedGeomapBackendPointsRequestAnyOfAnyOfItemRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeomapBackendPointsRequestAnyOfAnyOfItemRequestType) IsValid() bool {
	for _, existing := range allowedGeomapBackendPointsRequestAnyOfAnyOfItemRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeomapBackendPointsRequestAnyOfAnyOfItemRequestType value.
func (v GeomapBackendPointsRequestAnyOfAnyOfItemRequestType) Ptr() *GeomapBackendPointsRequestAnyOfAnyOfItemRequestType {
	return &v
}

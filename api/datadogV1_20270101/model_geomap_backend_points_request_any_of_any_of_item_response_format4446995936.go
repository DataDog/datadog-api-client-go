// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936 Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936 string

// List of GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936.
const (
	GEOMAPBACKENDPOINTSREQUESTANYOFANYOFITEMRESPONSEFORMAT4446995936_EVENT_LIST GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936 = "event_list"
)

var allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936EnumValues = []GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936{
	GEOMAPBACKENDPOINTSREQUESTANYOFANYOFITEMRESPONSEFORMAT4446995936_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936) GetAllowedValues() []GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936 {
	return allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936(value)
	return nil
}

// NewGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936FromValue returns a pointer to a valid GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936FromValue(v string) (*GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936, error) {
	ev := GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936: valid values are %v", v, allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936) IsValid() bool {
	for _, existing := range allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936 value.
func (v GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936) Ptr() *GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446995936 {
	return &v
}

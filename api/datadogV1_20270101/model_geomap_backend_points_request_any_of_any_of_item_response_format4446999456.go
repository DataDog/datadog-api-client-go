// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456 Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456 string

// List of GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456.
const (
	GEOMAPBACKENDPOINTSREQUESTANYOFANYOFITEMRESPONSEFORMAT4446999456_EVENT_LIST GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456 = "event_list"
)

var allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456EnumValues = []GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456{
	GEOMAPBACKENDPOINTSREQUESTANYOFANYOFITEMRESPONSEFORMAT4446999456_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456) GetAllowedValues() []GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456 {
	return allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456(value)
	return nil
}

// NewGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456FromValue returns a pointer to a valid GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456FromValue(v string) (*GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456, error) {
	ev := GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456: valid values are %v", v, allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456) IsValid() bool {
	for _, existing := range allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456 value.
func (v GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456) Ptr() *GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4446999456 {
	return &v
}

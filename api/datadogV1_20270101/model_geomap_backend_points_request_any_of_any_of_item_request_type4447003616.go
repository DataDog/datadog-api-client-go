// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616 Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616 string

// List of GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616.
const (
	GEOMAPBACKENDPOINTSREQUESTANYOFANYOFITEMREQUESTTYPE4447003616_DATA_PROJECTION GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616 = "data_projection"
)

var allowedGeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616EnumValues = []GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616{
	GEOMAPBACKENDPOINTSREQUESTANYOFANYOFITEMREQUESTTYPE4447003616_DATA_PROJECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616) GetAllowedValues() []GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616 {
	return allowedGeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616(value)
	return nil
}

// NewGeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616FromValue returns a pointer to a valid GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616FromValue(v string) (*GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616, error) {
	ev := GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616: valid values are %v", v, allowedGeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616) IsValid() bool {
	for _, existing := range allowedGeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616 value.
func (v GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616) Ptr() *GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616 {
	return &v
}

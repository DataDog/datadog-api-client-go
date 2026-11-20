// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376 Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376 string

// List of GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376.
const (
	GEOMAPBACKENDPOINTSREQUESTANYOFANYOFITEMRESPONSEFORMAT4447001376_TABULAR GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376 = "tabular"
)

var allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376EnumValues = []GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376{
	GEOMAPBACKENDPOINTSREQUESTANYOFANYOFITEMRESPONSEFORMAT4447001376_TABULAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376) GetAllowedValues() []GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376 {
	return allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376(value)
	return nil
}

// NewGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376FromValue returns a pointer to a valid GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376FromValue(v string) (*GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376, error) {
	ev := GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376: valid values are %v", v, allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376) IsValid() bool {
	for _, existing := range allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376 value.
func (v GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376) Ptr() *GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat4447001376 {
	return &v
}

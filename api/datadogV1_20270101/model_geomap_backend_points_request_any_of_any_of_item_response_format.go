// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat string

// List of GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat.
const (
	GEOMAPBACKENDPOINTSREQUESTANYOFANYOFITEMRESPONSEFORMAT_EVENT_LIST GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat = "event_list"
)

var allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormatEnumValues = []GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat{
	GEOMAPBACKENDPOINTSREQUESTANYOFANYOFITEMRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat) GetAllowedValues() []GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat {
	return allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat(value)
	return nil
}

// NewGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormatFromValue returns a pointer to a valid GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormatFromValue(v string) (*GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat, error) {
	ev := GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat: valid values are %v", v, allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat) IsValid() bool {
	for _, existing := range allowedGeomapBackendPointsRequestAnyOfAnyOfItemResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat value.
func (v GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat) Ptr() *GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat {
	return &v
}

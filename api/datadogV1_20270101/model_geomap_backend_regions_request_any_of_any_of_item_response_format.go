// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat string

// List of GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat.
const (
	GEOMAPBACKENDREGIONSREQUESTANYOFANYOFITEMRESPONSEFORMAT_SCALAR GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat = "scalar"
)

var allowedGeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormatEnumValues = []GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat{
	GEOMAPBACKENDREGIONSREQUESTANYOFANYOFITEMRESPONSEFORMAT_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat) GetAllowedValues() []GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat {
	return allowedGeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat(value)
	return nil
}

// NewGeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormatFromValue returns a pointer to a valid GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormatFromValue(v string) (*GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat, error) {
	ev := GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat: valid values are %v", v, allowedGeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat) IsValid() bool {
	for _, existing := range allowedGeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat value.
func (v GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat) Ptr() *GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat {
	return &v
}

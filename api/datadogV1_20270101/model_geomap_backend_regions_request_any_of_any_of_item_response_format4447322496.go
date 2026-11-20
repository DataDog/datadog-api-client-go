// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496 Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496 string

// List of GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496.
const (
	GEOMAPBACKENDREGIONSREQUESTANYOFANYOFITEMRESPONSEFORMAT4447322496_SCALAR GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496 = "scalar"
)

var allowedGeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496EnumValues = []GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496{
	GEOMAPBACKENDREGIONSREQUESTANYOFANYOFITEMRESPONSEFORMAT4447322496_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496) GetAllowedValues() []GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496 {
	return allowedGeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496(value)
	return nil
}

// NewGeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496FromValue returns a pointer to a valid GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496FromValue(v string) (*GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496, error) {
	ev := GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496: valid values are %v", v, allowedGeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496) IsValid() bool {
	for _, existing := range allowedGeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496 value.
func (v GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496) Ptr() *GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496 {
	return &v
}

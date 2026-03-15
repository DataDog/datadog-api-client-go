// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FacetType The JSON:API type for facets.
type FacetType string

// List of FacetType.
const (
	FACETTYPE_FACET FacetType = "facet"
)

var allowedFacetTypeEnumValues = []FacetType{
	FACETTYPE_FACET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FacetType) GetAllowedValues() []FacetType {
	return allowedFacetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FacetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FacetType(value)
	return nil
}

// NewFacetTypeFromValue returns a pointer to a valid FacetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFacetTypeFromValue(v string) (*FacetType, error) {
	ev := FacetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FacetType: valid values are %v", v, allowedFacetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FacetType) IsValid() bool {
	for _, existing := range allowedFacetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FacetType value.
func (v FacetType) Ptr() *FacetType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstLegendAnyOfAnyOfItemType Test-only generated schema for the dashboard versioning proof of concept.
type SunburstLegendAnyOfAnyOfItemType string

// List of SunburstLegendAnyOfAnyOfItemType.
const (
	SUNBURSTLEGENDANYOFANYOFITEMTYPE_AUTOMATIC SunburstLegendAnyOfAnyOfItemType = "automatic"
)

var allowedSunburstLegendAnyOfAnyOfItemTypeEnumValues = []SunburstLegendAnyOfAnyOfItemType{
	SUNBURSTLEGENDANYOFANYOFITEMTYPE_AUTOMATIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SunburstLegendAnyOfAnyOfItemType) GetAllowedValues() []SunburstLegendAnyOfAnyOfItemType {
	return allowedSunburstLegendAnyOfAnyOfItemTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SunburstLegendAnyOfAnyOfItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstLegendAnyOfAnyOfItemType(value)
	return nil
}

// NewSunburstLegendAnyOfAnyOfItemTypeFromValue returns a pointer to a valid SunburstLegendAnyOfAnyOfItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSunburstLegendAnyOfAnyOfItemTypeFromValue(v string) (*SunburstLegendAnyOfAnyOfItemType, error) {
	ev := SunburstLegendAnyOfAnyOfItemType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SunburstLegendAnyOfAnyOfItemType: valid values are %v", v, allowedSunburstLegendAnyOfAnyOfItemTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SunburstLegendAnyOfAnyOfItemType) IsValid() bool {
	for _, existing := range allowedSunburstLegendAnyOfAnyOfItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstLegendAnyOfAnyOfItemType value.
func (v SunburstLegendAnyOfAnyOfItemType) Ptr() *SunburstLegendAnyOfAnyOfItemType {
	return &v
}

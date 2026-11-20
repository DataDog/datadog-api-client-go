// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstKnownLegendAnyOfAnyOfItemType Test-only generated schema for the dashboard versioning proof of concept.
type SunburstKnownLegendAnyOfAnyOfItemType string

// List of SunburstKnownLegendAnyOfAnyOfItemType.
const (
	SUNBURSTKNOWNLEGENDANYOFANYOFITEMTYPE_TABLE SunburstKnownLegendAnyOfAnyOfItemType = "table"
)

var allowedSunburstKnownLegendAnyOfAnyOfItemTypeEnumValues = []SunburstKnownLegendAnyOfAnyOfItemType{
	SUNBURSTKNOWNLEGENDANYOFANYOFITEMTYPE_TABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SunburstKnownLegendAnyOfAnyOfItemType) GetAllowedValues() []SunburstKnownLegendAnyOfAnyOfItemType {
	return allowedSunburstKnownLegendAnyOfAnyOfItemTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SunburstKnownLegendAnyOfAnyOfItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstKnownLegendAnyOfAnyOfItemType(value)
	return nil
}

// NewSunburstKnownLegendAnyOfAnyOfItemTypeFromValue returns a pointer to a valid SunburstKnownLegendAnyOfAnyOfItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSunburstKnownLegendAnyOfAnyOfItemTypeFromValue(v string) (*SunburstKnownLegendAnyOfAnyOfItemType, error) {
	ev := SunburstKnownLegendAnyOfAnyOfItemType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SunburstKnownLegendAnyOfAnyOfItemType: valid values are %v", v, allowedSunburstKnownLegendAnyOfAnyOfItemTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SunburstKnownLegendAnyOfAnyOfItemType) IsValid() bool {
	for _, existing := range allowedSunburstKnownLegendAnyOfAnyOfItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstKnownLegendAnyOfAnyOfItemType value.
func (v SunburstKnownLegendAnyOfAnyOfItemType) Ptr() *SunburstKnownLegendAnyOfAnyOfItemType {
	return &v
}

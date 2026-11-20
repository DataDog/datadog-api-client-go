// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstKnownLegendAnyOfAnyOfItemType4452235232 Test-only generated schema for the dashboard versioning proof of concept.
type SunburstKnownLegendAnyOfAnyOfItemType4452235232 string

// List of SunburstKnownLegendAnyOfAnyOfItemType4452235232.
const (
	SUNBURSTKNOWNLEGENDANYOFANYOFITEMTYPE4452235232_INLINE SunburstKnownLegendAnyOfAnyOfItemType4452235232 = "inline"
)

var allowedSunburstKnownLegendAnyOfAnyOfItemType4452235232EnumValues = []SunburstKnownLegendAnyOfAnyOfItemType4452235232{
	SUNBURSTKNOWNLEGENDANYOFANYOFITEMTYPE4452235232_INLINE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SunburstKnownLegendAnyOfAnyOfItemType4452235232) GetAllowedValues() []SunburstKnownLegendAnyOfAnyOfItemType4452235232 {
	return allowedSunburstKnownLegendAnyOfAnyOfItemType4452235232EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SunburstKnownLegendAnyOfAnyOfItemType4452235232) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstKnownLegendAnyOfAnyOfItemType4452235232(value)
	return nil
}

// NewSunburstKnownLegendAnyOfAnyOfItemType4452235232FromValue returns a pointer to a valid SunburstKnownLegendAnyOfAnyOfItemType4452235232
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSunburstKnownLegendAnyOfAnyOfItemType4452235232FromValue(v string) (*SunburstKnownLegendAnyOfAnyOfItemType4452235232, error) {
	ev := SunburstKnownLegendAnyOfAnyOfItemType4452235232(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SunburstKnownLegendAnyOfAnyOfItemType4452235232: valid values are %v", v, allowedSunburstKnownLegendAnyOfAnyOfItemType4452235232EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SunburstKnownLegendAnyOfAnyOfItemType4452235232) IsValid() bool {
	for _, existing := range allowedSunburstKnownLegendAnyOfAnyOfItemType4452235232EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstKnownLegendAnyOfAnyOfItemType4452235232 value.
func (v SunburstKnownLegendAnyOfAnyOfItemType4452235232) Ptr() *SunburstKnownLegendAnyOfAnyOfItemType4452235232 {
	return &v
}

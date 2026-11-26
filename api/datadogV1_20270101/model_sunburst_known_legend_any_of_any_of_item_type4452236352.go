// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstKnownLegendAnyOfAnyOfItemType4452236352 Test-only generated schema for the dashboard versioning proof of concept.
type SunburstKnownLegendAnyOfAnyOfItemType4452236352 string

// List of SunburstKnownLegendAnyOfAnyOfItemType4452236352.
const (
	SUNBURSTKNOWNLEGENDANYOFANYOFITEMTYPE4452236352_NONE SunburstKnownLegendAnyOfAnyOfItemType4452236352 = "none"
)

var allowedSunburstKnownLegendAnyOfAnyOfItemType4452236352EnumValues = []SunburstKnownLegendAnyOfAnyOfItemType4452236352{
	SUNBURSTKNOWNLEGENDANYOFANYOFITEMTYPE4452236352_NONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SunburstKnownLegendAnyOfAnyOfItemType4452236352) GetAllowedValues() []SunburstKnownLegendAnyOfAnyOfItemType4452236352 {
	return allowedSunburstKnownLegendAnyOfAnyOfItemType4452236352EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SunburstKnownLegendAnyOfAnyOfItemType4452236352) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstKnownLegendAnyOfAnyOfItemType4452236352(value)
	return nil
}

// NewSunburstKnownLegendAnyOfAnyOfItemType4452236352FromValue returns a pointer to a valid SunburstKnownLegendAnyOfAnyOfItemType4452236352
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSunburstKnownLegendAnyOfAnyOfItemType4452236352FromValue(v string) (*SunburstKnownLegendAnyOfAnyOfItemType4452236352, error) {
	ev := SunburstKnownLegendAnyOfAnyOfItemType4452236352(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SunburstKnownLegendAnyOfAnyOfItemType4452236352: valid values are %v", v, allowedSunburstKnownLegendAnyOfAnyOfItemType4452236352EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SunburstKnownLegendAnyOfAnyOfItemType4452236352) IsValid() bool {
	for _, existing := range allowedSunburstKnownLegendAnyOfAnyOfItemType4452236352EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstKnownLegendAnyOfAnyOfItemType4452236352 value.
func (v SunburstKnownLegendAnyOfAnyOfItemType4452236352) Ptr() *SunburstKnownLegendAnyOfAnyOfItemType4452236352 {
	return &v
}

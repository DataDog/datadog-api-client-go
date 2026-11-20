// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ThresholdPalette Test-only generated schema for the dashboard versioning proof of concept.
type ThresholdPalette string

// List of ThresholdPalette.
const (
	THRESHOLDPALETTE_WHITE_ON_RED          ThresholdPalette = "white_on_red"
	THRESHOLDPALETTE_BLACK_ON_LIGHT_RED    ThresholdPalette = "black_on_light_red"
	THRESHOLDPALETTE_WHITE_ON_GREEN        ThresholdPalette = "white_on_green"
	THRESHOLDPALETTE_BLACK_ON_LIGHT_GREEN  ThresholdPalette = "black_on_light_green"
	THRESHOLDPALETTE_WHITE_ON_YELLOW       ThresholdPalette = "white_on_yellow"
	THRESHOLDPALETTE_BLACK_ON_LIGHT_YELLOW ThresholdPalette = "black_on_light_yellow"
	THRESHOLDPALETTE_WHITE_ON_GRAY         ThresholdPalette = "white_on_gray"
	THRESHOLDPALETTE_GREEN_ON_WHITE        ThresholdPalette = "green_on_white"
	THRESHOLDPALETTE_YELLOW_ON_WHITE       ThresholdPalette = "yellow_on_white"
	THRESHOLDPALETTE_RED_ON_WHITE          ThresholdPalette = "red_on_white"
	THRESHOLDPALETTE_GRAY_ON_WHITE         ThresholdPalette = "gray_on_white"
	THRESHOLDPALETTE_RED                   ThresholdPalette = "red"
	THRESHOLDPALETTE_GREEN                 ThresholdPalette = "green"
	THRESHOLDPALETTE_ORANGE                ThresholdPalette = "orange"
	THRESHOLDPALETTE_GRAY                  ThresholdPalette = "gray"
)

var allowedThresholdPaletteEnumValues = []ThresholdPalette{
	THRESHOLDPALETTE_WHITE_ON_RED,
	THRESHOLDPALETTE_BLACK_ON_LIGHT_RED,
	THRESHOLDPALETTE_WHITE_ON_GREEN,
	THRESHOLDPALETTE_BLACK_ON_LIGHT_GREEN,
	THRESHOLDPALETTE_WHITE_ON_YELLOW,
	THRESHOLDPALETTE_BLACK_ON_LIGHT_YELLOW,
	THRESHOLDPALETTE_WHITE_ON_GRAY,
	THRESHOLDPALETTE_GREEN_ON_WHITE,
	THRESHOLDPALETTE_YELLOW_ON_WHITE,
	THRESHOLDPALETTE_RED_ON_WHITE,
	THRESHOLDPALETTE_GRAY_ON_WHITE,
	THRESHOLDPALETTE_RED,
	THRESHOLDPALETTE_GREEN,
	THRESHOLDPALETTE_ORANGE,
	THRESHOLDPALETTE_GRAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ThresholdPalette) GetAllowedValues() []ThresholdPalette {
	return allowedThresholdPaletteEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ThresholdPalette) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ThresholdPalette(value)
	return nil
}

// NewThresholdPaletteFromValue returns a pointer to a valid ThresholdPalette
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewThresholdPaletteFromValue(v string) (*ThresholdPalette, error) {
	ev := ThresholdPalette(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ThresholdPalette: valid values are %v", v, allowedThresholdPaletteEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ThresholdPalette) IsValid() bool {
	for _, existing := range allowedThresholdPaletteEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ThresholdPalette value.
func (v ThresholdPalette) Ptr() *ThresholdPalette {
	return &v
}

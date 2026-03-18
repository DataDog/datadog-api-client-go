// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableThresholdPalette Color palette used by threshold-based conditional formatting and text formatting rules in a guided table.
type GuidedTableThresholdPalette string

// List of GuidedTableThresholdPalette.
const (
	GUIDEDTABLETHRESHOLDPALETTE_WHITE_ON_RED          GuidedTableThresholdPalette = "white_on_red"
	GUIDEDTABLETHRESHOLDPALETTE_BLACK_ON_LIGHT_RED    GuidedTableThresholdPalette = "black_on_light_red"
	GUIDEDTABLETHRESHOLDPALETTE_WHITE_ON_GREEN        GuidedTableThresholdPalette = "white_on_green"
	GUIDEDTABLETHRESHOLDPALETTE_BLACK_ON_LIGHT_GREEN  GuidedTableThresholdPalette = "black_on_light_green"
	GUIDEDTABLETHRESHOLDPALETTE_WHITE_ON_YELLOW       GuidedTableThresholdPalette = "white_on_yellow"
	GUIDEDTABLETHRESHOLDPALETTE_BLACK_ON_LIGHT_YELLOW GuidedTableThresholdPalette = "black_on_light_yellow"
	GUIDEDTABLETHRESHOLDPALETTE_WHITE_ON_GRAY         GuidedTableThresholdPalette = "white_on_gray"
	GUIDEDTABLETHRESHOLDPALETTE_GREEN_ON_WHITE        GuidedTableThresholdPalette = "green_on_white"
	GUIDEDTABLETHRESHOLDPALETTE_YELLOW_ON_WHITE       GuidedTableThresholdPalette = "yellow_on_white"
	GUIDEDTABLETHRESHOLDPALETTE_RED_ON_WHITE          GuidedTableThresholdPalette = "red_on_white"
	GUIDEDTABLETHRESHOLDPALETTE_GRAY_ON_WHITE         GuidedTableThresholdPalette = "gray_on_white"
	GUIDEDTABLETHRESHOLDPALETTE_RED                   GuidedTableThresholdPalette = "red"
	GUIDEDTABLETHRESHOLDPALETTE_GREEN                 GuidedTableThresholdPalette = "green"
	GUIDEDTABLETHRESHOLDPALETTE_ORANGE                GuidedTableThresholdPalette = "orange"
	GUIDEDTABLETHRESHOLDPALETTE_GRAY                  GuidedTableThresholdPalette = "gray"
	GUIDEDTABLETHRESHOLDPALETTE_CUSTOM_BG             GuidedTableThresholdPalette = "custom_bg"
	GUIDEDTABLETHRESHOLDPALETTE_CUSTOM_TEXT           GuidedTableThresholdPalette = "custom_text"
	GUIDEDTABLETHRESHOLDPALETTE_CUSTOM_IMAGE          GuidedTableThresholdPalette = "custom_image"
)

var allowedGuidedTableThresholdPaletteEnumValues = []GuidedTableThresholdPalette{
	GUIDEDTABLETHRESHOLDPALETTE_WHITE_ON_RED,
	GUIDEDTABLETHRESHOLDPALETTE_BLACK_ON_LIGHT_RED,
	GUIDEDTABLETHRESHOLDPALETTE_WHITE_ON_GREEN,
	GUIDEDTABLETHRESHOLDPALETTE_BLACK_ON_LIGHT_GREEN,
	GUIDEDTABLETHRESHOLDPALETTE_WHITE_ON_YELLOW,
	GUIDEDTABLETHRESHOLDPALETTE_BLACK_ON_LIGHT_YELLOW,
	GUIDEDTABLETHRESHOLDPALETTE_WHITE_ON_GRAY,
	GUIDEDTABLETHRESHOLDPALETTE_GREEN_ON_WHITE,
	GUIDEDTABLETHRESHOLDPALETTE_YELLOW_ON_WHITE,
	GUIDEDTABLETHRESHOLDPALETTE_RED_ON_WHITE,
	GUIDEDTABLETHRESHOLDPALETTE_GRAY_ON_WHITE,
	GUIDEDTABLETHRESHOLDPALETTE_RED,
	GUIDEDTABLETHRESHOLDPALETTE_GREEN,
	GUIDEDTABLETHRESHOLDPALETTE_ORANGE,
	GUIDEDTABLETHRESHOLDPALETTE_GRAY,
	GUIDEDTABLETHRESHOLDPALETTE_CUSTOM_BG,
	GUIDEDTABLETHRESHOLDPALETTE_CUSTOM_TEXT,
	GUIDEDTABLETHRESHOLDPALETTE_CUSTOM_IMAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GuidedTableThresholdPalette) GetAllowedValues() []GuidedTableThresholdPalette {
	return allowedGuidedTableThresholdPaletteEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GuidedTableThresholdPalette) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GuidedTableThresholdPalette(value)
	return nil
}

// NewGuidedTableThresholdPaletteFromValue returns a pointer to a valid GuidedTableThresholdPalette
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGuidedTableThresholdPaletteFromValue(v string) (*GuidedTableThresholdPalette, error) {
	ev := GuidedTableThresholdPalette(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GuidedTableThresholdPalette: valid values are %v", v, allowedGuidedTableThresholdPaletteEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GuidedTableThresholdPalette) IsValid() bool {
	for _, existing := range allowedGuidedTableThresholdPaletteEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuidedTableThresholdPalette value.
func (v GuidedTableThresholdPalette) Ptr() *GuidedTableThresholdPalette {
	return &v
}

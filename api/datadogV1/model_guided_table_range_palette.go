// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableRangePalette Color palette used by range-based conditional formatting rules.
type GuidedTableRangePalette string

// List of GuidedTableRangePalette.
const (
	GUIDEDTABLERANGEPALETTE_TRANSPARENT_TO_GREEN        GuidedTableRangePalette = "transparent_to_green"
	GUIDEDTABLERANGEPALETTE_TRANSPARENT_TO_ORANGE       GuidedTableRangePalette = "transparent_to_orange"
	GUIDEDTABLERANGEPALETTE_TRANSPARENT_TO_RED          GuidedTableRangePalette = "transparent_to_red"
	GUIDEDTABLERANGEPALETTE_TRANSPARENT_TO_BLUE         GuidedTableRangePalette = "transparent_to_blue"
	GUIDEDTABLERANGEPALETTE_RED_TRANSPARENT_GREEN       GuidedTableRangePalette = "red_transparent_green"
	GUIDEDTABLERANGEPALETTE_RED_TRANSPARENT_BLUE        GuidedTableRangePalette = "red_transparent_blue"
	GUIDEDTABLERANGEPALETTE_VIVID_TRANSPARENT_TO_GREEN  GuidedTableRangePalette = "vivid_transparent_to_green"
	GUIDEDTABLERANGEPALETTE_VIVID_TRANSPARENT_TO_ORANGE GuidedTableRangePalette = "vivid_transparent_to_orange"
	GUIDEDTABLERANGEPALETTE_VIVID_TRANSPARENT_TO_RED    GuidedTableRangePalette = "vivid_transparent_to_red"
	GUIDEDTABLERANGEPALETTE_VIVID_TRANSPARENT_TO_BLUE   GuidedTableRangePalette = "vivid_transparent_to_blue"
	GUIDEDTABLERANGEPALETTE_VIVID_RED_TRANSPARENT_GREEN GuidedTableRangePalette = "vivid_red_transparent_green"
	GUIDEDTABLERANGEPALETTE_VIVID_RED_TRANSPARENT_BLUE  GuidedTableRangePalette = "vivid_red_transparent_blue"
)

var allowedGuidedTableRangePaletteEnumValues = []GuidedTableRangePalette{
	GUIDEDTABLERANGEPALETTE_TRANSPARENT_TO_GREEN,
	GUIDEDTABLERANGEPALETTE_TRANSPARENT_TO_ORANGE,
	GUIDEDTABLERANGEPALETTE_TRANSPARENT_TO_RED,
	GUIDEDTABLERANGEPALETTE_TRANSPARENT_TO_BLUE,
	GUIDEDTABLERANGEPALETTE_RED_TRANSPARENT_GREEN,
	GUIDEDTABLERANGEPALETTE_RED_TRANSPARENT_BLUE,
	GUIDEDTABLERANGEPALETTE_VIVID_TRANSPARENT_TO_GREEN,
	GUIDEDTABLERANGEPALETTE_VIVID_TRANSPARENT_TO_ORANGE,
	GUIDEDTABLERANGEPALETTE_VIVID_TRANSPARENT_TO_RED,
	GUIDEDTABLERANGEPALETTE_VIVID_TRANSPARENT_TO_BLUE,
	GUIDEDTABLERANGEPALETTE_VIVID_RED_TRANSPARENT_GREEN,
	GUIDEDTABLERANGEPALETTE_VIVID_RED_TRANSPARENT_BLUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GuidedTableRangePalette) GetAllowedValues() []GuidedTableRangePalette {
	return allowedGuidedTableRangePaletteEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GuidedTableRangePalette) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GuidedTableRangePalette(value)
	return nil
}

// NewGuidedTableRangePaletteFromValue returns a pointer to a valid GuidedTableRangePalette
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGuidedTableRangePaletteFromValue(v string) (*GuidedTableRangePalette, error) {
	ev := GuidedTableRangePalette(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GuidedTableRangePalette: valid values are %v", v, allowedGuidedTableRangePaletteEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GuidedTableRangePalette) IsValid() bool {
	for _, existing := range allowedGuidedTableRangePaletteEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuidedTableRangePalette value.
func (v GuidedTableRangePalette) Ptr() *GuidedTableRangePalette {
	return &v
}

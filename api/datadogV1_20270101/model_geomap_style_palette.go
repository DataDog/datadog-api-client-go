// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapStylePalette Test-only generated schema for the dashboard versioning proof of concept.
type GeomapStylePalette string

// List of GeomapStylePalette.
const (
	GEOMAPSTYLEPALETTE_DOG_CLASSIC              GeomapStylePalette = "dog_classic"
	GEOMAPSTYLEPALETTE_DOG_CLASSIC_AREA         GeomapStylePalette = "dog_classic_area"
	GEOMAPSTYLEPALETTE_DOG_CLASSIC_LINE         GeomapStylePalette = "dog_classic_line"
	GEOMAPSTYLEPALETTE_DOG_CLASSIC_BARS         GeomapStylePalette = "dog_classic_bars"
	GEOMAPSTYLEPALETTE_DOG_HIGHLIGHT            GeomapStylePalette = "dog_highlight"
	GEOMAPSTYLEPALETTE_MONITOR_STATE            GeomapStylePalette = "monitor_state"
	GEOMAPSTYLEPALETTE_YELLOW_TO_GREEN          GeomapStylePalette = "yellow_to_green"
	GEOMAPSTYLEPALETTE_GREEN_TO_ORANGE          GeomapStylePalette = "green_to_orange"
	GEOMAPSTYLEPALETTE_HOSTMAP_BLUES            GeomapStylePalette = "hostmap_blues"
	GEOMAPSTYLEPALETTE_HOSTMAP_GREEN_RED        GeomapStylePalette = "hostmap_green_red"
	GEOMAPSTYLEPALETTE_HOSTMAP_GREEN_ORANGE_RED GeomapStylePalette = "hostmap_green_orange_red"
	GEOMAPSTYLEPALETTE_HOSTMAP_BLUE_RED         GeomapStylePalette = "hostmap_blue_red"
	GEOMAPSTYLEPALETTE_DARK_GREY                GeomapStylePalette = "dark_grey"
	GEOMAPSTYLEPALETTE_PRIDE                    GeomapStylePalette = "pride"
	GEOMAPSTYLEPALETTE_GNBU                     GeomapStylePalette = "GnBu"
	GEOMAPSTYLEPALETTE_YLORRD                   GeomapStylePalette = "YlOrRd"
	GEOMAPSTYLEPALETTE_PURPLES                  GeomapStylePalette = "Purples"
	GEOMAPSTYLEPALETTE_BLUES                    GeomapStylePalette = "Blues"
	GEOMAPSTYLEPALETTE_ORANGES                  GeomapStylePalette = "Oranges"
	GEOMAPSTYLEPALETTE_REDS                     GeomapStylePalette = "Reds"
	GEOMAPSTYLEPALETTE_GREENS                   GeomapStylePalette = "Greens"
	GEOMAPSTYLEPALETTE_GREYS                    GeomapStylePalette = "Greys"
	GEOMAPSTYLEPALETTE_STATUS_ERRORS            GeomapStylePalette = "status_errors"
	GEOMAPSTYLEPALETTE_STATUS_WARNINGS          GeomapStylePalette = "status_warnings"
	GEOMAPSTYLEPALETTE_STATUS_SUCCESSES         GeomapStylePalette = "status_successes"
	GEOMAPSTYLEPALETTE_PLASMA                   GeomapStylePalette = "Plasma"
	GEOMAPSTYLEPALETTE_VIRIDIS                  GeomapStylePalette = "Viridis"
	GEOMAPSTYLEPALETTE_DOGCAT                   GeomapStylePalette = "dogcat"
	GEOMAPSTYLEPALETTE_DATADOG16                GeomapStylePalette = "datadog16"
	GEOMAPSTYLEPALETTE_PA_CLASSIC_LINE          GeomapStylePalette = "pa_classic_line"
	GEOMAPSTYLEPALETTE_PA_CLASSIC_AREA          GeomapStylePalette = "pa_classic_area"
	GEOMAPSTYLEPALETTE_PA_CLASSIC_BARS          GeomapStylePalette = "pa_classic_bars"
	GEOMAPSTYLEPALETTE_PA_DATADOG16             GeomapStylePalette = "pa_datadog16"
)

var allowedGeomapStylePaletteEnumValues = []GeomapStylePalette{
	GEOMAPSTYLEPALETTE_DOG_CLASSIC,
	GEOMAPSTYLEPALETTE_DOG_CLASSIC_AREA,
	GEOMAPSTYLEPALETTE_DOG_CLASSIC_LINE,
	GEOMAPSTYLEPALETTE_DOG_CLASSIC_BARS,
	GEOMAPSTYLEPALETTE_DOG_HIGHLIGHT,
	GEOMAPSTYLEPALETTE_MONITOR_STATE,
	GEOMAPSTYLEPALETTE_YELLOW_TO_GREEN,
	GEOMAPSTYLEPALETTE_GREEN_TO_ORANGE,
	GEOMAPSTYLEPALETTE_HOSTMAP_BLUES,
	GEOMAPSTYLEPALETTE_HOSTMAP_GREEN_RED,
	GEOMAPSTYLEPALETTE_HOSTMAP_GREEN_ORANGE_RED,
	GEOMAPSTYLEPALETTE_HOSTMAP_BLUE_RED,
	GEOMAPSTYLEPALETTE_DARK_GREY,
	GEOMAPSTYLEPALETTE_PRIDE,
	GEOMAPSTYLEPALETTE_GNBU,
	GEOMAPSTYLEPALETTE_YLORRD,
	GEOMAPSTYLEPALETTE_PURPLES,
	GEOMAPSTYLEPALETTE_BLUES,
	GEOMAPSTYLEPALETTE_ORANGES,
	GEOMAPSTYLEPALETTE_REDS,
	GEOMAPSTYLEPALETTE_GREENS,
	GEOMAPSTYLEPALETTE_GREYS,
	GEOMAPSTYLEPALETTE_STATUS_ERRORS,
	GEOMAPSTYLEPALETTE_STATUS_WARNINGS,
	GEOMAPSTYLEPALETTE_STATUS_SUCCESSES,
	GEOMAPSTYLEPALETTE_PLASMA,
	GEOMAPSTYLEPALETTE_VIRIDIS,
	GEOMAPSTYLEPALETTE_DOGCAT,
	GEOMAPSTYLEPALETTE_DATADOG16,
	GEOMAPSTYLEPALETTE_PA_CLASSIC_LINE,
	GEOMAPSTYLEPALETTE_PA_CLASSIC_AREA,
	GEOMAPSTYLEPALETTE_PA_CLASSIC_BARS,
	GEOMAPSTYLEPALETTE_PA_DATADOG16,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeomapStylePalette) GetAllowedValues() []GeomapStylePalette {
	return allowedGeomapStylePaletteEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeomapStylePalette) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeomapStylePalette(value)
	return nil
}

// NewGeomapStylePaletteFromValue returns a pointer to a valid GeomapStylePalette
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeomapStylePaletteFromValue(v string) (*GeomapStylePalette, error) {
	ev := GeomapStylePalette(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeomapStylePalette: valid values are %v", v, allowedGeomapStylePaletteEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeomapStylePalette) IsValid() bool {
	for _, existing := range allowedGeomapStylePaletteEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeomapStylePalette value.
func (v GeomapStylePalette) Ptr() *GeomapStylePalette {
	return &v
}

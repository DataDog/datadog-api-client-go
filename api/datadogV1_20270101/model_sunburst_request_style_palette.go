// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstRequestStylePalette Test-only generated schema for the dashboard versioning proof of concept.
type SunburstRequestStylePalette string

// List of SunburstRequestStylePalette.
const (
	SUNBURSTREQUESTSTYLEPALETTE_SUCCESS                  SunburstRequestStylePalette = "success"
	SUNBURSTREQUESTSTYLEPALETTE_RED                      SunburstRequestStylePalette = "red"
	SUNBURSTREQUESTSTYLEPALETTE_GREEN                    SunburstRequestStylePalette = "green"
	SUNBURSTREQUESTSTYLEPALETTE_ORANGE                   SunburstRequestStylePalette = "orange"
	SUNBURSTREQUESTSTYLEPALETTE_GRAY                     SunburstRequestStylePalette = "gray"
	SUNBURSTREQUESTSTYLEPALETTE_DOG_CLASSIC              SunburstRequestStylePalette = "dog_classic"
	SUNBURSTREQUESTSTYLEPALETTE_DOG_CLASSIC_AREA         SunburstRequestStylePalette = "dog_classic_area"
	SUNBURSTREQUESTSTYLEPALETTE_DOG_CLASSIC_LINE         SunburstRequestStylePalette = "dog_classic_line"
	SUNBURSTREQUESTSTYLEPALETTE_DOG_CLASSIC_BARS         SunburstRequestStylePalette = "dog_classic_bars"
	SUNBURSTREQUESTSTYLEPALETTE_DOG_HIGHLIGHT            SunburstRequestStylePalette = "dog_highlight"
	SUNBURSTREQUESTSTYLEPALETTE_MONITOR_STATE            SunburstRequestStylePalette = "monitor_state"
	SUNBURSTREQUESTSTYLEPALETTE_YELLOW_TO_GREEN          SunburstRequestStylePalette = "yellow_to_green"
	SUNBURSTREQUESTSTYLEPALETTE_GREEN_TO_ORANGE          SunburstRequestStylePalette = "green_to_orange"
	SUNBURSTREQUESTSTYLEPALETTE_HOSTMAP_BLUES            SunburstRequestStylePalette = "hostmap_blues"
	SUNBURSTREQUESTSTYLEPALETTE_HOSTMAP_GREEN_RED        SunburstRequestStylePalette = "hostmap_green_red"
	SUNBURSTREQUESTSTYLEPALETTE_HOSTMAP_GREEN_ORANGE_RED SunburstRequestStylePalette = "hostmap_green_orange_red"
	SUNBURSTREQUESTSTYLEPALETTE_HOSTMAP_BLUE_RED         SunburstRequestStylePalette = "hostmap_blue_red"
	SUNBURSTREQUESTSTYLEPALETTE_DARK_GREY                SunburstRequestStylePalette = "dark_grey"
	SUNBURSTREQUESTSTYLEPALETTE_PRIDE                    SunburstRequestStylePalette = "pride"
	SUNBURSTREQUESTSTYLEPALETTE_GNBU                     SunburstRequestStylePalette = "GnBu"
	SUNBURSTREQUESTSTYLEPALETTE_YLORRD                   SunburstRequestStylePalette = "YlOrRd"
	SUNBURSTREQUESTSTYLEPALETTE_PURPLES                  SunburstRequestStylePalette = "Purples"
	SUNBURSTREQUESTSTYLEPALETTE_BLUES                    SunburstRequestStylePalette = "Blues"
	SUNBURSTREQUESTSTYLEPALETTE_ORANGES                  SunburstRequestStylePalette = "Oranges"
	SUNBURSTREQUESTSTYLEPALETTE_REDS                     SunburstRequestStylePalette = "Reds"
	SUNBURSTREQUESTSTYLEPALETTE_GREENS                   SunburstRequestStylePalette = "Greens"
	SUNBURSTREQUESTSTYLEPALETTE_GREYS                    SunburstRequestStylePalette = "Greys"
	SUNBURSTREQUESTSTYLEPALETTE_STATUS_ERRORS            SunburstRequestStylePalette = "status_errors"
	SUNBURSTREQUESTSTYLEPALETTE_STATUS_WARNINGS          SunburstRequestStylePalette = "status_warnings"
	SUNBURSTREQUESTSTYLEPALETTE_STATUS_SUCCESSES         SunburstRequestStylePalette = "status_successes"
	SUNBURSTREQUESTSTYLEPALETTE_PLASMA                   SunburstRequestStylePalette = "Plasma"
	SUNBURSTREQUESTSTYLEPALETTE_VIRIDIS                  SunburstRequestStylePalette = "Viridis"
	SUNBURSTREQUESTSTYLEPALETTE_DOGCAT                   SunburstRequestStylePalette = "dogcat"
	SUNBURSTREQUESTSTYLEPALETTE_DATADOG16                SunburstRequestStylePalette = "datadog16"
	SUNBURSTREQUESTSTYLEPALETTE_PA_CLASSIC_LINE          SunburstRequestStylePalette = "pa_classic_line"
	SUNBURSTREQUESTSTYLEPALETTE_PA_CLASSIC_AREA          SunburstRequestStylePalette = "pa_classic_area"
	SUNBURSTREQUESTSTYLEPALETTE_PA_CLASSIC_BARS          SunburstRequestStylePalette = "pa_classic_bars"
	SUNBURSTREQUESTSTYLEPALETTE_PA_DATADOG16             SunburstRequestStylePalette = "pa_datadog16"
	SUNBURSTREQUESTSTYLEPALETTE_CLASSIC                  SunburstRequestStylePalette = "classic"
	SUNBURSTREQUESTSTYLEPALETTE_COOL_HM                  SunburstRequestStylePalette = "Cool_hm"
	SUNBURSTREQUESTSTYLEPALETTE_WARM_HM                  SunburstRequestStylePalette = "Warm_hm"
	SUNBURSTREQUESTSTYLEPALETTE_YELLOW___GREEN           SunburstRequestStylePalette = "Yellow / Green"
	SUNBURSTREQUESTSTYLEPALETTE_GREEN___ORANGE           SunburstRequestStylePalette = "Green / Orange"
	SUNBURSTREQUESTSTYLEPALETTE_GREEN___RED              SunburstRequestStylePalette = "Green / Red"
	SUNBURSTREQUESTSTYLEPALETTE_GREEN___ORANGE___RED     SunburstRequestStylePalette = "Green / Orange / Red"
	SUNBURSTREQUESTSTYLEPALETTE_BLUE___RED               SunburstRequestStylePalette = "Blue / Red"
	SUNBURSTREQUESTSTYLEPALETTE_GOOD___BAD               SunburstRequestStylePalette = "Good - bad"
	SUNBURSTREQUESTSTYLEPALETTE_COOL                     SunburstRequestStylePalette = "cool"
	SUNBURSTREQUESTSTYLEPALETTE_OCEAN                    SunburstRequestStylePalette = "ocean"
	SUNBURSTREQUESTSTYLEPALETTE_WARM                     SunburstRequestStylePalette = "warm"
	SUNBURSTREQUESTSTYLEPALETTE_FIRE                     SunburstRequestStylePalette = "fire"
	SUNBURSTREQUESTSTYLEPALETTE_WARN                     SunburstRequestStylePalette = "warn"
	SUNBURSTREQUESTSTYLEPALETTE_PURPLE                   SunburstRequestStylePalette = "purple"
	SUNBURSTREQUESTSTYLEPALETTE_GREY                     SunburstRequestStylePalette = "grey"
	SUNBURSTREQUESTSTYLEPALETTE_BLUE                     SunburstRequestStylePalette = "blue"
	SUNBURSTREQUESTSTYLEPALETTE_DD20                     SunburstRequestStylePalette = "dd20"
	SUNBURSTREQUESTSTYLEPALETTE_SEMANTIC                 SunburstRequestStylePalette = "semantic"
)

var allowedSunburstRequestStylePaletteEnumValues = []SunburstRequestStylePalette{
	SUNBURSTREQUESTSTYLEPALETTE_SUCCESS,
	SUNBURSTREQUESTSTYLEPALETTE_RED,
	SUNBURSTREQUESTSTYLEPALETTE_GREEN,
	SUNBURSTREQUESTSTYLEPALETTE_ORANGE,
	SUNBURSTREQUESTSTYLEPALETTE_GRAY,
	SUNBURSTREQUESTSTYLEPALETTE_DOG_CLASSIC,
	SUNBURSTREQUESTSTYLEPALETTE_DOG_CLASSIC_AREA,
	SUNBURSTREQUESTSTYLEPALETTE_DOG_CLASSIC_LINE,
	SUNBURSTREQUESTSTYLEPALETTE_DOG_CLASSIC_BARS,
	SUNBURSTREQUESTSTYLEPALETTE_DOG_HIGHLIGHT,
	SUNBURSTREQUESTSTYLEPALETTE_MONITOR_STATE,
	SUNBURSTREQUESTSTYLEPALETTE_YELLOW_TO_GREEN,
	SUNBURSTREQUESTSTYLEPALETTE_GREEN_TO_ORANGE,
	SUNBURSTREQUESTSTYLEPALETTE_HOSTMAP_BLUES,
	SUNBURSTREQUESTSTYLEPALETTE_HOSTMAP_GREEN_RED,
	SUNBURSTREQUESTSTYLEPALETTE_HOSTMAP_GREEN_ORANGE_RED,
	SUNBURSTREQUESTSTYLEPALETTE_HOSTMAP_BLUE_RED,
	SUNBURSTREQUESTSTYLEPALETTE_DARK_GREY,
	SUNBURSTREQUESTSTYLEPALETTE_PRIDE,
	SUNBURSTREQUESTSTYLEPALETTE_GNBU,
	SUNBURSTREQUESTSTYLEPALETTE_YLORRD,
	SUNBURSTREQUESTSTYLEPALETTE_PURPLES,
	SUNBURSTREQUESTSTYLEPALETTE_BLUES,
	SUNBURSTREQUESTSTYLEPALETTE_ORANGES,
	SUNBURSTREQUESTSTYLEPALETTE_REDS,
	SUNBURSTREQUESTSTYLEPALETTE_GREENS,
	SUNBURSTREQUESTSTYLEPALETTE_GREYS,
	SUNBURSTREQUESTSTYLEPALETTE_STATUS_ERRORS,
	SUNBURSTREQUESTSTYLEPALETTE_STATUS_WARNINGS,
	SUNBURSTREQUESTSTYLEPALETTE_STATUS_SUCCESSES,
	SUNBURSTREQUESTSTYLEPALETTE_PLASMA,
	SUNBURSTREQUESTSTYLEPALETTE_VIRIDIS,
	SUNBURSTREQUESTSTYLEPALETTE_DOGCAT,
	SUNBURSTREQUESTSTYLEPALETTE_DATADOG16,
	SUNBURSTREQUESTSTYLEPALETTE_PA_CLASSIC_LINE,
	SUNBURSTREQUESTSTYLEPALETTE_PA_CLASSIC_AREA,
	SUNBURSTREQUESTSTYLEPALETTE_PA_CLASSIC_BARS,
	SUNBURSTREQUESTSTYLEPALETTE_PA_DATADOG16,
	SUNBURSTREQUESTSTYLEPALETTE_CLASSIC,
	SUNBURSTREQUESTSTYLEPALETTE_COOL_HM,
	SUNBURSTREQUESTSTYLEPALETTE_WARM_HM,
	SUNBURSTREQUESTSTYLEPALETTE_YELLOW___GREEN,
	SUNBURSTREQUESTSTYLEPALETTE_GREEN___ORANGE,
	SUNBURSTREQUESTSTYLEPALETTE_GREEN___RED,
	SUNBURSTREQUESTSTYLEPALETTE_GREEN___ORANGE___RED,
	SUNBURSTREQUESTSTYLEPALETTE_BLUE___RED,
	SUNBURSTREQUESTSTYLEPALETTE_GOOD___BAD,
	SUNBURSTREQUESTSTYLEPALETTE_COOL,
	SUNBURSTREQUESTSTYLEPALETTE_OCEAN,
	SUNBURSTREQUESTSTYLEPALETTE_WARM,
	SUNBURSTREQUESTSTYLEPALETTE_FIRE,
	SUNBURSTREQUESTSTYLEPALETTE_WARN,
	SUNBURSTREQUESTSTYLEPALETTE_PURPLE,
	SUNBURSTREQUESTSTYLEPALETTE_GREY,
	SUNBURSTREQUESTSTYLEPALETTE_BLUE,
	SUNBURSTREQUESTSTYLEPALETTE_DD20,
	SUNBURSTREQUESTSTYLEPALETTE_SEMANTIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SunburstRequestStylePalette) GetAllowedValues() []SunburstRequestStylePalette {
	return allowedSunburstRequestStylePaletteEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SunburstRequestStylePalette) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstRequestStylePalette(value)
	return nil
}

// NewSunburstRequestStylePaletteFromValue returns a pointer to a valid SunburstRequestStylePalette
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSunburstRequestStylePaletteFromValue(v string) (*SunburstRequestStylePalette, error) {
	ev := SunburstRequestStylePalette(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SunburstRequestStylePalette: valid values are %v", v, allowedSunburstRequestStylePaletteEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SunburstRequestStylePalette) IsValid() bool {
	for _, existing := range allowedSunburstRequestStylePaletteEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstRequestStylePalette value.
func (v SunburstRequestStylePalette) Ptr() *SunburstRequestStylePalette {
	return &v
}

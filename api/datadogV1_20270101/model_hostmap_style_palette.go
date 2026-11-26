// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapStylePalette Test-only generated schema for the dashboard versioning proof of concept.
type HostmapStylePalette string

// List of HostmapStylePalette.
const (
	HOSTMAPSTYLEPALETTE_SUCCESS                  HostmapStylePalette = "success"
	HOSTMAPSTYLEPALETTE_RED                      HostmapStylePalette = "red"
	HOSTMAPSTYLEPALETTE_GREEN                    HostmapStylePalette = "green"
	HOSTMAPSTYLEPALETTE_ORANGE                   HostmapStylePalette = "orange"
	HOSTMAPSTYLEPALETTE_GRAY                     HostmapStylePalette = "gray"
	HOSTMAPSTYLEPALETTE_DOG_CLASSIC              HostmapStylePalette = "dog_classic"
	HOSTMAPSTYLEPALETTE_DOG_CLASSIC_AREA         HostmapStylePalette = "dog_classic_area"
	HOSTMAPSTYLEPALETTE_DOG_CLASSIC_LINE         HostmapStylePalette = "dog_classic_line"
	HOSTMAPSTYLEPALETTE_DOG_CLASSIC_BARS         HostmapStylePalette = "dog_classic_bars"
	HOSTMAPSTYLEPALETTE_DOG_HIGHLIGHT            HostmapStylePalette = "dog_highlight"
	HOSTMAPSTYLEPALETTE_MONITOR_STATE            HostmapStylePalette = "monitor_state"
	HOSTMAPSTYLEPALETTE_YELLOW_TO_GREEN          HostmapStylePalette = "yellow_to_green"
	HOSTMAPSTYLEPALETTE_GREEN_TO_ORANGE          HostmapStylePalette = "green_to_orange"
	HOSTMAPSTYLEPALETTE_HOSTMAP_BLUES            HostmapStylePalette = "hostmap_blues"
	HOSTMAPSTYLEPALETTE_HOSTMAP_GREEN_RED        HostmapStylePalette = "hostmap_green_red"
	HOSTMAPSTYLEPALETTE_HOSTMAP_GREEN_ORANGE_RED HostmapStylePalette = "hostmap_green_orange_red"
	HOSTMAPSTYLEPALETTE_HOSTMAP_BLUE_RED         HostmapStylePalette = "hostmap_blue_red"
	HOSTMAPSTYLEPALETTE_DARK_GREY                HostmapStylePalette = "dark_grey"
	HOSTMAPSTYLEPALETTE_PRIDE                    HostmapStylePalette = "pride"
	HOSTMAPSTYLEPALETTE_GNBU                     HostmapStylePalette = "GnBu"
	HOSTMAPSTYLEPALETTE_YLORRD                   HostmapStylePalette = "YlOrRd"
	HOSTMAPSTYLEPALETTE_PURPLES                  HostmapStylePalette = "Purples"
	HOSTMAPSTYLEPALETTE_BLUES                    HostmapStylePalette = "Blues"
	HOSTMAPSTYLEPALETTE_ORANGES                  HostmapStylePalette = "Oranges"
	HOSTMAPSTYLEPALETTE_REDS                     HostmapStylePalette = "Reds"
	HOSTMAPSTYLEPALETTE_GREENS                   HostmapStylePalette = "Greens"
	HOSTMAPSTYLEPALETTE_GREYS                    HostmapStylePalette = "Greys"
	HOSTMAPSTYLEPALETTE_STATUS_ERRORS            HostmapStylePalette = "status_errors"
	HOSTMAPSTYLEPALETTE_STATUS_WARNINGS          HostmapStylePalette = "status_warnings"
	HOSTMAPSTYLEPALETTE_STATUS_SUCCESSES         HostmapStylePalette = "status_successes"
	HOSTMAPSTYLEPALETTE_PLASMA                   HostmapStylePalette = "Plasma"
	HOSTMAPSTYLEPALETTE_VIRIDIS                  HostmapStylePalette = "Viridis"
	HOSTMAPSTYLEPALETTE_DOGCAT                   HostmapStylePalette = "dogcat"
	HOSTMAPSTYLEPALETTE_DATADOG16                HostmapStylePalette = "datadog16"
	HOSTMAPSTYLEPALETTE_PA_CLASSIC_LINE          HostmapStylePalette = "pa_classic_line"
	HOSTMAPSTYLEPALETTE_PA_CLASSIC_AREA          HostmapStylePalette = "pa_classic_area"
	HOSTMAPSTYLEPALETTE_PA_CLASSIC_BARS          HostmapStylePalette = "pa_classic_bars"
	HOSTMAPSTYLEPALETTE_PA_DATADOG16             HostmapStylePalette = "pa_datadog16"
	HOSTMAPSTYLEPALETTE_CLASSIC                  HostmapStylePalette = "classic"
	HOSTMAPSTYLEPALETTE_COOL_HM                  HostmapStylePalette = "Cool_hm"
	HOSTMAPSTYLEPALETTE_WARM_HM                  HostmapStylePalette = "Warm_hm"
	HOSTMAPSTYLEPALETTE_YELLOW___GREEN           HostmapStylePalette = "Yellow / Green"
	HOSTMAPSTYLEPALETTE_GREEN___ORANGE           HostmapStylePalette = "Green / Orange"
	HOSTMAPSTYLEPALETTE_GREEN___RED              HostmapStylePalette = "Green / Red"
	HOSTMAPSTYLEPALETTE_GREEN___ORANGE___RED     HostmapStylePalette = "Green / Orange / Red"
	HOSTMAPSTYLEPALETTE_BLUE___RED               HostmapStylePalette = "Blue / Red"
	HOSTMAPSTYLEPALETTE_GOOD___BAD               HostmapStylePalette = "Good - bad"
	HOSTMAPSTYLEPALETTE_COOL                     HostmapStylePalette = "cool"
	HOSTMAPSTYLEPALETTE_OCEAN                    HostmapStylePalette = "ocean"
	HOSTMAPSTYLEPALETTE_WARM                     HostmapStylePalette = "warm"
	HOSTMAPSTYLEPALETTE_FIRE                     HostmapStylePalette = "fire"
	HOSTMAPSTYLEPALETTE_WARN                     HostmapStylePalette = "warn"
	HOSTMAPSTYLEPALETTE_PURPLE                   HostmapStylePalette = "purple"
	HOSTMAPSTYLEPALETTE_GREY                     HostmapStylePalette = "grey"
	HOSTMAPSTYLEPALETTE_BLUE                     HostmapStylePalette = "blue"
	HOSTMAPSTYLEPALETTE_DD20                     HostmapStylePalette = "dd20"
)

var allowedHostmapStylePaletteEnumValues = []HostmapStylePalette{
	HOSTMAPSTYLEPALETTE_SUCCESS,
	HOSTMAPSTYLEPALETTE_RED,
	HOSTMAPSTYLEPALETTE_GREEN,
	HOSTMAPSTYLEPALETTE_ORANGE,
	HOSTMAPSTYLEPALETTE_GRAY,
	HOSTMAPSTYLEPALETTE_DOG_CLASSIC,
	HOSTMAPSTYLEPALETTE_DOG_CLASSIC_AREA,
	HOSTMAPSTYLEPALETTE_DOG_CLASSIC_LINE,
	HOSTMAPSTYLEPALETTE_DOG_CLASSIC_BARS,
	HOSTMAPSTYLEPALETTE_DOG_HIGHLIGHT,
	HOSTMAPSTYLEPALETTE_MONITOR_STATE,
	HOSTMAPSTYLEPALETTE_YELLOW_TO_GREEN,
	HOSTMAPSTYLEPALETTE_GREEN_TO_ORANGE,
	HOSTMAPSTYLEPALETTE_HOSTMAP_BLUES,
	HOSTMAPSTYLEPALETTE_HOSTMAP_GREEN_RED,
	HOSTMAPSTYLEPALETTE_HOSTMAP_GREEN_ORANGE_RED,
	HOSTMAPSTYLEPALETTE_HOSTMAP_BLUE_RED,
	HOSTMAPSTYLEPALETTE_DARK_GREY,
	HOSTMAPSTYLEPALETTE_PRIDE,
	HOSTMAPSTYLEPALETTE_GNBU,
	HOSTMAPSTYLEPALETTE_YLORRD,
	HOSTMAPSTYLEPALETTE_PURPLES,
	HOSTMAPSTYLEPALETTE_BLUES,
	HOSTMAPSTYLEPALETTE_ORANGES,
	HOSTMAPSTYLEPALETTE_REDS,
	HOSTMAPSTYLEPALETTE_GREENS,
	HOSTMAPSTYLEPALETTE_GREYS,
	HOSTMAPSTYLEPALETTE_STATUS_ERRORS,
	HOSTMAPSTYLEPALETTE_STATUS_WARNINGS,
	HOSTMAPSTYLEPALETTE_STATUS_SUCCESSES,
	HOSTMAPSTYLEPALETTE_PLASMA,
	HOSTMAPSTYLEPALETTE_VIRIDIS,
	HOSTMAPSTYLEPALETTE_DOGCAT,
	HOSTMAPSTYLEPALETTE_DATADOG16,
	HOSTMAPSTYLEPALETTE_PA_CLASSIC_LINE,
	HOSTMAPSTYLEPALETTE_PA_CLASSIC_AREA,
	HOSTMAPSTYLEPALETTE_PA_CLASSIC_BARS,
	HOSTMAPSTYLEPALETTE_PA_DATADOG16,
	HOSTMAPSTYLEPALETTE_CLASSIC,
	HOSTMAPSTYLEPALETTE_COOL_HM,
	HOSTMAPSTYLEPALETTE_WARM_HM,
	HOSTMAPSTYLEPALETTE_YELLOW___GREEN,
	HOSTMAPSTYLEPALETTE_GREEN___ORANGE,
	HOSTMAPSTYLEPALETTE_GREEN___RED,
	HOSTMAPSTYLEPALETTE_GREEN___ORANGE___RED,
	HOSTMAPSTYLEPALETTE_BLUE___RED,
	HOSTMAPSTYLEPALETTE_GOOD___BAD,
	HOSTMAPSTYLEPALETTE_COOL,
	HOSTMAPSTYLEPALETTE_OCEAN,
	HOSTMAPSTYLEPALETTE_WARM,
	HOSTMAPSTYLEPALETTE_FIRE,
	HOSTMAPSTYLEPALETTE_WARN,
	HOSTMAPSTYLEPALETTE_PURPLE,
	HOSTMAPSTYLEPALETTE_GREY,
	HOSTMAPSTYLEPALETTE_BLUE,
	HOSTMAPSTYLEPALETTE_DD20,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HostmapStylePalette) GetAllowedValues() []HostmapStylePalette {
	return allowedHostmapStylePaletteEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostmapStylePalette) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostmapStylePalette(value)
	return nil
}

// NewHostmapStylePaletteFromValue returns a pointer to a valid HostmapStylePalette
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostmapStylePaletteFromValue(v string) (*HostmapStylePalette, error) {
	ev := HostmapStylePalette(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostmapStylePalette: valid values are %v", v, allowedHostmapStylePaletteEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostmapStylePalette) IsValid() bool {
	for _, existing := range allowedHostmapStylePaletteEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostmapStylePalette value.
func (v HostmapStylePalette) Ptr() *HostmapStylePalette {
	return &v
}

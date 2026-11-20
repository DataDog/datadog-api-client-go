// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapRequestStylePalette Test-only generated schema for the dashboard versioning proof of concept.
type TreemapRequestStylePalette string

// List of TreemapRequestStylePalette.
const (
	TREEMAPREQUESTSTYLEPALETTE_SUCCESS                  TreemapRequestStylePalette = "success"
	TREEMAPREQUESTSTYLEPALETTE_RED                      TreemapRequestStylePalette = "red"
	TREEMAPREQUESTSTYLEPALETTE_GREEN                    TreemapRequestStylePalette = "green"
	TREEMAPREQUESTSTYLEPALETTE_ORANGE                   TreemapRequestStylePalette = "orange"
	TREEMAPREQUESTSTYLEPALETTE_GRAY                     TreemapRequestStylePalette = "gray"
	TREEMAPREQUESTSTYLEPALETTE_DOG_CLASSIC              TreemapRequestStylePalette = "dog_classic"
	TREEMAPREQUESTSTYLEPALETTE_DOG_CLASSIC_AREA         TreemapRequestStylePalette = "dog_classic_area"
	TREEMAPREQUESTSTYLEPALETTE_DOG_CLASSIC_LINE         TreemapRequestStylePalette = "dog_classic_line"
	TREEMAPREQUESTSTYLEPALETTE_DOG_CLASSIC_BARS         TreemapRequestStylePalette = "dog_classic_bars"
	TREEMAPREQUESTSTYLEPALETTE_DOG_HIGHLIGHT            TreemapRequestStylePalette = "dog_highlight"
	TREEMAPREQUESTSTYLEPALETTE_MONITOR_STATE            TreemapRequestStylePalette = "monitor_state"
	TREEMAPREQUESTSTYLEPALETTE_YELLOW_TO_GREEN          TreemapRequestStylePalette = "yellow_to_green"
	TREEMAPREQUESTSTYLEPALETTE_GREEN_TO_ORANGE          TreemapRequestStylePalette = "green_to_orange"
	TREEMAPREQUESTSTYLEPALETTE_HOSTMAP_BLUES            TreemapRequestStylePalette = "hostmap_blues"
	TREEMAPREQUESTSTYLEPALETTE_HOSTMAP_GREEN_RED        TreemapRequestStylePalette = "hostmap_green_red"
	TREEMAPREQUESTSTYLEPALETTE_HOSTMAP_GREEN_ORANGE_RED TreemapRequestStylePalette = "hostmap_green_orange_red"
	TREEMAPREQUESTSTYLEPALETTE_HOSTMAP_BLUE_RED         TreemapRequestStylePalette = "hostmap_blue_red"
	TREEMAPREQUESTSTYLEPALETTE_DARK_GREY                TreemapRequestStylePalette = "dark_grey"
	TREEMAPREQUESTSTYLEPALETTE_PRIDE                    TreemapRequestStylePalette = "pride"
	TREEMAPREQUESTSTYLEPALETTE_GNBU                     TreemapRequestStylePalette = "GnBu"
	TREEMAPREQUESTSTYLEPALETTE_YLORRD                   TreemapRequestStylePalette = "YlOrRd"
	TREEMAPREQUESTSTYLEPALETTE_PURPLES                  TreemapRequestStylePalette = "Purples"
	TREEMAPREQUESTSTYLEPALETTE_BLUES                    TreemapRequestStylePalette = "Blues"
	TREEMAPREQUESTSTYLEPALETTE_ORANGES                  TreemapRequestStylePalette = "Oranges"
	TREEMAPREQUESTSTYLEPALETTE_REDS                     TreemapRequestStylePalette = "Reds"
	TREEMAPREQUESTSTYLEPALETTE_GREENS                   TreemapRequestStylePalette = "Greens"
	TREEMAPREQUESTSTYLEPALETTE_GREYS                    TreemapRequestStylePalette = "Greys"
	TREEMAPREQUESTSTYLEPALETTE_STATUS_ERRORS            TreemapRequestStylePalette = "status_errors"
	TREEMAPREQUESTSTYLEPALETTE_STATUS_WARNINGS          TreemapRequestStylePalette = "status_warnings"
	TREEMAPREQUESTSTYLEPALETTE_STATUS_SUCCESSES         TreemapRequestStylePalette = "status_successes"
	TREEMAPREQUESTSTYLEPALETTE_PLASMA                   TreemapRequestStylePalette = "Plasma"
	TREEMAPREQUESTSTYLEPALETTE_VIRIDIS                  TreemapRequestStylePalette = "Viridis"
	TREEMAPREQUESTSTYLEPALETTE_DOGCAT                   TreemapRequestStylePalette = "dogcat"
	TREEMAPREQUESTSTYLEPALETTE_DATADOG16                TreemapRequestStylePalette = "datadog16"
	TREEMAPREQUESTSTYLEPALETTE_PA_CLASSIC_LINE          TreemapRequestStylePalette = "pa_classic_line"
	TREEMAPREQUESTSTYLEPALETTE_PA_CLASSIC_AREA          TreemapRequestStylePalette = "pa_classic_area"
	TREEMAPREQUESTSTYLEPALETTE_PA_CLASSIC_BARS          TreemapRequestStylePalette = "pa_classic_bars"
	TREEMAPREQUESTSTYLEPALETTE_PA_DATADOG16             TreemapRequestStylePalette = "pa_datadog16"
	TREEMAPREQUESTSTYLEPALETTE_CLASSIC                  TreemapRequestStylePalette = "classic"
	TREEMAPREQUESTSTYLEPALETTE_COOL_HM                  TreemapRequestStylePalette = "Cool_hm"
	TREEMAPREQUESTSTYLEPALETTE_WARM_HM                  TreemapRequestStylePalette = "Warm_hm"
	TREEMAPREQUESTSTYLEPALETTE_YELLOW___GREEN           TreemapRequestStylePalette = "Yellow / Green"
	TREEMAPREQUESTSTYLEPALETTE_GREEN___ORANGE           TreemapRequestStylePalette = "Green / Orange"
	TREEMAPREQUESTSTYLEPALETTE_GREEN___RED              TreemapRequestStylePalette = "Green / Red"
	TREEMAPREQUESTSTYLEPALETTE_GREEN___ORANGE___RED     TreemapRequestStylePalette = "Green / Orange / Red"
	TREEMAPREQUESTSTYLEPALETTE_BLUE___RED               TreemapRequestStylePalette = "Blue / Red"
	TREEMAPREQUESTSTYLEPALETTE_GOOD___BAD               TreemapRequestStylePalette = "Good - bad"
	TREEMAPREQUESTSTYLEPALETTE_COOL                     TreemapRequestStylePalette = "cool"
	TREEMAPREQUESTSTYLEPALETTE_OCEAN                    TreemapRequestStylePalette = "ocean"
	TREEMAPREQUESTSTYLEPALETTE_WARM                     TreemapRequestStylePalette = "warm"
	TREEMAPREQUESTSTYLEPALETTE_FIRE                     TreemapRequestStylePalette = "fire"
	TREEMAPREQUESTSTYLEPALETTE_WARN                     TreemapRequestStylePalette = "warn"
	TREEMAPREQUESTSTYLEPALETTE_PURPLE                   TreemapRequestStylePalette = "purple"
	TREEMAPREQUESTSTYLEPALETTE_GREY                     TreemapRequestStylePalette = "grey"
	TREEMAPREQUESTSTYLEPALETTE_BLUE                     TreemapRequestStylePalette = "blue"
	TREEMAPREQUESTSTYLEPALETTE_DD20                     TreemapRequestStylePalette = "dd20"
	TREEMAPREQUESTSTYLEPALETTE_SEMANTIC                 TreemapRequestStylePalette = "semantic"
)

var allowedTreemapRequestStylePaletteEnumValues = []TreemapRequestStylePalette{
	TREEMAPREQUESTSTYLEPALETTE_SUCCESS,
	TREEMAPREQUESTSTYLEPALETTE_RED,
	TREEMAPREQUESTSTYLEPALETTE_GREEN,
	TREEMAPREQUESTSTYLEPALETTE_ORANGE,
	TREEMAPREQUESTSTYLEPALETTE_GRAY,
	TREEMAPREQUESTSTYLEPALETTE_DOG_CLASSIC,
	TREEMAPREQUESTSTYLEPALETTE_DOG_CLASSIC_AREA,
	TREEMAPREQUESTSTYLEPALETTE_DOG_CLASSIC_LINE,
	TREEMAPREQUESTSTYLEPALETTE_DOG_CLASSIC_BARS,
	TREEMAPREQUESTSTYLEPALETTE_DOG_HIGHLIGHT,
	TREEMAPREQUESTSTYLEPALETTE_MONITOR_STATE,
	TREEMAPREQUESTSTYLEPALETTE_YELLOW_TO_GREEN,
	TREEMAPREQUESTSTYLEPALETTE_GREEN_TO_ORANGE,
	TREEMAPREQUESTSTYLEPALETTE_HOSTMAP_BLUES,
	TREEMAPREQUESTSTYLEPALETTE_HOSTMAP_GREEN_RED,
	TREEMAPREQUESTSTYLEPALETTE_HOSTMAP_GREEN_ORANGE_RED,
	TREEMAPREQUESTSTYLEPALETTE_HOSTMAP_BLUE_RED,
	TREEMAPREQUESTSTYLEPALETTE_DARK_GREY,
	TREEMAPREQUESTSTYLEPALETTE_PRIDE,
	TREEMAPREQUESTSTYLEPALETTE_GNBU,
	TREEMAPREQUESTSTYLEPALETTE_YLORRD,
	TREEMAPREQUESTSTYLEPALETTE_PURPLES,
	TREEMAPREQUESTSTYLEPALETTE_BLUES,
	TREEMAPREQUESTSTYLEPALETTE_ORANGES,
	TREEMAPREQUESTSTYLEPALETTE_REDS,
	TREEMAPREQUESTSTYLEPALETTE_GREENS,
	TREEMAPREQUESTSTYLEPALETTE_GREYS,
	TREEMAPREQUESTSTYLEPALETTE_STATUS_ERRORS,
	TREEMAPREQUESTSTYLEPALETTE_STATUS_WARNINGS,
	TREEMAPREQUESTSTYLEPALETTE_STATUS_SUCCESSES,
	TREEMAPREQUESTSTYLEPALETTE_PLASMA,
	TREEMAPREQUESTSTYLEPALETTE_VIRIDIS,
	TREEMAPREQUESTSTYLEPALETTE_DOGCAT,
	TREEMAPREQUESTSTYLEPALETTE_DATADOG16,
	TREEMAPREQUESTSTYLEPALETTE_PA_CLASSIC_LINE,
	TREEMAPREQUESTSTYLEPALETTE_PA_CLASSIC_AREA,
	TREEMAPREQUESTSTYLEPALETTE_PA_CLASSIC_BARS,
	TREEMAPREQUESTSTYLEPALETTE_PA_DATADOG16,
	TREEMAPREQUESTSTYLEPALETTE_CLASSIC,
	TREEMAPREQUESTSTYLEPALETTE_COOL_HM,
	TREEMAPREQUESTSTYLEPALETTE_WARM_HM,
	TREEMAPREQUESTSTYLEPALETTE_YELLOW___GREEN,
	TREEMAPREQUESTSTYLEPALETTE_GREEN___ORANGE,
	TREEMAPREQUESTSTYLEPALETTE_GREEN___RED,
	TREEMAPREQUESTSTYLEPALETTE_GREEN___ORANGE___RED,
	TREEMAPREQUESTSTYLEPALETTE_BLUE___RED,
	TREEMAPREQUESTSTYLEPALETTE_GOOD___BAD,
	TREEMAPREQUESTSTYLEPALETTE_COOL,
	TREEMAPREQUESTSTYLEPALETTE_OCEAN,
	TREEMAPREQUESTSTYLEPALETTE_WARM,
	TREEMAPREQUESTSTYLEPALETTE_FIRE,
	TREEMAPREQUESTSTYLEPALETTE_WARN,
	TREEMAPREQUESTSTYLEPALETTE_PURPLE,
	TREEMAPREQUESTSTYLEPALETTE_GREY,
	TREEMAPREQUESTSTYLEPALETTE_BLUE,
	TREEMAPREQUESTSTYLEPALETTE_DD20,
	TREEMAPREQUESTSTYLEPALETTE_SEMANTIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreemapRequestStylePalette) GetAllowedValues() []TreemapRequestStylePalette {
	return allowedTreemapRequestStylePaletteEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreemapRequestStylePalette) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreemapRequestStylePalette(value)
	return nil
}

// NewTreemapRequestStylePaletteFromValue returns a pointer to a valid TreemapRequestStylePalette
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreemapRequestStylePaletteFromValue(v string) (*TreemapRequestStylePalette, error) {
	ev := TreemapRequestStylePalette(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreemapRequestStylePalette: valid values are %v", v, allowedTreemapRequestStylePaletteEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreemapRequestStylePalette) IsValid() bool {
	for _, existing := range allowedTreemapRequestStylePaletteEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreemapRequestStylePalette value.
func (v TreemapRequestStylePalette) Ptr() *TreemapRequestStylePalette {
	return &v
}

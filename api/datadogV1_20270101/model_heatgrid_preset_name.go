// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridPresetName Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridPresetName string

// List of HeatgridPresetName.
const (
	HEATGRIDPRESETNAME_DOG_CLASSIC              HeatgridPresetName = "dog_classic"
	HEATGRIDPRESETNAME_DOG_CLASSIC_AREA         HeatgridPresetName = "dog_classic_area"
	HEATGRIDPRESETNAME_DOG_CLASSIC_LINE         HeatgridPresetName = "dog_classic_line"
	HEATGRIDPRESETNAME_DOG_CLASSIC_BARS         HeatgridPresetName = "dog_classic_bars"
	HEATGRIDPRESETNAME_DOG_HIGHLIGHT            HeatgridPresetName = "dog_highlight"
	HEATGRIDPRESETNAME_MONITOR_STATE            HeatgridPresetName = "monitor_state"
	HEATGRIDPRESETNAME_YELLOW_TO_GREEN          HeatgridPresetName = "yellow_to_green"
	HEATGRIDPRESETNAME_GREEN_TO_ORANGE          HeatgridPresetName = "green_to_orange"
	HEATGRIDPRESETNAME_HOSTMAP_BLUES            HeatgridPresetName = "hostmap_blues"
	HEATGRIDPRESETNAME_HOSTMAP_GREEN_RED        HeatgridPresetName = "hostmap_green_red"
	HEATGRIDPRESETNAME_HOSTMAP_GREEN_ORANGE_RED HeatgridPresetName = "hostmap_green_orange_red"
	HEATGRIDPRESETNAME_HOSTMAP_BLUE_RED         HeatgridPresetName = "hostmap_blue_red"
	HEATGRIDPRESETNAME_DARK_GREY                HeatgridPresetName = "dark_grey"
	HEATGRIDPRESETNAME_PRIDE                    HeatgridPresetName = "pride"
	HEATGRIDPRESETNAME_GNBU                     HeatgridPresetName = "GnBu"
	HEATGRIDPRESETNAME_YLORRD                   HeatgridPresetName = "YlOrRd"
	HEATGRIDPRESETNAME_PURPLES                  HeatgridPresetName = "Purples"
	HEATGRIDPRESETNAME_BLUES                    HeatgridPresetName = "Blues"
	HEATGRIDPRESETNAME_ORANGES                  HeatgridPresetName = "Oranges"
	HEATGRIDPRESETNAME_REDS                     HeatgridPresetName = "Reds"
	HEATGRIDPRESETNAME_GREENS                   HeatgridPresetName = "Greens"
	HEATGRIDPRESETNAME_GREYS                    HeatgridPresetName = "Greys"
	HEATGRIDPRESETNAME_STATUS_ERRORS            HeatgridPresetName = "status_errors"
	HEATGRIDPRESETNAME_STATUS_WARNINGS          HeatgridPresetName = "status_warnings"
	HEATGRIDPRESETNAME_STATUS_SUCCESSES         HeatgridPresetName = "status_successes"
	HEATGRIDPRESETNAME_PLASMA                   HeatgridPresetName = "Plasma"
	HEATGRIDPRESETNAME_VIRIDIS                  HeatgridPresetName = "Viridis"
	HEATGRIDPRESETNAME_DOGCAT                   HeatgridPresetName = "dogcat"
	HEATGRIDPRESETNAME_DATADOG16                HeatgridPresetName = "datadog16"
	HEATGRIDPRESETNAME_PA_CLASSIC_LINE          HeatgridPresetName = "pa_classic_line"
	HEATGRIDPRESETNAME_PA_CLASSIC_AREA          HeatgridPresetName = "pa_classic_area"
	HEATGRIDPRESETNAME_PA_CLASSIC_BARS          HeatgridPresetName = "pa_classic_bars"
	HEATGRIDPRESETNAME_PA_DATADOG16             HeatgridPresetName = "pa_datadog16"
	HEATGRIDPRESETNAME_CLASSIC                  HeatgridPresetName = "classic"
)

var allowedHeatgridPresetNameEnumValues = []HeatgridPresetName{
	HEATGRIDPRESETNAME_DOG_CLASSIC,
	HEATGRIDPRESETNAME_DOG_CLASSIC_AREA,
	HEATGRIDPRESETNAME_DOG_CLASSIC_LINE,
	HEATGRIDPRESETNAME_DOG_CLASSIC_BARS,
	HEATGRIDPRESETNAME_DOG_HIGHLIGHT,
	HEATGRIDPRESETNAME_MONITOR_STATE,
	HEATGRIDPRESETNAME_YELLOW_TO_GREEN,
	HEATGRIDPRESETNAME_GREEN_TO_ORANGE,
	HEATGRIDPRESETNAME_HOSTMAP_BLUES,
	HEATGRIDPRESETNAME_HOSTMAP_GREEN_RED,
	HEATGRIDPRESETNAME_HOSTMAP_GREEN_ORANGE_RED,
	HEATGRIDPRESETNAME_HOSTMAP_BLUE_RED,
	HEATGRIDPRESETNAME_DARK_GREY,
	HEATGRIDPRESETNAME_PRIDE,
	HEATGRIDPRESETNAME_GNBU,
	HEATGRIDPRESETNAME_YLORRD,
	HEATGRIDPRESETNAME_PURPLES,
	HEATGRIDPRESETNAME_BLUES,
	HEATGRIDPRESETNAME_ORANGES,
	HEATGRIDPRESETNAME_REDS,
	HEATGRIDPRESETNAME_GREENS,
	HEATGRIDPRESETNAME_GREYS,
	HEATGRIDPRESETNAME_STATUS_ERRORS,
	HEATGRIDPRESETNAME_STATUS_WARNINGS,
	HEATGRIDPRESETNAME_STATUS_SUCCESSES,
	HEATGRIDPRESETNAME_PLASMA,
	HEATGRIDPRESETNAME_VIRIDIS,
	HEATGRIDPRESETNAME_DOGCAT,
	HEATGRIDPRESETNAME_DATADOG16,
	HEATGRIDPRESETNAME_PA_CLASSIC_LINE,
	HEATGRIDPRESETNAME_PA_CLASSIC_AREA,
	HEATGRIDPRESETNAME_PA_CLASSIC_BARS,
	HEATGRIDPRESETNAME_PA_DATADOG16,
	HEATGRIDPRESETNAME_CLASSIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridPresetName) GetAllowedValues() []HeatgridPresetName {
	return allowedHeatgridPresetNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridPresetName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridPresetName(value)
	return nil
}

// NewHeatgridPresetNameFromValue returns a pointer to a valid HeatgridPresetName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridPresetNameFromValue(v string) (*HeatgridPresetName, error) {
	ev := HeatgridPresetName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridPresetName: valid values are %v", v, allowedHeatgridPresetNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridPresetName) IsValid() bool {
	for _, existing := range allowedHeatgridPresetNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridPresetName value.
func (v HeatgridPresetName) Ptr() *HeatgridPresetName {
	return &v
}

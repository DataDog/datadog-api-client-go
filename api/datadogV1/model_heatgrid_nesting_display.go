// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridNestingDisplay Display groups as flat rows.
type HeatgridNestingDisplay string

// List of HeatgridNestingDisplay.
const (
	HEATGRIDNESTINGDISPLAY_FLAT HeatgridNestingDisplay = "flat"
)

var allowedHeatgridNestingDisplayEnumValues = []HeatgridNestingDisplay{
	HEATGRIDNESTINGDISPLAY_FLAT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridNestingDisplay) GetAllowedValues() []HeatgridNestingDisplay {
	return allowedHeatgridNestingDisplayEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridNestingDisplay) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridNestingDisplay(value)
	return nil
}

// NewHeatgridNestingDisplayFromValue returns a pointer to a valid HeatgridNestingDisplay
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridNestingDisplayFromValue(v string) (*HeatgridNestingDisplay, error) {
	ev := HeatgridNestingDisplay(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridNestingDisplay: valid values are %v", v, allowedHeatgridNestingDisplayEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridNestingDisplay) IsValid() bool {
	for _, existing := range allowedHeatgridNestingDisplayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridNestingDisplay value.
func (v HeatgridNestingDisplay) Ptr() *HeatgridNestingDisplay {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelGroupByDisplay Test-only generated schema for the dashboard versioning proof of concept.
type FunnelGroupByDisplay string

// List of FunnelGroupByDisplay.
const (
	FUNNELGROUPBYDISPLAY_STACKED      FunnelGroupByDisplay = "stacked"
	FUNNELGROUPBYDISPLAY_SIDE_BY_SIDE FunnelGroupByDisplay = "side_by_side"
)

var allowedFunnelGroupByDisplayEnumValues = []FunnelGroupByDisplay{
	FUNNELGROUPBYDISPLAY_STACKED,
	FUNNELGROUPBYDISPLAY_SIDE_BY_SIDE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FunnelGroupByDisplay) GetAllowedValues() []FunnelGroupByDisplay {
	return allowedFunnelGroupByDisplayEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FunnelGroupByDisplay) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FunnelGroupByDisplay(value)
	return nil
}

// NewFunnelGroupByDisplayFromValue returns a pointer to a valid FunnelGroupByDisplay
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFunnelGroupByDisplayFromValue(v string) (*FunnelGroupByDisplay, error) {
	ev := FunnelGroupByDisplay(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FunnelGroupByDisplay: valid values are %v", v, allowedFunnelGroupByDisplayEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FunnelGroupByDisplay) IsValid() bool {
	for _, existing := range allowedFunnelGroupByDisplayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FunnelGroupByDisplay value.
func (v FunnelGroupByDisplay) Ptr() *FunnelGroupByDisplay {
	return &v
}

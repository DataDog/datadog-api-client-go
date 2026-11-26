// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SloTimeWindow Test-only generated schema for the dashboard versioning proof of concept.
type SloTimeWindow string

// List of SloTimeWindow.
const (
	SLOTIMEWINDOW_SEVEN_D  SloTimeWindow = "7d"
	SLOTIMEWINDOW_THIRTY_D SloTimeWindow = "30d"
	SLOTIMEWINDOW_NINETY_D SloTimeWindow = "90d"
)

var allowedSloTimeWindowEnumValues = []SloTimeWindow{
	SLOTIMEWINDOW_SEVEN_D,
	SLOTIMEWINDOW_THIRTY_D,
	SLOTIMEWINDOW_NINETY_D,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SloTimeWindow) GetAllowedValues() []SloTimeWindow {
	return allowedSloTimeWindowEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SloTimeWindow) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SloTimeWindow(value)
	return nil
}

// NewSloTimeWindowFromValue returns a pointer to a valid SloTimeWindow
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSloTimeWindowFromValue(v string) (*SloTimeWindow, error) {
	ev := SloTimeWindow(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SloTimeWindow: valid values are %v", v, allowedSloTimeWindowEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SloTimeWindow) IsValid() bool {
	for _, existing := range allowedSloTimeWindowEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SloTimeWindow value.
func (v SloTimeWindow) Ptr() *SloTimeWindow {
	return &v
}

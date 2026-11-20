// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SloWidgetTimeWindow Test-only generated schema for the dashboard versioning proof of concept.
type SloWidgetTimeWindow string

// List of SloWidgetTimeWindow.
const (
	SLOWIDGETTIMEWINDOW_WEEK_TO_DATE   SloWidgetTimeWindow = "week_to_date"
	SLOWIDGETTIMEWINDOW_MONTH_TO_DATE  SloWidgetTimeWindow = "month_to_date"
	SLOWIDGETTIMEWINDOW_SEVEN_D        SloWidgetTimeWindow = "7d"
	SLOWIDGETTIMEWINDOW_THIRTY_D       SloWidgetTimeWindow = "30d"
	SLOWIDGETTIMEWINDOW_NINETY_D       SloWidgetTimeWindow = "90d"
	SLOWIDGETTIMEWINDOW_PREVIOUS_WEEK  SloWidgetTimeWindow = "previous_week"
	SLOWIDGETTIMEWINDOW_PREVIOUS_MONTH SloWidgetTimeWindow = "previous_month"
	SLOWIDGETTIMEWINDOW_GLOBAL_TIME    SloWidgetTimeWindow = "global_time"
)

var allowedSloWidgetTimeWindowEnumValues = []SloWidgetTimeWindow{
	SLOWIDGETTIMEWINDOW_WEEK_TO_DATE,
	SLOWIDGETTIMEWINDOW_MONTH_TO_DATE,
	SLOWIDGETTIMEWINDOW_SEVEN_D,
	SLOWIDGETTIMEWINDOW_THIRTY_D,
	SLOWIDGETTIMEWINDOW_NINETY_D,
	SLOWIDGETTIMEWINDOW_PREVIOUS_WEEK,
	SLOWIDGETTIMEWINDOW_PREVIOUS_MONTH,
	SLOWIDGETTIMEWINDOW_GLOBAL_TIME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SloWidgetTimeWindow) GetAllowedValues() []SloWidgetTimeWindow {
	return allowedSloWidgetTimeWindowEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SloWidgetTimeWindow) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SloWidgetTimeWindow(value)
	return nil
}

// NewSloWidgetTimeWindowFromValue returns a pointer to a valid SloWidgetTimeWindow
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSloWidgetTimeWindowFromValue(v string) (*SloWidgetTimeWindow, error) {
	ev := SloWidgetTimeWindow(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SloWidgetTimeWindow: valid values are %v", v, allowedSloWidgetTimeWindowEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SloWidgetTimeWindow) IsValid() bool {
	for _, existing := range allowedSloWidgetTimeWindowEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SloWidgetTimeWindow value.
func (v SloWidgetTimeWindow) Ptr() *SloWidgetTimeWindow {
	return &v
}

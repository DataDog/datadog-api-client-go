// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// DashboardReportScheduleRepeatOnDayOfWeek Defines which day of the week to send the weekly report. If the frequency field uses weeks (such as `1w`), this field is required and defines which day of the week to send the report. If the frequency field does not use weeks (for example, `1d`, `1mo`), then this field must be `null` (if provided). If this field is not provided in an update request, and the update request changes the frequency to use days or months (for example, `1d`, `1mo`), then this field is automatically set to `null`.  This field is mutually exclusive with `repeat_on_day_of_month` and cannot be defined in the same request that includes a non-null value for `repeat_on_day_of_month`.
type DashboardReportScheduleRepeatOnDayOfWeek string

// List of DashboardReportScheduleRepeatOnDayOfWeek.
const (
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFWEEK_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_WEEK_MONDAY    DashboardReportScheduleRepeatOnDayOfWeek = "Monday"
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFWEEK_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_WEEK_TUESDAY   DashboardReportScheduleRepeatOnDayOfWeek = "Tuesday"
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFWEEK_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_WEEK_WEDNESDAY DashboardReportScheduleRepeatOnDayOfWeek = "Wednesday"
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFWEEK_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_WEEK_THURSDAY  DashboardReportScheduleRepeatOnDayOfWeek = "Thursday"
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFWEEK_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_WEEK_FRIDAY    DashboardReportScheduleRepeatOnDayOfWeek = "Friday"
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFWEEK_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_WEEK_SATURDAY  DashboardReportScheduleRepeatOnDayOfWeek = "Saturday"
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFWEEK_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_WEEK_SUNDAY    DashboardReportScheduleRepeatOnDayOfWeek = "Sunday"
)

var allowedDashboardReportScheduleRepeatOnDayOfWeekEnumValues = []DashboardReportScheduleRepeatOnDayOfWeek{
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFWEEK_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_WEEK_MONDAY,
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFWEEK_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_WEEK_TUESDAY,
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFWEEK_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_WEEK_WEDNESDAY,
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFWEEK_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_WEEK_THURSDAY,
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFWEEK_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_WEEK_FRIDAY,
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFWEEK_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_WEEK_SATURDAY,
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFWEEK_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_WEEK_SUNDAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardReportScheduleRepeatOnDayOfWeek) GetAllowedValues() []DashboardReportScheduleRepeatOnDayOfWeek {
	return allowedDashboardReportScheduleRepeatOnDayOfWeekEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardReportScheduleRepeatOnDayOfWeek) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardReportScheduleRepeatOnDayOfWeek(value)
	return nil
}

// NewDashboardReportScheduleRepeatOnDayOfWeekFromValue returns a pointer to a valid DashboardReportScheduleRepeatOnDayOfWeek
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardReportScheduleRepeatOnDayOfWeekFromValue(v string) (*DashboardReportScheduleRepeatOnDayOfWeek, error) {
	ev := DashboardReportScheduleRepeatOnDayOfWeek(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardReportScheduleRepeatOnDayOfWeek: valid values are %v", v, allowedDashboardReportScheduleRepeatOnDayOfWeekEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardReportScheduleRepeatOnDayOfWeek) IsValid() bool {
	for _, existing := range allowedDashboardReportScheduleRepeatOnDayOfWeekEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardReportScheduleRepeatOnDayOfWeek value.
func (v DashboardReportScheduleRepeatOnDayOfWeek) Ptr() *DashboardReportScheduleRepeatOnDayOfWeek {
	return &v
}

// NullableDashboardReportScheduleRepeatOnDayOfWeek handles when a null is used for DashboardReportScheduleRepeatOnDayOfWeek.
type NullableDashboardReportScheduleRepeatOnDayOfWeek struct {
	value *DashboardReportScheduleRepeatOnDayOfWeek
	isSet bool
}

// Get returns the associated value.
func (v NullableDashboardReportScheduleRepeatOnDayOfWeek) Get() *DashboardReportScheduleRepeatOnDayOfWeek {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDashboardReportScheduleRepeatOnDayOfWeek) Set(val *DashboardReportScheduleRepeatOnDayOfWeek) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDashboardReportScheduleRepeatOnDayOfWeek) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableDashboardReportScheduleRepeatOnDayOfWeek) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDashboardReportScheduleRepeatOnDayOfWeek initializes the struct as if Set has been called.
func NewNullableDashboardReportScheduleRepeatOnDayOfWeek(val *DashboardReportScheduleRepeatOnDayOfWeek) *NullableDashboardReportScheduleRepeatOnDayOfWeek {
	return &NullableDashboardReportScheduleRepeatOnDayOfWeek{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDashboardReportScheduleRepeatOnDayOfWeek) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDashboardReportScheduleRepeatOnDayOfWeek) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

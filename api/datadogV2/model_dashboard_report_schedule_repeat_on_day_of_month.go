// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// DashboardReportScheduleRepeatOnDayOfMonth Defines which day of the month to send the monthly report. If the frequency field uses months (for example, `1mo`), this field is required. If the frequency field does not use months (for example `1d`, `1w`), then this field must be `null` (if provided). If this field is not provided in an update request, and the update request changes the frequency to use days or weeks (for example, `1d`, `1w`), then this field is automatically set to `null`. This field is mutually exclusive with `repeat_on_day_of_week`, and cannot be defined in the same request that includes a non-null value for `repeat_on_day_of_week`.
type DashboardReportScheduleRepeatOnDayOfMonth string

// List of DashboardReportScheduleRepeatOnDayOfMonth.
const (
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFMONTH_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_MONTH_1ST  DashboardReportScheduleRepeatOnDayOfMonth = "1st"
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFMONTH_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_MONTH_15TH DashboardReportScheduleRepeatOnDayOfMonth = "15th"
)

var allowedDashboardReportScheduleRepeatOnDayOfMonthEnumValues = []DashboardReportScheduleRepeatOnDayOfMonth{
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFMONTH_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_MONTH_1ST,
	DASHBOARDREPORTSCHEDULEREPEATONDAYOFMONTH_DASHBOARD_REPORT_SCHEDULE_REPEAT_ON_DAY_OF_MONTH_15TH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardReportScheduleRepeatOnDayOfMonth) GetAllowedValues() []DashboardReportScheduleRepeatOnDayOfMonth {
	return allowedDashboardReportScheduleRepeatOnDayOfMonthEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardReportScheduleRepeatOnDayOfMonth) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardReportScheduleRepeatOnDayOfMonth(value)
	return nil
}

// NewDashboardReportScheduleRepeatOnDayOfMonthFromValue returns a pointer to a valid DashboardReportScheduleRepeatOnDayOfMonth
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardReportScheduleRepeatOnDayOfMonthFromValue(v string) (*DashboardReportScheduleRepeatOnDayOfMonth, error) {
	ev := DashboardReportScheduleRepeatOnDayOfMonth(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardReportScheduleRepeatOnDayOfMonth: valid values are %v", v, allowedDashboardReportScheduleRepeatOnDayOfMonthEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardReportScheduleRepeatOnDayOfMonth) IsValid() bool {
	for _, existing := range allowedDashboardReportScheduleRepeatOnDayOfMonthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardReportScheduleRepeatOnDayOfMonth value.
func (v DashboardReportScheduleRepeatOnDayOfMonth) Ptr() *DashboardReportScheduleRepeatOnDayOfMonth {
	return &v
}

// NullableDashboardReportScheduleRepeatOnDayOfMonth handles when a null is used for DashboardReportScheduleRepeatOnDayOfMonth.
type NullableDashboardReportScheduleRepeatOnDayOfMonth struct {
	value *DashboardReportScheduleRepeatOnDayOfMonth
	isSet bool
}

// Get returns the associated value.
func (v NullableDashboardReportScheduleRepeatOnDayOfMonth) Get() *DashboardReportScheduleRepeatOnDayOfMonth {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDashboardReportScheduleRepeatOnDayOfMonth) Set(val *DashboardReportScheduleRepeatOnDayOfMonth) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDashboardReportScheduleRepeatOnDayOfMonth) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableDashboardReportScheduleRepeatOnDayOfMonth) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDashboardReportScheduleRepeatOnDayOfMonth initializes the struct as if Set has been called.
func NewNullableDashboardReportScheduleRepeatOnDayOfMonth(val *DashboardReportScheduleRepeatOnDayOfMonth) *NullableDashboardReportScheduleRepeatOnDayOfMonth {
	return &NullableDashboardReportScheduleRepeatOnDayOfMonth{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDashboardReportScheduleRepeatOnDayOfMonth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDashboardReportScheduleRepeatOnDayOfMonth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

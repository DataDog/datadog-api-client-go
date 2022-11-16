// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// DashboardReportScheduleFrequency Frequency with which to send the report.
type DashboardReportScheduleFrequency string

// List of DashboardReportScheduleFrequency.
const (
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_1D   DashboardReportScheduleFrequency = "1d"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_2D   DashboardReportScheduleFrequency = "2d"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_3D   DashboardReportScheduleFrequency = "3d"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_4D   DashboardReportScheduleFrequency = "4d"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_5D   DashboardReportScheduleFrequency = "5d"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_6D   DashboardReportScheduleFrequency = "6d"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_1W   DashboardReportScheduleFrequency = "1w"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_2W   DashboardReportScheduleFrequency = "2w"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_3W   DashboardReportScheduleFrequency = "3w"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_1MO  DashboardReportScheduleFrequency = "1mo"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_2MO  DashboardReportScheduleFrequency = "2mo"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_3MO  DashboardReportScheduleFrequency = "3mo"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_4MO  DashboardReportScheduleFrequency = "4mo"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_5MO  DashboardReportScheduleFrequency = "5mo"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_6MO  DashboardReportScheduleFrequency = "6mo"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_7MO  DashboardReportScheduleFrequency = "7mo"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_8MO  DashboardReportScheduleFrequency = "8mo"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_9MO  DashboardReportScheduleFrequency = "9mo"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_10MO DashboardReportScheduleFrequency = "10mo"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_11MO DashboardReportScheduleFrequency = "11mo"
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_12MO DashboardReportScheduleFrequency = "12mo"
)

var allowedDashboardReportScheduleFrequencyEnumValues = []DashboardReportScheduleFrequency{
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_1D,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_2D,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_3D,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_4D,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_5D,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_6D,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_1W,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_2W,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_3W,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_1MO,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_2MO,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_3MO,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_4MO,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_5MO,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_6MO,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_7MO,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_8MO,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_9MO,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_10MO,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_11MO,
	DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_12MO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardReportScheduleFrequency) GetAllowedValues() []DashboardReportScheduleFrequency {
	return allowedDashboardReportScheduleFrequencyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardReportScheduleFrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardReportScheduleFrequency(value)
	return nil
}

// NewDashboardReportScheduleFrequencyFromValue returns a pointer to a valid DashboardReportScheduleFrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardReportScheduleFrequencyFromValue(v string) (*DashboardReportScheduleFrequency, error) {
	ev := DashboardReportScheduleFrequency(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardReportScheduleFrequency: valid values are %v", v, allowedDashboardReportScheduleFrequencyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardReportScheduleFrequency) IsValid() bool {
	for _, existing := range allowedDashboardReportScheduleFrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardReportScheduleFrequency value.
func (v DashboardReportScheduleFrequency) Ptr() *DashboardReportScheduleFrequency {
	return &v
}

// NullableDashboardReportScheduleFrequency handles when a null is used for DashboardReportScheduleFrequency.
type NullableDashboardReportScheduleFrequency struct {
	value *DashboardReportScheduleFrequency
	isSet bool
}

// Get returns the associated value.
func (v NullableDashboardReportScheduleFrequency) Get() *DashboardReportScheduleFrequency {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDashboardReportScheduleFrequency) Set(val *DashboardReportScheduleFrequency) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDashboardReportScheduleFrequency) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableDashboardReportScheduleFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDashboardReportScheduleFrequency initializes the struct as if Set has been called.
func NewNullableDashboardReportScheduleFrequency(val *DashboardReportScheduleFrequency) *NullableDashboardReportScheduleFrequency {
	return &NullableDashboardReportScheduleFrequency{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDashboardReportScheduleFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDashboardReportScheduleFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

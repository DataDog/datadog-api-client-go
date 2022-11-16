// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// DashboardReportTimeframe Time period covered by the widgets in the dashboard report, at the time of report generation.
type DashboardReportTimeframe string

// List of DashboardReportTimeframe.
const (
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_5M  DashboardReportTimeframe = "5m"
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_15M DashboardReportTimeframe = "15m"
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_30M DashboardReportTimeframe = "30m"
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_1H  DashboardReportTimeframe = "1h"
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_4H  DashboardReportTimeframe = "4h"
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_1D  DashboardReportTimeframe = "1d"
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_2D  DashboardReportTimeframe = "2d"
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_1W  DashboardReportTimeframe = "1w"
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_1MO DashboardReportTimeframe = "1mo"
)

var allowedDashboardReportTimeframeEnumValues = []DashboardReportTimeframe{
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_5M,
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_15M,
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_30M,
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_1H,
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_4H,
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_1D,
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_2D,
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_1W,
	DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_1MO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardReportTimeframe) GetAllowedValues() []DashboardReportTimeframe {
	return allowedDashboardReportTimeframeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardReportTimeframe) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardReportTimeframe(value)
	return nil
}

// NewDashboardReportTimeframeFromValue returns a pointer to a valid DashboardReportTimeframe
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardReportTimeframeFromValue(v string) (*DashboardReportTimeframe, error) {
	ev := DashboardReportTimeframe(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardReportTimeframe: valid values are %v", v, allowedDashboardReportTimeframeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardReportTimeframe) IsValid() bool {
	for _, existing := range allowedDashboardReportTimeframeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardReportTimeframe value.
func (v DashboardReportTimeframe) Ptr() *DashboardReportTimeframe {
	return &v
}

// NullableDashboardReportTimeframe handles when a null is used for DashboardReportTimeframe.
type NullableDashboardReportTimeframe struct {
	value *DashboardReportTimeframe
	isSet bool
}

// Get returns the associated value.
func (v NullableDashboardReportTimeframe) Get() *DashboardReportTimeframe {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDashboardReportTimeframe) Set(val *DashboardReportTimeframe) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDashboardReportTimeframe) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableDashboardReportTimeframe) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDashboardReportTimeframe initializes the struct as if Set has been called.
func NewNullableDashboardReportTimeframe(val *DashboardReportTimeframe) *NullableDashboardReportTimeframe {
	return &NullableDashboardReportTimeframe{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDashboardReportTimeframe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDashboardReportTimeframe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

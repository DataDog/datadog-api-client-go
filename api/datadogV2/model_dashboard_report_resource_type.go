// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// DashboardReportResourceType JSON:API type of the report's associated dashboard.
type DashboardReportResourceType string

// List of DashboardReportResourceType.
const (
	DASHBOARDREPORTRESOURCETYPE_DASHBOARDS_TYPE DashboardReportResourceType = "dashboards"
)

var allowedDashboardReportResourceTypeEnumValues = []DashboardReportResourceType{
	DASHBOARDREPORTRESOURCETYPE_DASHBOARDS_TYPE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardReportResourceType) GetAllowedValues() []DashboardReportResourceType {
	return allowedDashboardReportResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardReportResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardReportResourceType(value)
	return nil
}

// NewDashboardReportResourceTypeFromValue returns a pointer to a valid DashboardReportResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardReportResourceTypeFromValue(v string) (*DashboardReportResourceType, error) {
	ev := DashboardReportResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardReportResourceType: valid values are %v", v, allowedDashboardReportResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardReportResourceType) IsValid() bool {
	for _, existing := range allowedDashboardReportResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardReportResourceType value.
func (v DashboardReportResourceType) Ptr() *DashboardReportResourceType {
	return &v
}

// NullableDashboardReportResourceType handles when a null is used for DashboardReportResourceType.
type NullableDashboardReportResourceType struct {
	value *DashboardReportResourceType
	isSet bool
}

// Get returns the associated value.
func (v NullableDashboardReportResourceType) Get() *DashboardReportResourceType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDashboardReportResourceType) Set(val *DashboardReportResourceType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDashboardReportResourceType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableDashboardReportResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDashboardReportResourceType initializes the struct as if Set has been called.
func NewNullableDashboardReportResourceType(val *DashboardReportResourceType) *NullableDashboardReportResourceType {
	return &NullableDashboardReportResourceType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDashboardReportResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDashboardReportResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

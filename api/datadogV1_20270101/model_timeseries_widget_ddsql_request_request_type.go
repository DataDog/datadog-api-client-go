// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetDdsqlRequestRequestType Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesWidgetDdsqlRequestRequestType string

// List of TimeseriesWidgetDdsqlRequestRequestType.
const (
	TIMESERIESWIDGETDDSQLREQUESTREQUESTTYPE_DDSQL TimeseriesWidgetDdsqlRequestRequestType = "ddsql"
)

var allowedTimeseriesWidgetDdsqlRequestRequestTypeEnumValues = []TimeseriesWidgetDdsqlRequestRequestType{
	TIMESERIESWIDGETDDSQLREQUESTREQUESTTYPE_DDSQL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesWidgetDdsqlRequestRequestType) GetAllowedValues() []TimeseriesWidgetDdsqlRequestRequestType {
	return allowedTimeseriesWidgetDdsqlRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesWidgetDdsqlRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesWidgetDdsqlRequestRequestType(value)
	return nil
}

// NewTimeseriesWidgetDdsqlRequestRequestTypeFromValue returns a pointer to a valid TimeseriesWidgetDdsqlRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesWidgetDdsqlRequestRequestTypeFromValue(v string) (*TimeseriesWidgetDdsqlRequestRequestType, error) {
	ev := TimeseriesWidgetDdsqlRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesWidgetDdsqlRequestRequestType: valid values are %v", v, allowedTimeseriesWidgetDdsqlRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesWidgetDdsqlRequestRequestType) IsValid() bool {
	for _, existing := range allowedTimeseriesWidgetDdsqlRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesWidgetDdsqlRequestRequestType value.
func (v TimeseriesWidgetDdsqlRequestRequestType) Ptr() *TimeseriesWidgetDdsqlRequestRequestType {
	return &v
}

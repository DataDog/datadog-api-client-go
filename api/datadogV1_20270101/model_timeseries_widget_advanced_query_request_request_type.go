// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetAdvancedQueryRequestRequestType Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesWidgetAdvancedQueryRequestRequestType string

// List of TimeseriesWidgetAdvancedQueryRequestRequestType.
const (
	TIMESERIESWIDGETADVANCEDQUERYREQUESTREQUESTTYPE_ADVANCED_QUERY TimeseriesWidgetAdvancedQueryRequestRequestType = "advanced_query"
)

var allowedTimeseriesWidgetAdvancedQueryRequestRequestTypeEnumValues = []TimeseriesWidgetAdvancedQueryRequestRequestType{
	TIMESERIESWIDGETADVANCEDQUERYREQUESTREQUESTTYPE_ADVANCED_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesWidgetAdvancedQueryRequestRequestType) GetAllowedValues() []TimeseriesWidgetAdvancedQueryRequestRequestType {
	return allowedTimeseriesWidgetAdvancedQueryRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesWidgetAdvancedQueryRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesWidgetAdvancedQueryRequestRequestType(value)
	return nil
}

// NewTimeseriesWidgetAdvancedQueryRequestRequestTypeFromValue returns a pointer to a valid TimeseriesWidgetAdvancedQueryRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesWidgetAdvancedQueryRequestRequestTypeFromValue(v string) (*TimeseriesWidgetAdvancedQueryRequestRequestType, error) {
	ev := TimeseriesWidgetAdvancedQueryRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesWidgetAdvancedQueryRequestRequestType: valid values are %v", v, allowedTimeseriesWidgetAdvancedQueryRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesWidgetAdvancedQueryRequestRequestType) IsValid() bool {
	for _, existing := range allowedTimeseriesWidgetAdvancedQueryRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesWidgetAdvancedQueryRequestRequestType value.
func (v TimeseriesWidgetAdvancedQueryRequestRequestType) Ptr() *TimeseriesWidgetAdvancedQueryRequestRequestType {
	return &v
}

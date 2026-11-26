// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetLocalDatasetRequestRequestType Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesWidgetLocalDatasetRequestRequestType string

// List of TimeseriesWidgetLocalDatasetRequestRequestType.
const (
	TIMESERIESWIDGETLOCALDATASETREQUESTREQUESTTYPE_LOCAL_DATASET TimeseriesWidgetLocalDatasetRequestRequestType = "local_dataset"
)

var allowedTimeseriesWidgetLocalDatasetRequestRequestTypeEnumValues = []TimeseriesWidgetLocalDatasetRequestRequestType{
	TIMESERIESWIDGETLOCALDATASETREQUESTREQUESTTYPE_LOCAL_DATASET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesWidgetLocalDatasetRequestRequestType) GetAllowedValues() []TimeseriesWidgetLocalDatasetRequestRequestType {
	return allowedTimeseriesWidgetLocalDatasetRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesWidgetLocalDatasetRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesWidgetLocalDatasetRequestRequestType(value)
	return nil
}

// NewTimeseriesWidgetLocalDatasetRequestRequestTypeFromValue returns a pointer to a valid TimeseriesWidgetLocalDatasetRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesWidgetLocalDatasetRequestRequestTypeFromValue(v string) (*TimeseriesWidgetLocalDatasetRequestRequestType, error) {
	ev := TimeseriesWidgetLocalDatasetRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesWidgetLocalDatasetRequestRequestType: valid values are %v", v, allowedTimeseriesWidgetLocalDatasetRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesWidgetLocalDatasetRequestRequestType) IsValid() bool {
	for _, existing := range allowedTimeseriesWidgetLocalDatasetRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesWidgetLocalDatasetRequestRequestType value.
func (v TimeseriesWidgetLocalDatasetRequestRequestType) Ptr() *TimeseriesWidgetLocalDatasetRequestRequestType {
	return &v
}

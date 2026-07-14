// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetReportScheduleResourceType The type of resource targeted by a dataset report schedule.
type DatasetReportScheduleResourceType string

// List of DatasetReportScheduleResourceType.
const (
	DATASETREPORTSCHEDULERESOURCETYPE_WIDGET_DATASET_LIST DatasetReportScheduleResourceType = "widget_dataset_list"
)

var allowedDatasetReportScheduleResourceTypeEnumValues = []DatasetReportScheduleResourceType{
	DATASETREPORTSCHEDULERESOURCETYPE_WIDGET_DATASET_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatasetReportScheduleResourceType) GetAllowedValues() []DatasetReportScheduleResourceType {
	return allowedDatasetReportScheduleResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatasetReportScheduleResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatasetReportScheduleResourceType(value)
	return nil
}

// NewDatasetReportScheduleResourceTypeFromValue returns a pointer to a valid DatasetReportScheduleResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatasetReportScheduleResourceTypeFromValue(v string) (*DatasetReportScheduleResourceType, error) {
	ev := DatasetReportScheduleResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatasetReportScheduleResourceType: valid values are %v", v, allowedDatasetReportScheduleResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatasetReportScheduleResourceType) IsValid() bool {
	for _, existing := range allowedDatasetReportScheduleResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatasetReportScheduleResourceType value.
func (v DatasetReportScheduleResourceType) Ptr() *DatasetReportScheduleResourceType {
	return &v
}

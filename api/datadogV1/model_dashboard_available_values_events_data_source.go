// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardAvailableValuesEventsDataSource The events-based data source for the query.
type DashboardAvailableValuesEventsDataSource string

// List of DashboardAvailableValuesEventsDataSource.
const (
	DASHBOARDAVAILABLEVALUESEVENTSDATASOURCE_SPANS DashboardAvailableValuesEventsDataSource = "spans"
	DASHBOARDAVAILABLEVALUESEVENTSDATASOURCE_LOGS  DashboardAvailableValuesEventsDataSource = "logs"
	DASHBOARDAVAILABLEVALUESEVENTSDATASOURCE_RUM   DashboardAvailableValuesEventsDataSource = "rum"
)

var allowedDashboardAvailableValuesEventsDataSourceEnumValues = []DashboardAvailableValuesEventsDataSource{
	DASHBOARDAVAILABLEVALUESEVENTSDATASOURCE_SPANS,
	DASHBOARDAVAILABLEVALUESEVENTSDATASOURCE_LOGS,
	DASHBOARDAVAILABLEVALUESEVENTSDATASOURCE_RUM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardAvailableValuesEventsDataSource) GetAllowedValues() []DashboardAvailableValuesEventsDataSource {
	return allowedDashboardAvailableValuesEventsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardAvailableValuesEventsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardAvailableValuesEventsDataSource(value)
	return nil
}

// NewDashboardAvailableValuesEventsDataSourceFromValue returns a pointer to a valid DashboardAvailableValuesEventsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardAvailableValuesEventsDataSourceFromValue(v string) (*DashboardAvailableValuesEventsDataSource, error) {
	ev := DashboardAvailableValuesEventsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardAvailableValuesEventsDataSource: valid values are %v", v, allowedDashboardAvailableValuesEventsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardAvailableValuesEventsDataSource) IsValid() bool {
	for _, existing := range allowedDashboardAvailableValuesEventsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardAvailableValuesEventsDataSource value.
func (v DashboardAvailableValuesEventsDataSource) Ptr() *DashboardAvailableValuesEventsDataSource {
	return &v
}

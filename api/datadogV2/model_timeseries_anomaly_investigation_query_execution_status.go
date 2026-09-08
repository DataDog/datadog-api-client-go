// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationQueryExecutionStatus Current execution status for a named query.
type TimeseriesAnomalyInvestigationQueryExecutionStatus string

// List of TimeseriesAnomalyInvestigationQueryExecutionStatus.
const (
	TIMESERIESANOMALYINVESTIGATIONQUERYEXECUTIONSTATUS_RUNNING TimeseriesAnomalyInvestigationQueryExecutionStatus = "running"
	TIMESERIESANOMALYINVESTIGATIONQUERYEXECUTIONSTATUS_DONE    TimeseriesAnomalyInvestigationQueryExecutionStatus = "done"
)

var allowedTimeseriesAnomalyInvestigationQueryExecutionStatusEnumValues = []TimeseriesAnomalyInvestigationQueryExecutionStatus{
	TIMESERIESANOMALYINVESTIGATIONQUERYEXECUTIONSTATUS_RUNNING,
	TIMESERIESANOMALYINVESTIGATIONQUERYEXECUTIONSTATUS_DONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesAnomalyInvestigationQueryExecutionStatus) GetAllowedValues() []TimeseriesAnomalyInvestigationQueryExecutionStatus {
	return allowedTimeseriesAnomalyInvestigationQueryExecutionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesAnomalyInvestigationQueryExecutionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesAnomalyInvestigationQueryExecutionStatus(value)
	return nil
}

// NewTimeseriesAnomalyInvestigationQueryExecutionStatusFromValue returns a pointer to a valid TimeseriesAnomalyInvestigationQueryExecutionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesAnomalyInvestigationQueryExecutionStatusFromValue(v string) (*TimeseriesAnomalyInvestigationQueryExecutionStatus, error) {
	ev := TimeseriesAnomalyInvestigationQueryExecutionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesAnomalyInvestigationQueryExecutionStatus: valid values are %v", v, allowedTimeseriesAnomalyInvestigationQueryExecutionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesAnomalyInvestigationQueryExecutionStatus) IsValid() bool {
	for _, existing := range allowedTimeseriesAnomalyInvestigationQueryExecutionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesAnomalyInvestigationQueryExecutionStatus value.
func (v TimeseriesAnomalyInvestigationQueryExecutionStatus) Ptr() *TimeseriesAnomalyInvestigationQueryExecutionStatus {
	return &v
}

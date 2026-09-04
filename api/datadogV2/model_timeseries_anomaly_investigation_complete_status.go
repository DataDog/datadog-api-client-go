// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationCompleteStatus Status value indicating successful completion.
type TimeseriesAnomalyInvestigationCompleteStatus string

// List of TimeseriesAnomalyInvestigationCompleteStatus.
const (
	TIMESERIESANOMALYINVESTIGATIONCOMPLETESTATUS_COMPLETE TimeseriesAnomalyInvestigationCompleteStatus = "complete"
)

var allowedTimeseriesAnomalyInvestigationCompleteStatusEnumValues = []TimeseriesAnomalyInvestigationCompleteStatus{
	TIMESERIESANOMALYINVESTIGATIONCOMPLETESTATUS_COMPLETE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesAnomalyInvestigationCompleteStatus) GetAllowedValues() []TimeseriesAnomalyInvestigationCompleteStatus {
	return allowedTimeseriesAnomalyInvestigationCompleteStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesAnomalyInvestigationCompleteStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesAnomalyInvestigationCompleteStatus(value)
	return nil
}

// NewTimeseriesAnomalyInvestigationCompleteStatusFromValue returns a pointer to a valid TimeseriesAnomalyInvestigationCompleteStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesAnomalyInvestigationCompleteStatusFromValue(v string) (*TimeseriesAnomalyInvestigationCompleteStatus, error) {
	ev := TimeseriesAnomalyInvestigationCompleteStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesAnomalyInvestigationCompleteStatus: valid values are %v", v, allowedTimeseriesAnomalyInvestigationCompleteStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesAnomalyInvestigationCompleteStatus) IsValid() bool {
	for _, existing := range allowedTimeseriesAnomalyInvestigationCompleteStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesAnomalyInvestigationCompleteStatus value.
func (v TimeseriesAnomalyInvestigationCompleteStatus) Ptr() *TimeseriesAnomalyInvestigationCompleteStatus {
	return &v
}

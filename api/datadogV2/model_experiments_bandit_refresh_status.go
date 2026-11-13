// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditRefreshStatus The status of the refresh.
type ExperimentsBanditRefreshStatus string

// List of ExperimentsBanditRefreshStatus.
const (
	EXPERIMENTSBANDITREFRESHSTATUS_PENDING   ExperimentsBanditRefreshStatus = "PENDING"
	EXPERIMENTSBANDITREFRESHSTATUS_RUNNING   ExperimentsBanditRefreshStatus = "RUNNING"
	EXPERIMENTSBANDITREFRESHSTATUS_SUCCEEDED ExperimentsBanditRefreshStatus = "SUCCEEDED"
	EXPERIMENTSBANDITREFRESHSTATUS_FAILED    ExperimentsBanditRefreshStatus = "FAILED"
	EXPERIMENTSBANDITREFRESHSTATUS_CANCELLED ExperimentsBanditRefreshStatus = "CANCELLED"
)

var allowedExperimentsBanditRefreshStatusEnumValues = []ExperimentsBanditRefreshStatus{
	EXPERIMENTSBANDITREFRESHSTATUS_PENDING,
	EXPERIMENTSBANDITREFRESHSTATUS_RUNNING,
	EXPERIMENTSBANDITREFRESHSTATUS_SUCCEEDED,
	EXPERIMENTSBANDITREFRESHSTATUS_FAILED,
	EXPERIMENTSBANDITREFRESHSTATUS_CANCELLED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditRefreshStatus) GetAllowedValues() []ExperimentsBanditRefreshStatus {
	return allowedExperimentsBanditRefreshStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditRefreshStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditRefreshStatus(value)
	return nil
}

// NewExperimentsBanditRefreshStatusFromValue returns a pointer to a valid ExperimentsBanditRefreshStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditRefreshStatusFromValue(v string) (*ExperimentsBanditRefreshStatus, error) {
	ev := ExperimentsBanditRefreshStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditRefreshStatus: valid values are %v", v, allowedExperimentsBanditRefreshStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditRefreshStatus) IsValid() bool {
	for _, existing := range allowedExperimentsBanditRefreshStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditRefreshStatus value.
func (v ExperimentsBanditRefreshStatus) Ptr() *ExperimentsBanditRefreshStatus {
	return &v
}

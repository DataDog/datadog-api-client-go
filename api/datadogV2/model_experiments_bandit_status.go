// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditStatus The bandit's status. STOPPED means the bandit was stopped and serves the variant chosen then. COMPLETED means it reached its scheduled end and serves its fallback.
type ExperimentsBanditStatus string

// List of ExperimentsBanditStatus.
const (
	EXPERIMENTSBANDITSTATUS_DRAFT     ExperimentsBanditStatus = "DRAFT"
	EXPERIMENTSBANDITSTATUS_SCHEDULED ExperimentsBanditStatus = "SCHEDULED"
	EXPERIMENTSBANDITSTATUS_RUNNING   ExperimentsBanditStatus = "RUNNING"
	EXPERIMENTSBANDITSTATUS_PAUSED    ExperimentsBanditStatus = "PAUSED"
	EXPERIMENTSBANDITSTATUS_STOPPED   ExperimentsBanditStatus = "STOPPED"
	EXPERIMENTSBANDITSTATUS_COMPLETED ExperimentsBanditStatus = "COMPLETED"
)

var allowedExperimentsBanditStatusEnumValues = []ExperimentsBanditStatus{
	EXPERIMENTSBANDITSTATUS_DRAFT,
	EXPERIMENTSBANDITSTATUS_SCHEDULED,
	EXPERIMENTSBANDITSTATUS_RUNNING,
	EXPERIMENTSBANDITSTATUS_PAUSED,
	EXPERIMENTSBANDITSTATUS_STOPPED,
	EXPERIMENTSBANDITSTATUS_COMPLETED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditStatus) GetAllowedValues() []ExperimentsBanditStatus {
	return allowedExperimentsBanditStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditStatus(value)
	return nil
}

// NewExperimentsBanditStatusFromValue returns a pointer to a valid ExperimentsBanditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditStatusFromValue(v string) (*ExperimentsBanditStatus, error) {
	ev := ExperimentsBanditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditStatus: valid values are %v", v, allowedExperimentsBanditStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditStatus) IsValid() bool {
	for _, existing := range allowedExperimentsBanditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditStatus value.
func (v ExperimentsBanditStatus) Ptr() *ExperimentsBanditStatus {
	return &v
}

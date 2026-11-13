// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditArmDataStatus Where the arm's reward variance came from, or why it's unavailable.
type ExperimentsBanditArmDataStatus string

// List of ExperimentsBanditArmDataStatus.
const (
	EXPERIMENTSBANDITARMDATASTATUS_USABLE                 ExperimentsBanditArmDataStatus = "usable"
	EXPERIMENTSBANDITARMDATASTATUS_INSUFFICIENT_DATA      ExperimentsBanditArmDataStatus = "insufficient_data"
	EXPERIMENTSBANDITARMDATASTATUS_ZERO_VARIANCE          ExperimentsBanditArmDataStatus = "zero_variance"
	EXPERIMENTSBANDITARMDATASTATUS_POOLED_WITHIN_VARIANCE ExperimentsBanditArmDataStatus = "pooled_within_variance"
	EXPERIMENTSBANDITARMDATASTATUS_POOLED_TOTAL_VARIANCE  ExperimentsBanditArmDataStatus = "pooled_total_variance"
)

var allowedExperimentsBanditArmDataStatusEnumValues = []ExperimentsBanditArmDataStatus{
	EXPERIMENTSBANDITARMDATASTATUS_USABLE,
	EXPERIMENTSBANDITARMDATASTATUS_INSUFFICIENT_DATA,
	EXPERIMENTSBANDITARMDATASTATUS_ZERO_VARIANCE,
	EXPERIMENTSBANDITARMDATASTATUS_POOLED_WITHIN_VARIANCE,
	EXPERIMENTSBANDITARMDATASTATUS_POOLED_TOTAL_VARIANCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditArmDataStatus) GetAllowedValues() []ExperimentsBanditArmDataStatus {
	return allowedExperimentsBanditArmDataStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditArmDataStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditArmDataStatus(value)
	return nil
}

// NewExperimentsBanditArmDataStatusFromValue returns a pointer to a valid ExperimentsBanditArmDataStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditArmDataStatusFromValue(v string) (*ExperimentsBanditArmDataStatus, error) {
	ev := ExperimentsBanditArmDataStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditArmDataStatus: valid values are %v", v, allowedExperimentsBanditArmDataStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditArmDataStatus) IsValid() bool {
	for _, existing := range allowedExperimentsBanditArmDataStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditArmDataStatus value.
func (v ExperimentsBanditArmDataStatus) Ptr() *ExperimentsBanditArmDataStatus {
	return &v
}

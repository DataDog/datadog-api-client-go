// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionStepStatus The current status of the step.
type ExecutionStepStatus string

// List of ExecutionStepStatus.
const (
	EXECUTIONSTEPSTATUS_PENDING   ExecutionStepStatus = "PENDING"
	EXECUTIONSTEPSTATUS_RUNNING   ExecutionStepStatus = "RUNNING"
	EXECUTIONSTEPSTATUS_COMPLETED ExecutionStepStatus = "COMPLETED"
	EXECUTIONSTEPSTATUS_FAILED    ExecutionStepStatus = "FAILED"
	EXECUTIONSTEPSTATUS_CANCELED  ExecutionStepStatus = "CANCELED"
)

var allowedExecutionStepStatusEnumValues = []ExecutionStepStatus{
	EXECUTIONSTEPSTATUS_PENDING,
	EXECUTIONSTEPSTATUS_RUNNING,
	EXECUTIONSTEPSTATUS_COMPLETED,
	EXECUTIONSTEPSTATUS_FAILED,
	EXECUTIONSTEPSTATUS_CANCELED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExecutionStepStatus) GetAllowedValues() []ExecutionStepStatus {
	return allowedExecutionStepStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExecutionStepStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExecutionStepStatus(value)
	return nil
}

// NewExecutionStepStatusFromValue returns a pointer to a valid ExecutionStepStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExecutionStepStatusFromValue(v string) (*ExecutionStepStatus, error) {
	ev := ExecutionStepStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExecutionStepStatus: valid values are %v", v, allowedExecutionStepStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExecutionStepStatus) IsValid() bool {
	for _, existing := range allowedExecutionStepStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExecutionStepStatus value.
func (v ExecutionStepStatus) Ptr() *ExecutionStepStatus {
	return &v
}

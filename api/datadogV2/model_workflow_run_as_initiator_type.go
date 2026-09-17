// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowRunAsInitiatorType The initiator run-as type.
type WorkflowRunAsInitiatorType string

// List of WorkflowRunAsInitiatorType.
const (
	WORKFLOWRUNASINITIATORTYPE_INITIATOR WorkflowRunAsInitiatorType = "initiator"
)

var allowedWorkflowRunAsInitiatorTypeEnumValues = []WorkflowRunAsInitiatorType{
	WORKFLOWRUNASINITIATORTYPE_INITIATOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WorkflowRunAsInitiatorType) GetAllowedValues() []WorkflowRunAsInitiatorType {
	return allowedWorkflowRunAsInitiatorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WorkflowRunAsInitiatorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WorkflowRunAsInitiatorType(value)
	return nil
}

// NewWorkflowRunAsInitiatorTypeFromValue returns a pointer to a valid WorkflowRunAsInitiatorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWorkflowRunAsInitiatorTypeFromValue(v string) (*WorkflowRunAsInitiatorType, error) {
	ev := WorkflowRunAsInitiatorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WorkflowRunAsInitiatorType: valid values are %v", v, allowedWorkflowRunAsInitiatorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WorkflowRunAsInitiatorType) IsValid() bool {
	for _, existing := range allowedWorkflowRunAsInitiatorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkflowRunAsInitiatorType value.
func (v WorkflowRunAsInitiatorType) Ptr() *WorkflowRunAsInitiatorType {
	return &v
}

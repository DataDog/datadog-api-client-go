// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowRunAsOwnerType The owner run-as type.
type WorkflowRunAsOwnerType string

// List of WorkflowRunAsOwnerType.
const (
	WORKFLOWRUNASOWNERTYPE_OWNER WorkflowRunAsOwnerType = "owner"
)

var allowedWorkflowRunAsOwnerTypeEnumValues = []WorkflowRunAsOwnerType{
	WORKFLOWRUNASOWNERTYPE_OWNER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WorkflowRunAsOwnerType) GetAllowedValues() []WorkflowRunAsOwnerType {
	return allowedWorkflowRunAsOwnerTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WorkflowRunAsOwnerType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WorkflowRunAsOwnerType(value)
	return nil
}

// NewWorkflowRunAsOwnerTypeFromValue returns a pointer to a valid WorkflowRunAsOwnerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWorkflowRunAsOwnerTypeFromValue(v string) (*WorkflowRunAsOwnerType, error) {
	ev := WorkflowRunAsOwnerType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WorkflowRunAsOwnerType: valid values are %v", v, allowedWorkflowRunAsOwnerTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WorkflowRunAsOwnerType) IsValid() bool {
	for _, existing := range allowedWorkflowRunAsOwnerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkflowRunAsOwnerType value.
func (v WorkflowRunAsOwnerType) Ptr() *WorkflowRunAsOwnerType {
	return &v
}

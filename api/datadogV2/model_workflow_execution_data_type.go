// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowExecutionDataType The resource type for workflow executions.
type WorkflowExecutionDataType string

// List of WorkflowExecutionDataType.
const (
	WORKFLOWEXECUTIONDATATYPE_WORKFLOW_EXECUTIONS WorkflowExecutionDataType = "workflow-executions"
)

var allowedWorkflowExecutionDataTypeEnumValues = []WorkflowExecutionDataType{
	WORKFLOWEXECUTIONDATATYPE_WORKFLOW_EXECUTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WorkflowExecutionDataType) GetAllowedValues() []WorkflowExecutionDataType {
	return allowedWorkflowExecutionDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WorkflowExecutionDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WorkflowExecutionDataType(value)
	return nil
}

// NewWorkflowExecutionDataTypeFromValue returns a pointer to a valid WorkflowExecutionDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWorkflowExecutionDataTypeFromValue(v string) (*WorkflowExecutionDataType, error) {
	ev := WorkflowExecutionDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WorkflowExecutionDataType: valid values are %v", v, allowedWorkflowExecutionDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WorkflowExecutionDataType) IsValid() bool {
	for _, existing := range allowedWorkflowExecutionDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkflowExecutionDataType value.
func (v WorkflowExecutionDataType) Ptr() *WorkflowExecutionDataType {
	return &v
}

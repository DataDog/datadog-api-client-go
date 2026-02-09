// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowHeadlessExecutionResponseType The type for workflow headless execution response
type WorkflowHeadlessExecutionResponseType string

// List of WorkflowHeadlessExecutionResponseType.
const (
	WORKFLOWHEADLESSEXECUTIONRESPONSETYPE_WORKFLOW_HEADLESS_EXECUTION WorkflowHeadlessExecutionResponseType = "workflow_headless_execution"
)

var allowedWorkflowHeadlessExecutionResponseTypeEnumValues = []WorkflowHeadlessExecutionResponseType{
	WORKFLOWHEADLESSEXECUTIONRESPONSETYPE_WORKFLOW_HEADLESS_EXECUTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WorkflowHeadlessExecutionResponseType) GetAllowedValues() []WorkflowHeadlessExecutionResponseType {
	return allowedWorkflowHeadlessExecutionResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WorkflowHeadlessExecutionResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WorkflowHeadlessExecutionResponseType(value)
	return nil
}

// NewWorkflowHeadlessExecutionResponseTypeFromValue returns a pointer to a valid WorkflowHeadlessExecutionResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWorkflowHeadlessExecutionResponseTypeFromValue(v string) (*WorkflowHeadlessExecutionResponseType, error) {
	ev := WorkflowHeadlessExecutionResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WorkflowHeadlessExecutionResponseType: valid values are %v", v, allowedWorkflowHeadlessExecutionResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WorkflowHeadlessExecutionResponseType) IsValid() bool {
	for _, existing := range allowedWorkflowHeadlessExecutionResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkflowHeadlessExecutionResponseType value.
func (v WorkflowHeadlessExecutionResponseType) Ptr() *WorkflowHeadlessExecutionResponseType {
	return &v
}

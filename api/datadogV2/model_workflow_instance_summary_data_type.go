// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowInstanceSummaryDataType The resource type for workflow instances.
type WorkflowInstanceSummaryDataType string

// List of WorkflowInstanceSummaryDataType.
const (
	WORKFLOWINSTANCESUMMARYDATATYPE_WORKFLOW_INSTANCES WorkflowInstanceSummaryDataType = "workflow-instances"
)

var allowedWorkflowInstanceSummaryDataTypeEnumValues = []WorkflowInstanceSummaryDataType{
	WORKFLOWINSTANCESUMMARYDATATYPE_WORKFLOW_INSTANCES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WorkflowInstanceSummaryDataType) GetAllowedValues() []WorkflowInstanceSummaryDataType {
	return allowedWorkflowInstanceSummaryDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WorkflowInstanceSummaryDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WorkflowInstanceSummaryDataType(value)
	return nil
}

// NewWorkflowInstanceSummaryDataTypeFromValue returns a pointer to a valid WorkflowInstanceSummaryDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWorkflowInstanceSummaryDataTypeFromValue(v string) (*WorkflowInstanceSummaryDataType, error) {
	ev := WorkflowInstanceSummaryDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WorkflowInstanceSummaryDataType: valid values are %v", v, allowedWorkflowInstanceSummaryDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WorkflowInstanceSummaryDataType) IsValid() bool {
	for _, existing := range allowedWorkflowInstanceSummaryDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkflowInstanceSummaryDataType value.
func (v WorkflowInstanceSummaryDataType) Ptr() *WorkflowInstanceSummaryDataType {
	return &v
}

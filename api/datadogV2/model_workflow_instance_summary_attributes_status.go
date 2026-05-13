// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowInstanceSummaryAttributesStatus The current status of the step.
type WorkflowInstanceSummaryAttributesStatus string

// List of WorkflowInstanceSummaryAttributesStatus.
const (
	WORKFLOWINSTANCESUMMARYATTRIBUTESSTATUS_PENDING   WorkflowInstanceSummaryAttributesStatus = "PENDING"
	WORKFLOWINSTANCESUMMARYATTRIBUTESSTATUS_RUNNING   WorkflowInstanceSummaryAttributesStatus = "RUNNING"
	WORKFLOWINSTANCESUMMARYATTRIBUTESSTATUS_COMPLETED WorkflowInstanceSummaryAttributesStatus = "COMPLETED"
	WORKFLOWINSTANCESUMMARYATTRIBUTESSTATUS_FAILED    WorkflowInstanceSummaryAttributesStatus = "FAILED"
	WORKFLOWINSTANCESUMMARYATTRIBUTESSTATUS_CANCELED  WorkflowInstanceSummaryAttributesStatus = "CANCELED"
)

var allowedWorkflowInstanceSummaryAttributesStatusEnumValues = []WorkflowInstanceSummaryAttributesStatus{
	WORKFLOWINSTANCESUMMARYATTRIBUTESSTATUS_PENDING,
	WORKFLOWINSTANCESUMMARYATTRIBUTESSTATUS_RUNNING,
	WORKFLOWINSTANCESUMMARYATTRIBUTESSTATUS_COMPLETED,
	WORKFLOWINSTANCESUMMARYATTRIBUTESSTATUS_FAILED,
	WORKFLOWINSTANCESUMMARYATTRIBUTESSTATUS_CANCELED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WorkflowInstanceSummaryAttributesStatus) GetAllowedValues() []WorkflowInstanceSummaryAttributesStatus {
	return allowedWorkflowInstanceSummaryAttributesStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WorkflowInstanceSummaryAttributesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WorkflowInstanceSummaryAttributesStatus(value)
	return nil
}

// NewWorkflowInstanceSummaryAttributesStatusFromValue returns a pointer to a valid WorkflowInstanceSummaryAttributesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWorkflowInstanceSummaryAttributesStatusFromValue(v string) (*WorkflowInstanceSummaryAttributesStatus, error) {
	ev := WorkflowInstanceSummaryAttributesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WorkflowInstanceSummaryAttributesStatus: valid values are %v", v, allowedWorkflowInstanceSummaryAttributesStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WorkflowInstanceSummaryAttributesStatus) IsValid() bool {
	for _, existing := range allowedWorkflowInstanceSummaryAttributesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkflowInstanceSummaryAttributesStatus value.
func (v WorkflowInstanceSummaryAttributesStatus) Ptr() *WorkflowInstanceSummaryAttributesStatus {
	return &v
}

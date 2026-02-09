// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowWebhookExecutionResponseType The type for workflow webhook execution response
type WorkflowWebhookExecutionResponseType string

// List of WorkflowWebhookExecutionResponseType.
const (
	WORKFLOWWEBHOOKEXECUTIONRESPONSETYPE_WORKFLOW_WEBHOOK_EXECUTION WorkflowWebhookExecutionResponseType = "workflow_webhook_execution"
)

var allowedWorkflowWebhookExecutionResponseTypeEnumValues = []WorkflowWebhookExecutionResponseType{
	WORKFLOWWEBHOOKEXECUTIONRESPONSETYPE_WORKFLOW_WEBHOOK_EXECUTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WorkflowWebhookExecutionResponseType) GetAllowedValues() []WorkflowWebhookExecutionResponseType {
	return allowedWorkflowWebhookExecutionResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WorkflowWebhookExecutionResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WorkflowWebhookExecutionResponseType(value)
	return nil
}

// NewWorkflowWebhookExecutionResponseTypeFromValue returns a pointer to a valid WorkflowWebhookExecutionResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWorkflowWebhookExecutionResponseTypeFromValue(v string) (*WorkflowWebhookExecutionResponseType, error) {
	ev := WorkflowWebhookExecutionResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WorkflowWebhookExecutionResponseType: valid values are %v", v, allowedWorkflowWebhookExecutionResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WorkflowWebhookExecutionResponseType) IsValid() bool {
	for _, existing := range allowedWorkflowWebhookExecutionResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkflowWebhookExecutionResponseType value.
func (v WorkflowWebhookExecutionResponseType) Ptr() *WorkflowWebhookExecutionResponseType {
	return &v
}

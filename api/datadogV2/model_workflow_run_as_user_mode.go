// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowRunAsUserMode The effective type of identity used to run the workflow.
type WorkflowRunAsUserMode string

// List of WorkflowRunAsUserMode.
const (
	WORKFLOWRUNASUSERMODE_OWNER           WorkflowRunAsUserMode = "owner"
	WORKFLOWRUNASUSERMODE_SERVICE_ACCOUNT WorkflowRunAsUserMode = "service_account"
	WORKFLOWRUNASUSERMODE_INITIATOR       WorkflowRunAsUserMode = "initiator"
)

var allowedWorkflowRunAsUserModeEnumValues = []WorkflowRunAsUserMode{
	WORKFLOWRUNASUSERMODE_OWNER,
	WORKFLOWRUNASUSERMODE_SERVICE_ACCOUNT,
	WORKFLOWRUNASUSERMODE_INITIATOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WorkflowRunAsUserMode) GetAllowedValues() []WorkflowRunAsUserMode {
	return allowedWorkflowRunAsUserModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WorkflowRunAsUserMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WorkflowRunAsUserMode(value)
	return nil
}

// NewWorkflowRunAsUserModeFromValue returns a pointer to a valid WorkflowRunAsUserMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWorkflowRunAsUserModeFromValue(v string) (*WorkflowRunAsUserMode, error) {
	ev := WorkflowRunAsUserMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WorkflowRunAsUserMode: valid values are %v", v, allowedWorkflowRunAsUserModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WorkflowRunAsUserMode) IsValid() bool {
	for _, existing := range allowedWorkflowRunAsUserModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkflowRunAsUserMode value.
func (v WorkflowRunAsUserMode) Ptr() *WorkflowRunAsUserMode {
	return &v
}

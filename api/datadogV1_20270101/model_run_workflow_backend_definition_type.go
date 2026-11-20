// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RunWorkflowBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type RunWorkflowBackendDefinitionType string

// List of RunWorkflowBackendDefinitionType.
const (
	RUNWORKFLOWBACKENDDEFINITIONTYPE_RUN_WORKFLOW RunWorkflowBackendDefinitionType = "run_workflow"
)

var allowedRunWorkflowBackendDefinitionTypeEnumValues = []RunWorkflowBackendDefinitionType{
	RUNWORKFLOWBACKENDDEFINITIONTYPE_RUN_WORKFLOW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RunWorkflowBackendDefinitionType) GetAllowedValues() []RunWorkflowBackendDefinitionType {
	return allowedRunWorkflowBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RunWorkflowBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RunWorkflowBackendDefinitionType(value)
	return nil
}

// NewRunWorkflowBackendDefinitionTypeFromValue returns a pointer to a valid RunWorkflowBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRunWorkflowBackendDefinitionTypeFromValue(v string) (*RunWorkflowBackendDefinitionType, error) {
	ev := RunWorkflowBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RunWorkflowBackendDefinitionType: valid values are %v", v, allowedRunWorkflowBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RunWorkflowBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedRunWorkflowBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RunWorkflowBackendDefinitionType value.
func (v RunWorkflowBackendDefinitionType) Ptr() *RunWorkflowBackendDefinitionType {
	return &v
}

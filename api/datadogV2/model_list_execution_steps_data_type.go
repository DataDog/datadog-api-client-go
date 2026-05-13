// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListExecutionStepsDataType The resource type for workflow execution steps.
type ListExecutionStepsDataType string

// List of ListExecutionStepsDataType.
const (
	LISTEXECUTIONSTEPSDATATYPE_WORKFLOW_EXECUTION_STEPS ListExecutionStepsDataType = "workflow-execution-steps"
)

var allowedListExecutionStepsDataTypeEnumValues = []ListExecutionStepsDataType{
	LISTEXECUTIONSTEPSDATATYPE_WORKFLOW_EXECUTION_STEPS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListExecutionStepsDataType) GetAllowedValues() []ListExecutionStepsDataType {
	return allowedListExecutionStepsDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListExecutionStepsDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListExecutionStepsDataType(value)
	return nil
}

// NewListExecutionStepsDataTypeFromValue returns a pointer to a valid ListExecutionStepsDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListExecutionStepsDataTypeFromValue(v string) (*ListExecutionStepsDataType, error) {
	ev := ListExecutionStepsDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListExecutionStepsDataType: valid values are %v", v, allowedListExecutionStepsDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListExecutionStepsDataType) IsValid() bool {
	for _, existing := range allowedListExecutionStepsDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListExecutionStepsDataType value.
func (v ListExecutionStepsDataType) Ptr() *ListExecutionStepsDataType {
	return &v
}

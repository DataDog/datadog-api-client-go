// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowFavoriteRequestType The type for workflow favorite request
type WorkflowFavoriteRequestType string

// List of WorkflowFavoriteRequestType.
const (
	WORKFLOWFAVORITEREQUESTTYPE_WORKFLOW_FAVORITE_REQUEST WorkflowFavoriteRequestType = "workflow_favorite_request"
)

var allowedWorkflowFavoriteRequestTypeEnumValues = []WorkflowFavoriteRequestType{
	WORKFLOWFAVORITEREQUESTTYPE_WORKFLOW_FAVORITE_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WorkflowFavoriteRequestType) GetAllowedValues() []WorkflowFavoriteRequestType {
	return allowedWorkflowFavoriteRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WorkflowFavoriteRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WorkflowFavoriteRequestType(value)
	return nil
}

// NewWorkflowFavoriteRequestTypeFromValue returns a pointer to a valid WorkflowFavoriteRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWorkflowFavoriteRequestTypeFromValue(v string) (*WorkflowFavoriteRequestType, error) {
	ev := WorkflowFavoriteRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WorkflowFavoriteRequestType: valid values are %v", v, allowedWorkflowFavoriteRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WorkflowFavoriteRequestType) IsValid() bool {
	for _, existing := range allowedWorkflowFavoriteRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkflowFavoriteRequestType value.
func (v WorkflowFavoriteRequestType) Ptr() *WorkflowFavoriteRequestType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowRunAsServiceAccountType The service account run-as type.
type WorkflowRunAsServiceAccountType string

// List of WorkflowRunAsServiceAccountType.
const (
	WORKFLOWRUNASSERVICEACCOUNTTYPE_SERVICE_ACCOUNT WorkflowRunAsServiceAccountType = "service_account"
)

var allowedWorkflowRunAsServiceAccountTypeEnumValues = []WorkflowRunAsServiceAccountType{
	WORKFLOWRUNASSERVICEACCOUNTTYPE_SERVICE_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WorkflowRunAsServiceAccountType) GetAllowedValues() []WorkflowRunAsServiceAccountType {
	return allowedWorkflowRunAsServiceAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WorkflowRunAsServiceAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WorkflowRunAsServiceAccountType(value)
	return nil
}

// NewWorkflowRunAsServiceAccountTypeFromValue returns a pointer to a valid WorkflowRunAsServiceAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWorkflowRunAsServiceAccountTypeFromValue(v string) (*WorkflowRunAsServiceAccountType, error) {
	ev := WorkflowRunAsServiceAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WorkflowRunAsServiceAccountType: valid values are %v", v, allowedWorkflowRunAsServiceAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WorkflowRunAsServiceAccountType) IsValid() bool {
	for _, existing := range allowedWorkflowRunAsServiceAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkflowRunAsServiceAccountType value.
func (v WorkflowRunAsServiceAccountType) Ptr() *WorkflowRunAsServiceAccountType {
	return &v
}

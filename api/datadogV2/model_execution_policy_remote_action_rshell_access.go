// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyRemoteActionRshellAccess The level of remote shell access granted for the target paths.
type ExecutionPolicyRemoteActionRshellAccess string

// List of ExecutionPolicyRemoteActionRshellAccess.
const (
	EXECUTIONPOLICYREMOTEACTIONRSHELLACCESS_READ_ONLY  ExecutionPolicyRemoteActionRshellAccess = "read_only"
	EXECUTIONPOLICYREMOTEACTIONRSHELLACCESS_READ_WRITE ExecutionPolicyRemoteActionRshellAccess = "read_write"
)

var allowedExecutionPolicyRemoteActionRshellAccessEnumValues = []ExecutionPolicyRemoteActionRshellAccess{
	EXECUTIONPOLICYREMOTEACTIONRSHELLACCESS_READ_ONLY,
	EXECUTIONPOLICYREMOTEACTIONRSHELLACCESS_READ_WRITE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExecutionPolicyRemoteActionRshellAccess) GetAllowedValues() []ExecutionPolicyRemoteActionRshellAccess {
	return allowedExecutionPolicyRemoteActionRshellAccessEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExecutionPolicyRemoteActionRshellAccess) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExecutionPolicyRemoteActionRshellAccess(value)
	return nil
}

// NewExecutionPolicyRemoteActionRshellAccessFromValue returns a pointer to a valid ExecutionPolicyRemoteActionRshellAccess
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExecutionPolicyRemoteActionRshellAccessFromValue(v string) (*ExecutionPolicyRemoteActionRshellAccess, error) {
	ev := ExecutionPolicyRemoteActionRshellAccess(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExecutionPolicyRemoteActionRshellAccess: valid values are %v", v, allowedExecutionPolicyRemoteActionRshellAccessEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExecutionPolicyRemoteActionRshellAccess) IsValid() bool {
	for _, existing := range allowedExecutionPolicyRemoteActionRshellAccessEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExecutionPolicyRemoteActionRshellAccess value.
func (v ExecutionPolicyRemoteActionRshellAccess) Ptr() *ExecutionPolicyRemoteActionRshellAccess {
	return &v
}

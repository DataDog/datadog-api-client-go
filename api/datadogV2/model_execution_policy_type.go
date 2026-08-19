// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyType The type of the resource. The value should always be `execution_policy`.
type ExecutionPolicyType string

// List of ExecutionPolicyType.
const (
	EXECUTIONPOLICYTYPE_EXECUTION_POLICY ExecutionPolicyType = "execution_policy"
)

var allowedExecutionPolicyTypeEnumValues = []ExecutionPolicyType{
	EXECUTIONPOLICYTYPE_EXECUTION_POLICY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExecutionPolicyType) GetAllowedValues() []ExecutionPolicyType {
	return allowedExecutionPolicyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExecutionPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExecutionPolicyType(value)
	return nil
}

// NewExecutionPolicyTypeFromValue returns a pointer to a valid ExecutionPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExecutionPolicyTypeFromValue(v string) (*ExecutionPolicyType, error) {
	ev := ExecutionPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExecutionPolicyType: valid values are %v", v, allowedExecutionPolicyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExecutionPolicyType) IsValid() bool {
	for _, existing := range allowedExecutionPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExecutionPolicyType value.
func (v ExecutionPolicyType) Ptr() *ExecutionPolicyType {
	return &v
}

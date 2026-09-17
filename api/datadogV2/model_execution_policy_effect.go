// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyEffect Whether the policy allows or denies matching actions.
type ExecutionPolicyEffect string

// List of ExecutionPolicyEffect.
const (
	EXECUTIONPOLICYEFFECT_ALLOW ExecutionPolicyEffect = "allow"
	EXECUTIONPOLICYEFFECT_DENY  ExecutionPolicyEffect = "deny"
)

var allowedExecutionPolicyEffectEnumValues = []ExecutionPolicyEffect{
	EXECUTIONPOLICYEFFECT_ALLOW,
	EXECUTIONPOLICYEFFECT_DENY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExecutionPolicyEffect) GetAllowedValues() []ExecutionPolicyEffect {
	return allowedExecutionPolicyEffectEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExecutionPolicyEffect) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExecutionPolicyEffect(value)
	return nil
}

// NewExecutionPolicyEffectFromValue returns a pointer to a valid ExecutionPolicyEffect
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExecutionPolicyEffectFromValue(v string) (*ExecutionPolicyEffect, error) {
	ev := ExecutionPolicyEffect(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExecutionPolicyEffect: valid values are %v", v, allowedExecutionPolicyEffectEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExecutionPolicyEffect) IsValid() bool {
	for _, existing := range allowedExecutionPolicyEffectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExecutionPolicyEffect value.
func (v ExecutionPolicyEffect) Ptr() *ExecutionPolicyEffect {
	return &v
}

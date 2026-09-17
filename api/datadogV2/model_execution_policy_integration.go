// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyIntegration The integration the action pattern applies to.
type ExecutionPolicyIntegration string

// List of ExecutionPolicyIntegration.
const (
	EXECUTIONPOLICYINTEGRATION_INTEGRATION_KUBERNETES    ExecutionPolicyIntegration = "INTEGRATION_KUBERNETES"
	EXECUTIONPOLICYINTEGRATION_INTEGRATION_SCRIPT        ExecutionPolicyIntegration = "INTEGRATION_SCRIPT"
	EXECUTIONPOLICYINTEGRATION_INTEGRATION_REMOTE_ACTION ExecutionPolicyIntegration = "INTEGRATION_REMOTE_ACTION"
)

var allowedExecutionPolicyIntegrationEnumValues = []ExecutionPolicyIntegration{
	EXECUTIONPOLICYINTEGRATION_INTEGRATION_KUBERNETES,
	EXECUTIONPOLICYINTEGRATION_INTEGRATION_SCRIPT,
	EXECUTIONPOLICYINTEGRATION_INTEGRATION_REMOTE_ACTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExecutionPolicyIntegration) GetAllowedValues() []ExecutionPolicyIntegration {
	return allowedExecutionPolicyIntegrationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExecutionPolicyIntegration) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExecutionPolicyIntegration(value)
	return nil
}

// NewExecutionPolicyIntegrationFromValue returns a pointer to a valid ExecutionPolicyIntegration
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExecutionPolicyIntegrationFromValue(v string) (*ExecutionPolicyIntegration, error) {
	ev := ExecutionPolicyIntegration(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExecutionPolicyIntegration: valid values are %v", v, allowedExecutionPolicyIntegrationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExecutionPolicyIntegration) IsValid() bool {
	for _, existing := range allowedExecutionPolicyIntegrationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExecutionPolicyIntegration value.
func (v ExecutionPolicyIntegration) Ptr() *ExecutionPolicyIntegration {
	return &v
}

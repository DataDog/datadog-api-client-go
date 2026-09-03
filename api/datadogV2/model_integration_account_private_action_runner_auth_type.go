// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAccountPrivateActionRunnerAuthType The authentication method type.
type IntegrationAccountPrivateActionRunnerAuthType string

// List of IntegrationAccountPrivateActionRunnerAuthType.
const (
	INTEGRATIONACCOUNTPRIVATEACTIONRUNNERAUTHTYPE_PRIVATE_ACTION_RUNNER IntegrationAccountPrivateActionRunnerAuthType = "private-action-runner"
)

var allowedIntegrationAccountPrivateActionRunnerAuthTypeEnumValues = []IntegrationAccountPrivateActionRunnerAuthType{
	INTEGRATIONACCOUNTPRIVATEACTIONRUNNERAUTHTYPE_PRIVATE_ACTION_RUNNER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IntegrationAccountPrivateActionRunnerAuthType) GetAllowedValues() []IntegrationAccountPrivateActionRunnerAuthType {
	return allowedIntegrationAccountPrivateActionRunnerAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IntegrationAccountPrivateActionRunnerAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IntegrationAccountPrivateActionRunnerAuthType(value)
	return nil
}

// NewIntegrationAccountPrivateActionRunnerAuthTypeFromValue returns a pointer to a valid IntegrationAccountPrivateActionRunnerAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIntegrationAccountPrivateActionRunnerAuthTypeFromValue(v string) (*IntegrationAccountPrivateActionRunnerAuthType, error) {
	ev := IntegrationAccountPrivateActionRunnerAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IntegrationAccountPrivateActionRunnerAuthType: valid values are %v", v, allowedIntegrationAccountPrivateActionRunnerAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IntegrationAccountPrivateActionRunnerAuthType) IsValid() bool {
	for _, existing := range allowedIntegrationAccountPrivateActionRunnerAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntegrationAccountPrivateActionRunnerAuthType value.
func (v IntegrationAccountPrivateActionRunnerAuthType) Ptr() *IntegrationAccountPrivateActionRunnerAuthType {
	return &v
}

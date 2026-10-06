// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountPrivateActionRunnerAuthType The authentication method type.
type DatabricksIntegrationAccountPrivateActionRunnerAuthType string

// List of DatabricksIntegrationAccountPrivateActionRunnerAuthType.
const (
	DATABRICKSINTEGRATIONACCOUNTPRIVATEACTIONRUNNERAUTHTYPE_PRIVATE_ACTION_RUNNER DatabricksIntegrationAccountPrivateActionRunnerAuthType = "private_action_runner"
)

var allowedDatabricksIntegrationAccountPrivateActionRunnerAuthTypeEnumValues = []DatabricksIntegrationAccountPrivateActionRunnerAuthType{
	DATABRICKSINTEGRATIONACCOUNTPRIVATEACTIONRUNNERAUTHTYPE_PRIVATE_ACTION_RUNNER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatabricksIntegrationAccountPrivateActionRunnerAuthType) GetAllowedValues() []DatabricksIntegrationAccountPrivateActionRunnerAuthType {
	return allowedDatabricksIntegrationAccountPrivateActionRunnerAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatabricksIntegrationAccountPrivateActionRunnerAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatabricksIntegrationAccountPrivateActionRunnerAuthType(value)
	return nil
}

// NewDatabricksIntegrationAccountPrivateActionRunnerAuthTypeFromValue returns a pointer to a valid DatabricksIntegrationAccountPrivateActionRunnerAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatabricksIntegrationAccountPrivateActionRunnerAuthTypeFromValue(v string) (*DatabricksIntegrationAccountPrivateActionRunnerAuthType, error) {
	ev := DatabricksIntegrationAccountPrivateActionRunnerAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatabricksIntegrationAccountPrivateActionRunnerAuthType: valid values are %v", v, allowedDatabricksIntegrationAccountPrivateActionRunnerAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatabricksIntegrationAccountPrivateActionRunnerAuthType) IsValid() bool {
	for _, existing := range allowedDatabricksIntegrationAccountPrivateActionRunnerAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatabricksIntegrationAccountPrivateActionRunnerAuthType value.
func (v DatabricksIntegrationAccountPrivateActionRunnerAuthType) Ptr() *DatabricksIntegrationAccountPrivateActionRunnerAuthType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceConsoleConfigResourceType Governance console config resource type.
type GovernanceConsoleConfigResourceType string

// List of GovernanceConsoleConfigResourceType.
const (
	GOVERNANCECONSOLECONFIGRESOURCETYPE_GOVERNANCE_CONSOLE_CONFIG GovernanceConsoleConfigResourceType = "governance_console_config"
)

var allowedGovernanceConsoleConfigResourceTypeEnumValues = []GovernanceConsoleConfigResourceType{
	GOVERNANCECONSOLECONFIGRESOURCETYPE_GOVERNANCE_CONSOLE_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GovernanceConsoleConfigResourceType) GetAllowedValues() []GovernanceConsoleConfigResourceType {
	return allowedGovernanceConsoleConfigResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GovernanceConsoleConfigResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GovernanceConsoleConfigResourceType(value)
	return nil
}

// NewGovernanceConsoleConfigResourceTypeFromValue returns a pointer to a valid GovernanceConsoleConfigResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGovernanceConsoleConfigResourceTypeFromValue(v string) (*GovernanceConsoleConfigResourceType, error) {
	ev := GovernanceConsoleConfigResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GovernanceConsoleConfigResourceType: valid values are %v", v, allowedGovernanceConsoleConfigResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GovernanceConsoleConfigResourceType) IsValid() bool {
	for _, existing := range allowedGovernanceConsoleConfigResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GovernanceConsoleConfigResourceType value.
func (v GovernanceConsoleConfigResourceType) Ptr() *GovernanceConsoleConfigResourceType {
	return &v
}

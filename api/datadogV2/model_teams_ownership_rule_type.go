// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsOwnershipRuleType The type of the resource. The value should always be teams_ownership_grouped_mappings.
type TeamsOwnershipRuleType string

// List of TeamsOwnershipRuleType.
const (
	TEAMSOWNERSHIPRULETYPE_TEAMS_OWNERSHIP_GROUPED_MAPPINGS TeamsOwnershipRuleType = "teams_ownership_grouped_mappings"
)

var allowedTeamsOwnershipRuleTypeEnumValues = []TeamsOwnershipRuleType{
	TEAMSOWNERSHIPRULETYPE_TEAMS_OWNERSHIP_GROUPED_MAPPINGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamsOwnershipRuleType) GetAllowedValues() []TeamsOwnershipRuleType {
	return allowedTeamsOwnershipRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamsOwnershipRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamsOwnershipRuleType(value)
	return nil
}

// NewTeamsOwnershipRuleTypeFromValue returns a pointer to a valid TeamsOwnershipRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamsOwnershipRuleTypeFromValue(v string) (*TeamsOwnershipRuleType, error) {
	ev := TeamsOwnershipRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamsOwnershipRuleType: valid values are %v", v, allowedTeamsOwnershipRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamsOwnershipRuleType) IsValid() bool {
	for _, existing := range allowedTeamsOwnershipRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamsOwnershipRuleType value.
func (v TeamsOwnershipRuleType) Ptr() *TeamsOwnershipRuleType {
	return &v
}

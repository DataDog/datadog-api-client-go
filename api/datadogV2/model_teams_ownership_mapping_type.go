// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsOwnershipMappingType The type of the resource. The value should always be teams_ownership_mappings.
type TeamsOwnershipMappingType string

// List of TeamsOwnershipMappingType.
const (
	TEAMSOWNERSHIPMAPPINGTYPE_TEAMS_OWNERSHIP_MAPPINGS TeamsOwnershipMappingType = "teams_ownership_mappings"
)

var allowedTeamsOwnershipMappingTypeEnumValues = []TeamsOwnershipMappingType{
	TEAMSOWNERSHIPMAPPINGTYPE_TEAMS_OWNERSHIP_MAPPINGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamsOwnershipMappingType) GetAllowedValues() []TeamsOwnershipMappingType {
	return allowedTeamsOwnershipMappingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamsOwnershipMappingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamsOwnershipMappingType(value)
	return nil
}

// NewTeamsOwnershipMappingTypeFromValue returns a pointer to a valid TeamsOwnershipMappingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamsOwnershipMappingTypeFromValue(v string) (*TeamsOwnershipMappingType, error) {
	ev := TeamsOwnershipMappingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamsOwnershipMappingType: valid values are %v", v, allowedTeamsOwnershipMappingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamsOwnershipMappingType) IsValid() bool {
	for _, existing := range allowedTeamsOwnershipMappingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamsOwnershipMappingType value.
func (v TeamsOwnershipMappingType) Ptr() *TeamsOwnershipMappingType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyMemberTeamType Requested team resource type.
type TeamHierarchyMemberTeamType string

// List of TeamHierarchyMemberTeamType.
const (
	TEAMHIERARCHYMEMBERTEAMTYPE_TEAMS TeamHierarchyMemberTeamType = "teams"
)

var allowedTeamHierarchyMemberTeamTypeEnumValues = []TeamHierarchyMemberTeamType{
	TEAMHIERARCHYMEMBERTEAMTYPE_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamHierarchyMemberTeamType) GetAllowedValues() []TeamHierarchyMemberTeamType {
	return allowedTeamHierarchyMemberTeamTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamHierarchyMemberTeamType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamHierarchyMemberTeamType(value)
	return nil
}

// NewTeamHierarchyMemberTeamTypeFromValue returns a pointer to a valid TeamHierarchyMemberTeamType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamHierarchyMemberTeamTypeFromValue(v string) (*TeamHierarchyMemberTeamType, error) {
	ev := TeamHierarchyMemberTeamType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamHierarchyMemberTeamType: valid values are %v", v, allowedTeamHierarchyMemberTeamTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamHierarchyMemberTeamType) IsValid() bool {
	for _, existing := range allowedTeamHierarchyMemberTeamTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamHierarchyMemberTeamType value.
func (v TeamHierarchyMemberTeamType) Ptr() *TeamHierarchyMemberTeamType {
	return &v
}

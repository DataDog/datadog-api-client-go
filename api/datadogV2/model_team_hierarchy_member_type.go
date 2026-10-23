// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyMemberType Team hierarchy membership resource type.
type TeamHierarchyMemberType string

// List of TeamHierarchyMemberType.
const (
	TEAMHIERARCHYMEMBERTYPE_USER_TEAM_RELATIONSHIP TeamHierarchyMemberType = "user_team_relationship"
)

var allowedTeamHierarchyMemberTypeEnumValues = []TeamHierarchyMemberType{
	TEAMHIERARCHYMEMBERTYPE_USER_TEAM_RELATIONSHIP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamHierarchyMemberType) GetAllowedValues() []TeamHierarchyMemberType {
	return allowedTeamHierarchyMemberTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamHierarchyMemberType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamHierarchyMemberType(value)
	return nil
}

// NewTeamHierarchyMemberTypeFromValue returns a pointer to a valid TeamHierarchyMemberType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamHierarchyMemberTypeFromValue(v string) (*TeamHierarchyMemberType, error) {
	ev := TeamHierarchyMemberType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamHierarchyMemberType: valid values are %v", v, allowedTeamHierarchyMemberTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamHierarchyMemberType) IsValid() bool {
	for _, existing := range allowedTeamHierarchyMemberTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamHierarchyMemberType value.
func (v TeamHierarchyMemberType) Ptr() *TeamHierarchyMemberType {
	return &v
}

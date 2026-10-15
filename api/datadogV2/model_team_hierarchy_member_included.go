// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyMemberIncluded - An included user or the requested team, distinguished by resource type.
type TeamHierarchyMemberIncluded struct {
	TeamHierarchyMemberUser *TeamHierarchyMemberUser
	TeamHierarchyMemberTeam *TeamHierarchyMemberTeam

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TeamHierarchyMemberUserAsTeamHierarchyMemberIncluded is a convenience function that returns TeamHierarchyMemberUser wrapped in TeamHierarchyMemberIncluded.
func TeamHierarchyMemberUserAsTeamHierarchyMemberIncluded(v *TeamHierarchyMemberUser) TeamHierarchyMemberIncluded {
	return TeamHierarchyMemberIncluded{TeamHierarchyMemberUser: v}
}

// TeamHierarchyMemberTeamAsTeamHierarchyMemberIncluded is a convenience function that returns TeamHierarchyMemberTeam wrapped in TeamHierarchyMemberIncluded.
func TeamHierarchyMemberTeamAsTeamHierarchyMemberIncluded(v *TeamHierarchyMemberTeam) TeamHierarchyMemberIncluded {
	return TeamHierarchyMemberIncluded{TeamHierarchyMemberTeam: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TeamHierarchyMemberIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TeamHierarchyMemberUser
	err = datadog.Unmarshal(data, &obj.TeamHierarchyMemberUser)
	if err == nil {
		if obj.TeamHierarchyMemberUser != nil && obj.TeamHierarchyMemberUser.UnparsedObject == nil {
			jsonTeamHierarchyMemberUser, _ := datadog.Marshal(obj.TeamHierarchyMemberUser)
			if string(jsonTeamHierarchyMemberUser) == "{}" { // empty struct
				obj.TeamHierarchyMemberUser = nil
			} else {
				match++
			}
		} else {
			obj.TeamHierarchyMemberUser = nil
		}
	} else {
		obj.TeamHierarchyMemberUser = nil
	}

	// try to unmarshal data into TeamHierarchyMemberTeam
	err = datadog.Unmarshal(data, &obj.TeamHierarchyMemberTeam)
	if err == nil {
		if obj.TeamHierarchyMemberTeam != nil && obj.TeamHierarchyMemberTeam.UnparsedObject == nil {
			jsonTeamHierarchyMemberTeam, _ := datadog.Marshal(obj.TeamHierarchyMemberTeam)
			if string(jsonTeamHierarchyMemberTeam) == "{}" { // empty struct
				obj.TeamHierarchyMemberTeam = nil
			} else {
				match++
			}
		} else {
			obj.TeamHierarchyMemberTeam = nil
		}
	} else {
		obj.TeamHierarchyMemberTeam = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TeamHierarchyMemberUser = nil
		obj.TeamHierarchyMemberTeam = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TeamHierarchyMemberIncluded) MarshalJSON() ([]byte, error) {
	if obj.TeamHierarchyMemberUser != nil {
		return datadog.Marshal(&obj.TeamHierarchyMemberUser)
	}

	if obj.TeamHierarchyMemberTeam != nil {
		return datadog.Marshal(&obj.TeamHierarchyMemberTeam)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TeamHierarchyMemberIncluded) GetActualInstance() interface{} {
	if obj.TeamHierarchyMemberUser != nil {
		return obj.TeamHierarchyMemberUser
	}

	if obj.TeamHierarchyMemberTeam != nil {
		return obj.TeamHierarchyMemberTeam
	}

	// all schemas are nil
	return nil
}

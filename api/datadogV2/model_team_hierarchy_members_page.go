// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyMembersPage - Pagination metadata for the selected pagination mode.
type TeamHierarchyMembersPage struct {
	TeamHierarchyMembersOffsetPage *TeamHierarchyMembersOffsetPage
	TeamHierarchyMembersNumberPage *TeamHierarchyMembersNumberPage

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TeamHierarchyMembersOffsetPageAsTeamHierarchyMembersPage is a convenience function that returns TeamHierarchyMembersOffsetPage wrapped in TeamHierarchyMembersPage.
func TeamHierarchyMembersOffsetPageAsTeamHierarchyMembersPage(v *TeamHierarchyMembersOffsetPage) TeamHierarchyMembersPage {
	return TeamHierarchyMembersPage{TeamHierarchyMembersOffsetPage: v}
}

// TeamHierarchyMembersNumberPageAsTeamHierarchyMembersPage is a convenience function that returns TeamHierarchyMembersNumberPage wrapped in TeamHierarchyMembersPage.
func TeamHierarchyMembersNumberPageAsTeamHierarchyMembersPage(v *TeamHierarchyMembersNumberPage) TeamHierarchyMembersPage {
	return TeamHierarchyMembersPage{TeamHierarchyMembersNumberPage: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TeamHierarchyMembersPage) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TeamHierarchyMembersOffsetPage
	err = datadog.Unmarshal(data, &obj.TeamHierarchyMembersOffsetPage)
	if err == nil {
		if obj.TeamHierarchyMembersOffsetPage != nil && obj.TeamHierarchyMembersOffsetPage.UnparsedObject == nil {
			jsonTeamHierarchyMembersOffsetPage, _ := datadog.Marshal(obj.TeamHierarchyMembersOffsetPage)
			if string(jsonTeamHierarchyMembersOffsetPage) == "{}" { // empty struct
				obj.TeamHierarchyMembersOffsetPage = nil
			} else {
				match++
			}
		} else {
			obj.TeamHierarchyMembersOffsetPage = nil
		}
	} else {
		obj.TeamHierarchyMembersOffsetPage = nil
	}

	// try to unmarshal data into TeamHierarchyMembersNumberPage
	err = datadog.Unmarshal(data, &obj.TeamHierarchyMembersNumberPage)
	if err == nil {
		if obj.TeamHierarchyMembersNumberPage != nil && obj.TeamHierarchyMembersNumberPage.UnparsedObject == nil {
			jsonTeamHierarchyMembersNumberPage, _ := datadog.Marshal(obj.TeamHierarchyMembersNumberPage)
			if string(jsonTeamHierarchyMembersNumberPage) == "{}" { // empty struct
				obj.TeamHierarchyMembersNumberPage = nil
			} else {
				match++
			}
		} else {
			obj.TeamHierarchyMembersNumberPage = nil
		}
	} else {
		obj.TeamHierarchyMembersNumberPage = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TeamHierarchyMembersOffsetPage = nil
		obj.TeamHierarchyMembersNumberPage = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TeamHierarchyMembersPage) MarshalJSON() ([]byte, error) {
	if obj.TeamHierarchyMembersOffsetPage != nil {
		return datadog.Marshal(&obj.TeamHierarchyMembersOffsetPage)
	}

	if obj.TeamHierarchyMembersNumberPage != nil {
		return datadog.Marshal(&obj.TeamHierarchyMembersNumberPage)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TeamHierarchyMembersPage) GetActualInstance() interface{} {
	if obj.TeamHierarchyMembersOffsetPage != nil {
		return obj.TeamHierarchyMembersOffsetPage
	}

	if obj.TeamHierarchyMembersNumberPage != nil {
		return obj.TeamHierarchyMembersNumberPage
	}

	// all schemas are nil
	return nil
}

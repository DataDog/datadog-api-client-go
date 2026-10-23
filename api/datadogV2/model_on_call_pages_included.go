// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPagesIncluded - A resource related to a page.
type OnCallPagesIncluded struct {
	ScheduleUser              *ScheduleUser
	TeamReference             *TeamReference
	OnCallPageScheduledAction *OnCallPageScheduledAction
	OnCallPageLink            *OnCallPageLink

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ScheduleUserAsOnCallPagesIncluded is a convenience function that returns ScheduleUser wrapped in OnCallPagesIncluded.
func ScheduleUserAsOnCallPagesIncluded(v *ScheduleUser) OnCallPagesIncluded {
	return OnCallPagesIncluded{ScheduleUser: v}
}

// TeamReferenceAsOnCallPagesIncluded is a convenience function that returns TeamReference wrapped in OnCallPagesIncluded.
func TeamReferenceAsOnCallPagesIncluded(v *TeamReference) OnCallPagesIncluded {
	return OnCallPagesIncluded{TeamReference: v}
}

// OnCallPageScheduledActionAsOnCallPagesIncluded is a convenience function that returns OnCallPageScheduledAction wrapped in OnCallPagesIncluded.
func OnCallPageScheduledActionAsOnCallPagesIncluded(v *OnCallPageScheduledAction) OnCallPagesIncluded {
	return OnCallPagesIncluded{OnCallPageScheduledAction: v}
}

// OnCallPageLinkAsOnCallPagesIncluded is a convenience function that returns OnCallPageLink wrapped in OnCallPagesIncluded.
func OnCallPageLinkAsOnCallPagesIncluded(v *OnCallPageLink) OnCallPagesIncluded {
	return OnCallPagesIncluded{OnCallPageLink: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OnCallPagesIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ScheduleUser
	err = datadog.Unmarshal(data, &obj.ScheduleUser)
	if err == nil {
		if obj.ScheduleUser != nil && obj.ScheduleUser.UnparsedObject == nil {
			jsonScheduleUser, _ := datadog.Marshal(obj.ScheduleUser)
			if string(jsonScheduleUser) == "{}" { // empty struct
				obj.ScheduleUser = nil
			} else {
				match++
			}
		} else {
			obj.ScheduleUser = nil
		}
	} else {
		obj.ScheduleUser = nil
	}

	// try to unmarshal data into TeamReference
	err = datadog.Unmarshal(data, &obj.TeamReference)
	if err == nil {
		if obj.TeamReference != nil && obj.TeamReference.UnparsedObject == nil {
			jsonTeamReference, _ := datadog.Marshal(obj.TeamReference)
			if string(jsonTeamReference) == "{}" { // empty struct
				obj.TeamReference = nil
			} else {
				match++
			}
		} else {
			obj.TeamReference = nil
		}
	} else {
		obj.TeamReference = nil
	}

	// try to unmarshal data into OnCallPageScheduledAction
	err = datadog.Unmarshal(data, &obj.OnCallPageScheduledAction)
	if err == nil {
		if obj.OnCallPageScheduledAction != nil && obj.OnCallPageScheduledAction.UnparsedObject == nil {
			jsonOnCallPageScheduledAction, _ := datadog.Marshal(obj.OnCallPageScheduledAction)
			if string(jsonOnCallPageScheduledAction) == "{}" { // empty struct
				obj.OnCallPageScheduledAction = nil
			} else {
				match++
			}
		} else {
			obj.OnCallPageScheduledAction = nil
		}
	} else {
		obj.OnCallPageScheduledAction = nil
	}

	// try to unmarshal data into OnCallPageLink
	err = datadog.Unmarshal(data, &obj.OnCallPageLink)
	if err == nil {
		if obj.OnCallPageLink != nil && obj.OnCallPageLink.UnparsedObject == nil {
			jsonOnCallPageLink, _ := datadog.Marshal(obj.OnCallPageLink)
			if string(jsonOnCallPageLink) == "{}" { // empty struct
				obj.OnCallPageLink = nil
			} else {
				match++
			}
		} else {
			obj.OnCallPageLink = nil
		}
	} else {
		obj.OnCallPageLink = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ScheduleUser = nil
		obj.TeamReference = nil
		obj.OnCallPageScheduledAction = nil
		obj.OnCallPageLink = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OnCallPagesIncluded) MarshalJSON() ([]byte, error) {
	if obj.ScheduleUser != nil {
		return datadog.Marshal(&obj.ScheduleUser)
	}

	if obj.TeamReference != nil {
		return datadog.Marshal(&obj.TeamReference)
	}

	if obj.OnCallPageScheduledAction != nil {
		return datadog.Marshal(&obj.OnCallPageScheduledAction)
	}

	if obj.OnCallPageLink != nil {
		return datadog.Marshal(&obj.OnCallPageLink)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OnCallPagesIncluded) GetActualInstance() interface{} {
	if obj.ScheduleUser != nil {
		return obj.ScheduleUser
	}

	if obj.TeamReference != nil {
		return obj.TeamReference
	}

	if obj.OnCallPageScheduledAction != nil {
		return obj.OnCallPageScheduledAction
	}

	if obj.OnCallPageLink != nil {
		return obj.OnCallPageLink
	}

	// all schemas are nil
	return nil
}

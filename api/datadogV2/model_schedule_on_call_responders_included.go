// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleOnCallRespondersIncluded - Represents a union of related resources included in the response, such as responder groups, shifts, schedules, and users.
type ScheduleOnCallRespondersIncluded struct {
	ScheduleOnCallResponderData *ScheduleOnCallResponderData
	ShiftData                   *ShiftData
	ScheduleData                *ScheduleData
	User                        *User

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ScheduleOnCallResponderDataAsScheduleOnCallRespondersIncluded is a convenience function that returns ScheduleOnCallResponderData wrapped in ScheduleOnCallRespondersIncluded.
func ScheduleOnCallResponderDataAsScheduleOnCallRespondersIncluded(v *ScheduleOnCallResponderData) ScheduleOnCallRespondersIncluded {
	return ScheduleOnCallRespondersIncluded{ScheduleOnCallResponderData: v}
}

// ShiftDataAsScheduleOnCallRespondersIncluded is a convenience function that returns ShiftData wrapped in ScheduleOnCallRespondersIncluded.
func ShiftDataAsScheduleOnCallRespondersIncluded(v *ShiftData) ScheduleOnCallRespondersIncluded {
	return ScheduleOnCallRespondersIncluded{ShiftData: v}
}

// ScheduleDataAsScheduleOnCallRespondersIncluded is a convenience function that returns ScheduleData wrapped in ScheduleOnCallRespondersIncluded.
func ScheduleDataAsScheduleOnCallRespondersIncluded(v *ScheduleData) ScheduleOnCallRespondersIncluded {
	return ScheduleOnCallRespondersIncluded{ScheduleData: v}
}

// UserAsScheduleOnCallRespondersIncluded is a convenience function that returns User wrapped in ScheduleOnCallRespondersIncluded.
func UserAsScheduleOnCallRespondersIncluded(v *User) ScheduleOnCallRespondersIncluded {
	return ScheduleOnCallRespondersIncluded{User: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ScheduleOnCallRespondersIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ScheduleOnCallResponderData
	err = datadog.Unmarshal(data, &obj.ScheduleOnCallResponderData)
	if err == nil {
		if obj.ScheduleOnCallResponderData != nil && obj.ScheduleOnCallResponderData.UnparsedObject == nil {
			jsonScheduleOnCallResponderData, _ := datadog.Marshal(obj.ScheduleOnCallResponderData)
			if string(jsonScheduleOnCallResponderData) == "{}" { // empty struct
				obj.ScheduleOnCallResponderData = nil
			} else {
				match++
			}
		} else {
			obj.ScheduleOnCallResponderData = nil
		}
	} else {
		obj.ScheduleOnCallResponderData = nil
	}

	// try to unmarshal data into ShiftData
	err = datadog.Unmarshal(data, &obj.ShiftData)
	if err == nil {
		if obj.ShiftData != nil && obj.ShiftData.UnparsedObject == nil {
			jsonShiftData, _ := datadog.Marshal(obj.ShiftData)
			if string(jsonShiftData) == "{}" { // empty struct
				obj.ShiftData = nil
			} else {
				match++
			}
		} else {
			obj.ShiftData = nil
		}
	} else {
		obj.ShiftData = nil
	}

	// try to unmarshal data into ScheduleData
	err = datadog.Unmarshal(data, &obj.ScheduleData)
	if err == nil {
		if obj.ScheduleData != nil && obj.ScheduleData.UnparsedObject == nil {
			jsonScheduleData, _ := datadog.Marshal(obj.ScheduleData)
			if string(jsonScheduleData) == "{}" { // empty struct
				obj.ScheduleData = nil
			} else {
				match++
			}
		} else {
			obj.ScheduleData = nil
		}
	} else {
		obj.ScheduleData = nil
	}

	// try to unmarshal data into User
	err = datadog.Unmarshal(data, &obj.User)
	if err == nil {
		if obj.User != nil && obj.User.UnparsedObject == nil {
			jsonUser, _ := datadog.Marshal(obj.User)
			if string(jsonUser) == "{}" && string(data) != "{}" { // empty struct
				obj.User = nil
			} else {
				match++
			}
		} else {
			obj.User = nil
		}
	} else {
		obj.User = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ScheduleOnCallResponderData = nil
		obj.ShiftData = nil
		obj.ScheduleData = nil
		obj.User = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ScheduleOnCallRespondersIncluded) MarshalJSON() ([]byte, error) {
	if obj.ScheduleOnCallResponderData != nil {
		return datadog.Marshal(&obj.ScheduleOnCallResponderData)
	}

	if obj.ShiftData != nil {
		return datadog.Marshal(&obj.ShiftData)
	}

	if obj.ScheduleData != nil {
		return datadog.Marshal(&obj.ScheduleData)
	}

	if obj.User != nil {
		return datadog.Marshal(&obj.User)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ScheduleOnCallRespondersIncluded) GetActualInstance() interface{} {
	if obj.ScheduleOnCallResponderData != nil {
		return obj.ScheduleOnCallResponderData
	}

	if obj.ShiftData != nil {
		return obj.ShiftData
	}

	if obj.ScheduleData != nil {
		return obj.ScheduleData
	}

	if obj.User != nil {
		return obj.User
	}

	// all schemas are nil
	return nil
}

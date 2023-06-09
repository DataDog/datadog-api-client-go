// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
)

// DowntimeAttributeScheduleEditRequest - Schedule for the downtime.
type DowntimeAttributeScheduleEditRequest struct {
	DowntimeAttributeScheduleRecurrencesEditRequest   *DowntimeAttributeScheduleRecurrencesEditRequest
	DowntimeAttributeScheduleOneTimeCreateEditRequest *DowntimeAttributeScheduleOneTimeCreateEditRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DowntimeAttributeScheduleRecurrencesEditRequestAsDowntimeAttributeScheduleEditRequest is a convenience function that returns DowntimeAttributeScheduleRecurrencesEditRequest wrapped in DowntimeAttributeScheduleEditRequest.
func DowntimeAttributeScheduleRecurrencesEditRequestAsDowntimeAttributeScheduleEditRequest(v *DowntimeAttributeScheduleRecurrencesEditRequest) DowntimeAttributeScheduleEditRequest {
	return DowntimeAttributeScheduleEditRequest{DowntimeAttributeScheduleRecurrencesEditRequest: v}
}

// DowntimeAttributeScheduleOneTimeCreateEditRequestAsDowntimeAttributeScheduleEditRequest is a convenience function that returns DowntimeAttributeScheduleOneTimeCreateEditRequest wrapped in DowntimeAttributeScheduleEditRequest.
func DowntimeAttributeScheduleOneTimeCreateEditRequestAsDowntimeAttributeScheduleEditRequest(v *DowntimeAttributeScheduleOneTimeCreateEditRequest) DowntimeAttributeScheduleEditRequest {
	return DowntimeAttributeScheduleEditRequest{DowntimeAttributeScheduleOneTimeCreateEditRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DowntimeAttributeScheduleEditRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DowntimeAttributeScheduleRecurrencesEditRequest
	err = json.Unmarshal(data, &obj.DowntimeAttributeScheduleRecurrencesEditRequest)
	if err == nil {
		if obj.DowntimeAttributeScheduleRecurrencesEditRequest != nil && obj.DowntimeAttributeScheduleRecurrencesEditRequest.UnparsedObject == nil {
			jsonDowntimeAttributeScheduleRecurrencesEditRequest, _ := json.Marshal(obj.DowntimeAttributeScheduleRecurrencesEditRequest)
			if string(jsonDowntimeAttributeScheduleRecurrencesEditRequest) == "{}" { // empty struct
				obj.DowntimeAttributeScheduleRecurrencesEditRequest = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeAttributeScheduleRecurrencesEditRequest = nil
		}
	} else {
		obj.DowntimeAttributeScheduleRecurrencesEditRequest = nil
	}

	// try to unmarshal data into DowntimeAttributeScheduleOneTimeCreateEditRequest
	err = json.Unmarshal(data, &obj.DowntimeAttributeScheduleOneTimeCreateEditRequest)
	if err == nil {
		if obj.DowntimeAttributeScheduleOneTimeCreateEditRequest != nil && obj.DowntimeAttributeScheduleOneTimeCreateEditRequest.UnparsedObject == nil {
			jsonDowntimeAttributeScheduleOneTimeCreateEditRequest, _ := json.Marshal(obj.DowntimeAttributeScheduleOneTimeCreateEditRequest)
			if string(jsonDowntimeAttributeScheduleOneTimeCreateEditRequest) == "{}" { // empty struct
				obj.DowntimeAttributeScheduleOneTimeCreateEditRequest = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeAttributeScheduleOneTimeCreateEditRequest = nil
		}
	} else {
		obj.DowntimeAttributeScheduleOneTimeCreateEditRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DowntimeAttributeScheduleRecurrencesEditRequest = nil
		obj.DowntimeAttributeScheduleOneTimeCreateEditRequest = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DowntimeAttributeScheduleEditRequest) MarshalJSON() ([]byte, error) {
	if obj.DowntimeAttributeScheduleRecurrencesEditRequest != nil {
		return json.Marshal(&obj.DowntimeAttributeScheduleRecurrencesEditRequest)
	}

	if obj.DowntimeAttributeScheduleOneTimeCreateEditRequest != nil {
		return json.Marshal(&obj.DowntimeAttributeScheduleOneTimeCreateEditRequest)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DowntimeAttributeScheduleEditRequest) GetActualInstance() interface{} {
	if obj.DowntimeAttributeScheduleRecurrencesEditRequest != nil {
		return obj.DowntimeAttributeScheduleRecurrencesEditRequest
	}

	if obj.DowntimeAttributeScheduleOneTimeCreateEditRequest != nil {
		return obj.DowntimeAttributeScheduleOneTimeCreateEditRequest
	}

	// all schemas are nil
	return nil
}

// NullableDowntimeAttributeScheduleEditRequest handles when a null is used for DowntimeAttributeScheduleEditRequest.
type NullableDowntimeAttributeScheduleEditRequest struct {
	value *DowntimeAttributeScheduleEditRequest
	isSet bool
}

// Get returns the associated value.
func (v NullableDowntimeAttributeScheduleEditRequest) Get() *DowntimeAttributeScheduleEditRequest {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDowntimeAttributeScheduleEditRequest) Set(val *DowntimeAttributeScheduleEditRequest) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDowntimeAttributeScheduleEditRequest) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableDowntimeAttributeScheduleEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDowntimeAttributeScheduleEditRequest initializes the struct as if Set has been called.
func NewNullableDowntimeAttributeScheduleEditRequest(val *DowntimeAttributeScheduleEditRequest) *NullableDowntimeAttributeScheduleEditRequest {
	return &NullableDowntimeAttributeScheduleEditRequest{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDowntimeAttributeScheduleEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDowntimeAttributeScheduleEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
)

// DowntimeAttributeScheduleCreateRequest - Schedule for the downtime.
type DowntimeAttributeScheduleCreateRequest struct {
	DowntimeAttributeScheduleRecurrencesCreateRequest *DowntimeAttributeScheduleRecurrencesCreateRequest
	DowntimeAttributeScheduleOneTimeCreateEditRequest *DowntimeAttributeScheduleOneTimeCreateEditRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DowntimeAttributeScheduleRecurrencesCreateRequestAsDowntimeAttributeScheduleCreateRequest is a convenience function that returns DowntimeAttributeScheduleRecurrencesCreateRequest wrapped in DowntimeAttributeScheduleCreateRequest.
func DowntimeAttributeScheduleRecurrencesCreateRequestAsDowntimeAttributeScheduleCreateRequest(v *DowntimeAttributeScheduleRecurrencesCreateRequest) DowntimeAttributeScheduleCreateRequest {
	return DowntimeAttributeScheduleCreateRequest{DowntimeAttributeScheduleRecurrencesCreateRequest: v}
}

// DowntimeAttributeScheduleOneTimeCreateEditRequestAsDowntimeAttributeScheduleCreateRequest is a convenience function that returns DowntimeAttributeScheduleOneTimeCreateEditRequest wrapped in DowntimeAttributeScheduleCreateRequest.
func DowntimeAttributeScheduleOneTimeCreateEditRequestAsDowntimeAttributeScheduleCreateRequest(v *DowntimeAttributeScheduleOneTimeCreateEditRequest) DowntimeAttributeScheduleCreateRequest {
	return DowntimeAttributeScheduleCreateRequest{DowntimeAttributeScheduleOneTimeCreateEditRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DowntimeAttributeScheduleCreateRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DowntimeAttributeScheduleRecurrencesCreateRequest
	err = json.Unmarshal(data, &obj.DowntimeAttributeScheduleRecurrencesCreateRequest)
	if err == nil {
		if obj.DowntimeAttributeScheduleRecurrencesCreateRequest != nil && obj.DowntimeAttributeScheduleRecurrencesCreateRequest.UnparsedObject == nil {
			jsonDowntimeAttributeScheduleRecurrencesCreateRequest, _ := json.Marshal(obj.DowntimeAttributeScheduleRecurrencesCreateRequest)
			if string(jsonDowntimeAttributeScheduleRecurrencesCreateRequest) == "{}" { // empty struct
				obj.DowntimeAttributeScheduleRecurrencesCreateRequest = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeAttributeScheduleRecurrencesCreateRequest = nil
		}
	} else {
		obj.DowntimeAttributeScheduleRecurrencesCreateRequest = nil
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
		obj.DowntimeAttributeScheduleRecurrencesCreateRequest = nil
		obj.DowntimeAttributeScheduleOneTimeCreateEditRequest = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DowntimeAttributeScheduleCreateRequest) MarshalJSON() ([]byte, error) {
	if obj.DowntimeAttributeScheduleRecurrencesCreateRequest != nil {
		return json.Marshal(&obj.DowntimeAttributeScheduleRecurrencesCreateRequest)
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
func (obj *DowntimeAttributeScheduleCreateRequest) GetActualInstance() interface{} {
	if obj.DowntimeAttributeScheduleRecurrencesCreateRequest != nil {
		return obj.DowntimeAttributeScheduleRecurrencesCreateRequest
	}

	if obj.DowntimeAttributeScheduleOneTimeCreateEditRequest != nil {
		return obj.DowntimeAttributeScheduleOneTimeCreateEditRequest
	}

	// all schemas are nil
	return nil
}

// NullableDowntimeAttributeScheduleCreateRequest handles when a null is used for DowntimeAttributeScheduleCreateRequest.
type NullableDowntimeAttributeScheduleCreateRequest struct {
	value *DowntimeAttributeScheduleCreateRequest
	isSet bool
}

// Get returns the associated value.
func (v NullableDowntimeAttributeScheduleCreateRequest) Get() *DowntimeAttributeScheduleCreateRequest {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDowntimeAttributeScheduleCreateRequest) Set(val *DowntimeAttributeScheduleCreateRequest) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDowntimeAttributeScheduleCreateRequest) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableDowntimeAttributeScheduleCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDowntimeAttributeScheduleCreateRequest initializes the struct as if Set has been called.
func NewNullableDowntimeAttributeScheduleCreateRequest(val *DowntimeAttributeScheduleCreateRequest) *NullableDowntimeAttributeScheduleCreateRequest {
	return &NullableDowntimeAttributeScheduleCreateRequest{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDowntimeAttributeScheduleCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDowntimeAttributeScheduleCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}

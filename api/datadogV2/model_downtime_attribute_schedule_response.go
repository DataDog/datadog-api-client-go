// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
)

// DowntimeAttributeScheduleResponse - The schedule that defines when the monitor starts, stops, and recurs. There are two types of schedules:
// one-time and recurring. Recurring schedules may have up to five RRULE-based recurrences. If no schedules is
// provided, the downtime will begin immediately and never end.
type DowntimeAttributeScheduleResponse struct {
	DowntimeAttributeScheduleRecurrencesResponse *DowntimeAttributeScheduleRecurrencesResponse
	DowntimeAttributeScheduleOneTimeResponse     *DowntimeAttributeScheduleOneTimeResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DowntimeAttributeScheduleRecurrencesResponseAsDowntimeAttributeScheduleResponse is a convenience function that returns DowntimeAttributeScheduleRecurrencesResponse wrapped in DowntimeAttributeScheduleResponse.
func DowntimeAttributeScheduleRecurrencesResponseAsDowntimeAttributeScheduleResponse(v *DowntimeAttributeScheduleRecurrencesResponse) DowntimeAttributeScheduleResponse {
	return DowntimeAttributeScheduleResponse{DowntimeAttributeScheduleRecurrencesResponse: v}
}

// DowntimeAttributeScheduleOneTimeResponseAsDowntimeAttributeScheduleResponse is a convenience function that returns DowntimeAttributeScheduleOneTimeResponse wrapped in DowntimeAttributeScheduleResponse.
func DowntimeAttributeScheduleOneTimeResponseAsDowntimeAttributeScheduleResponse(v *DowntimeAttributeScheduleOneTimeResponse) DowntimeAttributeScheduleResponse {
	return DowntimeAttributeScheduleResponse{DowntimeAttributeScheduleOneTimeResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DowntimeAttributeScheduleResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DowntimeAttributeScheduleRecurrencesResponse
	err = json.Unmarshal(data, &obj.DowntimeAttributeScheduleRecurrencesResponse)
	if err == nil {
		if obj.DowntimeAttributeScheduleRecurrencesResponse != nil && obj.DowntimeAttributeScheduleRecurrencesResponse.UnparsedObject == nil {
			jsonDowntimeAttributeScheduleRecurrencesResponse, _ := json.Marshal(obj.DowntimeAttributeScheduleRecurrencesResponse)
			if string(jsonDowntimeAttributeScheduleRecurrencesResponse) == "{}" { // empty struct
				obj.DowntimeAttributeScheduleRecurrencesResponse = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeAttributeScheduleRecurrencesResponse = nil
		}
	} else {
		obj.DowntimeAttributeScheduleRecurrencesResponse = nil
	}

	// try to unmarshal data into DowntimeAttributeScheduleOneTimeResponse
	err = json.Unmarshal(data, &obj.DowntimeAttributeScheduleOneTimeResponse)
	if err == nil {
		if obj.DowntimeAttributeScheduleOneTimeResponse != nil && obj.DowntimeAttributeScheduleOneTimeResponse.UnparsedObject == nil {
			jsonDowntimeAttributeScheduleOneTimeResponse, _ := json.Marshal(obj.DowntimeAttributeScheduleOneTimeResponse)
			if string(jsonDowntimeAttributeScheduleOneTimeResponse) == "{}" { // empty struct
				obj.DowntimeAttributeScheduleOneTimeResponse = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeAttributeScheduleOneTimeResponse = nil
		}
	} else {
		obj.DowntimeAttributeScheduleOneTimeResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DowntimeAttributeScheduleRecurrencesResponse = nil
		obj.DowntimeAttributeScheduleOneTimeResponse = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DowntimeAttributeScheduleResponse) MarshalJSON() ([]byte, error) {
	if obj.DowntimeAttributeScheduleRecurrencesResponse != nil {
		return json.Marshal(&obj.DowntimeAttributeScheduleRecurrencesResponse)
	}

	if obj.DowntimeAttributeScheduleOneTimeResponse != nil {
		return json.Marshal(&obj.DowntimeAttributeScheduleOneTimeResponse)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DowntimeAttributeScheduleResponse) GetActualInstance() interface{} {
	if obj.DowntimeAttributeScheduleRecurrencesResponse != nil {
		return obj.DowntimeAttributeScheduleRecurrencesResponse
	}

	if obj.DowntimeAttributeScheduleOneTimeResponse != nil {
		return obj.DowntimeAttributeScheduleOneTimeResponse
	}

	// all schemas are nil
	return nil
}

// NullableDowntimeAttributeScheduleResponse handles when a null is used for DowntimeAttributeScheduleResponse.
type NullableDowntimeAttributeScheduleResponse struct {
	value *DowntimeAttributeScheduleResponse
	isSet bool
}

// Get returns the associated value.
func (v NullableDowntimeAttributeScheduleResponse) Get() *DowntimeAttributeScheduleResponse {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDowntimeAttributeScheduleResponse) Set(val *DowntimeAttributeScheduleResponse) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDowntimeAttributeScheduleResponse) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableDowntimeAttributeScheduleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDowntimeAttributeScheduleResponse initializes the struct as if Set has been called.
func NewNullableDowntimeAttributeScheduleResponse(val *DowntimeAttributeScheduleResponse) *NullableDowntimeAttributeScheduleResponse {
	return &NullableDowntimeAttributeScheduleResponse{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDowntimeAttributeScheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDowntimeAttributeScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}

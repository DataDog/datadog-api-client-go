// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"
)

// DowntimeScheduleResponse - The schedule that defines when the monitor starts, stops, and recurs. There are two types of schedules:
// one-time and recurring. Recurring schedules may have up to five RRULE-based recurrences. If no schedules are
// provided, the downtime will begin immediately and never end.
type DowntimeScheduleResponse struct {
	DowntimeScheduleRecurrencesResponse *DowntimeScheduleRecurrencesResponse
	DowntimeScheduleOneTimeResponse     *DowntimeScheduleOneTimeResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DowntimeScheduleRecurrencesResponseAsDowntimeScheduleResponse is a convenience function that returns DowntimeScheduleRecurrencesResponse wrapped in DowntimeScheduleResponse.
func DowntimeScheduleRecurrencesResponseAsDowntimeScheduleResponse(v *DowntimeScheduleRecurrencesResponse) DowntimeScheduleResponse {
	return DowntimeScheduleResponse{DowntimeScheduleRecurrencesResponse: v}
}

// DowntimeScheduleOneTimeResponseAsDowntimeScheduleResponse is a convenience function that returns DowntimeScheduleOneTimeResponse wrapped in DowntimeScheduleResponse.
func DowntimeScheduleOneTimeResponseAsDowntimeScheduleResponse(v *DowntimeScheduleOneTimeResponse) DowntimeScheduleResponse {
	return DowntimeScheduleResponse{DowntimeScheduleOneTimeResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DowntimeScheduleResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DowntimeScheduleRecurrencesResponse
	err = json.Unmarshal(data, &obj.DowntimeScheduleRecurrencesResponse)
	if err == nil {
		if obj.DowntimeScheduleRecurrencesResponse != nil && obj.DowntimeScheduleRecurrencesResponse.UnparsedObject == nil {
			jsonDowntimeScheduleRecurrencesResponse, _ := json.Marshal(obj.DowntimeScheduleRecurrencesResponse)
			if string(jsonDowntimeScheduleRecurrencesResponse) == "{}" { // empty struct
				obj.DowntimeScheduleRecurrencesResponse = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeScheduleRecurrencesResponse = nil
		}
	} else {
		obj.DowntimeScheduleRecurrencesResponse = nil
	}

	// try to unmarshal data into DowntimeScheduleOneTimeResponse
	err = json.Unmarshal(data, &obj.DowntimeScheduleOneTimeResponse)
	if err == nil {
		if obj.DowntimeScheduleOneTimeResponse != nil && obj.DowntimeScheduleOneTimeResponse.UnparsedObject == nil {
			jsonDowntimeScheduleOneTimeResponse, _ := json.Marshal(obj.DowntimeScheduleOneTimeResponse)
			if string(jsonDowntimeScheduleOneTimeResponse) == "{}" { // empty struct
				obj.DowntimeScheduleOneTimeResponse = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeScheduleOneTimeResponse = nil
		}
	} else {
		obj.DowntimeScheduleOneTimeResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DowntimeScheduleRecurrencesResponse = nil
		obj.DowntimeScheduleOneTimeResponse = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DowntimeScheduleResponse) MarshalJSON() ([]byte, error) {
	if obj.DowntimeScheduleRecurrencesResponse != nil {
		return json.Marshal(&obj.DowntimeScheduleRecurrencesResponse)
	}

	if obj.DowntimeScheduleOneTimeResponse != nil {
		return json.Marshal(&obj.DowntimeScheduleOneTimeResponse)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DowntimeScheduleResponse) GetActualInstance() interface{} {
	if obj.DowntimeScheduleRecurrencesResponse != nil {
		return obj.DowntimeScheduleRecurrencesResponse
	}

	if obj.DowntimeScheduleOneTimeResponse != nil {
		return obj.DowntimeScheduleOneTimeResponse
	}

	// all schemas are nil
	return nil
}

// NullableDowntimeScheduleResponse handles when a null is used for DowntimeScheduleResponse.
type NullableDowntimeScheduleResponse struct {
	value *DowntimeScheduleResponse
	isSet bool
}

// Get returns the associated value.
func (v NullableDowntimeScheduleResponse) Get() *DowntimeScheduleResponse {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDowntimeScheduleResponse) Set(val *DowntimeScheduleResponse) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDowntimeScheduleResponse) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableDowntimeScheduleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDowntimeScheduleResponse initializes the struct as if Set has been called.
func NewNullableDowntimeScheduleResponse(val *DowntimeScheduleResponse) *NullableDowntimeScheduleResponse {
	return &NullableDowntimeScheduleResponse{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDowntimeScheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDowntimeScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}

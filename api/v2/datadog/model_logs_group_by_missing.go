/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// LogsGroupByMissing - The value to use for logs that don't have the facet used to group by
type LogsGroupByMissing struct {
	LogsGroupByMissingString *string
	LogsGroupByMissingNumber *float64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LogsGroupByMissingStringAsLogsGroupByMissing is a convenience function that returns string wrapped in LogsGroupByMissing
func LogsGroupByMissingStringAsLogsGroupByMissing(v *string) LogsGroupByMissing {
	return LogsGroupByMissing{LogsGroupByMissingString: v}
}

// LogsGroupByMissingNumberAsLogsGroupByMissing is a convenience function that returns float64 wrapped in LogsGroupByMissing
func LogsGroupByMissingNumberAsLogsGroupByMissing(v *float64) LogsGroupByMissing {
	return LogsGroupByMissing{LogsGroupByMissingNumber: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LogsGroupByMissing) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LogsGroupByMissingString
	err = json.Unmarshal(data, &dst.LogsGroupByMissingString)
	if err == nil {
		if dst.LogsGroupByMissingString != nil {
			jsonLogsGroupByMissingString, _ := json.Marshal(dst.LogsGroupByMissingString)
			if string(jsonLogsGroupByMissingString) == "{}" { // empty struct
				dst.LogsGroupByMissingString = nil
			} else {
				match++
			}
		} else {
			dst.LogsGroupByMissingString = nil
		}
	} else {
		dst.LogsGroupByMissingString = nil
	}

	// try to unmarshal data into LogsGroupByMissingNumber
	err = json.Unmarshal(data, &dst.LogsGroupByMissingNumber)
	if err == nil {
		if dst.LogsGroupByMissingNumber != nil {
			jsonLogsGroupByMissingNumber, _ := json.Marshal(dst.LogsGroupByMissingNumber)
			if string(jsonLogsGroupByMissingNumber) == "{}" { // empty struct
				dst.LogsGroupByMissingNumber = nil
			} else {
				match++
			}
		} else {
			dst.LogsGroupByMissingNumber = nil
		}
	} else {
		dst.LogsGroupByMissingNumber = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.LogsGroupByMissingString = nil
		dst.LogsGroupByMissingNumber = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LogsGroupByMissing) MarshalJSON() ([]byte, error) {
	if src.LogsGroupByMissingString != nil {
		return json.Marshal(&src.LogsGroupByMissingString)
	}

	if src.LogsGroupByMissingNumber != nil {
		return json.Marshal(&src.LogsGroupByMissingNumber)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LogsGroupByMissing) GetActualInstance() interface{} {
	if obj.LogsGroupByMissingString != nil {
		return obj.LogsGroupByMissingString
	}

	if obj.LogsGroupByMissingNumber != nil {
		return obj.LogsGroupByMissingNumber
	}

	// all schemas are nil
	return nil
}

type NullableLogsGroupByMissing struct {
	value *LogsGroupByMissing
	isSet bool
}

func (v NullableLogsGroupByMissing) Get() *LogsGroupByMissing {
	return v.value
}

func (v *NullableLogsGroupByMissing) Set(val *LogsGroupByMissing) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsGroupByMissing) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsGroupByMissing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsGroupByMissing(val *LogsGroupByMissing) *NullableLogsGroupByMissing {
	return &NullableLogsGroupByMissing{value: val, isSet: true}
}

func (v NullableLogsGroupByMissing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsGroupByMissing) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}

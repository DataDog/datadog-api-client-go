/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */


package datadog

import (
	"encoding/json"
	"fmt"

)


// LogsGroupByTotal - A resulting object to put the given computes in over all the matching records.
type LogsGroupByTotal struct {
	LogsGroupByTotalBoolean *bool
	LogsGroupByTotalString *string
	LogsGroupByTotalNumber *float64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LogsGroupByTotalBooleanAsLogsGroupByTotal is a convenience function that returns bool wrapped in LogsGroupByTotal
func LogsGroupByTotalBooleanAsLogsGroupByTotal(v *bool) LogsGroupByTotal {
	return LogsGroupByTotal{LogsGroupByTotalBoolean: v}
}

// LogsGroupByTotalStringAsLogsGroupByTotal is a convenience function that returns string wrapped in LogsGroupByTotal
func LogsGroupByTotalStringAsLogsGroupByTotal(v *string) LogsGroupByTotal {
	return LogsGroupByTotal{LogsGroupByTotalString: v}
}

// LogsGroupByTotalNumberAsLogsGroupByTotal is a convenience function that returns float64 wrapped in LogsGroupByTotal
func LogsGroupByTotalNumberAsLogsGroupByTotal(v *float64) LogsGroupByTotal {
	return LogsGroupByTotal{LogsGroupByTotalNumber: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LogsGroupByTotal) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LogsGroupByTotalBoolean
	err = json.Unmarshal(data, &dst.LogsGroupByTotalBoolean)
	if err == nil {
		if dst.LogsGroupByTotalBoolean != nil {
			jsonLogsGroupByTotalBoolean, _ := json.Marshal(dst.LogsGroupByTotalBoolean)
			if string(jsonLogsGroupByTotalBoolean) == "{}" { // empty struct
				dst.LogsGroupByTotalBoolean = nil
			} else {
				match++
			}
		} else {
			dst.LogsGroupByTotalBoolean = nil
		}
	} else {
		dst.LogsGroupByTotalBoolean = nil
	}

	// try to unmarshal data into LogsGroupByTotalString
	err = json.Unmarshal(data, &dst.LogsGroupByTotalString)
	if err == nil {
		if dst.LogsGroupByTotalString != nil {
			jsonLogsGroupByTotalString, _ := json.Marshal(dst.LogsGroupByTotalString)
			if string(jsonLogsGroupByTotalString) == "{}" { // empty struct
				dst.LogsGroupByTotalString = nil
			} else {
				match++
			}
		} else {
			dst.LogsGroupByTotalString = nil
		}
	} else {
		dst.LogsGroupByTotalString = nil
	}

	// try to unmarshal data into LogsGroupByTotalNumber
	err = json.Unmarshal(data, &dst.LogsGroupByTotalNumber)
	if err == nil {
		if dst.LogsGroupByTotalNumber != nil {
			jsonLogsGroupByTotalNumber, _ := json.Marshal(dst.LogsGroupByTotalNumber)
			if string(jsonLogsGroupByTotalNumber) == "{}" { // empty struct
				dst.LogsGroupByTotalNumber = nil
			} else {
				match++
			}
		} else {
			dst.LogsGroupByTotalNumber = nil
		}
	} else {
		dst.LogsGroupByTotalNumber = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.LogsGroupByTotalBoolean = nil
		dst.LogsGroupByTotalString = nil
		dst.LogsGroupByTotalNumber = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LogsGroupByTotal) MarshalJSON() ([]byte, error) {
	if src.LogsGroupByTotalBoolean != nil {
		return json.Marshal(&src.LogsGroupByTotalBoolean)
	}


	if src.LogsGroupByTotalString != nil {
		return json.Marshal(&src.LogsGroupByTotalString)
	}


	if src.LogsGroupByTotalNumber != nil {
		return json.Marshal(&src.LogsGroupByTotalNumber)
	}


	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LogsGroupByTotal) GetActualInstance() (interface{}) {
	if obj.LogsGroupByTotalBoolean != nil {
		return obj.LogsGroupByTotalBoolean
	}


	if obj.LogsGroupByTotalString != nil {
		return obj.LogsGroupByTotalString
	}


	if obj.LogsGroupByTotalNumber != nil {
		return obj.LogsGroupByTotalNumber
	}


	// all schemas are nil
	return nil
}

type NullableLogsGroupByTotal struct {
	value *LogsGroupByTotal
	isSet bool
}

func (v NullableLogsGroupByTotal) Get() *LogsGroupByTotal {
	return v.value
}

func (v *NullableLogsGroupByTotal) Set(val *LogsGroupByTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsGroupByTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsGroupByTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsGroupByTotal(val *LogsGroupByTotal) *NullableLogsGroupByTotal {
	return &NullableLogsGroupByTotal{value: val, isSet: true}
}

func (v NullableLogsGroupByTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsGroupByTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}

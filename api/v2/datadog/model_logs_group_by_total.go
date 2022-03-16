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
	Bool *bool
	String *string
	Float64 *float64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// BoolAsLogsGroupByTotal is a convenience function that returns bool wrapped in LogsGroupByTotal
func BoolAsLogsGroupByTotal(v *bool) LogsGroupByTotal {
	return LogsGroupByTotal{Bool: v}
}

// StringAsLogsGroupByTotal is a convenience function that returns string wrapped in LogsGroupByTotal
func StringAsLogsGroupByTotal(v *string) LogsGroupByTotal {
	return LogsGroupByTotal{String: v}
}

// Float64AsLogsGroupByTotal is a convenience function that returns float64 wrapped in LogsGroupByTotal
func Float64AsLogsGroupByTotal(v *float64) LogsGroupByTotal {
	return LogsGroupByTotal{Float64: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LogsGroupByTotal) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Bool
	err = json.Unmarshal(data, &dst.Bool)
	if err == nil {
		if dst.Bool != nil {
			jsonBool, _ := json.Marshal(dst.Bool)
			if string(jsonBool) == "{}" { // empty struct
				dst.Bool = nil
			} else {
				match++
			}
		} else {
			dst.Bool = nil
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		if dst.String != nil {
			jsonString, _ := json.Marshal(dst.String)
			if string(jsonString) == "{}" { // empty struct
				dst.String = nil
			} else {
				match++
			}
		} else {
			dst.String = nil
		}
	} else {
		dst.String = nil
	}

	// try to unmarshal data into Float64
	err = json.Unmarshal(data, &dst.Float64)
	if err == nil {
		if dst.Float64 != nil {
			jsonFloat64, _ := json.Marshal(dst.Float64)
			if string(jsonFloat64) == "{}" { // empty struct
				dst.Float64 = nil
			} else {
				match++
			}
		} else {
			dst.Float64 = nil
		}
	} else {
		dst.Float64 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.Bool = nil
		dst.String = nil
		dst.Float64 = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LogsGroupByTotal) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}


	if src.String != nil {
		return json.Marshal(&src.String)
	}


	if src.Float64 != nil {
		return json.Marshal(&src.Float64)
	}


	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LogsGroupByTotal) GetActualInstance() (interface{}) {
	if obj.Bool != nil {
		return obj.Bool
	}


	if obj.String != nil {
		return obj.String
	}


	if obj.Float64 != nil {
		return obj.Float64
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

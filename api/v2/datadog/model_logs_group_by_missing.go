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


// LogsGroupByMissing - The value to use for logs that don't have the facet used to group by
type LogsGroupByMissing struct {
	String *string
	Float64 *float64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsLogsGroupByMissing is a convenience function that returns string wrapped in LogsGroupByMissing
func StringAsLogsGroupByMissing(v *string) LogsGroupByMissing {
	return LogsGroupByMissing{String: v}
}

// Float64AsLogsGroupByMissing is a convenience function that returns float64 wrapped in LogsGroupByMissing
func Float64AsLogsGroupByMissing(v *float64) LogsGroupByMissing {
	return LogsGroupByMissing{Float64: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LogsGroupByMissing) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
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
		dst.String = nil
		dst.Float64 = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LogsGroupByMissing) MarshalJSON() ([]byte, error) {
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
func (obj *LogsGroupByMissing) GetActualInstance() (interface{}) {
	if obj.String != nil {
		return obj.String
	}


	if obj.Float64 != nil {
		return obj.Float64
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

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


// NotebookGlobalTime - Notebook global timeframe.
type NotebookGlobalTime struct {
	NotebookRelativeTime *NotebookRelativeTime
	NotebookAbsoluteTime *NotebookAbsoluteTime

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NotebookRelativeTimeAsNotebookGlobalTime is a convenience function that returns NotebookRelativeTime wrapped in NotebookGlobalTime
func NotebookRelativeTimeAsNotebookGlobalTime(v *NotebookRelativeTime) NotebookGlobalTime {
	return NotebookGlobalTime{NotebookRelativeTime: v}
}

// NotebookAbsoluteTimeAsNotebookGlobalTime is a convenience function that returns NotebookAbsoluteTime wrapped in NotebookGlobalTime
func NotebookAbsoluteTimeAsNotebookGlobalTime(v *NotebookAbsoluteTime) NotebookGlobalTime {
	return NotebookGlobalTime{NotebookAbsoluteTime: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NotebookGlobalTime) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NotebookRelativeTime
	err = json.Unmarshal(data, &dst.NotebookRelativeTime)
	if err == nil {
		if dst.NotebookRelativeTime != nil && dst.NotebookRelativeTime.UnparsedObject == nil {
			jsonNotebookRelativeTime, _ := json.Marshal(dst.NotebookRelativeTime)
			if string(jsonNotebookRelativeTime) == "{}" { // empty struct
				dst.NotebookRelativeTime = nil
			} else {
				match++
			}
		} else {
			dst.NotebookRelativeTime = nil
		}
	} else {
		dst.NotebookRelativeTime = nil
	}

	// try to unmarshal data into NotebookAbsoluteTime
	err = json.Unmarshal(data, &dst.NotebookAbsoluteTime)
	if err == nil {
		if dst.NotebookAbsoluteTime != nil && dst.NotebookAbsoluteTime.UnparsedObject == nil {
			jsonNotebookAbsoluteTime, _ := json.Marshal(dst.NotebookAbsoluteTime)
			if string(jsonNotebookAbsoluteTime) == "{}" { // empty struct
				dst.NotebookAbsoluteTime = nil
			} else {
				match++
			}
		} else {
			dst.NotebookAbsoluteTime = nil
		}
	} else {
		dst.NotebookAbsoluteTime = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.NotebookRelativeTime = nil
		dst.NotebookAbsoluteTime = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NotebookGlobalTime) MarshalJSON() ([]byte, error) {
	if src.NotebookRelativeTime != nil {
		return json.Marshal(&src.NotebookRelativeTime)
	}


	if src.NotebookAbsoluteTime != nil {
		return json.Marshal(&src.NotebookAbsoluteTime)
	}


	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NotebookGlobalTime) GetActualInstance() (interface{}) {
	if obj.NotebookRelativeTime != nil {
		return obj.NotebookRelativeTime
	}


	if obj.NotebookAbsoluteTime != nil {
		return obj.NotebookAbsoluteTime
	}


	// all schemas are nil
	return nil
}

type NullableNotebookGlobalTime struct {
	value *NotebookGlobalTime
	isSet bool
}

func (v NullableNotebookGlobalTime) Get() *NotebookGlobalTime {
	return v.value
}

func (v *NullableNotebookGlobalTime) Set(val *NotebookGlobalTime) {
	v.value = val
	v.isSet = true
}

func (v NullableNotebookGlobalTime) IsSet() bool {
	return v.isSet
}

func (v *NullableNotebookGlobalTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotebookGlobalTime(val *NotebookGlobalTime) *NullableNotebookGlobalTime {
	return &NullableNotebookGlobalTime{value: val, isSet: true}
}

func (v NullableNotebookGlobalTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotebookGlobalTime) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}

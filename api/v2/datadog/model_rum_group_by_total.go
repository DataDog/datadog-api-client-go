/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// RUMGroupByTotal - A resulting object to put the given computes in over all the matching records.
type RUMGroupByTotal struct {
	RUMGroupByTotalBoolean *bool
	RUMGroupByTotalString  *string
	RUMGroupByTotalNumber  *float64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// RUMGroupByTotalBooleanAsRUMGroupByTotal is a convenience function that returns bool wrapped in RUMGroupByTotal
func RUMGroupByTotalBooleanAsRUMGroupByTotal(v *bool) RUMGroupByTotal {
	return RUMGroupByTotal{RUMGroupByTotalBoolean: v}
}

// RUMGroupByTotalStringAsRUMGroupByTotal is a convenience function that returns string wrapped in RUMGroupByTotal
func RUMGroupByTotalStringAsRUMGroupByTotal(v *string) RUMGroupByTotal {
	return RUMGroupByTotal{RUMGroupByTotalString: v}
}

// RUMGroupByTotalNumberAsRUMGroupByTotal is a convenience function that returns float64 wrapped in RUMGroupByTotal
func RUMGroupByTotalNumberAsRUMGroupByTotal(v *float64) RUMGroupByTotal {
	return RUMGroupByTotal{RUMGroupByTotalNumber: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RUMGroupByTotal) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RUMGroupByTotalBoolean
	err = json.Unmarshal(data, &dst.RUMGroupByTotalBoolean)
	if err == nil {
		if dst.RUMGroupByTotalBoolean != nil {
			jsonRUMGroupByTotalBoolean, _ := json.Marshal(dst.RUMGroupByTotalBoolean)
			if string(jsonRUMGroupByTotalBoolean) == "{}" { // empty struct
				dst.RUMGroupByTotalBoolean = nil
			} else {
				match++
			}
		} else {
			dst.RUMGroupByTotalBoolean = nil
		}
	} else {
		dst.RUMGroupByTotalBoolean = nil
	}

	// try to unmarshal data into RUMGroupByTotalString
	err = json.Unmarshal(data, &dst.RUMGroupByTotalString)
	if err == nil {
		if dst.RUMGroupByTotalString != nil {
			jsonRUMGroupByTotalString, _ := json.Marshal(dst.RUMGroupByTotalString)
			if string(jsonRUMGroupByTotalString) == "{}" { // empty struct
				dst.RUMGroupByTotalString = nil
			} else {
				match++
			}
		} else {
			dst.RUMGroupByTotalString = nil
		}
	} else {
		dst.RUMGroupByTotalString = nil
	}

	// try to unmarshal data into RUMGroupByTotalNumber
	err = json.Unmarshal(data, &dst.RUMGroupByTotalNumber)
	if err == nil {
		if dst.RUMGroupByTotalNumber != nil {
			jsonRUMGroupByTotalNumber, _ := json.Marshal(dst.RUMGroupByTotalNumber)
			if string(jsonRUMGroupByTotalNumber) == "{}" { // empty struct
				dst.RUMGroupByTotalNumber = nil
			} else {
				match++
			}
		} else {
			dst.RUMGroupByTotalNumber = nil
		}
	} else {
		dst.RUMGroupByTotalNumber = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.RUMGroupByTotalBoolean = nil
		dst.RUMGroupByTotalString = nil
		dst.RUMGroupByTotalNumber = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RUMGroupByTotal) MarshalJSON() ([]byte, error) {
	if src.RUMGroupByTotalBoolean != nil {
		return json.Marshal(&src.RUMGroupByTotalBoolean)
	}

	if src.RUMGroupByTotalString != nil {
		return json.Marshal(&src.RUMGroupByTotalString)
	}

	if src.RUMGroupByTotalNumber != nil {
		return json.Marshal(&src.RUMGroupByTotalNumber)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RUMGroupByTotal) GetActualInstance() interface{} {
	if obj.RUMGroupByTotalBoolean != nil {
		return obj.RUMGroupByTotalBoolean
	}

	if obj.RUMGroupByTotalString != nil {
		return obj.RUMGroupByTotalString
	}

	if obj.RUMGroupByTotalNumber != nil {
		return obj.RUMGroupByTotalNumber
	}

	// all schemas are nil
	return nil
}

type NullableRUMGroupByTotal struct {
	value *RUMGroupByTotal
	isSet bool
}

func (v NullableRUMGroupByTotal) Get() *RUMGroupByTotal {
	return v.value
}

func (v *NullableRUMGroupByTotal) Set(val *RUMGroupByTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableRUMGroupByTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableRUMGroupByTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRUMGroupByTotal(val *RUMGroupByTotal) *NullableRUMGroupByTotal {
	return &NullableRUMGroupByTotal{value: val, isSet: true}
}

func (v NullableRUMGroupByTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRUMGroupByTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}

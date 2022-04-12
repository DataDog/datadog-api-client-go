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


// RUMGroupByMissing - The value to use for logs that don't have the facet used to group by.
type RUMGroupByMissing struct {
	RUMGroupByMissingString *string
	RUMGroupByMissingNumber *float64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// RUMGroupByMissingStringAsRUMGroupByMissing is a convenience function that returns string wrapped in RUMGroupByMissing
func RUMGroupByMissingStringAsRUMGroupByMissing(v *string) RUMGroupByMissing {
	return RUMGroupByMissing{RUMGroupByMissingString: v}
}

// RUMGroupByMissingNumberAsRUMGroupByMissing is a convenience function that returns float64 wrapped in RUMGroupByMissing
func RUMGroupByMissingNumberAsRUMGroupByMissing(v *float64) RUMGroupByMissing {
	return RUMGroupByMissing{RUMGroupByMissingNumber: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RUMGroupByMissing) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RUMGroupByMissingString
	err = json.Unmarshal(data, &dst.RUMGroupByMissingString)
	if err == nil {
		if dst.RUMGroupByMissingString != nil {
			jsonRUMGroupByMissingString, _ := json.Marshal(dst.RUMGroupByMissingString)
			if string(jsonRUMGroupByMissingString) == "{}" { // empty struct
				dst.RUMGroupByMissingString = nil
			} else {
				match++
			}
		} else {
			dst.RUMGroupByMissingString = nil
		}
	} else {
		dst.RUMGroupByMissingString = nil
	}

	// try to unmarshal data into RUMGroupByMissingNumber
	err = json.Unmarshal(data, &dst.RUMGroupByMissingNumber)
	if err == nil {
		if dst.RUMGroupByMissingNumber != nil {
			jsonRUMGroupByMissingNumber, _ := json.Marshal(dst.RUMGroupByMissingNumber)
			if string(jsonRUMGroupByMissingNumber) == "{}" { // empty struct
				dst.RUMGroupByMissingNumber = nil
			} else {
				match++
			}
		} else {
			dst.RUMGroupByMissingNumber = nil
		}
	} else {
		dst.RUMGroupByMissingNumber = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.RUMGroupByMissingString = nil
		dst.RUMGroupByMissingNumber = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RUMGroupByMissing) MarshalJSON() ([]byte, error) {
	if src.RUMGroupByMissingString != nil {
		return json.Marshal(&src.RUMGroupByMissingString)
	}


	if src.RUMGroupByMissingNumber != nil {
		return json.Marshal(&src.RUMGroupByMissingNumber)
	}


	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RUMGroupByMissing) GetActualInstance() (interface{}) {
	if obj.RUMGroupByMissingString != nil {
		return obj.RUMGroupByMissingString
	}


	if obj.RUMGroupByMissingNumber != nil {
		return obj.RUMGroupByMissingNumber
	}


	// all schemas are nil
	return nil
}

type NullableRUMGroupByMissing struct {
	value *RUMGroupByMissing
	isSet bool
}

func (v NullableRUMGroupByMissing) Get() *RUMGroupByMissing {
	return v.value
}

func (v *NullableRUMGroupByMissing) Set(val *RUMGroupByMissing) {
	v.value = val
	v.isSet = true
}

func (v NullableRUMGroupByMissing) IsSet() bool {
	return v.isSet
}

func (v *NullableRUMGroupByMissing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRUMGroupByMissing(val *RUMGroupByMissing) *NullableRUMGroupByMissing {
	return &NullableRUMGroupByMissing{value: val, isSet: true}
}

func (v NullableRUMGroupByMissing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRUMGroupByMissing) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}

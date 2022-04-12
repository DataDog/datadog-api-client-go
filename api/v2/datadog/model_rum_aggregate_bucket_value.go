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


// RUMAggregateBucketValue - A bucket value, can be either a timeseries or a single value.
type RUMAggregateBucketValue struct {
	RUMAggregateBucketValueSingleString *string
	RUMAggregateBucketValueSingleNumber *float64
	RUMAggregateBucketValueTimeseries *RUMAggregateBucketValueTimeseries

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// RUMAggregateBucketValueSingleStringAsRUMAggregateBucketValue is a convenience function that returns string wrapped in RUMAggregateBucketValue
func RUMAggregateBucketValueSingleStringAsRUMAggregateBucketValue(v *string) RUMAggregateBucketValue {
	return RUMAggregateBucketValue{RUMAggregateBucketValueSingleString: v}
}

// RUMAggregateBucketValueSingleNumberAsRUMAggregateBucketValue is a convenience function that returns float64 wrapped in RUMAggregateBucketValue
func RUMAggregateBucketValueSingleNumberAsRUMAggregateBucketValue(v *float64) RUMAggregateBucketValue {
	return RUMAggregateBucketValue{RUMAggregateBucketValueSingleNumber: v}
}

// RUMAggregateBucketValueTimeseriesAsRUMAggregateBucketValue is a convenience function that returns RUMAggregateBucketValueTimeseries wrapped in RUMAggregateBucketValue
func RUMAggregateBucketValueTimeseriesAsRUMAggregateBucketValue(v *RUMAggregateBucketValueTimeseries) RUMAggregateBucketValue {
	return RUMAggregateBucketValue{RUMAggregateBucketValueTimeseries: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RUMAggregateBucketValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RUMAggregateBucketValueSingleString
	err = json.Unmarshal(data, &dst.RUMAggregateBucketValueSingleString)
	if err == nil {
		if dst.RUMAggregateBucketValueSingleString != nil {
			jsonRUMAggregateBucketValueSingleString, _ := json.Marshal(dst.RUMAggregateBucketValueSingleString)
			if string(jsonRUMAggregateBucketValueSingleString) == "{}" { // empty struct
				dst.RUMAggregateBucketValueSingleString = nil
			} else {
				match++
			}
		} else {
			dst.RUMAggregateBucketValueSingleString = nil
		}
	} else {
		dst.RUMAggregateBucketValueSingleString = nil
	}

	// try to unmarshal data into RUMAggregateBucketValueSingleNumber
	err = json.Unmarshal(data, &dst.RUMAggregateBucketValueSingleNumber)
	if err == nil {
		if dst.RUMAggregateBucketValueSingleNumber != nil {
			jsonRUMAggregateBucketValueSingleNumber, _ := json.Marshal(dst.RUMAggregateBucketValueSingleNumber)
			if string(jsonRUMAggregateBucketValueSingleNumber) == "{}" { // empty struct
				dst.RUMAggregateBucketValueSingleNumber = nil
			} else {
				match++
			}
		} else {
			dst.RUMAggregateBucketValueSingleNumber = nil
		}
	} else {
		dst.RUMAggregateBucketValueSingleNumber = nil
	}

	// try to unmarshal data into RUMAggregateBucketValueTimeseries
	err = json.Unmarshal(data, &dst.RUMAggregateBucketValueTimeseries)
	if err == nil {
		if dst.RUMAggregateBucketValueTimeseries != nil {
			jsonRUMAggregateBucketValueTimeseries, _ := json.Marshal(dst.RUMAggregateBucketValueTimeseries)
			if string(jsonRUMAggregateBucketValueTimeseries) == "{}" { // empty struct
				dst.RUMAggregateBucketValueTimeseries = nil
			} else {
				match++
			}
		} else {
			dst.RUMAggregateBucketValueTimeseries = nil
		}
	} else {
		dst.RUMAggregateBucketValueTimeseries = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.RUMAggregateBucketValueSingleString = nil
		dst.RUMAggregateBucketValueSingleNumber = nil
		dst.RUMAggregateBucketValueTimeseries = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RUMAggregateBucketValue) MarshalJSON() ([]byte, error) {
	if src.RUMAggregateBucketValueSingleString != nil {
		return json.Marshal(&src.RUMAggregateBucketValueSingleString)
	}


	if src.RUMAggregateBucketValueSingleNumber != nil {
		return json.Marshal(&src.RUMAggregateBucketValueSingleNumber)
	}


	if src.RUMAggregateBucketValueTimeseries != nil {
		return json.Marshal(&src.RUMAggregateBucketValueTimeseries)
	}


	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RUMAggregateBucketValue) GetActualInstance() (interface{}) {
	if obj.RUMAggregateBucketValueSingleString != nil {
		return obj.RUMAggregateBucketValueSingleString
	}


	if obj.RUMAggregateBucketValueSingleNumber != nil {
		return obj.RUMAggregateBucketValueSingleNumber
	}


	if obj.RUMAggregateBucketValueTimeseries != nil {
		return obj.RUMAggregateBucketValueTimeseries
	}


	// all schemas are nil
	return nil
}

type NullableRUMAggregateBucketValue struct {
	value *RUMAggregateBucketValue
	isSet bool
}

func (v NullableRUMAggregateBucketValue) Get() *RUMAggregateBucketValue {
	return v.value
}

func (v *NullableRUMAggregateBucketValue) Set(val *RUMAggregateBucketValue) {
	v.value = val
	v.isSet = true
}

func (v NullableRUMAggregateBucketValue) IsSet() bool {
	return v.isSet
}

func (v *NullableRUMAggregateBucketValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRUMAggregateBucketValue(val *RUMAggregateBucketValue) *NullableRUMAggregateBucketValue {
	return &NullableRUMAggregateBucketValue{value: val, isSet: true}
}

func (v NullableRUMAggregateBucketValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRUMAggregateBucketValue) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}

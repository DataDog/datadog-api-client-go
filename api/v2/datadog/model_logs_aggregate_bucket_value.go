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


// LogsAggregateBucketValue - A bucket value, can be either a timeseries or a single value
type LogsAggregateBucketValue struct {
	LogsAggregateBucketValueSingleString *string
	LogsAggregateBucketValueSingleNumber *float64
	LogsAggregateBucketValueTimeseries *LogsAggregateBucketValueTimeseries

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LogsAggregateBucketValueSingleStringAsLogsAggregateBucketValue is a convenience function that returns string wrapped in LogsAggregateBucketValue
func LogsAggregateBucketValueSingleStringAsLogsAggregateBucketValue(v *string) LogsAggregateBucketValue {
	return LogsAggregateBucketValue{LogsAggregateBucketValueSingleString: v}
}

// LogsAggregateBucketValueSingleNumberAsLogsAggregateBucketValue is a convenience function that returns float64 wrapped in LogsAggregateBucketValue
func LogsAggregateBucketValueSingleNumberAsLogsAggregateBucketValue(v *float64) LogsAggregateBucketValue {
	return LogsAggregateBucketValue{LogsAggregateBucketValueSingleNumber: v}
}

// LogsAggregateBucketValueTimeseriesAsLogsAggregateBucketValue is a convenience function that returns LogsAggregateBucketValueTimeseries wrapped in LogsAggregateBucketValue
func LogsAggregateBucketValueTimeseriesAsLogsAggregateBucketValue(v *LogsAggregateBucketValueTimeseries) LogsAggregateBucketValue {
	return LogsAggregateBucketValue{LogsAggregateBucketValueTimeseries: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LogsAggregateBucketValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LogsAggregateBucketValueSingleString
	err = json.Unmarshal(data, &dst.LogsAggregateBucketValueSingleString)
	if err == nil {
		if dst.LogsAggregateBucketValueSingleString != nil {
			jsonLogsAggregateBucketValueSingleString, _ := json.Marshal(dst.LogsAggregateBucketValueSingleString)
			if string(jsonLogsAggregateBucketValueSingleString) == "{}" { // empty struct
				dst.LogsAggregateBucketValueSingleString = nil
			} else {
				match++
			}
		} else {
			dst.LogsAggregateBucketValueSingleString = nil
		}
	} else {
		dst.LogsAggregateBucketValueSingleString = nil
	}

	// try to unmarshal data into LogsAggregateBucketValueSingleNumber
	err = json.Unmarshal(data, &dst.LogsAggregateBucketValueSingleNumber)
	if err == nil {
		if dst.LogsAggregateBucketValueSingleNumber != nil {
			jsonLogsAggregateBucketValueSingleNumber, _ := json.Marshal(dst.LogsAggregateBucketValueSingleNumber)
			if string(jsonLogsAggregateBucketValueSingleNumber) == "{}" { // empty struct
				dst.LogsAggregateBucketValueSingleNumber = nil
			} else {
				match++
			}
		} else {
			dst.LogsAggregateBucketValueSingleNumber = nil
		}
	} else {
		dst.LogsAggregateBucketValueSingleNumber = nil
	}

	// try to unmarshal data into LogsAggregateBucketValueTimeseries
	err = json.Unmarshal(data, &dst.LogsAggregateBucketValueTimeseries)
	if err == nil {
		if dst.LogsAggregateBucketValueTimeseries != nil {
			jsonLogsAggregateBucketValueTimeseries, _ := json.Marshal(dst.LogsAggregateBucketValueTimeseries)
			if string(jsonLogsAggregateBucketValueTimeseries) == "{}" { // empty struct
				dst.LogsAggregateBucketValueTimeseries = nil
			} else {
				match++
			}
		} else {
			dst.LogsAggregateBucketValueTimeseries = nil
		}
	} else {
		dst.LogsAggregateBucketValueTimeseries = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.LogsAggregateBucketValueSingleString = nil
		dst.LogsAggregateBucketValueSingleNumber = nil
		dst.LogsAggregateBucketValueTimeseries = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LogsAggregateBucketValue) MarshalJSON() ([]byte, error) {
	if src.LogsAggregateBucketValueSingleString != nil {
		return json.Marshal(&src.LogsAggregateBucketValueSingleString)
	}


	if src.LogsAggregateBucketValueSingleNumber != nil {
		return json.Marshal(&src.LogsAggregateBucketValueSingleNumber)
	}


	if src.LogsAggregateBucketValueTimeseries != nil {
		return json.Marshal(&src.LogsAggregateBucketValueTimeseries)
	}


	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LogsAggregateBucketValue) GetActualInstance() (interface{}) {
	if obj.LogsAggregateBucketValueSingleString != nil {
		return obj.LogsAggregateBucketValueSingleString
	}


	if obj.LogsAggregateBucketValueSingleNumber != nil {
		return obj.LogsAggregateBucketValueSingleNumber
	}


	if obj.LogsAggregateBucketValueTimeseries != nil {
		return obj.LogsAggregateBucketValueTimeseries
	}


	// all schemas are nil
	return nil
}

type NullableLogsAggregateBucketValue struct {
	value *LogsAggregateBucketValue
	isSet bool
}

func (v NullableLogsAggregateBucketValue) Get() *LogsAggregateBucketValue {
	return v.value
}

func (v *NullableLogsAggregateBucketValue) Set(val *LogsAggregateBucketValue) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsAggregateBucketValue) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsAggregateBucketValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsAggregateBucketValue(val *LogsAggregateBucketValue) *NullableLogsAggregateBucketValue {
	return &NullableLogsAggregateBucketValue{value: val, isSet: true}
}

func (v NullableLogsAggregateBucketValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsAggregateBucketValue) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}

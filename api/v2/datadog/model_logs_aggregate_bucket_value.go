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
	String *string
	Float64 *float64
	LogsAggregateBucketValueTimeseries *LogsAggregateBucketValueTimeseries

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsLogsAggregateBucketValue is a convenience function that returns string wrapped in LogsAggregateBucketValue
func StringAsLogsAggregateBucketValue(v *string) LogsAggregateBucketValue {
	return LogsAggregateBucketValue{String: v}
}

// Float64AsLogsAggregateBucketValue is a convenience function that returns float64 wrapped in LogsAggregateBucketValue
func Float64AsLogsAggregateBucketValue(v *float64) LogsAggregateBucketValue {
	return LogsAggregateBucketValue{Float64: v}
}

// LogsAggregateBucketValueTimeseriesAsLogsAggregateBucketValue is a convenience function that returns LogsAggregateBucketValueTimeseries wrapped in LogsAggregateBucketValue
func LogsAggregateBucketValueTimeseriesAsLogsAggregateBucketValue(v *LogsAggregateBucketValueTimeseries) LogsAggregateBucketValue {
	return LogsAggregateBucketValue{LogsAggregateBucketValueTimeseries: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LogsAggregateBucketValue) UnmarshalJSON(data []byte) error {
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
		dst.String = nil
		dst.Float64 = nil
		dst.LogsAggregateBucketValueTimeseries = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LogsAggregateBucketValue) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}


	if src.Float64 != nil {
		return json.Marshal(&src.Float64)
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
	if obj.String != nil {
		return obj.String
	}


	if obj.Float64 != nil {
		return obj.Float64
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

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RunChartProjectionDimension - Test-only generated schema for the dashboard versioning proof of concept.
type RunChartProjectionDimension struct {
	RunChartXDimension_20270101     *RunChartXDimension
	RunChartValueDimension_20270101 *RunChartValueDimension
	RunChartGroupDimension_20270101 *RunChartGroupDimension

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// RunChartXDimension_20270101AsRunChartProjectionDimension is a convenience function that returns RunChartXDimension wrapped in RunChartProjectionDimension.
func RunChartXDimension_20270101AsRunChartProjectionDimension(v *RunChartXDimension) RunChartProjectionDimension {
	return RunChartProjectionDimension{RunChartXDimension_20270101: v}
}

// RunChartValueDimension_20270101AsRunChartProjectionDimension is a convenience function that returns RunChartValueDimension wrapped in RunChartProjectionDimension.
func RunChartValueDimension_20270101AsRunChartProjectionDimension(v *RunChartValueDimension) RunChartProjectionDimension {
	return RunChartProjectionDimension{RunChartValueDimension_20270101: v}
}

// RunChartGroupDimension_20270101AsRunChartProjectionDimension is a convenience function that returns RunChartGroupDimension wrapped in RunChartProjectionDimension.
func RunChartGroupDimension_20270101AsRunChartProjectionDimension(v *RunChartGroupDimension) RunChartProjectionDimension {
	return RunChartProjectionDimension{RunChartGroupDimension_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *RunChartProjectionDimension) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RunChartXDimension_20270101
	err = datadog.Unmarshal(data, &obj.RunChartXDimension_20270101)
	if err == nil {
		if obj.RunChartXDimension_20270101 != nil && obj.RunChartXDimension_20270101.UnparsedObject == nil {
			jsonRunChartXDimension_20270101, _ := datadog.Marshal(obj.RunChartXDimension_20270101)
			if string(jsonRunChartXDimension_20270101) == "{}" { // empty struct
				obj.RunChartXDimension_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.RunChartXDimension_20270101 = nil
		}
	} else {
		obj.RunChartXDimension_20270101 = nil
	}

	// try to unmarshal data into RunChartValueDimension_20270101
	err = datadog.Unmarshal(data, &obj.RunChartValueDimension_20270101)
	if err == nil {
		if obj.RunChartValueDimension_20270101 != nil && obj.RunChartValueDimension_20270101.UnparsedObject == nil {
			jsonRunChartValueDimension_20270101, _ := datadog.Marshal(obj.RunChartValueDimension_20270101)
			if string(jsonRunChartValueDimension_20270101) == "{}" { // empty struct
				obj.RunChartValueDimension_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.RunChartValueDimension_20270101 = nil
		}
	} else {
		obj.RunChartValueDimension_20270101 = nil
	}

	// try to unmarshal data into RunChartGroupDimension_20270101
	err = datadog.Unmarshal(data, &obj.RunChartGroupDimension_20270101)
	if err == nil {
		if obj.RunChartGroupDimension_20270101 != nil && obj.RunChartGroupDimension_20270101.UnparsedObject == nil {
			jsonRunChartGroupDimension_20270101, _ := datadog.Marshal(obj.RunChartGroupDimension_20270101)
			if string(jsonRunChartGroupDimension_20270101) == "{}" { // empty struct
				obj.RunChartGroupDimension_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.RunChartGroupDimension_20270101 = nil
		}
	} else {
		obj.RunChartGroupDimension_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.RunChartXDimension_20270101 = nil
		obj.RunChartValueDimension_20270101 = nil
		obj.RunChartGroupDimension_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj RunChartProjectionDimension) MarshalJSON() ([]byte, error) {
	if obj.RunChartXDimension_20270101 != nil {
		return datadog.Marshal(&obj.RunChartXDimension_20270101)
	}

	if obj.RunChartValueDimension_20270101 != nil {
		return datadog.Marshal(&obj.RunChartValueDimension_20270101)
	}

	if obj.RunChartGroupDimension_20270101 != nil {
		return datadog.Marshal(&obj.RunChartGroupDimension_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *RunChartProjectionDimension) GetActualInstance() interface{} {
	if obj.RunChartXDimension_20270101 != nil {
		return obj.RunChartXDimension_20270101
	}

	if obj.RunChartValueDimension_20270101 != nil {
		return obj.RunChartValueDimension_20270101
	}

	if obj.RunChartGroupDimension_20270101 != nil {
		return obj.RunChartGroupDimension_20270101
	}

	// all schemas are nil
	return nil
}

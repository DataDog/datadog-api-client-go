// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionEventQueryGroupByConfig - Test-only generated schema for the dashboard versioning proof of concept.
type FormulaAndFunctionEventQueryGroupByConfig struct {
	FormulaAndFunctionEventQueryGroupByList_20270101   *[]FormulaAndFunctionEventQueryGroupBy
	FormulaAndFunctionEventQueryGroupByFields_20270101 *FormulaAndFunctionEventQueryGroupByFields

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FormulaAndFunctionEventQueryGroupByList_20270101AsFormulaAndFunctionEventQueryGroupByConfig is a convenience function that returns []FormulaAndFunctionEventQueryGroupBy wrapped in FormulaAndFunctionEventQueryGroupByConfig.
func FormulaAndFunctionEventQueryGroupByList_20270101AsFormulaAndFunctionEventQueryGroupByConfig(v *[]FormulaAndFunctionEventQueryGroupBy) FormulaAndFunctionEventQueryGroupByConfig {
	return FormulaAndFunctionEventQueryGroupByConfig{FormulaAndFunctionEventQueryGroupByList_20270101: v}
}

// FormulaAndFunctionEventQueryGroupByFields_20270101AsFormulaAndFunctionEventQueryGroupByConfig is a convenience function that returns FormulaAndFunctionEventQueryGroupByFields wrapped in FormulaAndFunctionEventQueryGroupByConfig.
func FormulaAndFunctionEventQueryGroupByFields_20270101AsFormulaAndFunctionEventQueryGroupByConfig(v *FormulaAndFunctionEventQueryGroupByFields) FormulaAndFunctionEventQueryGroupByConfig {
	return FormulaAndFunctionEventQueryGroupByConfig{FormulaAndFunctionEventQueryGroupByFields_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *FormulaAndFunctionEventQueryGroupByConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FormulaAndFunctionEventQueryGroupByList_20270101
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionEventQueryGroupByList_20270101)
	if err == nil {
		if obj.FormulaAndFunctionEventQueryGroupByList_20270101 != nil {
			jsonFormulaAndFunctionEventQueryGroupByList_20270101, _ := datadog.Marshal(obj.FormulaAndFunctionEventQueryGroupByList_20270101)
			if string(jsonFormulaAndFunctionEventQueryGroupByList_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.FormulaAndFunctionEventQueryGroupByList_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionEventQueryGroupByList_20270101 = nil
		}
	} else {
		obj.FormulaAndFunctionEventQueryGroupByList_20270101 = nil
	}

	// try to unmarshal data into FormulaAndFunctionEventQueryGroupByFields_20270101
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionEventQueryGroupByFields_20270101)
	if err == nil {
		if obj.FormulaAndFunctionEventQueryGroupByFields_20270101 != nil && obj.FormulaAndFunctionEventQueryGroupByFields_20270101.UnparsedObject == nil {
			jsonFormulaAndFunctionEventQueryGroupByFields_20270101, _ := datadog.Marshal(obj.FormulaAndFunctionEventQueryGroupByFields_20270101)
			if string(jsonFormulaAndFunctionEventQueryGroupByFields_20270101) == "{}" { // empty struct
				obj.FormulaAndFunctionEventQueryGroupByFields_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionEventQueryGroupByFields_20270101 = nil
		}
	} else {
		obj.FormulaAndFunctionEventQueryGroupByFields_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FormulaAndFunctionEventQueryGroupByList_20270101 = nil
		obj.FormulaAndFunctionEventQueryGroupByFields_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj FormulaAndFunctionEventQueryGroupByConfig) MarshalJSON() ([]byte, error) {
	if obj.FormulaAndFunctionEventQueryGroupByList_20270101 != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionEventQueryGroupByList_20270101)
	}

	if obj.FormulaAndFunctionEventQueryGroupByFields_20270101 != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionEventQueryGroupByFields_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *FormulaAndFunctionEventQueryGroupByConfig) GetActualInstance() interface{} {
	if obj.FormulaAndFunctionEventQueryGroupByList_20270101 != nil {
		return obj.FormulaAndFunctionEventQueryGroupByList_20270101
	}

	if obj.FormulaAndFunctionEventQueryGroupByFields_20270101 != nil {
		return obj.FormulaAndFunctionEventQueryGroupByFields_20270101
	}

	// all schemas are nil
	return nil
}

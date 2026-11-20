// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScalarFormulaOrderBy - Test-only generated schema for the dashboard versioning proof of concept.
type ScalarFormulaOrderBy struct {
	FormulaOrderByGroup_20270101        *FormulaOrderByGroup
	FormulaOrderByFormulaIndex_20270101 *FormulaOrderByFormulaIndex

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FormulaOrderByGroup_20270101AsScalarFormulaOrderBy is a convenience function that returns FormulaOrderByGroup wrapped in ScalarFormulaOrderBy.
func FormulaOrderByGroup_20270101AsScalarFormulaOrderBy(v *FormulaOrderByGroup) ScalarFormulaOrderBy {
	return ScalarFormulaOrderBy{FormulaOrderByGroup_20270101: v}
}

// FormulaOrderByFormulaIndex_20270101AsScalarFormulaOrderBy is a convenience function that returns FormulaOrderByFormulaIndex wrapped in ScalarFormulaOrderBy.
func FormulaOrderByFormulaIndex_20270101AsScalarFormulaOrderBy(v *FormulaOrderByFormulaIndex) ScalarFormulaOrderBy {
	return ScalarFormulaOrderBy{FormulaOrderByFormulaIndex_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ScalarFormulaOrderBy) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FormulaOrderByGroup_20270101
	err = datadog.Unmarshal(data, &obj.FormulaOrderByGroup_20270101)
	if err == nil {
		if obj.FormulaOrderByGroup_20270101 != nil && obj.FormulaOrderByGroup_20270101.UnparsedObject == nil {
			jsonFormulaOrderByGroup_20270101, _ := datadog.Marshal(obj.FormulaOrderByGroup_20270101)
			if string(jsonFormulaOrderByGroup_20270101) == "{}" { // empty struct
				obj.FormulaOrderByGroup_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaOrderByGroup_20270101 = nil
		}
	} else {
		obj.FormulaOrderByGroup_20270101 = nil
	}

	// try to unmarshal data into FormulaOrderByFormulaIndex_20270101
	err = datadog.Unmarshal(data, &obj.FormulaOrderByFormulaIndex_20270101)
	if err == nil {
		if obj.FormulaOrderByFormulaIndex_20270101 != nil && obj.FormulaOrderByFormulaIndex_20270101.UnparsedObject == nil {
			jsonFormulaOrderByFormulaIndex_20270101, _ := datadog.Marshal(obj.FormulaOrderByFormulaIndex_20270101)
			if string(jsonFormulaOrderByFormulaIndex_20270101) == "{}" { // empty struct
				obj.FormulaOrderByFormulaIndex_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaOrderByFormulaIndex_20270101 = nil
		}
	} else {
		obj.FormulaOrderByFormulaIndex_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FormulaOrderByGroup_20270101 = nil
		obj.FormulaOrderByFormulaIndex_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ScalarFormulaOrderBy) MarshalJSON() ([]byte, error) {
	if obj.FormulaOrderByGroup_20270101 != nil {
		return datadog.Marshal(&obj.FormulaOrderByGroup_20270101)
	}

	if obj.FormulaOrderByFormulaIndex_20270101 != nil {
		return datadog.Marshal(&obj.FormulaOrderByFormulaIndex_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ScalarFormulaOrderBy) GetActualInstance() interface{} {
	if obj.FormulaOrderByGroup_20270101 != nil {
		return obj.FormulaOrderByGroup_20270101
	}

	if obj.FormulaOrderByFormulaIndex_20270101 != nil {
		return obj.FormulaOrderByFormulaIndex_20270101
	}

	// all schemas are nil
	return nil
}

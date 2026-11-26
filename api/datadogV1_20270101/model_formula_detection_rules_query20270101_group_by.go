// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV1_20270101

import (
	_io "io"
	"time"
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

)


// FormulaDetectionRulesQuery20270101GroupBy - Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDetectionRulesQuery20270101GroupBy struct {
	[]FormulaAndFunctionEventQueryGroupBy *[]FormulaAndFunctionEventQueryGroupBy
	FormulaAndFunctionEventQueryGroupByFields_20270101 *FormulaAndFunctionEventQueryGroupByFields

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// []FormulaAndFunctionEventQueryGroupByAsFormulaDetectionRulesQuery20270101GroupBy is a convenience function that returns []FormulaAndFunctionEventQueryGroupBy wrapped in FormulaDetectionRulesQuery20270101GroupBy.
func []FormulaAndFunctionEventQueryGroupByAsFormulaDetectionRulesQuery20270101GroupBy(v *[]FormulaAndFunctionEventQueryGroupBy) FormulaDetectionRulesQuery20270101GroupBy {
	return FormulaDetectionRulesQuery20270101GroupBy{[]FormulaAndFunctionEventQueryGroupBy: v}
}

// FormulaAndFunctionEventQueryGroupByFields_20270101AsFormulaDetectionRulesQuery20270101GroupBy is a convenience function that returns FormulaAndFunctionEventQueryGroupByFields wrapped in FormulaDetectionRulesQuery20270101GroupBy.
func FormulaAndFunctionEventQueryGroupByFields_20270101AsFormulaDetectionRulesQuery20270101GroupBy(v *FormulaAndFunctionEventQueryGroupByFields) FormulaDetectionRulesQuery20270101GroupBy {
	return FormulaDetectionRulesQuery20270101GroupBy{FormulaAndFunctionEventQueryGroupByFields_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *FormulaDetectionRulesQuery20270101GroupBy) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into []FormulaAndFunctionEventQueryGroupBy
	err = datadog.Unmarshal(data, &obj.[]FormulaAndFunctionEventQueryGroupBy)
	if err == nil {
		if obj.[]FormulaAndFunctionEventQueryGroupBy != nil {
			json[]FormulaAndFunctionEventQueryGroupBy, _ := datadog.Marshal(obj.[]FormulaAndFunctionEventQueryGroupBy)
			if string(json[]FormulaAndFunctionEventQueryGroupBy) == "{}" && string(data) != "{}"  { // empty struct
				obj.[]FormulaAndFunctionEventQueryGroupBy = nil
			} else {
				match++
			}
		} else {
			obj.[]FormulaAndFunctionEventQueryGroupBy = nil
		}
	} else {
		obj.[]FormulaAndFunctionEventQueryGroupBy = nil
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
		obj.[]FormulaAndFunctionEventQueryGroupBy = nil
		obj.FormulaAndFunctionEventQueryGroupByFields_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj FormulaDetectionRulesQuery20270101GroupBy) MarshalJSON() ([]byte, error) {
	if obj.[]FormulaAndFunctionEventQueryGroupBy != nil {
		return datadog.Marshal(&obj.[]FormulaAndFunctionEventQueryGroupBy)
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
func (obj *FormulaDetectionRulesQuery20270101GroupBy) GetActualInstance() (interface{}) {
	if obj.[]FormulaAndFunctionEventQueryGroupBy != nil {
		return obj.[]FormulaAndFunctionEventQueryGroupBy
	}


	if obj.FormulaAndFunctionEventQueryGroupByFields_20270101 != nil {
		return obj.FormulaAndFunctionEventQueryGroupByFields_20270101
	}


	// all schemas are nil
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaApmStatsQuery - Test-only generated schema for the dashboard versioning proof of concept.
type FormulaApmStatsQuery struct {
	FormulaApmDependencyStatsQuery_20270101 *FormulaApmDependencyStatsQuery
	FormulaApmResourceStatsQuery_20270101   *FormulaApmResourceStatsQuery

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FormulaApmDependencyStatsQuery_20270101AsFormulaApmStatsQuery is a convenience function that returns FormulaApmDependencyStatsQuery wrapped in FormulaApmStatsQuery.
func FormulaApmDependencyStatsQuery_20270101AsFormulaApmStatsQuery(v *FormulaApmDependencyStatsQuery) FormulaApmStatsQuery {
	return FormulaApmStatsQuery{FormulaApmDependencyStatsQuery_20270101: v}
}

// FormulaApmResourceStatsQuery_20270101AsFormulaApmStatsQuery is a convenience function that returns FormulaApmResourceStatsQuery wrapped in FormulaApmStatsQuery.
func FormulaApmResourceStatsQuery_20270101AsFormulaApmStatsQuery(v *FormulaApmResourceStatsQuery) FormulaApmStatsQuery {
	return FormulaApmStatsQuery{FormulaApmResourceStatsQuery_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *FormulaApmStatsQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FormulaApmDependencyStatsQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaApmDependencyStatsQuery_20270101)
	if err == nil {
		if obj.FormulaApmDependencyStatsQuery_20270101 != nil && obj.FormulaApmDependencyStatsQuery_20270101.UnparsedObject == nil {
			jsonFormulaApmDependencyStatsQuery_20270101, _ := datadog.Marshal(obj.FormulaApmDependencyStatsQuery_20270101)
			if string(jsonFormulaApmDependencyStatsQuery_20270101) == "{}" { // empty struct
				obj.FormulaApmDependencyStatsQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaApmDependencyStatsQuery_20270101 = nil
		}
	} else {
		obj.FormulaApmDependencyStatsQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaApmResourceStatsQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaApmResourceStatsQuery_20270101)
	if err == nil {
		if obj.FormulaApmResourceStatsQuery_20270101 != nil && obj.FormulaApmResourceStatsQuery_20270101.UnparsedObject == nil {
			jsonFormulaApmResourceStatsQuery_20270101, _ := datadog.Marshal(obj.FormulaApmResourceStatsQuery_20270101)
			if string(jsonFormulaApmResourceStatsQuery_20270101) == "{}" { // empty struct
				obj.FormulaApmResourceStatsQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaApmResourceStatsQuery_20270101 = nil
		}
	} else {
		obj.FormulaApmResourceStatsQuery_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FormulaApmDependencyStatsQuery_20270101 = nil
		obj.FormulaApmResourceStatsQuery_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj FormulaApmStatsQuery) MarshalJSON() ([]byte, error) {
	if obj.FormulaApmDependencyStatsQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaApmDependencyStatsQuery_20270101)
	}

	if obj.FormulaApmResourceStatsQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaApmResourceStatsQuery_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *FormulaApmStatsQuery) GetActualInstance() interface{} {
	if obj.FormulaApmDependencyStatsQuery_20270101 != nil {
		return obj.FormulaApmDependencyStatsQuery_20270101
	}

	if obj.FormulaApmResourceStatsQuery_20270101 != nil {
		return obj.FormulaApmResourceStatsQuery_20270101
	}

	// all schemas are nil
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ComputationScope - Test-only generated schema for the dashboard versioning proof of concept.
type ComputationScope struct {
	CohortScope_20270101       *CohortScope
	ReturnPeriodScope_20270101 *ReturnPeriodScope
	CellScope_20270101         *CellScope

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CohortScope_20270101AsComputationScope is a convenience function that returns CohortScope wrapped in ComputationScope.
func CohortScope_20270101AsComputationScope(v *CohortScope) ComputationScope {
	return ComputationScope{CohortScope_20270101: v}
}

// ReturnPeriodScope_20270101AsComputationScope is a convenience function that returns ReturnPeriodScope wrapped in ComputationScope.
func ReturnPeriodScope_20270101AsComputationScope(v *ReturnPeriodScope) ComputationScope {
	return ComputationScope{ReturnPeriodScope_20270101: v}
}

// CellScope_20270101AsComputationScope is a convenience function that returns CellScope wrapped in ComputationScope.
func CellScope_20270101AsComputationScope(v *CellScope) ComputationScope {
	return ComputationScope{CellScope_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ComputationScope) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CohortScope_20270101
	err = datadog.Unmarshal(data, &obj.CohortScope_20270101)
	if err == nil {
		if obj.CohortScope_20270101 != nil && obj.CohortScope_20270101.UnparsedObject == nil {
			jsonCohortScope_20270101, _ := datadog.Marshal(obj.CohortScope_20270101)
			if string(jsonCohortScope_20270101) == "{}" { // empty struct
				obj.CohortScope_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CohortScope_20270101 = nil
		}
	} else {
		obj.CohortScope_20270101 = nil
	}

	// try to unmarshal data into ReturnPeriodScope_20270101
	err = datadog.Unmarshal(data, &obj.ReturnPeriodScope_20270101)
	if err == nil {
		if obj.ReturnPeriodScope_20270101 != nil && obj.ReturnPeriodScope_20270101.UnparsedObject == nil {
			jsonReturnPeriodScope_20270101, _ := datadog.Marshal(obj.ReturnPeriodScope_20270101)
			if string(jsonReturnPeriodScope_20270101) == "{}" { // empty struct
				obj.ReturnPeriodScope_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ReturnPeriodScope_20270101 = nil
		}
	} else {
		obj.ReturnPeriodScope_20270101 = nil
	}

	// try to unmarshal data into CellScope_20270101
	err = datadog.Unmarshal(data, &obj.CellScope_20270101)
	if err == nil {
		if obj.CellScope_20270101 != nil && obj.CellScope_20270101.UnparsedObject == nil {
			jsonCellScope_20270101, _ := datadog.Marshal(obj.CellScope_20270101)
			if string(jsonCellScope_20270101) == "{}" { // empty struct
				obj.CellScope_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CellScope_20270101 = nil
		}
	} else {
		obj.CellScope_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CohortScope_20270101 = nil
		obj.ReturnPeriodScope_20270101 = nil
		obj.CellScope_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ComputationScope) MarshalJSON() ([]byte, error) {
	if obj.CohortScope_20270101 != nil {
		return datadog.Marshal(&obj.CohortScope_20270101)
	}

	if obj.ReturnPeriodScope_20270101 != nil {
		return datadog.Marshal(&obj.ReturnPeriodScope_20270101)
	}

	if obj.CellScope_20270101 != nil {
		return datadog.Marshal(&obj.CellScope_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ComputationScope) GetActualInstance() interface{} {
	if obj.CohortScope_20270101 != nil {
		return obj.CohortScope_20270101
	}

	if obj.ReturnPeriodScope_20270101 != nil {
		return obj.ReturnPeriodScope_20270101
	}

	if obj.CellScope_20270101 != nil {
		return obj.CellScope_20270101
	}

	// all schemas are nil
	return nil
}

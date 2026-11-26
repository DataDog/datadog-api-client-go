// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CellScope20270101CohortTarget - Test-only generated schema for the dashboard versioning proof of concept.
type CellScope20270101CohortTarget struct {
	CellScope20270101CohortTargetOneOf           *CellScope20270101CohortTargetOneOf
	CellScope20270101CohortTargetOneOf4454617856 *CellScope20270101CohortTargetOneOf4454617856

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CellScope20270101CohortTargetOneOfAsCellScope20270101CohortTarget is a convenience function that returns CellScope20270101CohortTargetOneOf wrapped in CellScope20270101CohortTarget.
func CellScope20270101CohortTargetOneOfAsCellScope20270101CohortTarget(v *CellScope20270101CohortTargetOneOf) CellScope20270101CohortTarget {
	return CellScope20270101CohortTarget{CellScope20270101CohortTargetOneOf: v}
}

// CellScope20270101CohortTargetOneOf4454617856AsCellScope20270101CohortTarget is a convenience function that returns CellScope20270101CohortTargetOneOf4454617856 wrapped in CellScope20270101CohortTarget.
func CellScope20270101CohortTargetOneOf4454617856AsCellScope20270101CohortTarget(v *CellScope20270101CohortTargetOneOf4454617856) CellScope20270101CohortTarget {
	return CellScope20270101CohortTarget{CellScope20270101CohortTargetOneOf4454617856: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CellScope20270101CohortTarget) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CellScope20270101CohortTargetOneOf
	err = datadog.Unmarshal(data, &obj.CellScope20270101CohortTargetOneOf)
	if err == nil {
		if obj.CellScope20270101CohortTargetOneOf != nil && obj.CellScope20270101CohortTargetOneOf.UnparsedObject == nil {
			jsonCellScope20270101CohortTargetOneOf, _ := datadog.Marshal(obj.CellScope20270101CohortTargetOneOf)
			if string(jsonCellScope20270101CohortTargetOneOf) == "{}" { // empty struct
				obj.CellScope20270101CohortTargetOneOf = nil
			} else {
				match++
			}
		} else {
			obj.CellScope20270101CohortTargetOneOf = nil
		}
	} else {
		obj.CellScope20270101CohortTargetOneOf = nil
	}

	// try to unmarshal data into CellScope20270101CohortTargetOneOf4454617856
	err = datadog.Unmarshal(data, &obj.CellScope20270101CohortTargetOneOf4454617856)
	if err == nil {
		if obj.CellScope20270101CohortTargetOneOf4454617856 != nil && obj.CellScope20270101CohortTargetOneOf4454617856.UnparsedObject == nil {
			jsonCellScope20270101CohortTargetOneOf4454617856, _ := datadog.Marshal(obj.CellScope20270101CohortTargetOneOf4454617856)
			if string(jsonCellScope20270101CohortTargetOneOf4454617856) == "{}" { // empty struct
				obj.CellScope20270101CohortTargetOneOf4454617856 = nil
			} else {
				match++
			}
		} else {
			obj.CellScope20270101CohortTargetOneOf4454617856 = nil
		}
	} else {
		obj.CellScope20270101CohortTargetOneOf4454617856 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CellScope20270101CohortTargetOneOf = nil
		obj.CellScope20270101CohortTargetOneOf4454617856 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CellScope20270101CohortTarget) MarshalJSON() ([]byte, error) {
	if obj.CellScope20270101CohortTargetOneOf != nil {
		return datadog.Marshal(&obj.CellScope20270101CohortTargetOneOf)
	}

	if obj.CellScope20270101CohortTargetOneOf4454617856 != nil {
		return datadog.Marshal(&obj.CellScope20270101CohortTargetOneOf4454617856)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CellScope20270101CohortTarget) GetActualInstance() interface{} {
	if obj.CellScope20270101CohortTargetOneOf != nil {
		return obj.CellScope20270101CohortTargetOneOf
	}

	if obj.CellScope20270101CohortTargetOneOf4454617856 != nil {
		return obj.CellScope20270101CohortTargetOneOf4454617856
	}

	// all schemas are nil
	return nil
}

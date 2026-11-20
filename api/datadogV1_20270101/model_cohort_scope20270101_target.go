// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CohortScope20270101Target - Test-only generated schema for the dashboard versioning proof of concept.
type CohortScope20270101Target struct {
	CohortScope20270101TargetOneOf           *CohortScope20270101TargetOneOf
	CohortScope20270101TargetOneOf4454865216 *CohortScope20270101TargetOneOf4454865216

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CohortScope20270101TargetOneOfAsCohortScope20270101Target is a convenience function that returns CohortScope20270101TargetOneOf wrapped in CohortScope20270101Target.
func CohortScope20270101TargetOneOfAsCohortScope20270101Target(v *CohortScope20270101TargetOneOf) CohortScope20270101Target {
	return CohortScope20270101Target{CohortScope20270101TargetOneOf: v}
}

// CohortScope20270101TargetOneOf4454865216AsCohortScope20270101Target is a convenience function that returns CohortScope20270101TargetOneOf4454865216 wrapped in CohortScope20270101Target.
func CohortScope20270101TargetOneOf4454865216AsCohortScope20270101Target(v *CohortScope20270101TargetOneOf4454865216) CohortScope20270101Target {
	return CohortScope20270101Target{CohortScope20270101TargetOneOf4454865216: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CohortScope20270101Target) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CohortScope20270101TargetOneOf
	err = datadog.Unmarshal(data, &obj.CohortScope20270101TargetOneOf)
	if err == nil {
		if obj.CohortScope20270101TargetOneOf != nil && obj.CohortScope20270101TargetOneOf.UnparsedObject == nil {
			jsonCohortScope20270101TargetOneOf, _ := datadog.Marshal(obj.CohortScope20270101TargetOneOf)
			if string(jsonCohortScope20270101TargetOneOf) == "{}" { // empty struct
				obj.CohortScope20270101TargetOneOf = nil
			} else {
				match++
			}
		} else {
			obj.CohortScope20270101TargetOneOf = nil
		}
	} else {
		obj.CohortScope20270101TargetOneOf = nil
	}

	// try to unmarshal data into CohortScope20270101TargetOneOf4454865216
	err = datadog.Unmarshal(data, &obj.CohortScope20270101TargetOneOf4454865216)
	if err == nil {
		if obj.CohortScope20270101TargetOneOf4454865216 != nil && obj.CohortScope20270101TargetOneOf4454865216.UnparsedObject == nil {
			jsonCohortScope20270101TargetOneOf4454865216, _ := datadog.Marshal(obj.CohortScope20270101TargetOneOf4454865216)
			if string(jsonCohortScope20270101TargetOneOf4454865216) == "{}" { // empty struct
				obj.CohortScope20270101TargetOneOf4454865216 = nil
			} else {
				match++
			}
		} else {
			obj.CohortScope20270101TargetOneOf4454865216 = nil
		}
	} else {
		obj.CohortScope20270101TargetOneOf4454865216 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CohortScope20270101TargetOneOf = nil
		obj.CohortScope20270101TargetOneOf4454865216 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CohortScope20270101Target) MarshalJSON() ([]byte, error) {
	if obj.CohortScope20270101TargetOneOf != nil {
		return datadog.Marshal(&obj.CohortScope20270101TargetOneOf)
	}

	if obj.CohortScope20270101TargetOneOf4454865216 != nil {
		return datadog.Marshal(&obj.CohortScope20270101TargetOneOf4454865216)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CohortScope20270101Target) GetActualInstance() interface{} {
	if obj.CohortScope20270101TargetOneOf != nil {
		return obj.CohortScope20270101TargetOneOf
	}

	if obj.CohortScope20270101TargetOneOf4454865216 != nil {
		return obj.CohortScope20270101TargetOneOf4454865216
	}

	// all schemas are nil
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotRequests - Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotRequests struct {
	ScatterplotRequests20270101OneOf           *ScatterplotRequests20270101OneOf
	ScatterplotRequests20270101OneOf4461761120 *ScatterplotRequests20270101OneOf4461761120

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ScatterplotRequests20270101OneOfAsScatterplotRequests is a convenience function that returns ScatterplotRequests20270101OneOf wrapped in ScatterplotRequests.
func ScatterplotRequests20270101OneOfAsScatterplotRequests(v *ScatterplotRequests20270101OneOf) ScatterplotRequests {
	return ScatterplotRequests{ScatterplotRequests20270101OneOf: v}
}

// ScatterplotRequests20270101OneOf4461761120AsScatterplotRequests is a convenience function that returns ScatterplotRequests20270101OneOf4461761120 wrapped in ScatterplotRequests.
func ScatterplotRequests20270101OneOf4461761120AsScatterplotRequests(v *ScatterplotRequests20270101OneOf4461761120) ScatterplotRequests {
	return ScatterplotRequests{ScatterplotRequests20270101OneOf4461761120: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ScatterplotRequests) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ScatterplotRequests20270101OneOf
	err = datadog.Unmarshal(data, &obj.ScatterplotRequests20270101OneOf)
	if err == nil {
		if obj.ScatterplotRequests20270101OneOf != nil && obj.ScatterplotRequests20270101OneOf.UnparsedObject == nil {
			jsonScatterplotRequests20270101OneOf, _ := datadog.Marshal(obj.ScatterplotRequests20270101OneOf)
			if string(jsonScatterplotRequests20270101OneOf) == "{}" { // empty struct
				obj.ScatterplotRequests20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.ScatterplotRequests20270101OneOf = nil
		}
	} else {
		obj.ScatterplotRequests20270101OneOf = nil
	}

	// try to unmarshal data into ScatterplotRequests20270101OneOf4461761120
	err = datadog.Unmarshal(data, &obj.ScatterplotRequests20270101OneOf4461761120)
	if err == nil {
		if obj.ScatterplotRequests20270101OneOf4461761120 != nil && obj.ScatterplotRequests20270101OneOf4461761120.UnparsedObject == nil {
			jsonScatterplotRequests20270101OneOf4461761120, _ := datadog.Marshal(obj.ScatterplotRequests20270101OneOf4461761120)
			if string(jsonScatterplotRequests20270101OneOf4461761120) == "{}" { // empty struct
				obj.ScatterplotRequests20270101OneOf4461761120 = nil
			} else {
				match++
			}
		} else {
			obj.ScatterplotRequests20270101OneOf4461761120 = nil
		}
	} else {
		obj.ScatterplotRequests20270101OneOf4461761120 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ScatterplotRequests20270101OneOf = nil
		obj.ScatterplotRequests20270101OneOf4461761120 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ScatterplotRequests) MarshalJSON() ([]byte, error) {
	if obj.ScatterplotRequests20270101OneOf != nil {
		return datadog.Marshal(&obj.ScatterplotRequests20270101OneOf)
	}

	if obj.ScatterplotRequests20270101OneOf4461761120 != nil {
		return datadog.Marshal(&obj.ScatterplotRequests20270101OneOf4461761120)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ScatterplotRequests) GetActualInstance() interface{} {
	if obj.ScatterplotRequests20270101OneOf != nil {
		return obj.ScatterplotRequests20270101OneOf
	}

	if obj.ScatterplotRequests20270101OneOf4461761120 != nil {
		return obj.ScatterplotRequests20270101OneOf4461761120
	}

	// all schemas are nil
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlatSortBy - Test-only generated schema for the dashboard versioning proof of concept.
type FlatSortBy struct {
	FlatSortBy20270101OneOf           *FlatSortBy20270101OneOf
	FlatSortBy20270101OneOf4456386848 *FlatSortBy20270101OneOf4456386848

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FlatSortBy20270101OneOfAsFlatSortBy is a convenience function that returns FlatSortBy20270101OneOf wrapped in FlatSortBy.
func FlatSortBy20270101OneOfAsFlatSortBy(v *FlatSortBy20270101OneOf) FlatSortBy {
	return FlatSortBy{FlatSortBy20270101OneOf: v}
}

// FlatSortBy20270101OneOf4456386848AsFlatSortBy is a convenience function that returns FlatSortBy20270101OneOf4456386848 wrapped in FlatSortBy.
func FlatSortBy20270101OneOf4456386848AsFlatSortBy(v *FlatSortBy20270101OneOf4456386848) FlatSortBy {
	return FlatSortBy{FlatSortBy20270101OneOf4456386848: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *FlatSortBy) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FlatSortBy20270101OneOf
	err = datadog.Unmarshal(data, &obj.FlatSortBy20270101OneOf)
	if err == nil {
		if obj.FlatSortBy20270101OneOf != nil && obj.FlatSortBy20270101OneOf.UnparsedObject == nil {
			jsonFlatSortBy20270101OneOf, _ := datadog.Marshal(obj.FlatSortBy20270101OneOf)
			if string(jsonFlatSortBy20270101OneOf) == "{}" { // empty struct
				obj.FlatSortBy20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.FlatSortBy20270101OneOf = nil
		}
	} else {
		obj.FlatSortBy20270101OneOf = nil
	}

	// try to unmarshal data into FlatSortBy20270101OneOf4456386848
	err = datadog.Unmarshal(data, &obj.FlatSortBy20270101OneOf4456386848)
	if err == nil {
		if obj.FlatSortBy20270101OneOf4456386848 != nil && obj.FlatSortBy20270101OneOf4456386848.UnparsedObject == nil {
			jsonFlatSortBy20270101OneOf4456386848, _ := datadog.Marshal(obj.FlatSortBy20270101OneOf4456386848)
			if string(jsonFlatSortBy20270101OneOf4456386848) == "{}" { // empty struct
				obj.FlatSortBy20270101OneOf4456386848 = nil
			} else {
				match++
			}
		} else {
			obj.FlatSortBy20270101OneOf4456386848 = nil
		}
	} else {
		obj.FlatSortBy20270101OneOf4456386848 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FlatSortBy20270101OneOf = nil
		obj.FlatSortBy20270101OneOf4456386848 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj FlatSortBy) MarshalJSON() ([]byte, error) {
	if obj.FlatSortBy20270101OneOf != nil {
		return datadog.Marshal(&obj.FlatSortBy20270101OneOf)
	}

	if obj.FlatSortBy20270101OneOf4456386848 != nil {
		return datadog.Marshal(&obj.FlatSortBy20270101OneOf4456386848)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *FlatSortBy) GetActualInstance() interface{} {
	if obj.FlatSortBy20270101OneOf != nil {
		return obj.FlatSortBy20270101OneOf
	}

	if obj.FlatSortBy20270101OneOf4456386848 != nil {
		return obj.FlatSortBy20270101OneOf4456386848
	}

	// all schemas are nil
	return nil
}

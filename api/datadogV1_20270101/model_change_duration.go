// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeDuration - Test-only generated schema for the dashboard versioning proof of concept.
type ChangeDuration struct {
	ChangeDuration20270101OneOf           *ChangeDuration20270101OneOf
	ChangeDuration20270101OneOf4454627136 *ChangeDuration20270101OneOf4454627136

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ChangeDuration20270101OneOfAsChangeDuration is a convenience function that returns ChangeDuration20270101OneOf wrapped in ChangeDuration.
func ChangeDuration20270101OneOfAsChangeDuration(v *ChangeDuration20270101OneOf) ChangeDuration {
	return ChangeDuration{ChangeDuration20270101OneOf: v}
}

// ChangeDuration20270101OneOf4454627136AsChangeDuration is a convenience function that returns ChangeDuration20270101OneOf4454627136 wrapped in ChangeDuration.
func ChangeDuration20270101OneOf4454627136AsChangeDuration(v *ChangeDuration20270101OneOf4454627136) ChangeDuration {
	return ChangeDuration{ChangeDuration20270101OneOf4454627136: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ChangeDuration) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChangeDuration20270101OneOf
	err = datadog.Unmarshal(data, &obj.ChangeDuration20270101OneOf)
	if err == nil {
		if obj.ChangeDuration20270101OneOf != nil && obj.ChangeDuration20270101OneOf.UnparsedObject == nil {
			jsonChangeDuration20270101OneOf, _ := datadog.Marshal(obj.ChangeDuration20270101OneOf)
			if string(jsonChangeDuration20270101OneOf) == "{}" { // empty struct
				obj.ChangeDuration20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.ChangeDuration20270101OneOf = nil
		}
	} else {
		obj.ChangeDuration20270101OneOf = nil
	}

	// try to unmarshal data into ChangeDuration20270101OneOf4454627136
	err = datadog.Unmarshal(data, &obj.ChangeDuration20270101OneOf4454627136)
	if err == nil {
		if obj.ChangeDuration20270101OneOf4454627136 != nil && obj.ChangeDuration20270101OneOf4454627136.UnparsedObject == nil {
			jsonChangeDuration20270101OneOf4454627136, _ := datadog.Marshal(obj.ChangeDuration20270101OneOf4454627136)
			if string(jsonChangeDuration20270101OneOf4454627136) == "{}" { // empty struct
				obj.ChangeDuration20270101OneOf4454627136 = nil
			} else {
				match++
			}
		} else {
			obj.ChangeDuration20270101OneOf4454627136 = nil
		}
	} else {
		obj.ChangeDuration20270101OneOf4454627136 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ChangeDuration20270101OneOf = nil
		obj.ChangeDuration20270101OneOf4454627136 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ChangeDuration) MarshalJSON() ([]byte, error) {
	if obj.ChangeDuration20270101OneOf != nil {
		return datadog.Marshal(&obj.ChangeDuration20270101OneOf)
	}

	if obj.ChangeDuration20270101OneOf4454627136 != nil {
		return datadog.Marshal(&obj.ChangeDuration20270101OneOf4454627136)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ChangeDuration) GetActualInstance() interface{} {
	if obj.ChangeDuration20270101OneOf != nil {
		return obj.ChangeDuration20270101OneOf
	}

	if obj.ChangeDuration20270101OneOf4454627136 != nil {
		return obj.ChangeDuration20270101OneOf4454627136
	}

	// all schemas are nil
	return nil
}

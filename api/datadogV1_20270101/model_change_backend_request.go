// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeBackendRequest - Test-only generated schema for the dashboard versioning proof of concept.
type ChangeBackendRequest struct {
	ChangeBackendRequest20270101OneOf           *ChangeBackendRequest20270101OneOf
	ChangeBackendRequest20270101OneOf4454622976 *ChangeBackendRequest20270101OneOf4454622976

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ChangeBackendRequest20270101OneOfAsChangeBackendRequest is a convenience function that returns ChangeBackendRequest20270101OneOf wrapped in ChangeBackendRequest.
func ChangeBackendRequest20270101OneOfAsChangeBackendRequest(v *ChangeBackendRequest20270101OneOf) ChangeBackendRequest {
	return ChangeBackendRequest{ChangeBackendRequest20270101OneOf: v}
}

// ChangeBackendRequest20270101OneOf4454622976AsChangeBackendRequest is a convenience function that returns ChangeBackendRequest20270101OneOf4454622976 wrapped in ChangeBackendRequest.
func ChangeBackendRequest20270101OneOf4454622976AsChangeBackendRequest(v *ChangeBackendRequest20270101OneOf4454622976) ChangeBackendRequest {
	return ChangeBackendRequest{ChangeBackendRequest20270101OneOf4454622976: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ChangeBackendRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChangeBackendRequest20270101OneOf
	err = datadog.Unmarshal(data, &obj.ChangeBackendRequest20270101OneOf)
	if err == nil {
		if obj.ChangeBackendRequest20270101OneOf != nil && obj.ChangeBackendRequest20270101OneOf.UnparsedObject == nil {
			jsonChangeBackendRequest20270101OneOf, _ := datadog.Marshal(obj.ChangeBackendRequest20270101OneOf)
			if string(jsonChangeBackendRequest20270101OneOf) == "{}" { // empty struct
				obj.ChangeBackendRequest20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.ChangeBackendRequest20270101OneOf = nil
		}
	} else {
		obj.ChangeBackendRequest20270101OneOf = nil
	}

	// try to unmarshal data into ChangeBackendRequest20270101OneOf4454622976
	err = datadog.Unmarshal(data, &obj.ChangeBackendRequest20270101OneOf4454622976)
	if err == nil {
		if obj.ChangeBackendRequest20270101OneOf4454622976 != nil && obj.ChangeBackendRequest20270101OneOf4454622976.UnparsedObject == nil {
			jsonChangeBackendRequest20270101OneOf4454622976, _ := datadog.Marshal(obj.ChangeBackendRequest20270101OneOf4454622976)
			if string(jsonChangeBackendRequest20270101OneOf4454622976) == "{}" { // empty struct
				obj.ChangeBackendRequest20270101OneOf4454622976 = nil
			} else {
				match++
			}
		} else {
			obj.ChangeBackendRequest20270101OneOf4454622976 = nil
		}
	} else {
		obj.ChangeBackendRequest20270101OneOf4454622976 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ChangeBackendRequest20270101OneOf = nil
		obj.ChangeBackendRequest20270101OneOf4454622976 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ChangeBackendRequest) MarshalJSON() ([]byte, error) {
	if obj.ChangeBackendRequest20270101OneOf != nil {
		return datadog.Marshal(&obj.ChangeBackendRequest20270101OneOf)
	}

	if obj.ChangeBackendRequest20270101OneOf4454622976 != nil {
		return datadog.Marshal(&obj.ChangeBackendRequest20270101OneOf4454622976)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ChangeBackendRequest) GetActualInstance() interface{} {
	if obj.ChangeBackendRequest20270101OneOf != nil {
		return obj.ChangeBackendRequest20270101OneOf
	}

	if obj.ChangeBackendRequest20270101OneOf4454622976 != nil {
		return obj.ChangeBackendRequest20270101OneOf4454622976
	}

	// all schemas are nil
	return nil
}

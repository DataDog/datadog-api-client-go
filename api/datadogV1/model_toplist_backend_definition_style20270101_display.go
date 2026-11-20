// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendDefinitionStyle20270101Display - Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendDefinitionStyle20270101Display struct {
	ToplistBackendDefinitionStyle20270101DisplayOneOf           *ToplistBackendDefinitionStyle20270101DisplayOneOf
	ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744 *ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ToplistBackendDefinitionStyle20270101DisplayOneOfAsToplistBackendDefinitionStyle20270101Display is a convenience function that returns ToplistBackendDefinitionStyle20270101DisplayOneOf wrapped in ToplistBackendDefinitionStyle20270101Display.
func ToplistBackendDefinitionStyle20270101DisplayOneOfAsToplistBackendDefinitionStyle20270101Display(v *ToplistBackendDefinitionStyle20270101DisplayOneOf) ToplistBackendDefinitionStyle20270101Display {
	return ToplistBackendDefinitionStyle20270101Display{ToplistBackendDefinitionStyle20270101DisplayOneOf: v}
}

// ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744AsToplistBackendDefinitionStyle20270101Display is a convenience function that returns ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744 wrapped in ToplistBackendDefinitionStyle20270101Display.
func ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744AsToplistBackendDefinitionStyle20270101Display(v *ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744) ToplistBackendDefinitionStyle20270101Display {
	return ToplistBackendDefinitionStyle20270101Display{ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ToplistBackendDefinitionStyle20270101Display) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ToplistBackendDefinitionStyle20270101DisplayOneOf
	err = datadog.Unmarshal(data, &obj.ToplistBackendDefinitionStyle20270101DisplayOneOf)
	if err == nil {
		if obj.ToplistBackendDefinitionStyle20270101DisplayOneOf != nil && obj.ToplistBackendDefinitionStyle20270101DisplayOneOf.UnparsedObject == nil {
			jsonToplistBackendDefinitionStyle20270101DisplayOneOf, _ := datadog.Marshal(obj.ToplistBackendDefinitionStyle20270101DisplayOneOf)
			if string(jsonToplistBackendDefinitionStyle20270101DisplayOneOf) == "{}" { // empty struct
				obj.ToplistBackendDefinitionStyle20270101DisplayOneOf = nil
			} else {
				match++
			}
		} else {
			obj.ToplistBackendDefinitionStyle20270101DisplayOneOf = nil
		}
	} else {
		obj.ToplistBackendDefinitionStyle20270101DisplayOneOf = nil
	}

	// try to unmarshal data into ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744
	err = datadog.Unmarshal(data, &obj.ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744)
	if err == nil {
		if obj.ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744 != nil && obj.ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744.UnparsedObject == nil {
			jsonToplistBackendDefinitionStyle20270101DisplayOneOf4453175744, _ := datadog.Marshal(obj.ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744)
			if string(jsonToplistBackendDefinitionStyle20270101DisplayOneOf4453175744) == "{}" { // empty struct
				obj.ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744 = nil
			} else {
				match++
			}
		} else {
			obj.ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744 = nil
		}
	} else {
		obj.ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ToplistBackendDefinitionStyle20270101DisplayOneOf = nil
		obj.ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ToplistBackendDefinitionStyle20270101Display) MarshalJSON() ([]byte, error) {
	if obj.ToplistBackendDefinitionStyle20270101DisplayOneOf != nil {
		return datadog.Marshal(&obj.ToplistBackendDefinitionStyle20270101DisplayOneOf)
	}

	if obj.ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744 != nil {
		return datadog.Marshal(&obj.ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ToplistBackendDefinitionStyle20270101Display) GetActualInstance() interface{} {
	if obj.ToplistBackendDefinitionStyle20270101DisplayOneOf != nil {
		return obj.ToplistBackendDefinitionStyle20270101DisplayOneOf
	}

	if obj.ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744 != nil {
		return obj.ToplistBackendDefinitionStyle20270101DisplayOneOf4453175744
	}

	// all schemas are nil
	return nil
}

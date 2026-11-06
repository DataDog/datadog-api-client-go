// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryCreateAttributes - Targeted recovery (the default) or query-based batch recovery.
type RumRecoveryCreateAttributes struct {
	RumRecoverySessionsCreateAttributes *RumRecoverySessionsCreateAttributes
	RumRecoveryBatchCreateAttributes    *RumRecoveryBatchCreateAttributes

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// RumRecoverySessionsCreateAttributesAsRumRecoveryCreateAttributes is a convenience function that returns RumRecoverySessionsCreateAttributes wrapped in RumRecoveryCreateAttributes.
func RumRecoverySessionsCreateAttributesAsRumRecoveryCreateAttributes(v *RumRecoverySessionsCreateAttributes) RumRecoveryCreateAttributes {
	return RumRecoveryCreateAttributes{RumRecoverySessionsCreateAttributes: v}
}

// RumRecoveryBatchCreateAttributesAsRumRecoveryCreateAttributes is a convenience function that returns RumRecoveryBatchCreateAttributes wrapped in RumRecoveryCreateAttributes.
func RumRecoveryBatchCreateAttributesAsRumRecoveryCreateAttributes(v *RumRecoveryBatchCreateAttributes) RumRecoveryCreateAttributes {
	return RumRecoveryCreateAttributes{RumRecoveryBatchCreateAttributes: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *RumRecoveryCreateAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RumRecoverySessionsCreateAttributes
	err = datadog.Unmarshal(data, &obj.RumRecoverySessionsCreateAttributes)
	if err == nil {
		if obj.RumRecoverySessionsCreateAttributes != nil && obj.RumRecoverySessionsCreateAttributes.UnparsedObject == nil {
			jsonRumRecoverySessionsCreateAttributes, _ := datadog.Marshal(obj.RumRecoverySessionsCreateAttributes)
			if string(jsonRumRecoverySessionsCreateAttributes) == "{}" { // empty struct
				obj.RumRecoverySessionsCreateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.RumRecoverySessionsCreateAttributes = nil
		}
	} else {
		obj.RumRecoverySessionsCreateAttributes = nil
	}

	// try to unmarshal data into RumRecoveryBatchCreateAttributes
	err = datadog.Unmarshal(data, &obj.RumRecoveryBatchCreateAttributes)
	if err == nil {
		if obj.RumRecoveryBatchCreateAttributes != nil && obj.RumRecoveryBatchCreateAttributes.UnparsedObject == nil {
			jsonRumRecoveryBatchCreateAttributes, _ := datadog.Marshal(obj.RumRecoveryBatchCreateAttributes)
			if string(jsonRumRecoveryBatchCreateAttributes) == "{}" { // empty struct
				obj.RumRecoveryBatchCreateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.RumRecoveryBatchCreateAttributes = nil
		}
	} else {
		obj.RumRecoveryBatchCreateAttributes = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.RumRecoverySessionsCreateAttributes = nil
		obj.RumRecoveryBatchCreateAttributes = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj RumRecoveryCreateAttributes) MarshalJSON() ([]byte, error) {
	if obj.RumRecoverySessionsCreateAttributes != nil {
		return datadog.Marshal(&obj.RumRecoverySessionsCreateAttributes)
	}

	if obj.RumRecoveryBatchCreateAttributes != nil {
		return datadog.Marshal(&obj.RumRecoveryBatchCreateAttributes)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *RumRecoveryCreateAttributes) GetActualInstance() interface{} {
	if obj.RumRecoverySessionsCreateAttributes != nil {
		return obj.RumRecoverySessionsCreateAttributes
	}

	if obj.RumRecoveryBatchCreateAttributes != nil {
		return obj.RumRecoveryBatchCreateAttributes
	}

	// all schemas are nil
	return nil
}

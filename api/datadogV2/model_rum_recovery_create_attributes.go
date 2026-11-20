// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRecoveryCreateAttributes - Targeted recovery (the default) or query-based batch recovery.
type RUMRecoveryCreateAttributes struct {
	RUMRecoverySessionsCreateAttributes *RUMRecoverySessionsCreateAttributes
	RUMRecoveryBatchCreateAttributes    *RUMRecoveryBatchCreateAttributes

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// RUMRecoverySessionsCreateAttributesAsRUMRecoveryCreateAttributes is a convenience function that returns RUMRecoverySessionsCreateAttributes wrapped in RUMRecoveryCreateAttributes.
func RUMRecoverySessionsCreateAttributesAsRUMRecoveryCreateAttributes(v *RUMRecoverySessionsCreateAttributes) RUMRecoveryCreateAttributes {
	return RUMRecoveryCreateAttributes{RUMRecoverySessionsCreateAttributes: v}
}

// RUMRecoveryBatchCreateAttributesAsRUMRecoveryCreateAttributes is a convenience function that returns RUMRecoveryBatchCreateAttributes wrapped in RUMRecoveryCreateAttributes.
func RUMRecoveryBatchCreateAttributesAsRUMRecoveryCreateAttributes(v *RUMRecoveryBatchCreateAttributes) RUMRecoveryCreateAttributes {
	return RUMRecoveryCreateAttributes{RUMRecoveryBatchCreateAttributes: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *RUMRecoveryCreateAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RUMRecoverySessionsCreateAttributes
	err = datadog.Unmarshal(data, &obj.RUMRecoverySessionsCreateAttributes)
	if err == nil {
		if obj.RUMRecoverySessionsCreateAttributes != nil && obj.RUMRecoverySessionsCreateAttributes.UnparsedObject == nil {
			jsonRUMRecoverySessionsCreateAttributes, _ := datadog.Marshal(obj.RUMRecoverySessionsCreateAttributes)
			if string(jsonRUMRecoverySessionsCreateAttributes) == "{}" { // empty struct
				obj.RUMRecoverySessionsCreateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.RUMRecoverySessionsCreateAttributes = nil
		}
	} else {
		obj.RUMRecoverySessionsCreateAttributes = nil
	}

	// try to unmarshal data into RUMRecoveryBatchCreateAttributes
	err = datadog.Unmarshal(data, &obj.RUMRecoveryBatchCreateAttributes)
	if err == nil {
		if obj.RUMRecoveryBatchCreateAttributes != nil && obj.RUMRecoveryBatchCreateAttributes.UnparsedObject == nil {
			jsonRUMRecoveryBatchCreateAttributes, _ := datadog.Marshal(obj.RUMRecoveryBatchCreateAttributes)
			if string(jsonRUMRecoveryBatchCreateAttributes) == "{}" { // empty struct
				obj.RUMRecoveryBatchCreateAttributes = nil
			} else {
				match++
			}
		} else {
			obj.RUMRecoveryBatchCreateAttributes = nil
		}
	} else {
		obj.RUMRecoveryBatchCreateAttributes = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.RUMRecoverySessionsCreateAttributes = nil
		obj.RUMRecoveryBatchCreateAttributes = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj RUMRecoveryCreateAttributes) MarshalJSON() ([]byte, error) {
	if obj.RUMRecoverySessionsCreateAttributes != nil {
		return datadog.Marshal(&obj.RUMRecoverySessionsCreateAttributes)
	}

	if obj.RUMRecoveryBatchCreateAttributes != nil {
		return datadog.Marshal(&obj.RUMRecoveryBatchCreateAttributes)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *RUMRecoveryCreateAttributes) GetActualInstance() interface{} {
	if obj.RUMRecoverySessionsCreateAttributes != nil {
		return obj.RUMRecoverySessionsCreateAttributes
	}

	if obj.RUMRecoveryBatchCreateAttributes != nil {
		return obj.RUMRecoveryBatchCreateAttributes
	}

	// all schemas are nil
	return nil
}

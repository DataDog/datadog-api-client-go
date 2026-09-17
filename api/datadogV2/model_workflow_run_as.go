// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowRunAs - Identity used to run the workflow.
type WorkflowRunAs struct {
	WorkflowRunAsOwner          *WorkflowRunAsOwner
	WorkflowRunAsServiceAccount *WorkflowRunAsServiceAccount
	WorkflowRunAsInitiator      *WorkflowRunAsInitiator

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// WorkflowRunAsOwnerAsWorkflowRunAs is a convenience function that returns WorkflowRunAsOwner wrapped in WorkflowRunAs.
func WorkflowRunAsOwnerAsWorkflowRunAs(v *WorkflowRunAsOwner) WorkflowRunAs {
	return WorkflowRunAs{WorkflowRunAsOwner: v}
}

// WorkflowRunAsServiceAccountAsWorkflowRunAs is a convenience function that returns WorkflowRunAsServiceAccount wrapped in WorkflowRunAs.
func WorkflowRunAsServiceAccountAsWorkflowRunAs(v *WorkflowRunAsServiceAccount) WorkflowRunAs {
	return WorkflowRunAs{WorkflowRunAsServiceAccount: v}
}

// WorkflowRunAsInitiatorAsWorkflowRunAs is a convenience function that returns WorkflowRunAsInitiator wrapped in WorkflowRunAs.
func WorkflowRunAsInitiatorAsWorkflowRunAs(v *WorkflowRunAsInitiator) WorkflowRunAs {
	return WorkflowRunAs{WorkflowRunAsInitiator: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *WorkflowRunAs) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into WorkflowRunAsOwner
	err = datadog.Unmarshal(data, &obj.WorkflowRunAsOwner)
	if err == nil {
		if obj.WorkflowRunAsOwner != nil && obj.WorkflowRunAsOwner.UnparsedObject == nil {
			jsonWorkflowRunAsOwner, _ := datadog.Marshal(obj.WorkflowRunAsOwner)
			if string(jsonWorkflowRunAsOwner) == "{}" { // empty struct
				obj.WorkflowRunAsOwner = nil
			} else {
				match++
			}
		} else {
			obj.WorkflowRunAsOwner = nil
		}
	} else {
		obj.WorkflowRunAsOwner = nil
	}

	// try to unmarshal data into WorkflowRunAsServiceAccount
	err = datadog.Unmarshal(data, &obj.WorkflowRunAsServiceAccount)
	if err == nil {
		if obj.WorkflowRunAsServiceAccount != nil && obj.WorkflowRunAsServiceAccount.UnparsedObject == nil {
			jsonWorkflowRunAsServiceAccount, _ := datadog.Marshal(obj.WorkflowRunAsServiceAccount)
			if string(jsonWorkflowRunAsServiceAccount) == "{}" { // empty struct
				obj.WorkflowRunAsServiceAccount = nil
			} else {
				match++
			}
		} else {
			obj.WorkflowRunAsServiceAccount = nil
		}
	} else {
		obj.WorkflowRunAsServiceAccount = nil
	}

	// try to unmarshal data into WorkflowRunAsInitiator
	err = datadog.Unmarshal(data, &obj.WorkflowRunAsInitiator)
	if err == nil {
		if obj.WorkflowRunAsInitiator != nil && obj.WorkflowRunAsInitiator.UnparsedObject == nil {
			jsonWorkflowRunAsInitiator, _ := datadog.Marshal(obj.WorkflowRunAsInitiator)
			if string(jsonWorkflowRunAsInitiator) == "{}" { // empty struct
				obj.WorkflowRunAsInitiator = nil
			} else {
				match++
			}
		} else {
			obj.WorkflowRunAsInitiator = nil
		}
	} else {
		obj.WorkflowRunAsInitiator = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.WorkflowRunAsOwner = nil
		obj.WorkflowRunAsServiceAccount = nil
		obj.WorkflowRunAsInitiator = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj WorkflowRunAs) MarshalJSON() ([]byte, error) {
	if obj.WorkflowRunAsOwner != nil {
		return datadog.Marshal(&obj.WorkflowRunAsOwner)
	}

	if obj.WorkflowRunAsServiceAccount != nil {
		return datadog.Marshal(&obj.WorkflowRunAsServiceAccount)
	}

	if obj.WorkflowRunAsInitiator != nil {
		return datadog.Marshal(&obj.WorkflowRunAsInitiator)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *WorkflowRunAs) GetActualInstance() interface{} {
	if obj.WorkflowRunAsOwner != nil {
		return obj.WorkflowRunAsOwner
	}

	if obj.WorkflowRunAsServiceAccount != nil {
		return obj.WorkflowRunAsServiceAccount
	}

	if obj.WorkflowRunAsInitiator != nil {
		return obj.WorkflowRunAsInitiator
	}

	// all schemas are nil
	return nil
}

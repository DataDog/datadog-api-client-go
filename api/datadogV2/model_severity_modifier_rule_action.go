// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityModifierRuleAction - The action to take when a severity modifier rule matches a finding. This is a discriminated union on `type`: `set` assigns a fixed severity, while `shift` moves the severity up or down by one rank.
//
// A severity modifier rule's `rule.query` must not filter on `@severity` or on the `@severity_details.user_adjusted.*` namespace.
//
// Use `@severity_details.adjusted.value` instead, which reflects the severity before user-defined adjustments.
type SeverityModifierRuleAction struct {
	SeverityModifierRuleSetAction   *SeverityModifierRuleSetAction
	SeverityModifierRuleShiftAction *SeverityModifierRuleShiftAction

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SeverityModifierRuleSetActionAsSeverityModifierRuleAction is a convenience function that returns SeverityModifierRuleSetAction wrapped in SeverityModifierRuleAction.
func SeverityModifierRuleSetActionAsSeverityModifierRuleAction(v *SeverityModifierRuleSetAction) SeverityModifierRuleAction {
	return SeverityModifierRuleAction{SeverityModifierRuleSetAction: v}
}

// SeverityModifierRuleShiftActionAsSeverityModifierRuleAction is a convenience function that returns SeverityModifierRuleShiftAction wrapped in SeverityModifierRuleAction.
func SeverityModifierRuleShiftActionAsSeverityModifierRuleAction(v *SeverityModifierRuleShiftAction) SeverityModifierRuleAction {
	return SeverityModifierRuleAction{SeverityModifierRuleShiftAction: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SeverityModifierRuleAction) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SeverityModifierRuleSetAction
	err = datadog.Unmarshal(data, &obj.SeverityModifierRuleSetAction)
	if err == nil {
		if obj.SeverityModifierRuleSetAction != nil && obj.SeverityModifierRuleSetAction.UnparsedObject == nil {
			jsonSeverityModifierRuleSetAction, _ := datadog.Marshal(obj.SeverityModifierRuleSetAction)
			if string(jsonSeverityModifierRuleSetAction) == "{}" { // empty struct
				obj.SeverityModifierRuleSetAction = nil
			} else {
				match++
			}
		} else {
			obj.SeverityModifierRuleSetAction = nil
		}
	} else {
		obj.SeverityModifierRuleSetAction = nil
	}

	// try to unmarshal data into SeverityModifierRuleShiftAction
	err = datadog.Unmarshal(data, &obj.SeverityModifierRuleShiftAction)
	if err == nil {
		if obj.SeverityModifierRuleShiftAction != nil && obj.SeverityModifierRuleShiftAction.UnparsedObject == nil {
			jsonSeverityModifierRuleShiftAction, _ := datadog.Marshal(obj.SeverityModifierRuleShiftAction)
			if string(jsonSeverityModifierRuleShiftAction) == "{}" { // empty struct
				obj.SeverityModifierRuleShiftAction = nil
			} else {
				match++
			}
		} else {
			obj.SeverityModifierRuleShiftAction = nil
		}
	} else {
		obj.SeverityModifierRuleShiftAction = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SeverityModifierRuleSetAction = nil
		obj.SeverityModifierRuleShiftAction = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SeverityModifierRuleAction) MarshalJSON() ([]byte, error) {
	if obj.SeverityModifierRuleSetAction != nil {
		return datadog.Marshal(&obj.SeverityModifierRuleSetAction)
	}

	if obj.SeverityModifierRuleShiftAction != nil {
		return datadog.Marshal(&obj.SeverityModifierRuleShiftAction)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SeverityModifierRuleAction) GetActualInstance() interface{} {
	if obj.SeverityModifierRuleSetAction != nil {
		return obj.SeverityModifierRuleSetAction
	}

	if obj.SeverityModifierRuleShiftAction != nil {
		return obj.SeverityModifierRuleShiftAction
	}

	// all schemas are nil
	return nil
}

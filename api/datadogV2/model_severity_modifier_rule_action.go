// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityModifierRuleAction - The action to take when a severity modifier rule matches a finding. This is a discriminated union on `type`: `set` assigns a fixed severity, while `shift` moves the severity up or down by one rank.
//
// A severity modifier rule's `rule.query` must not filter on `@severity` or on the `@severity_details.user_adjusted.*` namespace; use `@severity_details.adjusted.value` to filter on the Datadog-adjusted severity instead.
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
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = datadog.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup.")
	}
	// check if the discriminator value is 'set'
	if jsonDict["type"] == "set" {
		// try to unmarshal JSON data into SeverityModifierRuleSetAction
		err = datadog.Unmarshal(data, &obj.SeverityModifierRuleSetAction)
		if err == nil {
			return nil // data stored in obj.SeverityModifierRuleSetAction, return on the first match
		} else {
			obj.SeverityModifierRuleSetAction = nil
			return fmt.Errorf("failed to unmarshal SeverityModifierRuleAction as SeverityModifierRuleSetAction: %s", err.Error())
		}
	}
	// check if the discriminator value is 'shift'
	if jsonDict["type"] == "shift" {
		// try to unmarshal JSON data into SeverityModifierRuleShiftAction
		err = datadog.Unmarshal(data, &obj.SeverityModifierRuleShiftAction)
		if err == nil {
			return nil // data stored in obj.SeverityModifierRuleShiftAction, return on the first match
		} else {
			obj.SeverityModifierRuleShiftAction = nil
			return fmt.Errorf("failed to unmarshal SeverityModifierRuleAction as SeverityModifierRuleShiftAction: %s", err.Error())
		}
	}
	return nil
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

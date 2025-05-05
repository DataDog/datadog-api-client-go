// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RoutingRuleAction - Defines an action that is executed when a routing rule matches certain criteria.
type RoutingRuleAction struct {
	SlackAction *SlackAction
	TeamsAction *TeamsAction

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SlackActionAsRoutingRuleAction is a convenience function that returns SlackAction wrapped in RoutingRuleAction.
func SlackActionAsRoutingRuleAction(v *SlackAction) RoutingRuleAction {
	return RoutingRuleAction{SlackAction: v}
}

// TeamsActionAsRoutingRuleAction is a convenience function that returns TeamsAction wrapped in RoutingRuleAction.
func TeamsActionAsRoutingRuleAction(v *TeamsAction) RoutingRuleAction {
	return RoutingRuleAction{TeamsAction: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *RoutingRuleAction) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SlackAction
	err = datadog.Unmarshal(data, &obj.SlackAction)
	if err == nil {
		if obj.SlackAction != nil && obj.SlackAction.UnparsedObject == nil {
			jsonSlackAction, _ := datadog.Marshal(obj.SlackAction)
			if string(jsonSlackAction) == "{}" { // empty struct
				obj.SlackAction = nil
			} else {
				match++
			}
		} else {
			obj.SlackAction = nil
		}
	} else {
		obj.SlackAction = nil
	}

	// try to unmarshal data into TeamsAction
	err = datadog.Unmarshal(data, &obj.TeamsAction)
	if err == nil {
		if obj.TeamsAction != nil && obj.TeamsAction.UnparsedObject == nil {
			jsonTeamsAction, _ := datadog.Marshal(obj.TeamsAction)
			if string(jsonTeamsAction) == "{}" { // empty struct
				obj.TeamsAction = nil
			} else {
				match++
			}
		} else {
			obj.TeamsAction = nil
		}
	} else {
		obj.TeamsAction = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SlackAction = nil
		obj.TeamsAction = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj RoutingRuleAction) MarshalJSON() ([]byte, error) {
	if obj.SlackAction != nil {
		return datadog.Marshal(&obj.SlackAction)
	}

	if obj.TeamsAction != nil {
		return datadog.Marshal(&obj.TeamsAction)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *RoutingRuleAction) GetActualInstance() interface{} {
	if obj.SlackAction != nil {
		return obj.SlackAction
	}

	if obj.TeamsAction != nil {
		return obj.TeamsAction
	}

	// all schemas are nil
	return nil
}

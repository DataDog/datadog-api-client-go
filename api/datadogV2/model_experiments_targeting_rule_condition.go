// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsTargetingRuleCondition - A saved-filter condition or a complete inline condition. The two forms cannot be combined.
type ExperimentsTargetingRuleCondition struct {
	ExperimentsSavedFilterCondition *ExperimentsSavedFilterCondition
	ExperimentsInlineCondition      *ExperimentsInlineCondition

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ExperimentsSavedFilterConditionAsExperimentsTargetingRuleCondition is a convenience function that returns ExperimentsSavedFilterCondition wrapped in ExperimentsTargetingRuleCondition.
func ExperimentsSavedFilterConditionAsExperimentsTargetingRuleCondition(v *ExperimentsSavedFilterCondition) ExperimentsTargetingRuleCondition {
	return ExperimentsTargetingRuleCondition{ExperimentsSavedFilterCondition: v}
}

// ExperimentsInlineConditionAsExperimentsTargetingRuleCondition is a convenience function that returns ExperimentsInlineCondition wrapped in ExperimentsTargetingRuleCondition.
func ExperimentsInlineConditionAsExperimentsTargetingRuleCondition(v *ExperimentsInlineCondition) ExperimentsTargetingRuleCondition {
	return ExperimentsTargetingRuleCondition{ExperimentsInlineCondition: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ExperimentsTargetingRuleCondition) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ExperimentsSavedFilterCondition
	err = datadog.Unmarshal(data, &obj.ExperimentsSavedFilterCondition)
	if err == nil {
		if obj.ExperimentsSavedFilterCondition != nil && obj.ExperimentsSavedFilterCondition.UnparsedObject == nil {
			jsonExperimentsSavedFilterCondition, _ := datadog.Marshal(obj.ExperimentsSavedFilterCondition)
			if string(jsonExperimentsSavedFilterCondition) == "{}" { // empty struct
				obj.ExperimentsSavedFilterCondition = nil
			} else {
				match++
			}
		} else {
			obj.ExperimentsSavedFilterCondition = nil
		}
	} else {
		obj.ExperimentsSavedFilterCondition = nil
	}

	// try to unmarshal data into ExperimentsInlineCondition
	err = datadog.Unmarshal(data, &obj.ExperimentsInlineCondition)
	if err == nil {
		if obj.ExperimentsInlineCondition != nil && obj.ExperimentsInlineCondition.UnparsedObject == nil {
			jsonExperimentsInlineCondition, _ := datadog.Marshal(obj.ExperimentsInlineCondition)
			if string(jsonExperimentsInlineCondition) == "{}" { // empty struct
				obj.ExperimentsInlineCondition = nil
			} else {
				match++
			}
		} else {
			obj.ExperimentsInlineCondition = nil
		}
	} else {
		obj.ExperimentsInlineCondition = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ExperimentsSavedFilterCondition = nil
		obj.ExperimentsInlineCondition = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ExperimentsTargetingRuleCondition) MarshalJSON() ([]byte, error) {
	if obj.ExperimentsSavedFilterCondition != nil {
		return datadog.Marshal(&obj.ExperimentsSavedFilterCondition)
	}

	if obj.ExperimentsInlineCondition != nil {
		return datadog.Marshal(&obj.ExperimentsInlineCondition)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ExperimentsTargetingRuleCondition) GetActualInstance() interface{} {
	if obj.ExperimentsSavedFilterCondition != nil {
		return obj.ExperimentsSavedFilterCondition
	}

	if obj.ExperimentsInlineCondition != nil {
		return obj.ExperimentsInlineCondition
	}

	// all schemas are nil
	return nil
}

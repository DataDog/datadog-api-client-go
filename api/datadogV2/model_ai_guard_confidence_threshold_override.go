// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardConfidenceThresholdOverride - Request-only confidence threshold override. A number sets the threshold
// for every threat category; an object supplies a global threshold and
// optional per-category overrides. Omit this field to use the stored
// sensitivity policy. This field is not returned in evaluation responses.
type AIGuardConfidenceThresholdOverride struct {
	AIGuardGlobalConfidenceThreshold *float64
	AIGuardConfidenceThreshold       *AIGuardConfidenceThreshold

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AIGuardGlobalConfidenceThresholdAsAIGuardConfidenceThresholdOverride is a convenience function that returns float64 wrapped in AIGuardConfidenceThresholdOverride.
func AIGuardGlobalConfidenceThresholdAsAIGuardConfidenceThresholdOverride(v *float64) AIGuardConfidenceThresholdOverride {
	return AIGuardConfidenceThresholdOverride{AIGuardGlobalConfidenceThreshold: v}
}

// AIGuardConfidenceThresholdAsAIGuardConfidenceThresholdOverride is a convenience function that returns AIGuardConfidenceThreshold wrapped in AIGuardConfidenceThresholdOverride.
func AIGuardConfidenceThresholdAsAIGuardConfidenceThresholdOverride(v *AIGuardConfidenceThreshold) AIGuardConfidenceThresholdOverride {
	return AIGuardConfidenceThresholdOverride{AIGuardConfidenceThreshold: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AIGuardConfidenceThresholdOverride) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AIGuardGlobalConfidenceThreshold
	err = datadog.Unmarshal(data, &obj.AIGuardGlobalConfidenceThreshold)
	if err == nil {
		if obj.AIGuardGlobalConfidenceThreshold != nil {
			jsonAIGuardGlobalConfidenceThreshold, _ := datadog.Marshal(obj.AIGuardGlobalConfidenceThreshold)
			if string(jsonAIGuardGlobalConfidenceThreshold) == "{}" { // empty struct
				obj.AIGuardGlobalConfidenceThreshold = nil
			} else {
				match++
			}
		} else {
			obj.AIGuardGlobalConfidenceThreshold = nil
		}
	} else {
		obj.AIGuardGlobalConfidenceThreshold = nil
	}

	// try to unmarshal data into AIGuardConfidenceThreshold
	err = datadog.Unmarshal(data, &obj.AIGuardConfidenceThreshold)
	if err == nil {
		if obj.AIGuardConfidenceThreshold != nil && obj.AIGuardConfidenceThreshold.UnparsedObject == nil {
			jsonAIGuardConfidenceThreshold, _ := datadog.Marshal(obj.AIGuardConfidenceThreshold)
			if string(jsonAIGuardConfidenceThreshold) == "{}" && string(data) != "{}" { // empty struct
				obj.AIGuardConfidenceThreshold = nil
			} else {
				match++
			}
		} else {
			obj.AIGuardConfidenceThreshold = nil
		}
	} else {
		obj.AIGuardConfidenceThreshold = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AIGuardGlobalConfidenceThreshold = nil
		obj.AIGuardConfidenceThreshold = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AIGuardConfidenceThresholdOverride) MarshalJSON() ([]byte, error) {
	if obj.AIGuardGlobalConfidenceThreshold != nil {
		return datadog.Marshal(&obj.AIGuardGlobalConfidenceThreshold)
	}

	if obj.AIGuardConfidenceThreshold != nil {
		return datadog.Marshal(&obj.AIGuardConfidenceThreshold)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AIGuardConfidenceThresholdOverride) GetActualInstance() interface{} {
	if obj.AIGuardGlobalConfidenceThreshold != nil {
		return obj.AIGuardGlobalConfidenceThreshold
	}

	if obj.AIGuardConfidenceThreshold != nil {
		return obj.AIGuardConfidenceThreshold
	}

	// all schemas are nil
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineParseGrokProcessorRuleItem - A single Grok parsing rule, selected by either source field or include query.
type ObservabilityPipelineParseGrokProcessorRuleItem struct {
	ObservabilityPipelineParseGrokProcessorRule        *ObservabilityPipelineParseGrokProcessorRule
	ObservabilityPipelineParseGrokProcessorIncludeRule *ObservabilityPipelineParseGrokProcessorIncludeRule

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineParseGrokProcessorRuleAsObservabilityPipelineParseGrokProcessorRuleItem is a convenience function that returns ObservabilityPipelineParseGrokProcessorRule wrapped in ObservabilityPipelineParseGrokProcessorRuleItem.
func ObservabilityPipelineParseGrokProcessorRuleAsObservabilityPipelineParseGrokProcessorRuleItem(v *ObservabilityPipelineParseGrokProcessorRule) ObservabilityPipelineParseGrokProcessorRuleItem {
	return ObservabilityPipelineParseGrokProcessorRuleItem{ObservabilityPipelineParseGrokProcessorRule: v}
}

// ObservabilityPipelineParseGrokProcessorIncludeRuleAsObservabilityPipelineParseGrokProcessorRuleItem is a convenience function that returns ObservabilityPipelineParseGrokProcessorIncludeRule wrapped in ObservabilityPipelineParseGrokProcessorRuleItem.
func ObservabilityPipelineParseGrokProcessorIncludeRuleAsObservabilityPipelineParseGrokProcessorRuleItem(v *ObservabilityPipelineParseGrokProcessorIncludeRule) ObservabilityPipelineParseGrokProcessorRuleItem {
	return ObservabilityPipelineParseGrokProcessorRuleItem{ObservabilityPipelineParseGrokProcessorIncludeRule: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineParseGrokProcessorRuleItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineParseGrokProcessorRule
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineParseGrokProcessorRule)
	if err == nil {
		if obj.ObservabilityPipelineParseGrokProcessorRule != nil && obj.ObservabilityPipelineParseGrokProcessorRule.UnparsedObject == nil {
			jsonObservabilityPipelineParseGrokProcessorRule, _ := datadog.Marshal(obj.ObservabilityPipelineParseGrokProcessorRule)
			if string(jsonObservabilityPipelineParseGrokProcessorRule) == "{}" { // empty struct
				obj.ObservabilityPipelineParseGrokProcessorRule = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineParseGrokProcessorRule = nil
		}
	} else {
		obj.ObservabilityPipelineParseGrokProcessorRule = nil
	}

	// try to unmarshal data into ObservabilityPipelineParseGrokProcessorIncludeRule
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineParseGrokProcessorIncludeRule)
	if err == nil {
		if obj.ObservabilityPipelineParseGrokProcessorIncludeRule != nil && obj.ObservabilityPipelineParseGrokProcessorIncludeRule.UnparsedObject == nil {
			jsonObservabilityPipelineParseGrokProcessorIncludeRule, _ := datadog.Marshal(obj.ObservabilityPipelineParseGrokProcessorIncludeRule)
			if string(jsonObservabilityPipelineParseGrokProcessorIncludeRule) == "{}" { // empty struct
				obj.ObservabilityPipelineParseGrokProcessorIncludeRule = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineParseGrokProcessorIncludeRule = nil
		}
	} else {
		obj.ObservabilityPipelineParseGrokProcessorIncludeRule = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineParseGrokProcessorRule = nil
		obj.ObservabilityPipelineParseGrokProcessorIncludeRule = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineParseGrokProcessorRuleItem) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineParseGrokProcessorRule != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineParseGrokProcessorRule)
	}

	if obj.ObservabilityPipelineParseGrokProcessorIncludeRule != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineParseGrokProcessorIncludeRule)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineParseGrokProcessorRuleItem) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineParseGrokProcessorRule != nil {
		return obj.ObservabilityPipelineParseGrokProcessorRule
	}

	if obj.ObservabilityPipelineParseGrokProcessorIncludeRule != nil {
		return obj.ObservabilityPipelineParseGrokProcessorIncludeRule
	}

	// all schemas are nil
	return nil
}

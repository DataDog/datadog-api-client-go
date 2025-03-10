// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineDataAttributesConfigProcessorsItem - The `config` `processors`.
type PipelineDataAttributesConfigProcessorsItem struct {
	FilterProcessor    *FilterProcessor
	ParseJSONProcessor *ParseJSONProcessor

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FilterProcessorAsPipelineDataAttributesConfigProcessorsItem is a convenience function that returns FilterProcessor wrapped in PipelineDataAttributesConfigProcessorsItem.
func FilterProcessorAsPipelineDataAttributesConfigProcessorsItem(v *FilterProcessor) PipelineDataAttributesConfigProcessorsItem {
	return PipelineDataAttributesConfigProcessorsItem{FilterProcessor: v}
}

// ParseJSONProcessorAsPipelineDataAttributesConfigProcessorsItem is a convenience function that returns ParseJSONProcessor wrapped in PipelineDataAttributesConfigProcessorsItem.
func ParseJSONProcessorAsPipelineDataAttributesConfigProcessorsItem(v *ParseJSONProcessor) PipelineDataAttributesConfigProcessorsItem {
	return PipelineDataAttributesConfigProcessorsItem{ParseJSONProcessor: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *PipelineDataAttributesConfigProcessorsItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FilterProcessor
	err = datadog.Unmarshal(data, &obj.FilterProcessor)
	if err == nil {
		if obj.FilterProcessor != nil && obj.FilterProcessor.UnparsedObject == nil {
			jsonFilterProcessor, _ := datadog.Marshal(obj.FilterProcessor)
			if string(jsonFilterProcessor) == "{}" { // empty struct
				obj.FilterProcessor = nil
			} else {
				match++
			}
		} else {
			obj.FilterProcessor = nil
		}
	} else {
		obj.FilterProcessor = nil
	}

	// try to unmarshal data into ParseJSONProcessor
	err = datadog.Unmarshal(data, &obj.ParseJSONProcessor)
	if err == nil {
		if obj.ParseJSONProcessor != nil && obj.ParseJSONProcessor.UnparsedObject == nil {
			jsonParseJSONProcessor, _ := datadog.Marshal(obj.ParseJSONProcessor)
			if string(jsonParseJSONProcessor) == "{}" { // empty struct
				obj.ParseJSONProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ParseJSONProcessor = nil
		}
	} else {
		obj.ParseJSONProcessor = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FilterProcessor = nil
		obj.ParseJSONProcessor = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj PipelineDataAttributesConfigProcessorsItem) MarshalJSON() ([]byte, error) {
	if obj.FilterProcessor != nil {
		return datadog.Marshal(&obj.FilterProcessor)
	}

	if obj.ParseJSONProcessor != nil {
		return datadog.Marshal(&obj.ParseJSONProcessor)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *PipelineDataAttributesConfigProcessorsItem) GetActualInstance() interface{} {
	if obj.FilterProcessor != nil {
		return obj.FilterProcessor
	}

	if obj.ParseJSONProcessor != nil {
		return obj.ParseJSONProcessor
	}

	// all schemas are nil
	return nil
}

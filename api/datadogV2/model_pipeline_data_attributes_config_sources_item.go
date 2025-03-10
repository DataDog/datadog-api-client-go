// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineDataAttributesConfigSourcesItem - The `config` `sources`.
type PipelineDataAttributesConfigSourcesItem struct {
	DatadogAgentSource *DatadogAgentSource

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DatadogAgentSourceAsPipelineDataAttributesConfigSourcesItem is a convenience function that returns DatadogAgentSource wrapped in PipelineDataAttributesConfigSourcesItem.
func DatadogAgentSourceAsPipelineDataAttributesConfigSourcesItem(v *DatadogAgentSource) PipelineDataAttributesConfigSourcesItem {
	return PipelineDataAttributesConfigSourcesItem{DatadogAgentSource: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *PipelineDataAttributesConfigSourcesItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DatadogAgentSource
	err = datadog.Unmarshal(data, &obj.DatadogAgentSource)
	if err == nil {
		if obj.DatadogAgentSource != nil && obj.DatadogAgentSource.UnparsedObject == nil {
			jsonDatadogAgentSource, _ := datadog.Marshal(obj.DatadogAgentSource)
			if string(jsonDatadogAgentSource) == "{}" { // empty struct
				obj.DatadogAgentSource = nil
			} else {
				match++
			}
		} else {
			obj.DatadogAgentSource = nil
		}
	} else {
		obj.DatadogAgentSource = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DatadogAgentSource = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj PipelineDataAttributesConfigSourcesItem) MarshalJSON() ([]byte, error) {
	if obj.DatadogAgentSource != nil {
		return datadog.Marshal(&obj.DatadogAgentSource)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *PipelineDataAttributesConfigSourcesItem) GetActualInstance() interface{} {
	if obj.DatadogAgentSource != nil {
		return obj.DatadogAgentSource
	}

	// all schemas are nil
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineDataAttributesConfigDestinationsItem - The `config` `destinations`.
type PipelineDataAttributesConfigDestinationsItem struct {
	DatadogLogsDestination *DatadogLogsDestination

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DatadogLogsDestinationAsPipelineDataAttributesConfigDestinationsItem is a convenience function that returns DatadogLogsDestination wrapped in PipelineDataAttributesConfigDestinationsItem.
func DatadogLogsDestinationAsPipelineDataAttributesConfigDestinationsItem(v *DatadogLogsDestination) PipelineDataAttributesConfigDestinationsItem {
	return PipelineDataAttributesConfigDestinationsItem{DatadogLogsDestination: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *PipelineDataAttributesConfigDestinationsItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DatadogLogsDestination
	err = datadog.Unmarshal(data, &obj.DatadogLogsDestination)
	if err == nil {
		if obj.DatadogLogsDestination != nil && obj.DatadogLogsDestination.UnparsedObject == nil {
			jsonDatadogLogsDestination, _ := datadog.Marshal(obj.DatadogLogsDestination)
			if string(jsonDatadogLogsDestination) == "{}" { // empty struct
				obj.DatadogLogsDestination = nil
			} else {
				match++
			}
		} else {
			obj.DatadogLogsDestination = nil
		}
	} else {
		obj.DatadogLogsDestination = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DatadogLogsDestination = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj PipelineDataAttributesConfigDestinationsItem) MarshalJSON() ([]byte, error) {
	if obj.DatadogLogsDestination != nil {
		return datadog.Marshal(&obj.DatadogLogsDestination)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *PipelineDataAttributesConfigDestinationsItem) GetActualInstance() interface{} {
	if obj.DatadogLogsDestination != nil {
		return obj.DatadogLogsDestination
	}

	// all schemas are nil
	return nil
}

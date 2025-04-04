// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineConfigDestination - A destination for the pipeline.
type PipelineConfigDestination struct {
	PipelineDatadogLogsDestination *PipelineDatadogLogsDestination

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// PipelineDatadogLogsDestinationAsPipelineConfigDestination is a convenience function that returns PipelineDatadogLogsDestination wrapped in PipelineConfigDestination.
func PipelineDatadogLogsDestinationAsPipelineConfigDestination(v *PipelineDatadogLogsDestination) PipelineConfigDestination {
	return PipelineConfigDestination{PipelineDatadogLogsDestination: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *PipelineConfigDestination) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PipelineDatadogLogsDestination
	err = datadog.Unmarshal(data, &obj.PipelineDatadogLogsDestination)
	if err == nil {
		if obj.PipelineDatadogLogsDestination != nil && obj.PipelineDatadogLogsDestination.UnparsedObject == nil {
			jsonPipelineDatadogLogsDestination, _ := datadog.Marshal(obj.PipelineDatadogLogsDestination)
			if string(jsonPipelineDatadogLogsDestination) == "{}" { // empty struct
				obj.PipelineDatadogLogsDestination = nil
			} else {
				match++
			}
		} else {
			obj.PipelineDatadogLogsDestination = nil
		}
	} else {
		obj.PipelineDatadogLogsDestination = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.PipelineDatadogLogsDestination = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj PipelineConfigDestination) MarshalJSON() ([]byte, error) {
	if obj.PipelineDatadogLogsDestination != nil {
		return datadog.Marshal(&obj.PipelineDatadogLogsDestination)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *PipelineConfigDestination) GetActualInstance() interface{} {
	if obj.PipelineDatadogLogsDestination != nil {
		return obj.PipelineDatadogLogsDestination
	}

	// all schemas are nil
	return nil
}

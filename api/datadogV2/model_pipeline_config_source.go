// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineConfigSource - A data source for the pipeline.
type PipelineConfigSource struct {
	PipelineKafkaSource        *PipelineKafkaSource
	PipelineDatadogAgentSource *PipelineDatadogAgentSource

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// PipelineKafkaSourceAsPipelineConfigSource is a convenience function that returns PipelineKafkaSource wrapped in PipelineConfigSource.
func PipelineKafkaSourceAsPipelineConfigSource(v *PipelineKafkaSource) PipelineConfigSource {
	return PipelineConfigSource{PipelineKafkaSource: v}
}

// PipelineDatadogAgentSourceAsPipelineConfigSource is a convenience function that returns PipelineDatadogAgentSource wrapped in PipelineConfigSource.
func PipelineDatadogAgentSourceAsPipelineConfigSource(v *PipelineDatadogAgentSource) PipelineConfigSource {
	return PipelineConfigSource{PipelineDatadogAgentSource: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *PipelineConfigSource) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PipelineKafkaSource
	err = datadog.Unmarshal(data, &obj.PipelineKafkaSource)
	if err == nil {
		if obj.PipelineKafkaSource != nil && obj.PipelineKafkaSource.UnparsedObject == nil {
			jsonPipelineKafkaSource, _ := datadog.Marshal(obj.PipelineKafkaSource)
			if string(jsonPipelineKafkaSource) == "{}" { // empty struct
				obj.PipelineKafkaSource = nil
			} else {
				match++
			}
		} else {
			obj.PipelineKafkaSource = nil
		}
	} else {
		obj.PipelineKafkaSource = nil
	}

	// try to unmarshal data into PipelineDatadogAgentSource
	err = datadog.Unmarshal(data, &obj.PipelineDatadogAgentSource)
	if err == nil {
		if obj.PipelineDatadogAgentSource != nil && obj.PipelineDatadogAgentSource.UnparsedObject == nil {
			jsonPipelineDatadogAgentSource, _ := datadog.Marshal(obj.PipelineDatadogAgentSource)
			if string(jsonPipelineDatadogAgentSource) == "{}" { // empty struct
				obj.PipelineDatadogAgentSource = nil
			} else {
				match++
			}
		} else {
			obj.PipelineDatadogAgentSource = nil
		}
	} else {
		obj.PipelineDatadogAgentSource = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.PipelineKafkaSource = nil
		obj.PipelineDatadogAgentSource = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj PipelineConfigSource) MarshalJSON() ([]byte, error) {
	if obj.PipelineKafkaSource != nil {
		return datadog.Marshal(&obj.PipelineKafkaSource)
	}

	if obj.PipelineDatadogAgentSource != nil {
		return datadog.Marshal(&obj.PipelineDatadogAgentSource)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *PipelineConfigSource) GetActualInstance() interface{} {
	if obj.PipelineKafkaSource != nil {
		return obj.PipelineKafkaSource
	}

	if obj.PipelineDatadogAgentSource != nil {
		return obj.PipelineDatadogAgentSource
	}

	// all schemas are nil
	return nil
}

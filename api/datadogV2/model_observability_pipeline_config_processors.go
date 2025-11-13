// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineConfigProcessors - A list of processors that transform or enrich log data, or a list of grouped processor configurations.
type ObservabilityPipelineConfigProcessors struct {
	ObservabilityPipelineConfigProcessorsList  *[]ObservabilityPipelineConfigProcessorItem
	ObservabilityPipelineConfigProcessorGroups *[]ObservabilityPipelineConfigProcessorGroup

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineConfigProcessorsListAsObservabilityPipelineConfigProcessors is a convenience function that returns []ObservabilityPipelineConfigProcessorItem wrapped in ObservabilityPipelineConfigProcessors.
func ObservabilityPipelineConfigProcessorsListAsObservabilityPipelineConfigProcessors(v *[]ObservabilityPipelineConfigProcessorItem) ObservabilityPipelineConfigProcessors {
	return ObservabilityPipelineConfigProcessors{ObservabilityPipelineConfigProcessorsList: v}
}

// ObservabilityPipelineConfigProcessorGroupsAsObservabilityPipelineConfigProcessors is a convenience function that returns []ObservabilityPipelineConfigProcessorGroup wrapped in ObservabilityPipelineConfigProcessors.
func ObservabilityPipelineConfigProcessorGroupsAsObservabilityPipelineConfigProcessors(v *[]ObservabilityPipelineConfigProcessorGroup) ObservabilityPipelineConfigProcessors {
	return ObservabilityPipelineConfigProcessors{ObservabilityPipelineConfigProcessorGroups: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineConfigProcessors) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineConfigProcessorsList
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineConfigProcessorsList)
	if err == nil {
		if obj.ObservabilityPipelineConfigProcessorsList != nil {
			jsonObservabilityPipelineConfigProcessorsList, _ := datadog.Marshal(obj.ObservabilityPipelineConfigProcessorsList)
			if string(jsonObservabilityPipelineConfigProcessorsList) == "{}" && string(data) != "{}" { // empty struct
				obj.ObservabilityPipelineConfigProcessorsList = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineConfigProcessorsList = nil
		}
	} else {
		obj.ObservabilityPipelineConfigProcessorsList = nil
	}

	// try to unmarshal data into ObservabilityPipelineConfigProcessorGroups
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineConfigProcessorGroups)
	if err == nil {
		if obj.ObservabilityPipelineConfigProcessorGroups != nil {
			jsonObservabilityPipelineConfigProcessorGroups, _ := datadog.Marshal(obj.ObservabilityPipelineConfigProcessorGroups)
			if string(jsonObservabilityPipelineConfigProcessorGroups) == "{}" && string(data) != "{}" { // empty struct
				obj.ObservabilityPipelineConfigProcessorGroups = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineConfigProcessorGroups = nil
		}
	} else {
		obj.ObservabilityPipelineConfigProcessorGroups = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineConfigProcessorsList = nil
		obj.ObservabilityPipelineConfigProcessorGroups = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineConfigProcessors) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineConfigProcessorsList != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineConfigProcessorsList)
	}

	if obj.ObservabilityPipelineConfigProcessorGroups != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineConfigProcessorGroups)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineConfigProcessors) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineConfigProcessorsList != nil {
		return obj.ObservabilityPipelineConfigProcessorsList
	}

	if obj.ObservabilityPipelineConfigProcessorGroups != nil {
		return obj.ObservabilityPipelineConfigProcessorGroups
	}

	// all schemas are nil
	return nil
}

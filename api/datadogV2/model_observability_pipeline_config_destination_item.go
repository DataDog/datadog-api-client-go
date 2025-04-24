// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineConfigDestinationItem - A destination for the pipeline.
type ObservabilityPipelineConfigDestinationItem struct {
	ObservabilityPipelineDatadogLogsDestination     *ObservabilityPipelineDatadogLogsDestination
	ObservabilityPipelineGoogleChronicleDestination *ObservabilityPipelineGoogleChronicleDestination
	ObservabilityPipelineNewRelicDestination        *ObservabilityPipelineNewRelicDestination
	ObservabilityPipelineSentinelOneDestination     *ObservabilityPipelineSentinelOneDestination

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineDatadogLogsDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineDatadogLogsDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineDatadogLogsDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineDatadogLogsDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineDatadogLogsDestination: v}
}

// ObservabilityPipelineGoogleChronicleDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineGoogleChronicleDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineGoogleChronicleDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineGoogleChronicleDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineGoogleChronicleDestination: v}
}

// ObservabilityPipelineNewRelicDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineNewRelicDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineNewRelicDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineNewRelicDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineNewRelicDestination: v}
}

// ObservabilityPipelineSentinelOneDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineSentinelOneDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineSentinelOneDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineSentinelOneDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineSentinelOneDestination: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineConfigDestinationItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineDatadogLogsDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineDatadogLogsDestination)
	if err == nil {
		if obj.ObservabilityPipelineDatadogLogsDestination != nil && obj.ObservabilityPipelineDatadogLogsDestination.UnparsedObject == nil {
			jsonObservabilityPipelineDatadogLogsDestination, _ := datadog.Marshal(obj.ObservabilityPipelineDatadogLogsDestination)
			if string(jsonObservabilityPipelineDatadogLogsDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineDatadogLogsDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineDatadogLogsDestination = nil
		}
	} else {
		obj.ObservabilityPipelineDatadogLogsDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineGoogleChronicleDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineGoogleChronicleDestination)
	if err == nil {
		if obj.ObservabilityPipelineGoogleChronicleDestination != nil && obj.ObservabilityPipelineGoogleChronicleDestination.UnparsedObject == nil {
			jsonObservabilityPipelineGoogleChronicleDestination, _ := datadog.Marshal(obj.ObservabilityPipelineGoogleChronicleDestination)
			if string(jsonObservabilityPipelineGoogleChronicleDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineGoogleChronicleDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineGoogleChronicleDestination = nil
		}
	} else {
		obj.ObservabilityPipelineGoogleChronicleDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineNewRelicDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineNewRelicDestination)
	if err == nil {
		if obj.ObservabilityPipelineNewRelicDestination != nil && obj.ObservabilityPipelineNewRelicDestination.UnparsedObject == nil {
			jsonObservabilityPipelineNewRelicDestination, _ := datadog.Marshal(obj.ObservabilityPipelineNewRelicDestination)
			if string(jsonObservabilityPipelineNewRelicDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineNewRelicDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineNewRelicDestination = nil
		}
	} else {
		obj.ObservabilityPipelineNewRelicDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineSentinelOneDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSentinelOneDestination)
	if err == nil {
		if obj.ObservabilityPipelineSentinelOneDestination != nil && obj.ObservabilityPipelineSentinelOneDestination.UnparsedObject == nil {
			jsonObservabilityPipelineSentinelOneDestination, _ := datadog.Marshal(obj.ObservabilityPipelineSentinelOneDestination)
			if string(jsonObservabilityPipelineSentinelOneDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineSentinelOneDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSentinelOneDestination = nil
		}
	} else {
		obj.ObservabilityPipelineSentinelOneDestination = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineDatadogLogsDestination = nil
		obj.ObservabilityPipelineGoogleChronicleDestination = nil
		obj.ObservabilityPipelineNewRelicDestination = nil
		obj.ObservabilityPipelineSentinelOneDestination = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineConfigDestinationItem) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineDatadogLogsDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineDatadogLogsDestination)
	}

	if obj.ObservabilityPipelineGoogleChronicleDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineGoogleChronicleDestination)
	}

	if obj.ObservabilityPipelineNewRelicDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineNewRelicDestination)
	}

	if obj.ObservabilityPipelineSentinelOneDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSentinelOneDestination)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineConfigDestinationItem) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineDatadogLogsDestination != nil {
		return obj.ObservabilityPipelineDatadogLogsDestination
	}

	if obj.ObservabilityPipelineGoogleChronicleDestination != nil {
		return obj.ObservabilityPipelineGoogleChronicleDestination
	}

	if obj.ObservabilityPipelineNewRelicDestination != nil {
		return obj.ObservabilityPipelineNewRelicDestination
	}

	if obj.ObservabilityPipelineSentinelOneDestination != nil {
		return obj.ObservabilityPipelineSentinelOneDestination
	}

	// all schemas are nil
	return nil
}

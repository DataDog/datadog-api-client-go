// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineConfigDestinationItem - A destination for the pipeline.
type ObservabilityPipelineConfigDestinationItem struct {
	ObservabilityPipelineDatadogLogsDestination        *ObservabilityPipelineDatadogLogsDestination
	ObservabilityPipelineAmazonS3Destination           *ObservabilityPipelineAmazonS3Destination
	ObservabilityPipelineGoogleCloudStorageDestination *ObservabilityPipelineGoogleCloudStorageDestination
	ObservabilityPipelineSplunkHecDestination          *ObservabilityPipelineSplunkHecDestination

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineDatadogLogsDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineDatadogLogsDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineDatadogLogsDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineDatadogLogsDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineDatadogLogsDestination: v}
}

// ObservabilityPipelineAmazonS3DestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineAmazonS3Destination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineAmazonS3DestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineAmazonS3Destination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineAmazonS3Destination: v}
}

// ObservabilityPipelineGoogleCloudStorageDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineGoogleCloudStorageDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineGoogleCloudStorageDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineGoogleCloudStorageDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineGoogleCloudStorageDestination: v}
}

// ObservabilityPipelineSplunkHecDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineSplunkHecDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineSplunkHecDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineSplunkHecDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineSplunkHecDestination: v}
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

	// try to unmarshal data into ObservabilityPipelineAmazonS3Destination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAmazonS3Destination)
	if err == nil {
		if obj.ObservabilityPipelineAmazonS3Destination != nil && obj.ObservabilityPipelineAmazonS3Destination.UnparsedObject == nil {
			jsonObservabilityPipelineAmazonS3Destination, _ := datadog.Marshal(obj.ObservabilityPipelineAmazonS3Destination)
			if string(jsonObservabilityPipelineAmazonS3Destination) == "{}" { // empty struct
				obj.ObservabilityPipelineAmazonS3Destination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAmazonS3Destination = nil
		}
	} else {
		obj.ObservabilityPipelineAmazonS3Destination = nil
	}

	// try to unmarshal data into ObservabilityPipelineGoogleCloudStorageDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineGoogleCloudStorageDestination)
	if err == nil {
		if obj.ObservabilityPipelineGoogleCloudStorageDestination != nil && obj.ObservabilityPipelineGoogleCloudStorageDestination.UnparsedObject == nil {
			jsonObservabilityPipelineGoogleCloudStorageDestination, _ := datadog.Marshal(obj.ObservabilityPipelineGoogleCloudStorageDestination)
			if string(jsonObservabilityPipelineGoogleCloudStorageDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineGoogleCloudStorageDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineGoogleCloudStorageDestination = nil
		}
	} else {
		obj.ObservabilityPipelineGoogleCloudStorageDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineSplunkHecDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSplunkHecDestination)
	if err == nil {
		if obj.ObservabilityPipelineSplunkHecDestination != nil && obj.ObservabilityPipelineSplunkHecDestination.UnparsedObject == nil {
			jsonObservabilityPipelineSplunkHecDestination, _ := datadog.Marshal(obj.ObservabilityPipelineSplunkHecDestination)
			if string(jsonObservabilityPipelineSplunkHecDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineSplunkHecDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSplunkHecDestination = nil
		}
	} else {
		obj.ObservabilityPipelineSplunkHecDestination = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineDatadogLogsDestination = nil
		obj.ObservabilityPipelineAmazonS3Destination = nil
		obj.ObservabilityPipelineGoogleCloudStorageDestination = nil
		obj.ObservabilityPipelineSplunkHecDestination = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineConfigDestinationItem) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineDatadogLogsDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineDatadogLogsDestination)
	}

	if obj.ObservabilityPipelineAmazonS3Destination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAmazonS3Destination)
	}

	if obj.ObservabilityPipelineGoogleCloudStorageDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineGoogleCloudStorageDestination)
	}

	if obj.ObservabilityPipelineSplunkHecDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSplunkHecDestination)
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

	if obj.ObservabilityPipelineAmazonS3Destination != nil {
		return obj.ObservabilityPipelineAmazonS3Destination
	}

	if obj.ObservabilityPipelineGoogleCloudStorageDestination != nil {
		return obj.ObservabilityPipelineGoogleCloudStorageDestination
	}

	if obj.ObservabilityPipelineSplunkHecDestination != nil {
		return obj.ObservabilityPipelineSplunkHecDestination
	}

	// all schemas are nil
	return nil
}

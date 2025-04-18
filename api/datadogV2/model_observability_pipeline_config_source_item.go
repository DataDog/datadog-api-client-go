// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineConfigSourceItem - A data source for the pipeline.
type ObservabilityPipelineConfigSourceItem struct {
	ObservabilityPipelineKafkaSource        *ObservabilityPipelineKafkaSource
	ObservabilityPipelineDatadogAgentSource *ObservabilityPipelineDatadogAgentSource
	ObservabilityPipelineSplunkTcpSource    *ObservabilityPipelineSplunkTcpSource
	ObservabilityPipelineSplunkHecSource    *ObservabilityPipelineSplunkHecSource
	ObservabilityPipelineAmazonS3Source     *ObservabilityPipelineAmazonS3Source

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineKafkaSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineKafkaSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineKafkaSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineKafkaSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineKafkaSource: v}
}

// ObservabilityPipelineDatadogAgentSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineDatadogAgentSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineDatadogAgentSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineDatadogAgentSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineDatadogAgentSource: v}
}

// ObservabilityPipelineSplunkTcpSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineSplunkTcpSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineSplunkTcpSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineSplunkTcpSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineSplunkTcpSource: v}
}

// ObservabilityPipelineSplunkHecSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineSplunkHecSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineSplunkHecSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineSplunkHecSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineSplunkHecSource: v}
}

// ObservabilityPipelineAmazonS3SourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineAmazonS3Source wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineAmazonS3SourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineAmazonS3Source) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineAmazonS3Source: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineConfigSourceItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineKafkaSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineKafkaSource)
	if err == nil {
		if obj.ObservabilityPipelineKafkaSource != nil && obj.ObservabilityPipelineKafkaSource.UnparsedObject == nil {
			jsonObservabilityPipelineKafkaSource, _ := datadog.Marshal(obj.ObservabilityPipelineKafkaSource)
			if string(jsonObservabilityPipelineKafkaSource) == "{}" { // empty struct
				obj.ObservabilityPipelineKafkaSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineKafkaSource = nil
		}
	} else {
		obj.ObservabilityPipelineKafkaSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineDatadogAgentSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineDatadogAgentSource)
	if err == nil {
		if obj.ObservabilityPipelineDatadogAgentSource != nil && obj.ObservabilityPipelineDatadogAgentSource.UnparsedObject == nil {
			jsonObservabilityPipelineDatadogAgentSource, _ := datadog.Marshal(obj.ObservabilityPipelineDatadogAgentSource)
			if string(jsonObservabilityPipelineDatadogAgentSource) == "{}" { // empty struct
				obj.ObservabilityPipelineDatadogAgentSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineDatadogAgentSource = nil
		}
	} else {
		obj.ObservabilityPipelineDatadogAgentSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineSplunkTcpSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSplunkTcpSource)
	if err == nil {
		if obj.ObservabilityPipelineSplunkTcpSource != nil && obj.ObservabilityPipelineSplunkTcpSource.UnparsedObject == nil {
			jsonObservabilityPipelineSplunkTcpSource, _ := datadog.Marshal(obj.ObservabilityPipelineSplunkTcpSource)
			if string(jsonObservabilityPipelineSplunkTcpSource) == "{}" { // empty struct
				obj.ObservabilityPipelineSplunkTcpSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSplunkTcpSource = nil
		}
	} else {
		obj.ObservabilityPipelineSplunkTcpSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineSplunkHecSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSplunkHecSource)
	if err == nil {
		if obj.ObservabilityPipelineSplunkHecSource != nil && obj.ObservabilityPipelineSplunkHecSource.UnparsedObject == nil {
			jsonObservabilityPipelineSplunkHecSource, _ := datadog.Marshal(obj.ObservabilityPipelineSplunkHecSource)
			if string(jsonObservabilityPipelineSplunkHecSource) == "{}" { // empty struct
				obj.ObservabilityPipelineSplunkHecSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSplunkHecSource = nil
		}
	} else {
		obj.ObservabilityPipelineSplunkHecSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineAmazonS3Source
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAmazonS3Source)
	if err == nil {
		if obj.ObservabilityPipelineAmazonS3Source != nil && obj.ObservabilityPipelineAmazonS3Source.UnparsedObject == nil {
			jsonObservabilityPipelineAmazonS3Source, _ := datadog.Marshal(obj.ObservabilityPipelineAmazonS3Source)
			if string(jsonObservabilityPipelineAmazonS3Source) == "{}" { // empty struct
				obj.ObservabilityPipelineAmazonS3Source = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAmazonS3Source = nil
		}
	} else {
		obj.ObservabilityPipelineAmazonS3Source = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineKafkaSource = nil
		obj.ObservabilityPipelineDatadogAgentSource = nil
		obj.ObservabilityPipelineSplunkTcpSource = nil
		obj.ObservabilityPipelineSplunkHecSource = nil
		obj.ObservabilityPipelineAmazonS3Source = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineConfigSourceItem) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineKafkaSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineKafkaSource)
	}

	if obj.ObservabilityPipelineDatadogAgentSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineDatadogAgentSource)
	}

	if obj.ObservabilityPipelineSplunkTcpSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSplunkTcpSource)
	}

	if obj.ObservabilityPipelineSplunkHecSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSplunkHecSource)
	}

	if obj.ObservabilityPipelineAmazonS3Source != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAmazonS3Source)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineConfigSourceItem) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineKafkaSource != nil {
		return obj.ObservabilityPipelineKafkaSource
	}

	if obj.ObservabilityPipelineDatadogAgentSource != nil {
		return obj.ObservabilityPipelineDatadogAgentSource
	}

	if obj.ObservabilityPipelineSplunkTcpSource != nil {
		return obj.ObservabilityPipelineSplunkTcpSource
	}

	if obj.ObservabilityPipelineSplunkHecSource != nil {
		return obj.ObservabilityPipelineSplunkHecSource
	}

	if obj.ObservabilityPipelineAmazonS3Source != nil {
		return obj.ObservabilityPipelineAmazonS3Source
	}

	// all schemas are nil
	return nil
}

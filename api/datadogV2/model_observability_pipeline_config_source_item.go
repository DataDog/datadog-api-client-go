// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineConfigSourceItem - A data source for the pipeline.
type ObservabilityPipelineConfigSourceItem struct {
	ObservabilityPipelineKafkaSource              *ObservabilityPipelineKafkaSource
	ObservabilityPipelineDatadogAgentSource       *ObservabilityPipelineDatadogAgentSource
	ObservabilityPipelineAmazonDataFirehoseSource *ObservabilityPipelineAmazonDataFirehoseSource
	ObservabilityPipelineGooglePubSubSource       *ObservabilityPipelineGooglePubSubSource
	ObservabilityPipelineHttpClientSource         *ObservabilityPipelineHttpClientSource
	ObservabilityPipelineLogstashSource           *ObservabilityPipelineLogstashSource

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

// ObservabilityPipelineAmazonDataFirehoseSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineAmazonDataFirehoseSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineAmazonDataFirehoseSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineAmazonDataFirehoseSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineAmazonDataFirehoseSource: v}
}

// ObservabilityPipelineGooglePubSubSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineGooglePubSubSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineGooglePubSubSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineGooglePubSubSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineGooglePubSubSource: v}
}

// ObservabilityPipelineHttpClientSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineHttpClientSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineHttpClientSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineHttpClientSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineHttpClientSource: v}
}

// ObservabilityPipelineLogstashSourceAsObservabilityPipelineConfigSourceItem is a convenience function that returns ObservabilityPipelineLogstashSource wrapped in ObservabilityPipelineConfigSourceItem.
func ObservabilityPipelineLogstashSourceAsObservabilityPipelineConfigSourceItem(v *ObservabilityPipelineLogstashSource) ObservabilityPipelineConfigSourceItem {
	return ObservabilityPipelineConfigSourceItem{ObservabilityPipelineLogstashSource: v}
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

	// try to unmarshal data into ObservabilityPipelineAmazonDataFirehoseSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAmazonDataFirehoseSource)
	if err == nil {
		if obj.ObservabilityPipelineAmazonDataFirehoseSource != nil && obj.ObservabilityPipelineAmazonDataFirehoseSource.UnparsedObject == nil {
			jsonObservabilityPipelineAmazonDataFirehoseSource, _ := datadog.Marshal(obj.ObservabilityPipelineAmazonDataFirehoseSource)
			if string(jsonObservabilityPipelineAmazonDataFirehoseSource) == "{}" { // empty struct
				obj.ObservabilityPipelineAmazonDataFirehoseSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAmazonDataFirehoseSource = nil
		}
	} else {
		obj.ObservabilityPipelineAmazonDataFirehoseSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineGooglePubSubSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineGooglePubSubSource)
	if err == nil {
		if obj.ObservabilityPipelineGooglePubSubSource != nil && obj.ObservabilityPipelineGooglePubSubSource.UnparsedObject == nil {
			jsonObservabilityPipelineGooglePubSubSource, _ := datadog.Marshal(obj.ObservabilityPipelineGooglePubSubSource)
			if string(jsonObservabilityPipelineGooglePubSubSource) == "{}" { // empty struct
				obj.ObservabilityPipelineGooglePubSubSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineGooglePubSubSource = nil
		}
	} else {
		obj.ObservabilityPipelineGooglePubSubSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineHttpClientSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineHttpClientSource)
	if err == nil {
		if obj.ObservabilityPipelineHttpClientSource != nil && obj.ObservabilityPipelineHttpClientSource.UnparsedObject == nil {
			jsonObservabilityPipelineHttpClientSource, _ := datadog.Marshal(obj.ObservabilityPipelineHttpClientSource)
			if string(jsonObservabilityPipelineHttpClientSource) == "{}" { // empty struct
				obj.ObservabilityPipelineHttpClientSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineHttpClientSource = nil
		}
	} else {
		obj.ObservabilityPipelineHttpClientSource = nil
	}

	// try to unmarshal data into ObservabilityPipelineLogstashSource
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineLogstashSource)
	if err == nil {
		if obj.ObservabilityPipelineLogstashSource != nil && obj.ObservabilityPipelineLogstashSource.UnparsedObject == nil {
			jsonObservabilityPipelineLogstashSource, _ := datadog.Marshal(obj.ObservabilityPipelineLogstashSource)
			if string(jsonObservabilityPipelineLogstashSource) == "{}" { // empty struct
				obj.ObservabilityPipelineLogstashSource = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineLogstashSource = nil
		}
	} else {
		obj.ObservabilityPipelineLogstashSource = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineKafkaSource = nil
		obj.ObservabilityPipelineDatadogAgentSource = nil
		obj.ObservabilityPipelineAmazonDataFirehoseSource = nil
		obj.ObservabilityPipelineGooglePubSubSource = nil
		obj.ObservabilityPipelineHttpClientSource = nil
		obj.ObservabilityPipelineLogstashSource = nil
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

	if obj.ObservabilityPipelineAmazonDataFirehoseSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAmazonDataFirehoseSource)
	}

	if obj.ObservabilityPipelineGooglePubSubSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineGooglePubSubSource)
	}

	if obj.ObservabilityPipelineHttpClientSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineHttpClientSource)
	}

	if obj.ObservabilityPipelineLogstashSource != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineLogstashSource)
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

	if obj.ObservabilityPipelineAmazonDataFirehoseSource != nil {
		return obj.ObservabilityPipelineAmazonDataFirehoseSource
	}

	if obj.ObservabilityPipelineGooglePubSubSource != nil {
		return obj.ObservabilityPipelineGooglePubSubSource
	}

	if obj.ObservabilityPipelineHttpClientSource != nil {
		return obj.ObservabilityPipelineHttpClientSource
	}

	if obj.ObservabilityPipelineLogstashSource != nil {
		return obj.ObservabilityPipelineLogstashSource
	}

	// all schemas are nil
	return nil
}

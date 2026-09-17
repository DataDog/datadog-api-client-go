// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3DestinationCompression - Compression configuration for archived logs. When omitted, logs are compressed with gzip
// for backward compatibility.
type ObservabilityPipelineAmazonS3DestinationCompression struct {
	ObservabilityPipelineAmazonS3DestinationCompressionZstd *ObservabilityPipelineAmazonS3DestinationCompressionZstd
	ObservabilityPipelineAmazonS3DestinationCompressionGzip *ObservabilityPipelineAmazonS3DestinationCompressionGzip

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineAmazonS3DestinationCompressionZstdAsObservabilityPipelineAmazonS3DestinationCompression is a convenience function that returns ObservabilityPipelineAmazonS3DestinationCompressionZstd wrapped in ObservabilityPipelineAmazonS3DestinationCompression.
func ObservabilityPipelineAmazonS3DestinationCompressionZstdAsObservabilityPipelineAmazonS3DestinationCompression(v *ObservabilityPipelineAmazonS3DestinationCompressionZstd) ObservabilityPipelineAmazonS3DestinationCompression {
	return ObservabilityPipelineAmazonS3DestinationCompression{ObservabilityPipelineAmazonS3DestinationCompressionZstd: v}
}

// ObservabilityPipelineAmazonS3DestinationCompressionGzipAsObservabilityPipelineAmazonS3DestinationCompression is a convenience function that returns ObservabilityPipelineAmazonS3DestinationCompressionGzip wrapped in ObservabilityPipelineAmazonS3DestinationCompression.
func ObservabilityPipelineAmazonS3DestinationCompressionGzipAsObservabilityPipelineAmazonS3DestinationCompression(v *ObservabilityPipelineAmazonS3DestinationCompressionGzip) ObservabilityPipelineAmazonS3DestinationCompression {
	return ObservabilityPipelineAmazonS3DestinationCompression{ObservabilityPipelineAmazonS3DestinationCompressionGzip: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineAmazonS3DestinationCompression) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineAmazonS3DestinationCompressionZstd
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAmazonS3DestinationCompressionZstd)
	if err == nil {
		if obj.ObservabilityPipelineAmazonS3DestinationCompressionZstd != nil && obj.ObservabilityPipelineAmazonS3DestinationCompressionZstd.UnparsedObject == nil {
			jsonObservabilityPipelineAmazonS3DestinationCompressionZstd, _ := datadog.Marshal(obj.ObservabilityPipelineAmazonS3DestinationCompressionZstd)
			if string(jsonObservabilityPipelineAmazonS3DestinationCompressionZstd) == "{}" { // empty struct
				obj.ObservabilityPipelineAmazonS3DestinationCompressionZstd = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAmazonS3DestinationCompressionZstd = nil
		}
	} else {
		obj.ObservabilityPipelineAmazonS3DestinationCompressionZstd = nil
	}

	// try to unmarshal data into ObservabilityPipelineAmazonS3DestinationCompressionGzip
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAmazonS3DestinationCompressionGzip)
	if err == nil {
		if obj.ObservabilityPipelineAmazonS3DestinationCompressionGzip != nil && obj.ObservabilityPipelineAmazonS3DestinationCompressionGzip.UnparsedObject == nil {
			jsonObservabilityPipelineAmazonS3DestinationCompressionGzip, _ := datadog.Marshal(obj.ObservabilityPipelineAmazonS3DestinationCompressionGzip)
			if string(jsonObservabilityPipelineAmazonS3DestinationCompressionGzip) == "{}" { // empty struct
				obj.ObservabilityPipelineAmazonS3DestinationCompressionGzip = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAmazonS3DestinationCompressionGzip = nil
		}
	} else {
		obj.ObservabilityPipelineAmazonS3DestinationCompressionGzip = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineAmazonS3DestinationCompressionZstd = nil
		obj.ObservabilityPipelineAmazonS3DestinationCompressionGzip = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineAmazonS3DestinationCompression) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineAmazonS3DestinationCompressionZstd != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAmazonS3DestinationCompressionZstd)
	}

	if obj.ObservabilityPipelineAmazonS3DestinationCompressionGzip != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAmazonS3DestinationCompressionGzip)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineAmazonS3DestinationCompression) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineAmazonS3DestinationCompressionZstd != nil {
		return obj.ObservabilityPipelineAmazonS3DestinationCompressionZstd
	}

	if obj.ObservabilityPipelineAmazonS3DestinationCompressionGzip != nil {
		return obj.ObservabilityPipelineAmazonS3DestinationCompressionGzip
	}

	// all schemas are nil
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureStorageDestinationCompression - Compression configuration for archived logs. When omitted, logs are compressed with gzip
// for backward compatibility.
type ObservabilityPipelineAzureStorageDestinationCompression struct {
	ObservabilityPipelineAzureStorageDestinationCompressionZstd *ObservabilityPipelineAzureStorageDestinationCompressionZstd
	ObservabilityPipelineAzureStorageDestinationCompressionGzip *ObservabilityPipelineAzureStorageDestinationCompressionGzip

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineAzureStorageDestinationCompressionZstdAsObservabilityPipelineAzureStorageDestinationCompression is a convenience function that returns ObservabilityPipelineAzureStorageDestinationCompressionZstd wrapped in ObservabilityPipelineAzureStorageDestinationCompression.
func ObservabilityPipelineAzureStorageDestinationCompressionZstdAsObservabilityPipelineAzureStorageDestinationCompression(v *ObservabilityPipelineAzureStorageDestinationCompressionZstd) ObservabilityPipelineAzureStorageDestinationCompression {
	return ObservabilityPipelineAzureStorageDestinationCompression{ObservabilityPipelineAzureStorageDestinationCompressionZstd: v}
}

// ObservabilityPipelineAzureStorageDestinationCompressionGzipAsObservabilityPipelineAzureStorageDestinationCompression is a convenience function that returns ObservabilityPipelineAzureStorageDestinationCompressionGzip wrapped in ObservabilityPipelineAzureStorageDestinationCompression.
func ObservabilityPipelineAzureStorageDestinationCompressionGzipAsObservabilityPipelineAzureStorageDestinationCompression(v *ObservabilityPipelineAzureStorageDestinationCompressionGzip) ObservabilityPipelineAzureStorageDestinationCompression {
	return ObservabilityPipelineAzureStorageDestinationCompression{ObservabilityPipelineAzureStorageDestinationCompressionGzip: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineAzureStorageDestinationCompression) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineAzureStorageDestinationCompressionZstd
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAzureStorageDestinationCompressionZstd)
	if err == nil {
		if obj.ObservabilityPipelineAzureStorageDestinationCompressionZstd != nil && obj.ObservabilityPipelineAzureStorageDestinationCompressionZstd.UnparsedObject == nil {
			jsonObservabilityPipelineAzureStorageDestinationCompressionZstd, _ := datadog.Marshal(obj.ObservabilityPipelineAzureStorageDestinationCompressionZstd)
			if string(jsonObservabilityPipelineAzureStorageDestinationCompressionZstd) == "{}" { // empty struct
				obj.ObservabilityPipelineAzureStorageDestinationCompressionZstd = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAzureStorageDestinationCompressionZstd = nil
		}
	} else {
		obj.ObservabilityPipelineAzureStorageDestinationCompressionZstd = nil
	}

	// try to unmarshal data into ObservabilityPipelineAzureStorageDestinationCompressionGzip
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAzureStorageDestinationCompressionGzip)
	if err == nil {
		if obj.ObservabilityPipelineAzureStorageDestinationCompressionGzip != nil && obj.ObservabilityPipelineAzureStorageDestinationCompressionGzip.UnparsedObject == nil {
			jsonObservabilityPipelineAzureStorageDestinationCompressionGzip, _ := datadog.Marshal(obj.ObservabilityPipelineAzureStorageDestinationCompressionGzip)
			if string(jsonObservabilityPipelineAzureStorageDestinationCompressionGzip) == "{}" { // empty struct
				obj.ObservabilityPipelineAzureStorageDestinationCompressionGzip = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAzureStorageDestinationCompressionGzip = nil
		}
	} else {
		obj.ObservabilityPipelineAzureStorageDestinationCompressionGzip = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineAzureStorageDestinationCompressionZstd = nil
		obj.ObservabilityPipelineAzureStorageDestinationCompressionGzip = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineAzureStorageDestinationCompression) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineAzureStorageDestinationCompressionZstd != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAzureStorageDestinationCompressionZstd)
	}

	if obj.ObservabilityPipelineAzureStorageDestinationCompressionGzip != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAzureStorageDestinationCompressionGzip)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineAzureStorageDestinationCompression) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineAzureStorageDestinationCompressionZstd != nil {
		return obj.ObservabilityPipelineAzureStorageDestinationCompressionZstd
	}

	if obj.ObservabilityPipelineAzureStorageDestinationCompressionGzip != nil {
		return obj.ObservabilityPipelineAzureStorageDestinationCompressionGzip
	}

	// all schemas are nil
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGoogleCloudStorageDestinationCompression - Compression configuration for archived logs. When omitted, logs are compressed with gzip
// for backward compatibility.
type ObservabilityPipelineGoogleCloudStorageDestinationCompression struct {
	ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd *ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd
	ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip *ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdAsObservabilityPipelineGoogleCloudStorageDestinationCompression is a convenience function that returns ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd wrapped in ObservabilityPipelineGoogleCloudStorageDestinationCompression.
func ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdAsObservabilityPipelineGoogleCloudStorageDestinationCompression(v *ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd) ObservabilityPipelineGoogleCloudStorageDestinationCompression {
	return ObservabilityPipelineGoogleCloudStorageDestinationCompression{ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd: v}
}

// ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzipAsObservabilityPipelineGoogleCloudStorageDestinationCompression is a convenience function that returns ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip wrapped in ObservabilityPipelineGoogleCloudStorageDestinationCompression.
func ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzipAsObservabilityPipelineGoogleCloudStorageDestinationCompression(v *ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip) ObservabilityPipelineGoogleCloudStorageDestinationCompression {
	return ObservabilityPipelineGoogleCloudStorageDestinationCompression{ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineGoogleCloudStorageDestinationCompression) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd)
	if err == nil {
		if obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd != nil && obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd.UnparsedObject == nil {
			jsonObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd, _ := datadog.Marshal(obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd)
			if string(jsonObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd) == "{}" { // empty struct
				obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd = nil
		}
	} else {
		obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd = nil
	}

	// try to unmarshal data into ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip)
	if err == nil {
		if obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip != nil && obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip.UnparsedObject == nil {
			jsonObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip, _ := datadog.Marshal(obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip)
			if string(jsonObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip) == "{}" { // empty struct
				obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip = nil
		}
	} else {
		obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd = nil
		obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineGoogleCloudStorageDestinationCompression) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd)
	}

	if obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineGoogleCloudStorageDestinationCompression) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd != nil {
		return obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd
	}

	if obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip != nil {
		return obj.ObservabilityPipelineGoogleCloudStorageDestinationCompressionGzip
	}

	// all schemas are nil
	return nil
}

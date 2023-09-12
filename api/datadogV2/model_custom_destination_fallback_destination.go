// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"
)

// CustomDestinationFallbackDestination - The archiving destination to fall back to.
type CustomDestinationFallbackDestination struct {
	AzureFallbackDestination *AzureFallbackDestination
	GCSFallbackDestination   *GCSFallbackDestination
	S3FallbackDestination    *S3FallbackDestination

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AzureFallbackDestinationAsCustomDestinationFallbackDestination is a convenience function that returns AzureFallbackDestination wrapped in CustomDestinationFallbackDestination.
func AzureFallbackDestinationAsCustomDestinationFallbackDestination(v *AzureFallbackDestination) CustomDestinationFallbackDestination {
	return CustomDestinationFallbackDestination{AzureFallbackDestination: v}
}

// GCSFallbackDestinationAsCustomDestinationFallbackDestination is a convenience function that returns GCSFallbackDestination wrapped in CustomDestinationFallbackDestination.
func GCSFallbackDestinationAsCustomDestinationFallbackDestination(v *GCSFallbackDestination) CustomDestinationFallbackDestination {
	return CustomDestinationFallbackDestination{GCSFallbackDestination: v}
}

// S3FallbackDestinationAsCustomDestinationFallbackDestination is a convenience function that returns S3FallbackDestination wrapped in CustomDestinationFallbackDestination.
func S3FallbackDestinationAsCustomDestinationFallbackDestination(v *S3FallbackDestination) CustomDestinationFallbackDestination {
	return CustomDestinationFallbackDestination{S3FallbackDestination: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CustomDestinationFallbackDestination) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AzureFallbackDestination
	err = json.Unmarshal(data, &obj.AzureFallbackDestination)
	if err == nil {
		if obj.AzureFallbackDestination != nil && obj.AzureFallbackDestination.UnparsedObject == nil {
			jsonAzureFallbackDestination, _ := json.Marshal(obj.AzureFallbackDestination)
			if string(jsonAzureFallbackDestination) == "{}" { // empty struct
				obj.AzureFallbackDestination = nil
			} else {
				match++
			}
		} else {
			obj.AzureFallbackDestination = nil
		}
	} else {
		obj.AzureFallbackDestination = nil
	}

	// try to unmarshal data into GCSFallbackDestination
	err = json.Unmarshal(data, &obj.GCSFallbackDestination)
	if err == nil {
		if obj.GCSFallbackDestination != nil && obj.GCSFallbackDestination.UnparsedObject == nil {
			jsonGCSFallbackDestination, _ := json.Marshal(obj.GCSFallbackDestination)
			if string(jsonGCSFallbackDestination) == "{}" { // empty struct
				obj.GCSFallbackDestination = nil
			} else {
				match++
			}
		} else {
			obj.GCSFallbackDestination = nil
		}
	} else {
		obj.GCSFallbackDestination = nil
	}

	// try to unmarshal data into S3FallbackDestination
	err = json.Unmarshal(data, &obj.S3FallbackDestination)
	if err == nil {
		if obj.S3FallbackDestination != nil && obj.S3FallbackDestination.UnparsedObject == nil {
			jsonS3FallbackDestination, _ := json.Marshal(obj.S3FallbackDestination)
			if string(jsonS3FallbackDestination) == "{}" { // empty struct
				obj.S3FallbackDestination = nil
			} else {
				match++
			}
		} else {
			obj.S3FallbackDestination = nil
		}
	} else {
		obj.S3FallbackDestination = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AzureFallbackDestination = nil
		obj.GCSFallbackDestination = nil
		obj.S3FallbackDestination = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CustomDestinationFallbackDestination) MarshalJSON() ([]byte, error) {
	if obj.AzureFallbackDestination != nil {
		return json.Marshal(&obj.AzureFallbackDestination)
	}

	if obj.GCSFallbackDestination != nil {
		return json.Marshal(&obj.GCSFallbackDestination)
	}

	if obj.S3FallbackDestination != nil {
		return json.Marshal(&obj.S3FallbackDestination)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CustomDestinationFallbackDestination) GetActualInstance() interface{} {
	if obj.AzureFallbackDestination != nil {
		return obj.AzureFallbackDestination
	}

	if obj.GCSFallbackDestination != nil {
		return obj.GCSFallbackDestination
	}

	if obj.S3FallbackDestination != nil {
		return obj.S3FallbackDestination
	}

	// all schemas are nil
	return nil
}

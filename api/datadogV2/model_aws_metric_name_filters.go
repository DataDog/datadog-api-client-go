// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSMetricNameFilters - AWS CloudWatch metric name filter for a single namespace.
// Exactly one of `include_only` or `exclude_only` must be set.
type AWSMetricNameFilters struct {
	AWSMetricNameFiltersIncludeOnly *AWSMetricNameFiltersIncludeOnly
	AWSMetricNameFiltersExcludeOnly *AWSMetricNameFiltersExcludeOnly

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AWSMetricNameFiltersIncludeOnlyAsAWSMetricNameFilters is a convenience function that returns AWSMetricNameFiltersIncludeOnly wrapped in AWSMetricNameFilters.
func AWSMetricNameFiltersIncludeOnlyAsAWSMetricNameFilters(v *AWSMetricNameFiltersIncludeOnly) AWSMetricNameFilters {
	return AWSMetricNameFilters{AWSMetricNameFiltersIncludeOnly: v}
}

// AWSMetricNameFiltersExcludeOnlyAsAWSMetricNameFilters is a convenience function that returns AWSMetricNameFiltersExcludeOnly wrapped in AWSMetricNameFilters.
func AWSMetricNameFiltersExcludeOnlyAsAWSMetricNameFilters(v *AWSMetricNameFiltersExcludeOnly) AWSMetricNameFilters {
	return AWSMetricNameFilters{AWSMetricNameFiltersExcludeOnly: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AWSMetricNameFilters) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSMetricNameFiltersIncludeOnly
	err = datadog.Unmarshal(data, &obj.AWSMetricNameFiltersIncludeOnly)
	if err == nil {
		if obj.AWSMetricNameFiltersIncludeOnly != nil && obj.AWSMetricNameFiltersIncludeOnly.UnparsedObject == nil {
			jsonAWSMetricNameFiltersIncludeOnly, _ := datadog.Marshal(obj.AWSMetricNameFiltersIncludeOnly)
			if string(jsonAWSMetricNameFiltersIncludeOnly) == "{}" { // empty struct
				obj.AWSMetricNameFiltersIncludeOnly = nil
			} else {
				match++
			}
		} else {
			obj.AWSMetricNameFiltersIncludeOnly = nil
		}
	} else {
		obj.AWSMetricNameFiltersIncludeOnly = nil
	}

	// try to unmarshal data into AWSMetricNameFiltersExcludeOnly
	err = datadog.Unmarshal(data, &obj.AWSMetricNameFiltersExcludeOnly)
	if err == nil {
		if obj.AWSMetricNameFiltersExcludeOnly != nil && obj.AWSMetricNameFiltersExcludeOnly.UnparsedObject == nil {
			jsonAWSMetricNameFiltersExcludeOnly, _ := datadog.Marshal(obj.AWSMetricNameFiltersExcludeOnly)
			if string(jsonAWSMetricNameFiltersExcludeOnly) == "{}" { // empty struct
				obj.AWSMetricNameFiltersExcludeOnly = nil
			} else {
				match++
			}
		} else {
			obj.AWSMetricNameFiltersExcludeOnly = nil
		}
	} else {
		obj.AWSMetricNameFiltersExcludeOnly = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AWSMetricNameFiltersIncludeOnly = nil
		obj.AWSMetricNameFiltersExcludeOnly = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AWSMetricNameFilters) MarshalJSON() ([]byte, error) {
	if obj.AWSMetricNameFiltersIncludeOnly != nil {
		return datadog.Marshal(&obj.AWSMetricNameFiltersIncludeOnly)
	}

	if obj.AWSMetricNameFiltersExcludeOnly != nil {
		return datadog.Marshal(&obj.AWSMetricNameFiltersExcludeOnly)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AWSMetricNameFilters) GetActualInstance() interface{} {
	if obj.AWSMetricNameFiltersIncludeOnly != nil {
		return obj.AWSMetricNameFiltersIncludeOnly
	}

	if obj.AWSMetricNameFiltersExcludeOnly != nil {
		return obj.AWSMetricNameFiltersExcludeOnly
	}

	// all schemas are nil
	return nil
}

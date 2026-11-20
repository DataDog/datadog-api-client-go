// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BackendDistributionYAxis20270101Min - Test-only generated schema for the dashboard versioning proof of concept.
type BackendDistributionYAxis20270101Min struct {
	String  *string
	Float64 *float64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsBackendDistributionYAxis20270101Min is a convenience function that returns string wrapped in BackendDistributionYAxis20270101Min.
func StringAsBackendDistributionYAxis20270101Min(v *string) BackendDistributionYAxis20270101Min {
	return BackendDistributionYAxis20270101Min{String: v}
}

// Float64AsBackendDistributionYAxis20270101Min is a convenience function that returns float64 wrapped in BackendDistributionYAxis20270101Min.
func Float64AsBackendDistributionYAxis20270101Min(v *float64) BackendDistributionYAxis20270101Min {
	return BackendDistributionYAxis20270101Min{Float64: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *BackendDistributionYAxis20270101Min) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = datadog.Unmarshal(data, &obj.String)
	if err == nil {
		if obj.String != nil {
			jsonString, _ := datadog.Marshal(obj.String)
			if string(jsonString) == "{}" { // empty struct
				obj.String = nil
			} else {
				match++
			}
		} else {
			obj.String = nil
		}
	} else {
		obj.String = nil
	}

	// try to unmarshal data into Float64
	err = datadog.Unmarshal(data, &obj.Float64)
	if err == nil {
		if obj.Float64 != nil {
			jsonFloat64, _ := datadog.Marshal(obj.Float64)
			if string(jsonFloat64) == "{}" { // empty struct
				obj.Float64 = nil
			} else {
				match++
			}
		} else {
			obj.Float64 = nil
		}
	} else {
		obj.Float64 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.String = nil
		obj.Float64 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj BackendDistributionYAxis20270101Min) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}

	if obj.Float64 != nil {
		return datadog.Marshal(&obj.Float64)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *BackendDistributionYAxis20270101Min) GetActualInstance() interface{} {
	if obj.String != nil {
		return obj.String
	}

	if obj.Float64 != nil {
		return obj.Float64
	}

	// all schemas are nil
	return nil
}

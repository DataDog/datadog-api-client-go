// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendRequest20270101OneOfQuery - Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendRequest20270101OneOfQuery struct {
	StructuredAnalysisQuery_20270101 *StructuredAnalysisQuery_20270101
	SampledAnalysisQuery_20270101    *SampledAnalysisQuery_20270101

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StructuredAnalysisQuery_20270101AsToplistBackendRequest20270101OneOfQuery is a convenience function that returns StructuredAnalysisQuery_20270101 wrapped in ToplistBackendRequest20270101OneOfQuery.
func StructuredAnalysisQuery_20270101AsToplistBackendRequest20270101OneOfQuery(v *StructuredAnalysisQuery_20270101) ToplistBackendRequest20270101OneOfQuery {
	return ToplistBackendRequest20270101OneOfQuery{StructuredAnalysisQuery_20270101: v}
}

// SampledAnalysisQuery_20270101AsToplistBackendRequest20270101OneOfQuery is a convenience function that returns SampledAnalysisQuery_20270101 wrapped in ToplistBackendRequest20270101OneOfQuery.
func SampledAnalysisQuery_20270101AsToplistBackendRequest20270101OneOfQuery(v *SampledAnalysisQuery_20270101) ToplistBackendRequest20270101OneOfQuery {
	return ToplistBackendRequest20270101OneOfQuery{SampledAnalysisQuery_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ToplistBackendRequest20270101OneOfQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into StructuredAnalysisQuery_20270101
	err = datadog.Unmarshal(data, &obj.StructuredAnalysisQuery_20270101)
	if err == nil {
		if obj.StructuredAnalysisQuery_20270101 != nil && obj.StructuredAnalysisQuery_20270101.UnparsedObject == nil {
			jsonStructuredAnalysisQuery_20270101, _ := datadog.Marshal(obj.StructuredAnalysisQuery_20270101)
			if string(jsonStructuredAnalysisQuery_20270101) == "{}" { // empty struct
				obj.StructuredAnalysisQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.StructuredAnalysisQuery_20270101 = nil
		}
	} else {
		obj.StructuredAnalysisQuery_20270101 = nil
	}

	// try to unmarshal data into SampledAnalysisQuery_20270101
	err = datadog.Unmarshal(data, &obj.SampledAnalysisQuery_20270101)
	if err == nil {
		if obj.SampledAnalysisQuery_20270101 != nil && obj.SampledAnalysisQuery_20270101.UnparsedObject == nil {
			jsonSampledAnalysisQuery_20270101, _ := datadog.Marshal(obj.SampledAnalysisQuery_20270101)
			if string(jsonSampledAnalysisQuery_20270101) == "{}" { // empty struct
				obj.SampledAnalysisQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.SampledAnalysisQuery_20270101 = nil
		}
	} else {
		obj.SampledAnalysisQuery_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.StructuredAnalysisQuery_20270101 = nil
		obj.SampledAnalysisQuery_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ToplistBackendRequest20270101OneOfQuery) MarshalJSON() ([]byte, error) {
	if obj.StructuredAnalysisQuery_20270101 != nil {
		return datadog.Marshal(&obj.StructuredAnalysisQuery_20270101)
	}

	if obj.SampledAnalysisQuery_20270101 != nil {
		return datadog.Marshal(&obj.SampledAnalysisQuery_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ToplistBackendRequest20270101OneOfQuery) GetActualInstance() interface{} {
	if obj.StructuredAnalysisQuery_20270101 != nil {
		return obj.StructuredAnalysisQuery_20270101
	}

	if obj.SampledAnalysisQuery_20270101 != nil {
		return obj.SampledAnalysisQuery_20270101
	}

	// all schemas are nil
	return nil
}

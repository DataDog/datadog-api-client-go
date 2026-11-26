// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisDataset20270101Query - Test-only generated schema for the dashboard versioning proof of concept.
type AnalysisDataset20270101Query struct {
	SqlAnalysisQuery_20270101        *SqlAnalysisQuery_20270101
	StructuredAnalysisQuery_20270101 *StructuredAnalysisQuery_20270101
	SampledAnalysisQuery_20270101    *SampledAnalysisQuery_20270101

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SqlAnalysisQuery_20270101AsAnalysisDataset20270101Query is a convenience function that returns SqlAnalysisQuery_20270101 wrapped in AnalysisDataset20270101Query.
func SqlAnalysisQuery_20270101AsAnalysisDataset20270101Query(v *SqlAnalysisQuery_20270101) AnalysisDataset20270101Query {
	return AnalysisDataset20270101Query{SqlAnalysisQuery_20270101: v}
}

// StructuredAnalysisQuery_20270101AsAnalysisDataset20270101Query is a convenience function that returns StructuredAnalysisQuery_20270101 wrapped in AnalysisDataset20270101Query.
func StructuredAnalysisQuery_20270101AsAnalysisDataset20270101Query(v *StructuredAnalysisQuery_20270101) AnalysisDataset20270101Query {
	return AnalysisDataset20270101Query{StructuredAnalysisQuery_20270101: v}
}

// SampledAnalysisQuery_20270101AsAnalysisDataset20270101Query is a convenience function that returns SampledAnalysisQuery_20270101 wrapped in AnalysisDataset20270101Query.
func SampledAnalysisQuery_20270101AsAnalysisDataset20270101Query(v *SampledAnalysisQuery_20270101) AnalysisDataset20270101Query {
	return AnalysisDataset20270101Query{SampledAnalysisQuery_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AnalysisDataset20270101Query) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SqlAnalysisQuery_20270101
	err = datadog.Unmarshal(data, &obj.SqlAnalysisQuery_20270101)
	if err == nil {
		if obj.SqlAnalysisQuery_20270101 != nil && obj.SqlAnalysisQuery_20270101.UnparsedObject == nil {
			jsonSqlAnalysisQuery_20270101, _ := datadog.Marshal(obj.SqlAnalysisQuery_20270101)
			if string(jsonSqlAnalysisQuery_20270101) == "{}" { // empty struct
				obj.SqlAnalysisQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.SqlAnalysisQuery_20270101 = nil
		}
	} else {
		obj.SqlAnalysisQuery_20270101 = nil
	}

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
		obj.SqlAnalysisQuery_20270101 = nil
		obj.StructuredAnalysisQuery_20270101 = nil
		obj.SampledAnalysisQuery_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AnalysisDataset20270101Query) MarshalJSON() ([]byte, error) {
	if obj.SqlAnalysisQuery_20270101 != nil {
		return datadog.Marshal(&obj.SqlAnalysisQuery_20270101)
	}

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
func (obj *AnalysisDataset20270101Query) GetActualInstance() interface{} {
	if obj.SqlAnalysisQuery_20270101 != nil {
		return obj.SqlAnalysisQuery_20270101
	}

	if obj.StructuredAnalysisQuery_20270101 != nil {
		return obj.StructuredAnalysisQuery_20270101
	}

	if obj.SampledAnalysisQuery_20270101 != nil {
		return obj.SampledAnalysisQuery_20270101
	}

	// all schemas are nil
	return nil
}

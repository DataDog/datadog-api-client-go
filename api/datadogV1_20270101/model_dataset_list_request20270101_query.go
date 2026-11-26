// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetListRequest20270101Query - Test-only generated schema for the dashboard versioning proof of concept.
type DatasetListRequest20270101Query struct {
	DatasetListQuery_20270101    *DatasetListQuery
	RelationalDataQuery_20270101 *RelationalDataQuery

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DatasetListQuery_20270101AsDatasetListRequest20270101Query is a convenience function that returns DatasetListQuery wrapped in DatasetListRequest20270101Query.
func DatasetListQuery_20270101AsDatasetListRequest20270101Query(v *DatasetListQuery) DatasetListRequest20270101Query {
	return DatasetListRequest20270101Query{DatasetListQuery_20270101: v}
}

// RelationalDataQuery_20270101AsDatasetListRequest20270101Query is a convenience function that returns RelationalDataQuery wrapped in DatasetListRequest20270101Query.
func RelationalDataQuery_20270101AsDatasetListRequest20270101Query(v *RelationalDataQuery) DatasetListRequest20270101Query {
	return DatasetListRequest20270101Query{RelationalDataQuery_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DatasetListRequest20270101Query) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DatasetListQuery_20270101
	err = datadog.Unmarshal(data, &obj.DatasetListQuery_20270101)
	if err == nil {
		if obj.DatasetListQuery_20270101 != nil && obj.DatasetListQuery_20270101.UnparsedObject == nil {
			jsonDatasetListQuery_20270101, _ := datadog.Marshal(obj.DatasetListQuery_20270101)
			if string(jsonDatasetListQuery_20270101) == "{}" { // empty struct
				obj.DatasetListQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DatasetListQuery_20270101 = nil
		}
	} else {
		obj.DatasetListQuery_20270101 = nil
	}

	// try to unmarshal data into RelationalDataQuery_20270101
	err = datadog.Unmarshal(data, &obj.RelationalDataQuery_20270101)
	if err == nil {
		if obj.RelationalDataQuery_20270101 != nil && obj.RelationalDataQuery_20270101.UnparsedObject == nil {
			jsonRelationalDataQuery_20270101, _ := datadog.Marshal(obj.RelationalDataQuery_20270101)
			if string(jsonRelationalDataQuery_20270101) == "{}" { // empty struct
				obj.RelationalDataQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.RelationalDataQuery_20270101 = nil
		}
	} else {
		obj.RelationalDataQuery_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DatasetListQuery_20270101 = nil
		obj.RelationalDataQuery_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DatasetListRequest20270101Query) MarshalJSON() ([]byte, error) {
	if obj.DatasetListQuery_20270101 != nil {
		return datadog.Marshal(&obj.DatasetListQuery_20270101)
	}

	if obj.RelationalDataQuery_20270101 != nil {
		return datadog.Marshal(&obj.RelationalDataQuery_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DatasetListRequest20270101Query) GetActualInstance() interface{} {
	if obj.DatasetListQuery_20270101 != nil {
		return obj.DatasetListQuery_20270101
	}

	if obj.RelationalDataQuery_20270101 != nil {
		return obj.RelationalDataQuery_20270101
	}

	// all schemas are nil
	return nil
}

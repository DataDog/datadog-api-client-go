// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataProjectionQuery - Test-only generated schema for the dashboard versioning proof of concept.
type DataProjectionQuery struct {
	AggregateAugmentedQuery_20270101                     *AggregateAugmentedQuery
	EventsPlatformDataProjectionQueryDefinition_20270101 *EventsPlatformDataProjectionQueryDefinition
	DatasetListQuery_20270101                            *DatasetListQuery

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AggregateAugmentedQuery_20270101AsDataProjectionQuery is a convenience function that returns AggregateAugmentedQuery wrapped in DataProjectionQuery.
func AggregateAugmentedQuery_20270101AsDataProjectionQuery(v *AggregateAugmentedQuery) DataProjectionQuery {
	return DataProjectionQuery{AggregateAugmentedQuery_20270101: v}
}

// EventsPlatformDataProjectionQueryDefinition_20270101AsDataProjectionQuery is a convenience function that returns EventsPlatformDataProjectionQueryDefinition wrapped in DataProjectionQuery.
func EventsPlatformDataProjectionQueryDefinition_20270101AsDataProjectionQuery(v *EventsPlatformDataProjectionQueryDefinition) DataProjectionQuery {
	return DataProjectionQuery{EventsPlatformDataProjectionQueryDefinition_20270101: v}
}

// DatasetListQuery_20270101AsDataProjectionQuery is a convenience function that returns DatasetListQuery wrapped in DataProjectionQuery.
func DatasetListQuery_20270101AsDataProjectionQuery(v *DatasetListQuery) DataProjectionQuery {
	return DataProjectionQuery{DatasetListQuery_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DataProjectionQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AggregateAugmentedQuery_20270101
	err = datadog.Unmarshal(data, &obj.AggregateAugmentedQuery_20270101)
	if err == nil {
		if obj.AggregateAugmentedQuery_20270101 != nil && obj.AggregateAugmentedQuery_20270101.UnparsedObject == nil {
			jsonAggregateAugmentedQuery_20270101, _ := datadog.Marshal(obj.AggregateAugmentedQuery_20270101)
			if string(jsonAggregateAugmentedQuery_20270101) == "{}" { // empty struct
				obj.AggregateAugmentedQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.AggregateAugmentedQuery_20270101 = nil
		}
	} else {
		obj.AggregateAugmentedQuery_20270101 = nil
	}

	// try to unmarshal data into EventsPlatformDataProjectionQueryDefinition_20270101
	err = datadog.Unmarshal(data, &obj.EventsPlatformDataProjectionQueryDefinition_20270101)
	if err == nil {
		if obj.EventsPlatformDataProjectionQueryDefinition_20270101 != nil && obj.EventsPlatformDataProjectionQueryDefinition_20270101.UnparsedObject == nil {
			jsonEventsPlatformDataProjectionQueryDefinition_20270101, _ := datadog.Marshal(obj.EventsPlatformDataProjectionQueryDefinition_20270101)
			if string(jsonEventsPlatformDataProjectionQueryDefinition_20270101) == "{}" { // empty struct
				obj.EventsPlatformDataProjectionQueryDefinition_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.EventsPlatformDataProjectionQueryDefinition_20270101 = nil
		}
	} else {
		obj.EventsPlatformDataProjectionQueryDefinition_20270101 = nil
	}

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

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AggregateAugmentedQuery_20270101 = nil
		obj.EventsPlatformDataProjectionQueryDefinition_20270101 = nil
		obj.DatasetListQuery_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DataProjectionQuery) MarshalJSON() ([]byte, error) {
	if obj.AggregateAugmentedQuery_20270101 != nil {
		return datadog.Marshal(&obj.AggregateAugmentedQuery_20270101)
	}

	if obj.EventsPlatformDataProjectionQueryDefinition_20270101 != nil {
		return datadog.Marshal(&obj.EventsPlatformDataProjectionQueryDefinition_20270101)
	}

	if obj.DatasetListQuery_20270101 != nil {
		return datadog.Marshal(&obj.DatasetListQuery_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DataProjectionQuery) GetActualInstance() interface{} {
	if obj.AggregateAugmentedQuery_20270101 != nil {
		return obj.AggregateAugmentedQuery_20270101
	}

	if obj.EventsPlatformDataProjectionQueryDefinition_20270101 != nil {
		return obj.EventsPlatformDataProjectionQueryDefinition_20270101
	}

	if obj.DatasetListQuery_20270101 != nil {
		return obj.DatasetListQuery_20270101
	}

	// all schemas are nil
	return nil
}

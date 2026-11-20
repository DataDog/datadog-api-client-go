// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetCompute20270101Aggregation - Test-only generated schema for the dashboard versioning proof of concept.
type DatasetCompute20270101Aggregation struct {
	DashboardInlineSchema11_20270101 *DashboardInlineSchema11_20270101
	EventsAggregator_20270101        *EventsAggregator_20270101

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DashboardInlineSchema11_20270101AsDatasetCompute20270101Aggregation is a convenience function that returns DashboardInlineSchema11_20270101 wrapped in DatasetCompute20270101Aggregation.
func DashboardInlineSchema11_20270101AsDatasetCompute20270101Aggregation(v *DashboardInlineSchema11_20270101) DatasetCompute20270101Aggregation {
	return DatasetCompute20270101Aggregation{DashboardInlineSchema11_20270101: v}
}

// EventsAggregator_20270101AsDatasetCompute20270101Aggregation is a convenience function that returns EventsAggregator_20270101 wrapped in DatasetCompute20270101Aggregation.
func EventsAggregator_20270101AsDatasetCompute20270101Aggregation(v *EventsAggregator_20270101) DatasetCompute20270101Aggregation {
	return DatasetCompute20270101Aggregation{EventsAggregator_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DatasetCompute20270101Aggregation) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DashboardInlineSchema11_20270101
	err = datadog.Unmarshal(data, &obj.DashboardInlineSchema11_20270101)
	if err == nil {
		if obj.DashboardInlineSchema11_20270101 != nil {
			jsonDashboardInlineSchema11_20270101, _ := datadog.Marshal(obj.DashboardInlineSchema11_20270101)
			if string(jsonDashboardInlineSchema11_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.DashboardInlineSchema11_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DashboardInlineSchema11_20270101 = nil
		}
	} else {
		obj.DashboardInlineSchema11_20270101 = nil
	}

	// try to unmarshal data into EventsAggregator_20270101
	err = datadog.Unmarshal(data, &obj.EventsAggregator_20270101)
	if err == nil {
		if obj.EventsAggregator_20270101 != nil && obj.EventsAggregator_20270101.UnparsedObject == nil {
			jsonEventsAggregator_20270101, _ := datadog.Marshal(obj.EventsAggregator_20270101)
			if string(jsonEventsAggregator_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.EventsAggregator_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.EventsAggregator_20270101 = nil
		}
	} else {
		obj.EventsAggregator_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DashboardInlineSchema11_20270101 = nil
		obj.EventsAggregator_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DatasetCompute20270101Aggregation) MarshalJSON() ([]byte, error) {
	if obj.DashboardInlineSchema11_20270101 != nil {
		return datadog.Marshal(&obj.DashboardInlineSchema11_20270101)
	}

	if obj.EventsAggregator_20270101 != nil {
		return datadog.Marshal(&obj.EventsAggregator_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DatasetCompute20270101Aggregation) GetActualInstance() interface{} {
	if obj.DashboardInlineSchema11_20270101 != nil {
		return obj.DashboardInlineSchema11_20270101
	}

	if obj.EventsAggregator_20270101 != nil {
		return obj.EventsAggregator_20270101
	}

	// all schemas are nil
	return nil
}

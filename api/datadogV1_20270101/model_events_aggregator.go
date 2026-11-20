// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsAggregator - Test-only generated schema for the dashboard versioning proof of concept.
type EventsAggregator struct {
	DashboardInlineSchema39_20270101   *DashboardInlineSchema39
	PercentileAggregationType_20270101 *string

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DashboardInlineSchema39_20270101AsEventsAggregator is a convenience function that returns DashboardInlineSchema39 wrapped in EventsAggregator.
func DashboardInlineSchema39_20270101AsEventsAggregator(v *DashboardInlineSchema39) EventsAggregator {
	return EventsAggregator{DashboardInlineSchema39_20270101: v}
}

// PercentileAggregationType_20270101AsEventsAggregator is a convenience function that returns string wrapped in EventsAggregator.
func PercentileAggregationType_20270101AsEventsAggregator(v *string) EventsAggregator {
	return EventsAggregator{PercentileAggregationType_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *EventsAggregator) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DashboardInlineSchema39_20270101
	err = datadog.Unmarshal(data, &obj.DashboardInlineSchema39_20270101)
	if err == nil {
		if obj.DashboardInlineSchema39_20270101 != nil {
			jsonDashboardInlineSchema39_20270101, _ := datadog.Marshal(obj.DashboardInlineSchema39_20270101)
			if string(jsonDashboardInlineSchema39_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.DashboardInlineSchema39_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DashboardInlineSchema39_20270101 = nil
		}
	} else {
		obj.DashboardInlineSchema39_20270101 = nil
	}

	// try to unmarshal data into PercentileAggregationType_20270101
	err = datadog.Unmarshal(data, &obj.PercentileAggregationType_20270101)
	if err == nil {
		if obj.PercentileAggregationType_20270101 != nil {
			jsonPercentileAggregationType_20270101, _ := datadog.Marshal(obj.PercentileAggregationType_20270101)
			if string(jsonPercentileAggregationType_20270101) == "{}" { // empty struct
				obj.PercentileAggregationType_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.PercentileAggregationType_20270101 = nil
		}
	} else {
		obj.PercentileAggregationType_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DashboardInlineSchema39_20270101 = nil
		obj.PercentileAggregationType_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj EventsAggregator) MarshalJSON() ([]byte, error) {
	if obj.DashboardInlineSchema39_20270101 != nil {
		return datadog.Marshal(&obj.DashboardInlineSchema39_20270101)
	}

	if obj.PercentileAggregationType_20270101 != nil {
		return datadog.Marshal(&obj.PercentileAggregationType_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *EventsAggregator) GetActualInstance() interface{} {
	if obj.DashboardInlineSchema39_20270101 != nil {
		return obj.DashboardInlineSchema39_20270101
	}

	if obj.PercentileAggregationType_20270101 != nil {
		return obj.PercentileAggregationType_20270101
	}

	// all schemas are nil
	return nil
}

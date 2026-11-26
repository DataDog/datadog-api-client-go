// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HistogramRequest - Test-only generated schema for the dashboard versioning proof of concept.
type HistogramRequest struct {
	MetricsHistogramRequest_20270101    *MetricsHistogramRequest
	ApmStatsHistogramRequest_20270101   *ApmStatsHistogramRequest
	ApmMetricsHistogramRequest_20270101 *ApmMetricsHistogramRequest
	EventsHistogramRequest_20270101     *EventsHistogramRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MetricsHistogramRequest_20270101AsHistogramRequest is a convenience function that returns MetricsHistogramRequest wrapped in HistogramRequest.
func MetricsHistogramRequest_20270101AsHistogramRequest(v *MetricsHistogramRequest) HistogramRequest {
	return HistogramRequest{MetricsHistogramRequest_20270101: v}
}

// ApmStatsHistogramRequest_20270101AsHistogramRequest is a convenience function that returns ApmStatsHistogramRequest wrapped in HistogramRequest.
func ApmStatsHistogramRequest_20270101AsHistogramRequest(v *ApmStatsHistogramRequest) HistogramRequest {
	return HistogramRequest{ApmStatsHistogramRequest_20270101: v}
}

// ApmMetricsHistogramRequest_20270101AsHistogramRequest is a convenience function that returns ApmMetricsHistogramRequest wrapped in HistogramRequest.
func ApmMetricsHistogramRequest_20270101AsHistogramRequest(v *ApmMetricsHistogramRequest) HistogramRequest {
	return HistogramRequest{ApmMetricsHistogramRequest_20270101: v}
}

// EventsHistogramRequest_20270101AsHistogramRequest is a convenience function that returns EventsHistogramRequest wrapped in HistogramRequest.
func EventsHistogramRequest_20270101AsHistogramRequest(v *EventsHistogramRequest) HistogramRequest {
	return HistogramRequest{EventsHistogramRequest_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *HistogramRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MetricsHistogramRequest_20270101
	err = datadog.Unmarshal(data, &obj.MetricsHistogramRequest_20270101)
	if err == nil {
		if obj.MetricsHistogramRequest_20270101 != nil && obj.MetricsHistogramRequest_20270101.UnparsedObject == nil {
			jsonMetricsHistogramRequest_20270101, _ := datadog.Marshal(obj.MetricsHistogramRequest_20270101)
			if string(jsonMetricsHistogramRequest_20270101) == "{}" { // empty struct
				obj.MetricsHistogramRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.MetricsHistogramRequest_20270101 = nil
		}
	} else {
		obj.MetricsHistogramRequest_20270101 = nil
	}

	// try to unmarshal data into ApmStatsHistogramRequest_20270101
	err = datadog.Unmarshal(data, &obj.ApmStatsHistogramRequest_20270101)
	if err == nil {
		if obj.ApmStatsHistogramRequest_20270101 != nil && obj.ApmStatsHistogramRequest_20270101.UnparsedObject == nil {
			jsonApmStatsHistogramRequest_20270101, _ := datadog.Marshal(obj.ApmStatsHistogramRequest_20270101)
			if string(jsonApmStatsHistogramRequest_20270101) == "{}" { // empty struct
				obj.ApmStatsHistogramRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ApmStatsHistogramRequest_20270101 = nil
		}
	} else {
		obj.ApmStatsHistogramRequest_20270101 = nil
	}

	// try to unmarshal data into ApmMetricsHistogramRequest_20270101
	err = datadog.Unmarshal(data, &obj.ApmMetricsHistogramRequest_20270101)
	if err == nil {
		if obj.ApmMetricsHistogramRequest_20270101 != nil && obj.ApmMetricsHistogramRequest_20270101.UnparsedObject == nil {
			jsonApmMetricsHistogramRequest_20270101, _ := datadog.Marshal(obj.ApmMetricsHistogramRequest_20270101)
			if string(jsonApmMetricsHistogramRequest_20270101) == "{}" { // empty struct
				obj.ApmMetricsHistogramRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ApmMetricsHistogramRequest_20270101 = nil
		}
	} else {
		obj.ApmMetricsHistogramRequest_20270101 = nil
	}

	// try to unmarshal data into EventsHistogramRequest_20270101
	err = datadog.Unmarshal(data, &obj.EventsHistogramRequest_20270101)
	if err == nil {
		if obj.EventsHistogramRequest_20270101 != nil && obj.EventsHistogramRequest_20270101.UnparsedObject == nil {
			jsonEventsHistogramRequest_20270101, _ := datadog.Marshal(obj.EventsHistogramRequest_20270101)
			if string(jsonEventsHistogramRequest_20270101) == "{}" { // empty struct
				obj.EventsHistogramRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.EventsHistogramRequest_20270101 = nil
		}
	} else {
		obj.EventsHistogramRequest_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MetricsHistogramRequest_20270101 = nil
		obj.ApmStatsHistogramRequest_20270101 = nil
		obj.ApmMetricsHistogramRequest_20270101 = nil
		obj.EventsHistogramRequest_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj HistogramRequest) MarshalJSON() ([]byte, error) {
	if obj.MetricsHistogramRequest_20270101 != nil {
		return datadog.Marshal(&obj.MetricsHistogramRequest_20270101)
	}

	if obj.ApmStatsHistogramRequest_20270101 != nil {
		return datadog.Marshal(&obj.ApmStatsHistogramRequest_20270101)
	}

	if obj.ApmMetricsHistogramRequest_20270101 != nil {
		return datadog.Marshal(&obj.ApmMetricsHistogramRequest_20270101)
	}

	if obj.EventsHistogramRequest_20270101 != nil {
		return datadog.Marshal(&obj.EventsHistogramRequest_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *HistogramRequest) GetActualInstance() interface{} {
	if obj.MetricsHistogramRequest_20270101 != nil {
		return obj.MetricsHistogramRequest_20270101
	}

	if obj.ApmStatsHistogramRequest_20270101 != nil {
		return obj.ApmStatsHistogramRequest_20270101
	}

	if obj.ApmMetricsHistogramRequest_20270101 != nil {
		return obj.ApmMetricsHistogramRequest_20270101
	}

	if obj.EventsHistogramRequest_20270101 != nil {
		return obj.EventsHistogramRequest_20270101
	}

	// all schemas are nil
	return nil
}

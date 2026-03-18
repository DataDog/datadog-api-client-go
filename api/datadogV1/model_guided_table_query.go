// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableQuery - A base query used as source data for a guided table widget. Either a metrics query or an events-platform query.
type GuidedTableQuery struct {
	GuidedTableMetricsQuery *GuidedTableMetricsQuery
	GuidedTableEventsQuery  *GuidedTableEventsQuery

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GuidedTableMetricsQueryAsGuidedTableQuery is a convenience function that returns GuidedTableMetricsQuery wrapped in GuidedTableQuery.
func GuidedTableMetricsQueryAsGuidedTableQuery(v *GuidedTableMetricsQuery) GuidedTableQuery {
	return GuidedTableQuery{GuidedTableMetricsQuery: v}
}

// GuidedTableEventsQueryAsGuidedTableQuery is a convenience function that returns GuidedTableEventsQuery wrapped in GuidedTableQuery.
func GuidedTableEventsQueryAsGuidedTableQuery(v *GuidedTableEventsQuery) GuidedTableQuery {
	return GuidedTableQuery{GuidedTableEventsQuery: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GuidedTableQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GuidedTableMetricsQuery
	err = datadog.Unmarshal(data, &obj.GuidedTableMetricsQuery)
	if err == nil {
		if obj.GuidedTableMetricsQuery != nil && obj.GuidedTableMetricsQuery.UnparsedObject == nil {
			jsonGuidedTableMetricsQuery, _ := datadog.Marshal(obj.GuidedTableMetricsQuery)
			if string(jsonGuidedTableMetricsQuery) == "{}" { // empty struct
				obj.GuidedTableMetricsQuery = nil
			} else {
				match++
			}
		} else {
			obj.GuidedTableMetricsQuery = nil
		}
	} else {
		obj.GuidedTableMetricsQuery = nil
	}

	// try to unmarshal data into GuidedTableEventsQuery
	err = datadog.Unmarshal(data, &obj.GuidedTableEventsQuery)
	if err == nil {
		if obj.GuidedTableEventsQuery != nil && obj.GuidedTableEventsQuery.UnparsedObject == nil {
			jsonGuidedTableEventsQuery, _ := datadog.Marshal(obj.GuidedTableEventsQuery)
			if string(jsonGuidedTableEventsQuery) == "{}" { // empty struct
				obj.GuidedTableEventsQuery = nil
			} else {
				match++
			}
		} else {
			obj.GuidedTableEventsQuery = nil
		}
	} else {
		obj.GuidedTableEventsQuery = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GuidedTableMetricsQuery = nil
		obj.GuidedTableEventsQuery = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GuidedTableQuery) MarshalJSON() ([]byte, error) {
	if obj.GuidedTableMetricsQuery != nil {
		return datadog.Marshal(&obj.GuidedTableMetricsQuery)
	}

	if obj.GuidedTableEventsQuery != nil {
		return datadog.Marshal(&obj.GuidedTableEventsQuery)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GuidedTableQuery) GetActualInstance() interface{} {
	if obj.GuidedTableMetricsQuery != nil {
		return obj.GuidedTableMetricsQuery
	}

	if obj.GuidedTableEventsQuery != nil {
		return obj.GuidedTableEventsQuery
	}

	// all schemas are nil
	return nil
}

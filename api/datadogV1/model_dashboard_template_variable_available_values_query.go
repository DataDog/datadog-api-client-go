// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardTemplateVariableAvailableValuesQuery - A query that dynamically computes the list of values available for this template variable.
type DashboardTemplateVariableAvailableValuesQuery struct {
	DashboardAvailableValuesEventsQuery  *DashboardAvailableValuesEventsQuery
	DashboardAvailableValuesMetricsQuery *DashboardAvailableValuesMetricsQuery

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DashboardAvailableValuesEventsQueryAsDashboardTemplateVariableAvailableValuesQuery is a convenience function that returns DashboardAvailableValuesEventsQuery wrapped in DashboardTemplateVariableAvailableValuesQuery.
func DashboardAvailableValuesEventsQueryAsDashboardTemplateVariableAvailableValuesQuery(v *DashboardAvailableValuesEventsQuery) DashboardTemplateVariableAvailableValuesQuery {
	return DashboardTemplateVariableAvailableValuesQuery{DashboardAvailableValuesEventsQuery: v}
}

// DashboardAvailableValuesMetricsQueryAsDashboardTemplateVariableAvailableValuesQuery is a convenience function that returns DashboardAvailableValuesMetricsQuery wrapped in DashboardTemplateVariableAvailableValuesQuery.
func DashboardAvailableValuesMetricsQueryAsDashboardTemplateVariableAvailableValuesQuery(v *DashboardAvailableValuesMetricsQuery) DashboardTemplateVariableAvailableValuesQuery {
	return DashboardTemplateVariableAvailableValuesQuery{DashboardAvailableValuesMetricsQuery: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DashboardTemplateVariableAvailableValuesQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DashboardAvailableValuesEventsQuery
	err = datadog.Unmarshal(data, &obj.DashboardAvailableValuesEventsQuery)
	if err == nil {
		if obj.DashboardAvailableValuesEventsQuery != nil && obj.DashboardAvailableValuesEventsQuery.UnparsedObject == nil {
			jsonDashboardAvailableValuesEventsQuery, _ := datadog.Marshal(obj.DashboardAvailableValuesEventsQuery)
			if string(jsonDashboardAvailableValuesEventsQuery) == "{}" { // empty struct
				obj.DashboardAvailableValuesEventsQuery = nil
			} else {
				match++
			}
		} else {
			obj.DashboardAvailableValuesEventsQuery = nil
		}
	} else {
		obj.DashboardAvailableValuesEventsQuery = nil
	}

	// try to unmarshal data into DashboardAvailableValuesMetricsQuery
	err = datadog.Unmarshal(data, &obj.DashboardAvailableValuesMetricsQuery)
	if err == nil {
		if obj.DashboardAvailableValuesMetricsQuery != nil && obj.DashboardAvailableValuesMetricsQuery.UnparsedObject == nil {
			jsonDashboardAvailableValuesMetricsQuery, _ := datadog.Marshal(obj.DashboardAvailableValuesMetricsQuery)
			if string(jsonDashboardAvailableValuesMetricsQuery) == "{}" { // empty struct
				obj.DashboardAvailableValuesMetricsQuery = nil
			} else {
				match++
			}
		} else {
			obj.DashboardAvailableValuesMetricsQuery = nil
		}
	} else {
		obj.DashboardAvailableValuesMetricsQuery = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DashboardAvailableValuesEventsQuery = nil
		obj.DashboardAvailableValuesMetricsQuery = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DashboardTemplateVariableAvailableValuesQuery) MarshalJSON() ([]byte, error) {
	if obj.DashboardAvailableValuesEventsQuery != nil {
		return datadog.Marshal(&obj.DashboardAvailableValuesEventsQuery)
	}

	if obj.DashboardAvailableValuesMetricsQuery != nil {
		return datadog.Marshal(&obj.DashboardAvailableValuesMetricsQuery)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DashboardTemplateVariableAvailableValuesQuery) GetActualInstance() interface{} {
	if obj.DashboardAvailableValuesEventsQuery != nil {
		return obj.DashboardAvailableValuesEventsQuery
	}

	if obj.DashboardAvailableValuesMetricsQuery != nil {
		return obj.DashboardAvailableValuesMetricsQuery
	}

	// all schemas are nil
	return nil
}

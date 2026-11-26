// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionSearchCohortCriteriaTimeInterval20270101Value - Test-only generated schema for the dashboard versioning proof of concept.
type RetentionSearchCohortCriteriaTimeInterval20270101Value struct {
	CalendarIntervalDay_20270101   *CalendarIntervalDay_20270101
	CalendarIntervalWeek_20270101  *CalendarIntervalWeek_20270101
	CalendarIntervalMonth_20270101 *CalendarIntervalMonth_20270101

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CalendarIntervalDay_20270101AsRetentionSearchCohortCriteriaTimeInterval20270101Value is a convenience function that returns CalendarIntervalDay_20270101 wrapped in RetentionSearchCohortCriteriaTimeInterval20270101Value.
func CalendarIntervalDay_20270101AsRetentionSearchCohortCriteriaTimeInterval20270101Value(v *CalendarIntervalDay_20270101) RetentionSearchCohortCriteriaTimeInterval20270101Value {
	return RetentionSearchCohortCriteriaTimeInterval20270101Value{CalendarIntervalDay_20270101: v}
}

// CalendarIntervalWeek_20270101AsRetentionSearchCohortCriteriaTimeInterval20270101Value is a convenience function that returns CalendarIntervalWeek_20270101 wrapped in RetentionSearchCohortCriteriaTimeInterval20270101Value.
func CalendarIntervalWeek_20270101AsRetentionSearchCohortCriteriaTimeInterval20270101Value(v *CalendarIntervalWeek_20270101) RetentionSearchCohortCriteriaTimeInterval20270101Value {
	return RetentionSearchCohortCriteriaTimeInterval20270101Value{CalendarIntervalWeek_20270101: v}
}

// CalendarIntervalMonth_20270101AsRetentionSearchCohortCriteriaTimeInterval20270101Value is a convenience function that returns CalendarIntervalMonth_20270101 wrapped in RetentionSearchCohortCriteriaTimeInterval20270101Value.
func CalendarIntervalMonth_20270101AsRetentionSearchCohortCriteriaTimeInterval20270101Value(v *CalendarIntervalMonth_20270101) RetentionSearchCohortCriteriaTimeInterval20270101Value {
	return RetentionSearchCohortCriteriaTimeInterval20270101Value{CalendarIntervalMonth_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *RetentionSearchCohortCriteriaTimeInterval20270101Value) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CalendarIntervalDay_20270101
	err = datadog.Unmarshal(data, &obj.CalendarIntervalDay_20270101)
	if err == nil {
		if obj.CalendarIntervalDay_20270101 != nil && obj.CalendarIntervalDay_20270101.UnparsedObject == nil {
			jsonCalendarIntervalDay_20270101, _ := datadog.Marshal(obj.CalendarIntervalDay_20270101)
			if string(jsonCalendarIntervalDay_20270101) == "{}" { // empty struct
				obj.CalendarIntervalDay_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CalendarIntervalDay_20270101 = nil
		}
	} else {
		obj.CalendarIntervalDay_20270101 = nil
	}

	// try to unmarshal data into CalendarIntervalWeek_20270101
	err = datadog.Unmarshal(data, &obj.CalendarIntervalWeek_20270101)
	if err == nil {
		if obj.CalendarIntervalWeek_20270101 != nil && obj.CalendarIntervalWeek_20270101.UnparsedObject == nil {
			jsonCalendarIntervalWeek_20270101, _ := datadog.Marshal(obj.CalendarIntervalWeek_20270101)
			if string(jsonCalendarIntervalWeek_20270101) == "{}" { // empty struct
				obj.CalendarIntervalWeek_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CalendarIntervalWeek_20270101 = nil
		}
	} else {
		obj.CalendarIntervalWeek_20270101 = nil
	}

	// try to unmarshal data into CalendarIntervalMonth_20270101
	err = datadog.Unmarshal(data, &obj.CalendarIntervalMonth_20270101)
	if err == nil {
		if obj.CalendarIntervalMonth_20270101 != nil && obj.CalendarIntervalMonth_20270101.UnparsedObject == nil {
			jsonCalendarIntervalMonth_20270101, _ := datadog.Marshal(obj.CalendarIntervalMonth_20270101)
			if string(jsonCalendarIntervalMonth_20270101) == "{}" { // empty struct
				obj.CalendarIntervalMonth_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CalendarIntervalMonth_20270101 = nil
		}
	} else {
		obj.CalendarIntervalMonth_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CalendarIntervalDay_20270101 = nil
		obj.CalendarIntervalWeek_20270101 = nil
		obj.CalendarIntervalMonth_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj RetentionSearchCohortCriteriaTimeInterval20270101Value) MarshalJSON() ([]byte, error) {
	if obj.CalendarIntervalDay_20270101 != nil {
		return datadog.Marshal(&obj.CalendarIntervalDay_20270101)
	}

	if obj.CalendarIntervalWeek_20270101 != nil {
		return datadog.Marshal(&obj.CalendarIntervalWeek_20270101)
	}

	if obj.CalendarIntervalMonth_20270101 != nil {
		return datadog.Marshal(&obj.CalendarIntervalMonth_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *RetentionSearchCohortCriteriaTimeInterval20270101Value) GetActualInstance() interface{} {
	if obj.CalendarIntervalDay_20270101 != nil {
		return obj.CalendarIntervalDay_20270101
	}

	if obj.CalendarIntervalWeek_20270101 != nil {
		return obj.CalendarIntervalWeek_20270101
	}

	if obj.CalendarIntervalMonth_20270101 != nil {
		return obj.CalendarIntervalMonth_20270101
	}

	// all schemas are nil
	return nil
}

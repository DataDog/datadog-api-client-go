// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardSummaryID - ID of the dashboard.
type DashboardSummaryID struct {
	String *string
	Int64  *int64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsDashboardSummaryID is a convenience function that returns string wrapped in DashboardSummaryID.
func StringAsDashboardSummaryID(v *string) DashboardSummaryID {
	return DashboardSummaryID{String: v}
}

// Int64AsDashboardSummaryID is a convenience function that returns int64 wrapped in DashboardSummaryID.
func Int64AsDashboardSummaryID(v *int64) DashboardSummaryID {
	return DashboardSummaryID{Int64: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DashboardSummaryID) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = datadog.Unmarshal(data, &obj.String)
	if err == nil {
		if obj.String != nil {
			jsonString, _ := datadog.Marshal(obj.String)
			if string(jsonString) == "{}" { // empty struct
				obj.String = nil
			} else {
				match++
			}
		} else {
			obj.String = nil
		}
	} else {
		obj.String = nil
	}

	// try to unmarshal data into Int64
	err = datadog.Unmarshal(data, &obj.Int64)
	if err == nil {
		if obj.Int64 != nil {
			jsonInt64, _ := datadog.Marshal(obj.Int64)
			if string(jsonInt64) == "{}" { // empty struct
				obj.Int64 = nil
			} else {
				match++
			}
		} else {
			obj.Int64 = nil
		}
	} else {
		obj.Int64 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.String = nil
		obj.Int64 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DashboardSummaryID) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}

	if obj.Int64 != nil {
		return datadog.Marshal(&obj.Int64)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DashboardSummaryID) GetActualInstance() interface{} {
	if obj.String != nil {
		return obj.String
	}

	if obj.Int64 != nil {
		return obj.Int64
	}

	// all schemas are nil
	return nil
}

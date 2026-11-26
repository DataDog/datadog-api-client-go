// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapNodeDimension - Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapNodeDimension struct {
	DashboardInlineSchema38_20270101          *DashboardInlineSchema38
	CustomTopologyMapValueColumnName_20270101 *string

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DashboardInlineSchema38_20270101AsCustomTopologyMapNodeDimension is a convenience function that returns DashboardInlineSchema38 wrapped in CustomTopologyMapNodeDimension.
func DashboardInlineSchema38_20270101AsCustomTopologyMapNodeDimension(v *DashboardInlineSchema38) CustomTopologyMapNodeDimension {
	return CustomTopologyMapNodeDimension{DashboardInlineSchema38_20270101: v}
}

// CustomTopologyMapValueColumnName_20270101AsCustomTopologyMapNodeDimension is a convenience function that returns string wrapped in CustomTopologyMapNodeDimension.
func CustomTopologyMapValueColumnName_20270101AsCustomTopologyMapNodeDimension(v *string) CustomTopologyMapNodeDimension {
	return CustomTopologyMapNodeDimension{CustomTopologyMapValueColumnName_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CustomTopologyMapNodeDimension) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DashboardInlineSchema38_20270101
	err = datadog.Unmarshal(data, &obj.DashboardInlineSchema38_20270101)
	if err == nil {
		if obj.DashboardInlineSchema38_20270101 != nil {
			jsonDashboardInlineSchema38_20270101, _ := datadog.Marshal(obj.DashboardInlineSchema38_20270101)
			if string(jsonDashboardInlineSchema38_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.DashboardInlineSchema38_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DashboardInlineSchema38_20270101 = nil
		}
	} else {
		obj.DashboardInlineSchema38_20270101 = nil
	}

	// try to unmarshal data into CustomTopologyMapValueColumnName_20270101
	err = datadog.Unmarshal(data, &obj.CustomTopologyMapValueColumnName_20270101)
	if err == nil {
		if obj.CustomTopologyMapValueColumnName_20270101 != nil {
			jsonCustomTopologyMapValueColumnName_20270101, _ := datadog.Marshal(obj.CustomTopologyMapValueColumnName_20270101)
			if string(jsonCustomTopologyMapValueColumnName_20270101) == "{}" { // empty struct
				obj.CustomTopologyMapValueColumnName_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CustomTopologyMapValueColumnName_20270101 = nil
		}
	} else {
		obj.CustomTopologyMapValueColumnName_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DashboardInlineSchema38_20270101 = nil
		obj.CustomTopologyMapValueColumnName_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CustomTopologyMapNodeDimension) MarshalJSON() ([]byte, error) {
	if obj.DashboardInlineSchema38_20270101 != nil {
		return datadog.Marshal(&obj.DashboardInlineSchema38_20270101)
	}

	if obj.CustomTopologyMapValueColumnName_20270101 != nil {
		return datadog.Marshal(&obj.CustomTopologyMapValueColumnName_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CustomTopologyMapNodeDimension) GetActualInstance() interface{} {
	if obj.DashboardInlineSchema38_20270101 != nil {
		return obj.DashboardInlineSchema38_20270101
	}

	if obj.CustomTopologyMapValueColumnName_20270101 != nil {
		return obj.CustomTopologyMapValueColumnName_20270101
	}

	// all schemas are nil
	return nil
}

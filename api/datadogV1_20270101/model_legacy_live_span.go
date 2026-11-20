// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LegacyLiveSpan - Test-only generated schema for the dashboard versioning proof of concept.
type LegacyLiveSpan struct {
	TimeSpan_20270101                *TimeSpan
	DashboardInlineSchema43_20270101 *DashboardInlineSchema43

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TimeSpan_20270101AsLegacyLiveSpan is a convenience function that returns TimeSpan wrapped in LegacyLiveSpan.
func TimeSpan_20270101AsLegacyLiveSpan(v *TimeSpan) LegacyLiveSpan {
	return LegacyLiveSpan{TimeSpan_20270101: v}
}

// DashboardInlineSchema43_20270101AsLegacyLiveSpan is a convenience function that returns DashboardInlineSchema43 wrapped in LegacyLiveSpan.
func DashboardInlineSchema43_20270101AsLegacyLiveSpan(v *DashboardInlineSchema43) LegacyLiveSpan {
	return LegacyLiveSpan{DashboardInlineSchema43_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LegacyLiveSpan) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TimeSpan_20270101
	err = datadog.Unmarshal(data, &obj.TimeSpan_20270101)
	if err == nil {
		if obj.TimeSpan_20270101 != nil {
			jsonTimeSpan_20270101, _ := datadog.Marshal(obj.TimeSpan_20270101)
			if string(jsonTimeSpan_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.TimeSpan_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeSpan_20270101 = nil
		}
	} else {
		obj.TimeSpan_20270101 = nil
	}

	// try to unmarshal data into DashboardInlineSchema43_20270101
	err = datadog.Unmarshal(data, &obj.DashboardInlineSchema43_20270101)
	if err == nil {
		if obj.DashboardInlineSchema43_20270101 != nil {
			jsonDashboardInlineSchema43_20270101, _ := datadog.Marshal(obj.DashboardInlineSchema43_20270101)
			if string(jsonDashboardInlineSchema43_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.DashboardInlineSchema43_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DashboardInlineSchema43_20270101 = nil
		}
	} else {
		obj.DashboardInlineSchema43_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TimeSpan_20270101 = nil
		obj.DashboardInlineSchema43_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LegacyLiveSpan) MarshalJSON() ([]byte, error) {
	if obj.TimeSpan_20270101 != nil {
		return datadog.Marshal(&obj.TimeSpan_20270101)
	}

	if obj.DashboardInlineSchema43_20270101 != nil {
		return datadog.Marshal(&obj.DashboardInlineSchema43_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LegacyLiveSpan) GetActualInstance() interface{} {
	if obj.TimeSpan_20270101 != nil {
		return obj.TimeSpan_20270101
	}

	if obj.DashboardInlineSchema43_20270101 != nil {
		return obj.DashboardInlineSchema43_20270101
	}

	// all schemas are nil
	return nil
}

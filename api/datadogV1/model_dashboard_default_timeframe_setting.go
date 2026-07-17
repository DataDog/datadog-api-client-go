// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardDefaultTimeframeSetting - The default timeframe applied when opening the dashboard. Set to `null` to clear the dashboard's default timeframe.
type DashboardDefaultTimeframeSetting struct {
	DashboardLiveTimeframe  *DashboardLiveTimeframe
	DashboardFixedTimeframe *DashboardFixedTimeframe

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DashboardLiveTimeframeAsDashboardDefaultTimeframeSetting is a convenience function that returns DashboardLiveTimeframe wrapped in DashboardDefaultTimeframeSetting.
func DashboardLiveTimeframeAsDashboardDefaultTimeframeSetting(v *DashboardLiveTimeframe) DashboardDefaultTimeframeSetting {
	return DashboardDefaultTimeframeSetting{DashboardLiveTimeframe: v}
}

// DashboardFixedTimeframeAsDashboardDefaultTimeframeSetting is a convenience function that returns DashboardFixedTimeframe wrapped in DashboardDefaultTimeframeSetting.
func DashboardFixedTimeframeAsDashboardDefaultTimeframeSetting(v *DashboardFixedTimeframe) DashboardDefaultTimeframeSetting {
	return DashboardDefaultTimeframeSetting{DashboardFixedTimeframe: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DashboardDefaultTimeframeSetting) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DashboardLiveTimeframe
	err = datadog.Unmarshal(data, &obj.DashboardLiveTimeframe)
	if err == nil {
		if obj.DashboardLiveTimeframe != nil && obj.DashboardLiveTimeframe.UnparsedObject == nil {
			jsonDashboardLiveTimeframe, _ := datadog.Marshal(obj.DashboardLiveTimeframe)
			if string(jsonDashboardLiveTimeframe) == "{}" { // empty struct
				obj.DashboardLiveTimeframe = nil
			} else {
				match++
			}
		} else {
			obj.DashboardLiveTimeframe = nil
		}
	} else {
		obj.DashboardLiveTimeframe = nil
	}

	// try to unmarshal data into DashboardFixedTimeframe
	err = datadog.Unmarshal(data, &obj.DashboardFixedTimeframe)
	if err == nil {
		if obj.DashboardFixedTimeframe != nil && obj.DashboardFixedTimeframe.UnparsedObject == nil {
			jsonDashboardFixedTimeframe, _ := datadog.Marshal(obj.DashboardFixedTimeframe)
			if string(jsonDashboardFixedTimeframe) == "{}" { // empty struct
				obj.DashboardFixedTimeframe = nil
			} else {
				match++
			}
		} else {
			obj.DashboardFixedTimeframe = nil
		}
	} else {
		obj.DashboardFixedTimeframe = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DashboardLiveTimeframe = nil
		obj.DashboardFixedTimeframe = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DashboardDefaultTimeframeSetting) MarshalJSON() ([]byte, error) {
	if obj.DashboardLiveTimeframe != nil {
		return datadog.Marshal(&obj.DashboardLiveTimeframe)
	}

	if obj.DashboardFixedTimeframe != nil {
		return datadog.Marshal(&obj.DashboardFixedTimeframe)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DashboardDefaultTimeframeSetting) GetActualInstance() interface{} {
	if obj.DashboardLiveTimeframe != nil {
		return obj.DashboardLiveTimeframe
	}

	if obj.DashboardFixedTimeframe != nil {
		return obj.DashboardFixedTimeframe
	}

	// all schemas are nil
	return nil
}

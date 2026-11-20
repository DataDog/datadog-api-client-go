// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SloMeasure - Test-only generated schema for the dashboard versioning proof of concept.
type SloMeasure struct {
	DashboardInlineSchema46_20270101 *DashboardInlineSchema46
	SloMeasureEvents_20270101        *SloMeasureEvents
	SloMeasureMinutes_20270101       *SloMeasureMinutes

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DashboardInlineSchema46_20270101AsSloMeasure is a convenience function that returns DashboardInlineSchema46 wrapped in SloMeasure.
func DashboardInlineSchema46_20270101AsSloMeasure(v *DashboardInlineSchema46) SloMeasure {
	return SloMeasure{DashboardInlineSchema46_20270101: v}
}

// SloMeasureEvents_20270101AsSloMeasure is a convenience function that returns SloMeasureEvents wrapped in SloMeasure.
func SloMeasureEvents_20270101AsSloMeasure(v *SloMeasureEvents) SloMeasure {
	return SloMeasure{SloMeasureEvents_20270101: v}
}

// SloMeasureMinutes_20270101AsSloMeasure is a convenience function that returns SloMeasureMinutes wrapped in SloMeasure.
func SloMeasureMinutes_20270101AsSloMeasure(v *SloMeasureMinutes) SloMeasure {
	return SloMeasure{SloMeasureMinutes_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SloMeasure) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DashboardInlineSchema46_20270101
	err = datadog.Unmarshal(data, &obj.DashboardInlineSchema46_20270101)
	if err == nil {
		if obj.DashboardInlineSchema46_20270101 != nil {
			jsonDashboardInlineSchema46_20270101, _ := datadog.Marshal(obj.DashboardInlineSchema46_20270101)
			if string(jsonDashboardInlineSchema46_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.DashboardInlineSchema46_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DashboardInlineSchema46_20270101 = nil
		}
	} else {
		obj.DashboardInlineSchema46_20270101 = nil
	}

	// try to unmarshal data into SloMeasureEvents_20270101
	err = datadog.Unmarshal(data, &obj.SloMeasureEvents_20270101)
	if err == nil {
		if obj.SloMeasureEvents_20270101 != nil {
			jsonSloMeasureEvents_20270101, _ := datadog.Marshal(obj.SloMeasureEvents_20270101)
			if string(jsonSloMeasureEvents_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.SloMeasureEvents_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.SloMeasureEvents_20270101 = nil
		}
	} else {
		obj.SloMeasureEvents_20270101 = nil
	}

	// try to unmarshal data into SloMeasureMinutes_20270101
	err = datadog.Unmarshal(data, &obj.SloMeasureMinutes_20270101)
	if err == nil {
		if obj.SloMeasureMinutes_20270101 != nil {
			jsonSloMeasureMinutes_20270101, _ := datadog.Marshal(obj.SloMeasureMinutes_20270101)
			if string(jsonSloMeasureMinutes_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.SloMeasureMinutes_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.SloMeasureMinutes_20270101 = nil
		}
	} else {
		obj.SloMeasureMinutes_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DashboardInlineSchema46_20270101 = nil
		obj.SloMeasureEvents_20270101 = nil
		obj.SloMeasureMinutes_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SloMeasure) MarshalJSON() ([]byte, error) {
	if obj.DashboardInlineSchema46_20270101 != nil {
		return datadog.Marshal(&obj.DashboardInlineSchema46_20270101)
	}

	if obj.SloMeasureEvents_20270101 != nil {
		return datadog.Marshal(&obj.SloMeasureEvents_20270101)
	}

	if obj.SloMeasureMinutes_20270101 != nil {
		return datadog.Marshal(&obj.SloMeasureMinutes_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SloMeasure) GetActualInstance() interface{} {
	if obj.DashboardInlineSchema46_20270101 != nil {
		return obj.DashboardInlineSchema46_20270101
	}

	if obj.SloMeasureEvents_20270101 != nil {
		return obj.SloMeasureEvents_20270101
	}

	if obj.SloMeasureMinutes_20270101 != nil {
		return obj.SloMeasureMinutes_20270101
	}

	// all schemas are nil
	return nil
}

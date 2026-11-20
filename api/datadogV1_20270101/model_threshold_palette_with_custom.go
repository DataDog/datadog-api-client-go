// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ThresholdPaletteWithCustom - Test-only generated schema for the dashboard versioning proof of concept.
type ThresholdPaletteWithCustom struct {
	ThresholdPalette_20270101        *ThresholdPalette
	DashboardInlineSchema49_20270101 *DashboardInlineSchema49

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ThresholdPalette_20270101AsThresholdPaletteWithCustom is a convenience function that returns ThresholdPalette wrapped in ThresholdPaletteWithCustom.
func ThresholdPalette_20270101AsThresholdPaletteWithCustom(v *ThresholdPalette) ThresholdPaletteWithCustom {
	return ThresholdPaletteWithCustom{ThresholdPalette_20270101: v}
}

// DashboardInlineSchema49_20270101AsThresholdPaletteWithCustom is a convenience function that returns DashboardInlineSchema49 wrapped in ThresholdPaletteWithCustom.
func DashboardInlineSchema49_20270101AsThresholdPaletteWithCustom(v *DashboardInlineSchema49) ThresholdPaletteWithCustom {
	return ThresholdPaletteWithCustom{DashboardInlineSchema49_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ThresholdPaletteWithCustom) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ThresholdPalette_20270101
	err = datadog.Unmarshal(data, &obj.ThresholdPalette_20270101)
	if err == nil {
		if obj.ThresholdPalette_20270101 != nil {
			jsonThresholdPalette_20270101, _ := datadog.Marshal(obj.ThresholdPalette_20270101)
			if string(jsonThresholdPalette_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.ThresholdPalette_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ThresholdPalette_20270101 = nil
		}
	} else {
		obj.ThresholdPalette_20270101 = nil
	}

	// try to unmarshal data into DashboardInlineSchema49_20270101
	err = datadog.Unmarshal(data, &obj.DashboardInlineSchema49_20270101)
	if err == nil {
		if obj.DashboardInlineSchema49_20270101 != nil {
			jsonDashboardInlineSchema49_20270101, _ := datadog.Marshal(obj.DashboardInlineSchema49_20270101)
			if string(jsonDashboardInlineSchema49_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.DashboardInlineSchema49_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DashboardInlineSchema49_20270101 = nil
		}
	} else {
		obj.DashboardInlineSchema49_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ThresholdPalette_20270101 = nil
		obj.DashboardInlineSchema49_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ThresholdPaletteWithCustom) MarshalJSON() ([]byte, error) {
	if obj.ThresholdPalette_20270101 != nil {
		return datadog.Marshal(&obj.ThresholdPalette_20270101)
	}

	if obj.DashboardInlineSchema49_20270101 != nil {
		return datadog.Marshal(&obj.DashboardInlineSchema49_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ThresholdPaletteWithCustom) GetActualInstance() interface{} {
	if obj.ThresholdPalette_20270101 != nil {
		return obj.ThresholdPalette_20270101
	}

	if obj.DashboardInlineSchema49_20270101 != nil {
		return obj.DashboardInlineSchema49_20270101
	}

	// all schemas are nil
	return nil
}

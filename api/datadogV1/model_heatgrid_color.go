// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridColor - A color string, or two color strings for the light and dark themes, in that order.
type HeatgridColor struct {
	String              *string
	HeatgridThemeColors *[]string

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsHeatgridColor is a convenience function that returns string wrapped in HeatgridColor.
func StringAsHeatgridColor(v *string) HeatgridColor {
	return HeatgridColor{String: v}
}

// HeatgridThemeColorsAsHeatgridColor is a convenience function that returns []string wrapped in HeatgridColor.
func HeatgridThemeColorsAsHeatgridColor(v *[]string) HeatgridColor {
	return HeatgridColor{HeatgridThemeColors: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *HeatgridColor) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into HeatgridThemeColors
	err = datadog.Unmarshal(data, &obj.HeatgridThemeColors)
	if err == nil {
		if obj.HeatgridThemeColors != nil {
			jsonHeatgridThemeColors, _ := datadog.Marshal(obj.HeatgridThemeColors)
			if string(jsonHeatgridThemeColors) == "{}" && string(data) != "{}" { // empty struct
				obj.HeatgridThemeColors = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridThemeColors = nil
		}
	} else {
		obj.HeatgridThemeColors = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.String = nil
		obj.HeatgridThemeColors = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj HeatgridColor) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}

	if obj.HeatgridThemeColors != nil {
		return datadog.Marshal(&obj.HeatgridThemeColors)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *HeatgridColor) GetActualInstance() interface{} {
	if obj.String != nil {
		return obj.String
	}

	if obj.HeatgridThemeColors != nil {
		return obj.HeatgridThemeColors
	}

	// all schemas are nil
	return nil
}

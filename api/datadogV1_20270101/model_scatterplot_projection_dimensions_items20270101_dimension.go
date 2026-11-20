// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotProjectionDimensionsItems20270101Dimension - Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotProjectionDimensionsItems20270101Dimension struct {
	DashboardInlineSchema32_20270101 *DashboardInlineSchema32
	ScatterPlotDimension_20270101    *ScatterPlotDimension

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DashboardInlineSchema32_20270101AsScatterplotProjectionDimensionsItems20270101Dimension is a convenience function that returns DashboardInlineSchema32 wrapped in ScatterplotProjectionDimensionsItems20270101Dimension.
func DashboardInlineSchema32_20270101AsScatterplotProjectionDimensionsItems20270101Dimension(v *DashboardInlineSchema32) ScatterplotProjectionDimensionsItems20270101Dimension {
	return ScatterplotProjectionDimensionsItems20270101Dimension{DashboardInlineSchema32_20270101: v}
}

// ScatterPlotDimension_20270101AsScatterplotProjectionDimensionsItems20270101Dimension is a convenience function that returns ScatterPlotDimension wrapped in ScatterplotProjectionDimensionsItems20270101Dimension.
func ScatterPlotDimension_20270101AsScatterplotProjectionDimensionsItems20270101Dimension(v *ScatterPlotDimension) ScatterplotProjectionDimensionsItems20270101Dimension {
	return ScatterplotProjectionDimensionsItems20270101Dimension{ScatterPlotDimension_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ScatterplotProjectionDimensionsItems20270101Dimension) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DashboardInlineSchema32_20270101
	err = datadog.Unmarshal(data, &obj.DashboardInlineSchema32_20270101)
	if err == nil {
		if obj.DashboardInlineSchema32_20270101 != nil {
			jsonDashboardInlineSchema32_20270101, _ := datadog.Marshal(obj.DashboardInlineSchema32_20270101)
			if string(jsonDashboardInlineSchema32_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.DashboardInlineSchema32_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DashboardInlineSchema32_20270101 = nil
		}
	} else {
		obj.DashboardInlineSchema32_20270101 = nil
	}

	// try to unmarshal data into ScatterPlotDimension_20270101
	err = datadog.Unmarshal(data, &obj.ScatterPlotDimension_20270101)
	if err == nil {
		if obj.ScatterPlotDimension_20270101 != nil {
			jsonScatterPlotDimension_20270101, _ := datadog.Marshal(obj.ScatterPlotDimension_20270101)
			if string(jsonScatterPlotDimension_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.ScatterPlotDimension_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ScatterPlotDimension_20270101 = nil
		}
	} else {
		obj.ScatterPlotDimension_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DashboardInlineSchema32_20270101 = nil
		obj.ScatterPlotDimension_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ScatterplotProjectionDimensionsItems20270101Dimension) MarshalJSON() ([]byte, error) {
	if obj.DashboardInlineSchema32_20270101 != nil {
		return datadog.Marshal(&obj.DashboardInlineSchema32_20270101)
	}

	if obj.ScatterPlotDimension_20270101 != nil {
		return datadog.Marshal(&obj.ScatterPlotDimension_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ScatterplotProjectionDimensionsItems20270101Dimension) GetActualInstance() interface{} {
	if obj.DashboardInlineSchema32_20270101 != nil {
		return obj.DashboardInlineSchema32_20270101
	}

	if obj.ScatterPlotDimension_20270101 != nil {
		return obj.ScatterPlotDimension_20270101
	}

	// all schemas are nil
	return nil
}

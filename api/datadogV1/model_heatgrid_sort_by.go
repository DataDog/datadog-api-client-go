// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridSortBy - Sort rows by aggregated value or group label.
type HeatgridSortBy struct {
	HeatgridSortByValue *HeatgridSortByValue
	HeatgridSortByLabel *HeatgridSortByLabel

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// HeatgridSortByValueAsHeatgridSortBy is a convenience function that returns HeatgridSortByValue wrapped in HeatgridSortBy.
func HeatgridSortByValueAsHeatgridSortBy(v *HeatgridSortByValue) HeatgridSortBy {
	return HeatgridSortBy{HeatgridSortByValue: v}
}

// HeatgridSortByLabelAsHeatgridSortBy is a convenience function that returns HeatgridSortByLabel wrapped in HeatgridSortBy.
func HeatgridSortByLabelAsHeatgridSortBy(v *HeatgridSortByLabel) HeatgridSortBy {
	return HeatgridSortBy{HeatgridSortByLabel: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *HeatgridSortBy) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HeatgridSortByValue
	err = datadog.Unmarshal(data, &obj.HeatgridSortByValue)
	if err == nil {
		if obj.HeatgridSortByValue != nil && obj.HeatgridSortByValue.UnparsedObject == nil {
			jsonHeatgridSortByValue, _ := datadog.Marshal(obj.HeatgridSortByValue)
			if string(jsonHeatgridSortByValue) == "{}" { // empty struct
				obj.HeatgridSortByValue = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridSortByValue = nil
		}
	} else {
		obj.HeatgridSortByValue = nil
	}

	// try to unmarshal data into HeatgridSortByLabel
	err = datadog.Unmarshal(data, &obj.HeatgridSortByLabel)
	if err == nil {
		if obj.HeatgridSortByLabel != nil && obj.HeatgridSortByLabel.UnparsedObject == nil {
			jsonHeatgridSortByLabel, _ := datadog.Marshal(obj.HeatgridSortByLabel)
			if string(jsonHeatgridSortByLabel) == "{}" { // empty struct
				obj.HeatgridSortByLabel = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridSortByLabel = nil
		}
	} else {
		obj.HeatgridSortByLabel = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.HeatgridSortByValue = nil
		obj.HeatgridSortByLabel = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj HeatgridSortBy) MarshalJSON() ([]byte, error) {
	if obj.HeatgridSortByValue != nil {
		return datadog.Marshal(&obj.HeatgridSortByValue)
	}

	if obj.HeatgridSortByLabel != nil {
		return datadog.Marshal(&obj.HeatgridSortByLabel)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *HeatgridSortBy) GetActualInstance() interface{} {
	if obj.HeatgridSortByValue != nil {
		return obj.HeatgridSortByValue
	}

	if obj.HeatgridSortByLabel != nil {
		return obj.HeatgridSortByLabel
	}

	// all schemas are nil
	return nil
}

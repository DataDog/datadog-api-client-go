// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridTextFormatReplace - Replace all or part of a matching group label.
type HeatgridTextFormatReplace struct {
	HeatgridTextFormatReplaceAll       *HeatgridTextFormatReplaceAll
	HeatgridTextFormatReplaceSubstring *HeatgridTextFormatReplaceSubstring

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// HeatgridTextFormatReplaceAllAsHeatgridTextFormatReplace is a convenience function that returns HeatgridTextFormatReplaceAll wrapped in HeatgridTextFormatReplace.
func HeatgridTextFormatReplaceAllAsHeatgridTextFormatReplace(v *HeatgridTextFormatReplaceAll) HeatgridTextFormatReplace {
	return HeatgridTextFormatReplace{HeatgridTextFormatReplaceAll: v}
}

// HeatgridTextFormatReplaceSubstringAsHeatgridTextFormatReplace is a convenience function that returns HeatgridTextFormatReplaceSubstring wrapped in HeatgridTextFormatReplace.
func HeatgridTextFormatReplaceSubstringAsHeatgridTextFormatReplace(v *HeatgridTextFormatReplaceSubstring) HeatgridTextFormatReplace {
	return HeatgridTextFormatReplace{HeatgridTextFormatReplaceSubstring: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *HeatgridTextFormatReplace) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HeatgridTextFormatReplaceAll
	err = datadog.Unmarshal(data, &obj.HeatgridTextFormatReplaceAll)
	if err == nil {
		if obj.HeatgridTextFormatReplaceAll != nil && obj.HeatgridTextFormatReplaceAll.UnparsedObject == nil {
			jsonHeatgridTextFormatReplaceAll, _ := datadog.Marshal(obj.HeatgridTextFormatReplaceAll)
			if string(jsonHeatgridTextFormatReplaceAll) == "{}" { // empty struct
				obj.HeatgridTextFormatReplaceAll = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridTextFormatReplaceAll = nil
		}
	} else {
		obj.HeatgridTextFormatReplaceAll = nil
	}

	// try to unmarshal data into HeatgridTextFormatReplaceSubstring
	err = datadog.Unmarshal(data, &obj.HeatgridTextFormatReplaceSubstring)
	if err == nil {
		if obj.HeatgridTextFormatReplaceSubstring != nil && obj.HeatgridTextFormatReplaceSubstring.UnparsedObject == nil {
			jsonHeatgridTextFormatReplaceSubstring, _ := datadog.Marshal(obj.HeatgridTextFormatReplaceSubstring)
			if string(jsonHeatgridTextFormatReplaceSubstring) == "{}" { // empty struct
				obj.HeatgridTextFormatReplaceSubstring = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridTextFormatReplaceSubstring = nil
		}
	} else {
		obj.HeatgridTextFormatReplaceSubstring = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.HeatgridTextFormatReplaceAll = nil
		obj.HeatgridTextFormatReplaceSubstring = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj HeatgridTextFormatReplace) MarshalJSON() ([]byte, error) {
	if obj.HeatgridTextFormatReplaceAll != nil {
		return datadog.Marshal(&obj.HeatgridTextFormatReplaceAll)
	}

	if obj.HeatgridTextFormatReplaceSubstring != nil {
		return datadog.Marshal(&obj.HeatgridTextFormatReplaceSubstring)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *HeatgridTextFormatReplace) GetActualInstance() interface{} {
	if obj.HeatgridTextFormatReplaceAll != nil {
		return obj.HeatgridTextFormatReplaceAll
	}

	if obj.HeatgridTextFormatReplaceSubstring != nil {
		return obj.HeatgridTextFormatReplaceSubstring
	}

	// all schemas are nil
	return nil
}

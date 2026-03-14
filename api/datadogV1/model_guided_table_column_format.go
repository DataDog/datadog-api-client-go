// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableColumnFormat - Visual formatting applied to values in a guided table column. Supports number/bar mode and trend mode.
type GuidedTableColumnFormat struct {
	GuidedTableNumberBarColumnFormat *GuidedTableNumberBarColumnFormat
	GuidedTableTrendColumnFormat     *GuidedTableTrendColumnFormat

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GuidedTableNumberBarColumnFormatAsGuidedTableColumnFormat is a convenience function that returns GuidedTableNumberBarColumnFormat wrapped in GuidedTableColumnFormat.
func GuidedTableNumberBarColumnFormatAsGuidedTableColumnFormat(v *GuidedTableNumberBarColumnFormat) GuidedTableColumnFormat {
	return GuidedTableColumnFormat{GuidedTableNumberBarColumnFormat: v}
}

// GuidedTableTrendColumnFormatAsGuidedTableColumnFormat is a convenience function that returns GuidedTableTrendColumnFormat wrapped in GuidedTableColumnFormat.
func GuidedTableTrendColumnFormatAsGuidedTableColumnFormat(v *GuidedTableTrendColumnFormat) GuidedTableColumnFormat {
	return GuidedTableColumnFormat{GuidedTableTrendColumnFormat: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GuidedTableColumnFormat) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GuidedTableNumberBarColumnFormat
	err = datadog.Unmarshal(data, &obj.GuidedTableNumberBarColumnFormat)
	if err == nil {
		if obj.GuidedTableNumberBarColumnFormat != nil && obj.GuidedTableNumberBarColumnFormat.UnparsedObject == nil {
			jsonGuidedTableNumberBarColumnFormat, _ := datadog.Marshal(obj.GuidedTableNumberBarColumnFormat)
			if string(jsonGuidedTableNumberBarColumnFormat) == "{}" { // empty struct
				obj.GuidedTableNumberBarColumnFormat = nil
			} else {
				match++
			}
		} else {
			obj.GuidedTableNumberBarColumnFormat = nil
		}
	} else {
		obj.GuidedTableNumberBarColumnFormat = nil
	}

	// try to unmarshal data into GuidedTableTrendColumnFormat
	err = datadog.Unmarshal(data, &obj.GuidedTableTrendColumnFormat)
	if err == nil {
		if obj.GuidedTableTrendColumnFormat != nil && obj.GuidedTableTrendColumnFormat.UnparsedObject == nil {
			jsonGuidedTableTrendColumnFormat, _ := datadog.Marshal(obj.GuidedTableTrendColumnFormat)
			if string(jsonGuidedTableTrendColumnFormat) == "{}" { // empty struct
				obj.GuidedTableTrendColumnFormat = nil
			} else {
				match++
			}
		} else {
			obj.GuidedTableTrendColumnFormat = nil
		}
	} else {
		obj.GuidedTableTrendColumnFormat = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GuidedTableNumberBarColumnFormat = nil
		obj.GuidedTableTrendColumnFormat = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GuidedTableColumnFormat) MarshalJSON() ([]byte, error) {
	if obj.GuidedTableNumberBarColumnFormat != nil {
		return datadog.Marshal(&obj.GuidedTableNumberBarColumnFormat)
	}

	if obj.GuidedTableTrendColumnFormat != nil {
		return datadog.Marshal(&obj.GuidedTableTrendColumnFormat)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GuidedTableColumnFormat) GetActualInstance() interface{} {
	if obj.GuidedTableNumberBarColumnFormat != nil {
		return obj.GuidedTableNumberBarColumnFormat
	}

	if obj.GuidedTableTrendColumnFormat != nil {
		return obj.GuidedTableTrendColumnFormat
	}

	// all schemas are nil
	return nil
}

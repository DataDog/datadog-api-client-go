// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableColumnFormat - Test-only generated schema for the dashboard versioning proof of concept.
type TableColumnFormat struct {
	NumberBarTableColumnFormatNumber_20270101 *NumberBarTableColumnFormatNumber
	NumberBarTableColumnFormatBar_20270101    *NumberBarTableColumnFormatBar
	TrendTableColumnFormat_20270101           *TrendTableColumnFormat
	StringTableColumnFormat_20270101          *StringTableColumnFormat

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NumberBarTableColumnFormatNumber_20270101AsTableColumnFormat is a convenience function that returns NumberBarTableColumnFormatNumber wrapped in TableColumnFormat.
func NumberBarTableColumnFormatNumber_20270101AsTableColumnFormat(v *NumberBarTableColumnFormatNumber) TableColumnFormat {
	return TableColumnFormat{NumberBarTableColumnFormatNumber_20270101: v}
}

// NumberBarTableColumnFormatBar_20270101AsTableColumnFormat is a convenience function that returns NumberBarTableColumnFormatBar wrapped in TableColumnFormat.
func NumberBarTableColumnFormatBar_20270101AsTableColumnFormat(v *NumberBarTableColumnFormatBar) TableColumnFormat {
	return TableColumnFormat{NumberBarTableColumnFormatBar_20270101: v}
}

// TrendTableColumnFormat_20270101AsTableColumnFormat is a convenience function that returns TrendTableColumnFormat wrapped in TableColumnFormat.
func TrendTableColumnFormat_20270101AsTableColumnFormat(v *TrendTableColumnFormat) TableColumnFormat {
	return TableColumnFormat{TrendTableColumnFormat_20270101: v}
}

// StringTableColumnFormat_20270101AsTableColumnFormat is a convenience function that returns StringTableColumnFormat wrapped in TableColumnFormat.
func StringTableColumnFormat_20270101AsTableColumnFormat(v *StringTableColumnFormat) TableColumnFormat {
	return TableColumnFormat{StringTableColumnFormat_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TableColumnFormat) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NumberBarTableColumnFormatNumber_20270101
	err = datadog.Unmarshal(data, &obj.NumberBarTableColumnFormatNumber_20270101)
	if err == nil {
		if obj.NumberBarTableColumnFormatNumber_20270101 != nil && obj.NumberBarTableColumnFormatNumber_20270101.UnparsedObject == nil {
			jsonNumberBarTableColumnFormatNumber_20270101, _ := datadog.Marshal(obj.NumberBarTableColumnFormatNumber_20270101)
			if string(jsonNumberBarTableColumnFormatNumber_20270101) == "{}" { // empty struct
				obj.NumberBarTableColumnFormatNumber_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.NumberBarTableColumnFormatNumber_20270101 = nil
		}
	} else {
		obj.NumberBarTableColumnFormatNumber_20270101 = nil
	}

	// try to unmarshal data into NumberBarTableColumnFormatBar_20270101
	err = datadog.Unmarshal(data, &obj.NumberBarTableColumnFormatBar_20270101)
	if err == nil {
		if obj.NumberBarTableColumnFormatBar_20270101 != nil && obj.NumberBarTableColumnFormatBar_20270101.UnparsedObject == nil {
			jsonNumberBarTableColumnFormatBar_20270101, _ := datadog.Marshal(obj.NumberBarTableColumnFormatBar_20270101)
			if string(jsonNumberBarTableColumnFormatBar_20270101) == "{}" { // empty struct
				obj.NumberBarTableColumnFormatBar_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.NumberBarTableColumnFormatBar_20270101 = nil
		}
	} else {
		obj.NumberBarTableColumnFormatBar_20270101 = nil
	}

	// try to unmarshal data into TrendTableColumnFormat_20270101
	err = datadog.Unmarshal(data, &obj.TrendTableColumnFormat_20270101)
	if err == nil {
		if obj.TrendTableColumnFormat_20270101 != nil && obj.TrendTableColumnFormat_20270101.UnparsedObject == nil {
			jsonTrendTableColumnFormat_20270101, _ := datadog.Marshal(obj.TrendTableColumnFormat_20270101)
			if string(jsonTrendTableColumnFormat_20270101) == "{}" { // empty struct
				obj.TrendTableColumnFormat_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TrendTableColumnFormat_20270101 = nil
		}
	} else {
		obj.TrendTableColumnFormat_20270101 = nil
	}

	// try to unmarshal data into StringTableColumnFormat_20270101
	err = datadog.Unmarshal(data, &obj.StringTableColumnFormat_20270101)
	if err == nil {
		if obj.StringTableColumnFormat_20270101 != nil && obj.StringTableColumnFormat_20270101.UnparsedObject == nil {
			jsonStringTableColumnFormat_20270101, _ := datadog.Marshal(obj.StringTableColumnFormat_20270101)
			if string(jsonStringTableColumnFormat_20270101) == "{}" { // empty struct
				obj.StringTableColumnFormat_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.StringTableColumnFormat_20270101 = nil
		}
	} else {
		obj.StringTableColumnFormat_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.NumberBarTableColumnFormatNumber_20270101 = nil
		obj.NumberBarTableColumnFormatBar_20270101 = nil
		obj.TrendTableColumnFormat_20270101 = nil
		obj.StringTableColumnFormat_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TableColumnFormat) MarshalJSON() ([]byte, error) {
	if obj.NumberBarTableColumnFormatNumber_20270101 != nil {
		return datadog.Marshal(&obj.NumberBarTableColumnFormatNumber_20270101)
	}

	if obj.NumberBarTableColumnFormatBar_20270101 != nil {
		return datadog.Marshal(&obj.NumberBarTableColumnFormatBar_20270101)
	}

	if obj.TrendTableColumnFormat_20270101 != nil {
		return datadog.Marshal(&obj.TrendTableColumnFormat_20270101)
	}

	if obj.StringTableColumnFormat_20270101 != nil {
		return datadog.Marshal(&obj.StringTableColumnFormat_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TableColumnFormat) GetActualInstance() interface{} {
	if obj.NumberBarTableColumnFormatNumber_20270101 != nil {
		return obj.NumberBarTableColumnFormatNumber_20270101
	}

	if obj.NumberBarTableColumnFormatBar_20270101 != nil {
		return obj.NumberBarTableColumnFormatBar_20270101
	}

	if obj.TrendTableColumnFormat_20270101 != nil {
		return obj.TrendTableColumnFormat_20270101
	}

	if obj.StringTableColumnFormat_20270101 != nil {
		return obj.StringTableColumnFormat_20270101
	}

	// all schemas are nil
	return nil
}

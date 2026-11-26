// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableDateFormat - Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableDateFormat struct {
	QueryTableDatePattern_20270101     *QueryTableDatePattern
	QueryTableTimestampFormat_20270101 *QueryTableTimestampFormat

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// QueryTableDatePattern_20270101AsQueryTableDateFormat is a convenience function that returns QueryTableDatePattern wrapped in QueryTableDateFormat.
func QueryTableDatePattern_20270101AsQueryTableDateFormat(v *QueryTableDatePattern) QueryTableDateFormat {
	return QueryTableDateFormat{QueryTableDatePattern_20270101: v}
}

// QueryTableTimestampFormat_20270101AsQueryTableDateFormat is a convenience function that returns QueryTableTimestampFormat wrapped in QueryTableDateFormat.
func QueryTableTimestampFormat_20270101AsQueryTableDateFormat(v *QueryTableTimestampFormat) QueryTableDateFormat {
	return QueryTableDateFormat{QueryTableTimestampFormat_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *QueryTableDateFormat) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into QueryTableDatePattern_20270101
	err = datadog.Unmarshal(data, &obj.QueryTableDatePattern_20270101)
	if err == nil {
		if obj.QueryTableDatePattern_20270101 != nil {
			jsonQueryTableDatePattern_20270101, _ := datadog.Marshal(obj.QueryTableDatePattern_20270101)
			if string(jsonQueryTableDatePattern_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.QueryTableDatePattern_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.QueryTableDatePattern_20270101 = nil
		}
	} else {
		obj.QueryTableDatePattern_20270101 = nil
	}

	// try to unmarshal data into QueryTableTimestampFormat_20270101
	err = datadog.Unmarshal(data, &obj.QueryTableTimestampFormat_20270101)
	if err == nil {
		if obj.QueryTableTimestampFormat_20270101 != nil {
			jsonQueryTableTimestampFormat_20270101, _ := datadog.Marshal(obj.QueryTableTimestampFormat_20270101)
			if string(jsonQueryTableTimestampFormat_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.QueryTableTimestampFormat_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.QueryTableTimestampFormat_20270101 = nil
		}
	} else {
		obj.QueryTableTimestampFormat_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.QueryTableDatePattern_20270101 = nil
		obj.QueryTableTimestampFormat_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj QueryTableDateFormat) MarshalJSON() ([]byte, error) {
	if obj.QueryTableDatePattern_20270101 != nil {
		return datadog.Marshal(&obj.QueryTableDatePattern_20270101)
	}

	if obj.QueryTableTimestampFormat_20270101 != nil {
		return datadog.Marshal(&obj.QueryTableTimestampFormat_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *QueryTableDateFormat) GetActualInstance() interface{} {
	if obj.QueryTableDatePattern_20270101 != nil {
		return obj.QueryTableDatePattern_20270101
	}

	if obj.QueryTableTimestampFormat_20270101 != nil {
		return obj.QueryTableTimestampFormat_20270101
	}

	// all schemas are nil
	return nil
}

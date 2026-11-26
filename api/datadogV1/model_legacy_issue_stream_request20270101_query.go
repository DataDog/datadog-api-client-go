// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LegacyIssueStreamRequest20270101Query - Test-only generated schema for the dashboard versioning proof of concept.
type LegacyIssueStreamRequest20270101Query struct {
	APMIssueStreamQuery_20270101  *APMIssueStreamQuery_20270101
	RUMIssueStreamQuery_20270101  *RUMIssueStreamQuery_20270101
	LogsIssueStreamQuery_20270101 *LogsIssueStreamQuery_20270101

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// APMIssueStreamQuery_20270101AsLegacyIssueStreamRequest20270101Query is a convenience function that returns APMIssueStreamQuery_20270101 wrapped in LegacyIssueStreamRequest20270101Query.
func APMIssueStreamQuery_20270101AsLegacyIssueStreamRequest20270101Query(v *APMIssueStreamQuery_20270101) LegacyIssueStreamRequest20270101Query {
	return LegacyIssueStreamRequest20270101Query{APMIssueStreamQuery_20270101: v}
}

// RUMIssueStreamQuery_20270101AsLegacyIssueStreamRequest20270101Query is a convenience function that returns RUMIssueStreamQuery_20270101 wrapped in LegacyIssueStreamRequest20270101Query.
func RUMIssueStreamQuery_20270101AsLegacyIssueStreamRequest20270101Query(v *RUMIssueStreamQuery_20270101) LegacyIssueStreamRequest20270101Query {
	return LegacyIssueStreamRequest20270101Query{RUMIssueStreamQuery_20270101: v}
}

// LogsIssueStreamQuery_20270101AsLegacyIssueStreamRequest20270101Query is a convenience function that returns LogsIssueStreamQuery_20270101 wrapped in LegacyIssueStreamRequest20270101Query.
func LogsIssueStreamQuery_20270101AsLegacyIssueStreamRequest20270101Query(v *LogsIssueStreamQuery_20270101) LegacyIssueStreamRequest20270101Query {
	return LegacyIssueStreamRequest20270101Query{LogsIssueStreamQuery_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LegacyIssueStreamRequest20270101Query) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into APMIssueStreamQuery_20270101
	err = datadog.Unmarshal(data, &obj.APMIssueStreamQuery_20270101)
	if err == nil {
		if obj.APMIssueStreamQuery_20270101 != nil && obj.APMIssueStreamQuery_20270101.UnparsedObject == nil {
			jsonAPMIssueStreamQuery_20270101, _ := datadog.Marshal(obj.APMIssueStreamQuery_20270101)
			if string(jsonAPMIssueStreamQuery_20270101) == "{}" { // empty struct
				obj.APMIssueStreamQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.APMIssueStreamQuery_20270101 = nil
		}
	} else {
		obj.APMIssueStreamQuery_20270101 = nil
	}

	// try to unmarshal data into RUMIssueStreamQuery_20270101
	err = datadog.Unmarshal(data, &obj.RUMIssueStreamQuery_20270101)
	if err == nil {
		if obj.RUMIssueStreamQuery_20270101 != nil && obj.RUMIssueStreamQuery_20270101.UnparsedObject == nil {
			jsonRUMIssueStreamQuery_20270101, _ := datadog.Marshal(obj.RUMIssueStreamQuery_20270101)
			if string(jsonRUMIssueStreamQuery_20270101) == "{}" { // empty struct
				obj.RUMIssueStreamQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.RUMIssueStreamQuery_20270101 = nil
		}
	} else {
		obj.RUMIssueStreamQuery_20270101 = nil
	}

	// try to unmarshal data into LogsIssueStreamQuery_20270101
	err = datadog.Unmarshal(data, &obj.LogsIssueStreamQuery_20270101)
	if err == nil {
		if obj.LogsIssueStreamQuery_20270101 != nil && obj.LogsIssueStreamQuery_20270101.UnparsedObject == nil {
			jsonLogsIssueStreamQuery_20270101, _ := datadog.Marshal(obj.LogsIssueStreamQuery_20270101)
			if string(jsonLogsIssueStreamQuery_20270101) == "{}" { // empty struct
				obj.LogsIssueStreamQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.LogsIssueStreamQuery_20270101 = nil
		}
	} else {
		obj.LogsIssueStreamQuery_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.APMIssueStreamQuery_20270101 = nil
		obj.RUMIssueStreamQuery_20270101 = nil
		obj.LogsIssueStreamQuery_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LegacyIssueStreamRequest20270101Query) MarshalJSON() ([]byte, error) {
	if obj.APMIssueStreamQuery_20270101 != nil {
		return datadog.Marshal(&obj.APMIssueStreamQuery_20270101)
	}

	if obj.RUMIssueStreamQuery_20270101 != nil {
		return datadog.Marshal(&obj.RUMIssueStreamQuery_20270101)
	}

	if obj.LogsIssueStreamQuery_20270101 != nil {
		return datadog.Marshal(&obj.LogsIssueStreamQuery_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LegacyIssueStreamRequest20270101Query) GetActualInstance() interface{} {
	if obj.APMIssueStreamQuery_20270101 != nil {
		return obj.APMIssueStreamQuery_20270101
	}

	if obj.RUMIssueStreamQuery_20270101 != nil {
		return obj.RUMIssueStreamQuery_20270101
	}

	if obj.LogsIssueStreamQuery_20270101 != nil {
		return obj.LogsIssueStreamQuery_20270101
	}

	// all schemas are nil
	return nil
}

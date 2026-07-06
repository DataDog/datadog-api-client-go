// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookTemplateVariableAvailableValuesQuery - Query used to dynamically populate the list of available values for the template variable.
type NotebookTemplateVariableAvailableValuesQuery struct {
	NotebookTemplateVariableAvailableValuesQueryLogRumSpans *NotebookTemplateVariableAvailableValuesQueryLogRumSpans
	NotebookTemplateVariableAvailableValuesQueryMetrics     *NotebookTemplateVariableAvailableValuesQueryMetrics

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NotebookTemplateVariableAvailableValuesQueryLogRumSpansAsNotebookTemplateVariableAvailableValuesQuery is a convenience function that returns NotebookTemplateVariableAvailableValuesQueryLogRumSpans wrapped in NotebookTemplateVariableAvailableValuesQuery.
func NotebookTemplateVariableAvailableValuesQueryLogRumSpansAsNotebookTemplateVariableAvailableValuesQuery(v *NotebookTemplateVariableAvailableValuesQueryLogRumSpans) NotebookTemplateVariableAvailableValuesQuery {
	return NotebookTemplateVariableAvailableValuesQuery{NotebookTemplateVariableAvailableValuesQueryLogRumSpans: v}
}

// NotebookTemplateVariableAvailableValuesQueryMetricsAsNotebookTemplateVariableAvailableValuesQuery is a convenience function that returns NotebookTemplateVariableAvailableValuesQueryMetrics wrapped in NotebookTemplateVariableAvailableValuesQuery.
func NotebookTemplateVariableAvailableValuesQueryMetricsAsNotebookTemplateVariableAvailableValuesQuery(v *NotebookTemplateVariableAvailableValuesQueryMetrics) NotebookTemplateVariableAvailableValuesQuery {
	return NotebookTemplateVariableAvailableValuesQuery{NotebookTemplateVariableAvailableValuesQueryMetrics: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *NotebookTemplateVariableAvailableValuesQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NotebookTemplateVariableAvailableValuesQueryLogRumSpans
	err = datadog.Unmarshal(data, &obj.NotebookTemplateVariableAvailableValuesQueryLogRumSpans)
	if err == nil {
		if obj.NotebookTemplateVariableAvailableValuesQueryLogRumSpans != nil && obj.NotebookTemplateVariableAvailableValuesQueryLogRumSpans.UnparsedObject == nil {
			jsonNotebookTemplateVariableAvailableValuesQueryLogRumSpans, _ := datadog.Marshal(obj.NotebookTemplateVariableAvailableValuesQueryLogRumSpans)
			if string(jsonNotebookTemplateVariableAvailableValuesQueryLogRumSpans) == "{}" { // empty struct
				obj.NotebookTemplateVariableAvailableValuesQueryLogRumSpans = nil
			} else {
				match++
			}
		} else {
			obj.NotebookTemplateVariableAvailableValuesQueryLogRumSpans = nil
		}
	} else {
		obj.NotebookTemplateVariableAvailableValuesQueryLogRumSpans = nil
	}

	// try to unmarshal data into NotebookTemplateVariableAvailableValuesQueryMetrics
	err = datadog.Unmarshal(data, &obj.NotebookTemplateVariableAvailableValuesQueryMetrics)
	if err == nil {
		if obj.NotebookTemplateVariableAvailableValuesQueryMetrics != nil && obj.NotebookTemplateVariableAvailableValuesQueryMetrics.UnparsedObject == nil {
			jsonNotebookTemplateVariableAvailableValuesQueryMetrics, _ := datadog.Marshal(obj.NotebookTemplateVariableAvailableValuesQueryMetrics)
			if string(jsonNotebookTemplateVariableAvailableValuesQueryMetrics) == "{}" { // empty struct
				obj.NotebookTemplateVariableAvailableValuesQueryMetrics = nil
			} else {
				match++
			}
		} else {
			obj.NotebookTemplateVariableAvailableValuesQueryMetrics = nil
		}
	} else {
		obj.NotebookTemplateVariableAvailableValuesQueryMetrics = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.NotebookTemplateVariableAvailableValuesQueryLogRumSpans = nil
		obj.NotebookTemplateVariableAvailableValuesQueryMetrics = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj NotebookTemplateVariableAvailableValuesQuery) MarshalJSON() ([]byte, error) {
	if obj.NotebookTemplateVariableAvailableValuesQueryLogRumSpans != nil {
		return datadog.Marshal(&obj.NotebookTemplateVariableAvailableValuesQueryLogRumSpans)
	}

	if obj.NotebookTemplateVariableAvailableValuesQueryMetrics != nil {
		return datadog.Marshal(&obj.NotebookTemplateVariableAvailableValuesQueryMetrics)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *NotebookTemplateVariableAvailableValuesQuery) GetActualInstance() interface{} {
	if obj.NotebookTemplateVariableAvailableValuesQueryLogRumSpans != nil {
		return obj.NotebookTemplateVariableAvailableValuesQueryLogRumSpans
	}

	if obj.NotebookTemplateVariableAvailableValuesQueryMetrics != nil {
		return obj.NotebookTemplateVariableAvailableValuesQueryMetrics
	}

	// all schemas are nil
	return nil
}

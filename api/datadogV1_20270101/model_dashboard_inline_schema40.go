// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema40 - Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema40 struct {
	FormulaDatasetMetricQuery_20270101     *FormulaDatasetMetricQuery
	FormulaDatasetApmMetricsQuery_20270101 *FormulaDatasetApmMetricsQuery

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FormulaDatasetMetricQuery_20270101AsDashboardInlineSchema40 is a convenience function that returns FormulaDatasetMetricQuery wrapped in DashboardInlineSchema40.
func FormulaDatasetMetricQuery_20270101AsDashboardInlineSchema40(v *FormulaDatasetMetricQuery) DashboardInlineSchema40 {
	return DashboardInlineSchema40{FormulaDatasetMetricQuery_20270101: v}
}

// FormulaDatasetApmMetricsQuery_20270101AsDashboardInlineSchema40 is a convenience function that returns FormulaDatasetApmMetricsQuery wrapped in DashboardInlineSchema40.
func FormulaDatasetApmMetricsQuery_20270101AsDashboardInlineSchema40(v *FormulaDatasetApmMetricsQuery) DashboardInlineSchema40 {
	return DashboardInlineSchema40{FormulaDatasetApmMetricsQuery_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DashboardInlineSchema40) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FormulaDatasetMetricQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaDatasetMetricQuery_20270101)
	if err == nil {
		if obj.FormulaDatasetMetricQuery_20270101 != nil && obj.FormulaDatasetMetricQuery_20270101.UnparsedObject == nil {
			jsonFormulaDatasetMetricQuery_20270101, _ := datadog.Marshal(obj.FormulaDatasetMetricQuery_20270101)
			if string(jsonFormulaDatasetMetricQuery_20270101) == "{}" { // empty struct
				obj.FormulaDatasetMetricQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaDatasetMetricQuery_20270101 = nil
		}
	} else {
		obj.FormulaDatasetMetricQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaDatasetApmMetricsQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaDatasetApmMetricsQuery_20270101)
	if err == nil {
		if obj.FormulaDatasetApmMetricsQuery_20270101 != nil && obj.FormulaDatasetApmMetricsQuery_20270101.UnparsedObject == nil {
			jsonFormulaDatasetApmMetricsQuery_20270101, _ := datadog.Marshal(obj.FormulaDatasetApmMetricsQuery_20270101)
			if string(jsonFormulaDatasetApmMetricsQuery_20270101) == "{}" { // empty struct
				obj.FormulaDatasetApmMetricsQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaDatasetApmMetricsQuery_20270101 = nil
		}
	} else {
		obj.FormulaDatasetApmMetricsQuery_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FormulaDatasetMetricQuery_20270101 = nil
		obj.FormulaDatasetApmMetricsQuery_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DashboardInlineSchema40) MarshalJSON() ([]byte, error) {
	if obj.FormulaDatasetMetricQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaDatasetMetricQuery_20270101)
	}

	if obj.FormulaDatasetApmMetricsQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaDatasetApmMetricsQuery_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DashboardInlineSchema40) GetActualInstance() interface{} {
	if obj.FormulaDatasetMetricQuery_20270101 != nil {
		return obj.FormulaDatasetMetricQuery_20270101
	}

	if obj.FormulaDatasetApmMetricsQuery_20270101 != nil {
		return obj.FormulaDatasetApmMetricsQuery_20270101
	}

	// all schemas are nil
	return nil
}

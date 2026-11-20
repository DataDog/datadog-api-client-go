// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TabularDataRequest - Test-only generated schema for the dashboard versioning proof of concept.
type TabularDataRequest struct {
	AllListRequestVariants_20270101      *AllListRequestVariants
	TabularAdvancedQueryRequest_20270101 *TabularAdvancedQueryRequest
	TimeseriesFormulaRequest_20270101    *TimeseriesFormulaRequest
	ScalarFormulaRequest_20270101        *ScalarFormulaRequest
	TabularLocalDatasetRequest_20270101  *TabularLocalDatasetRequest
	HistogramRequest_20270101            *HistogramRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AllListRequestVariants_20270101AsTabularDataRequest is a convenience function that returns AllListRequestVariants wrapped in TabularDataRequest.
func AllListRequestVariants_20270101AsTabularDataRequest(v *AllListRequestVariants) TabularDataRequest {
	return TabularDataRequest{AllListRequestVariants_20270101: v}
}

// TabularAdvancedQueryRequest_20270101AsTabularDataRequest is a convenience function that returns TabularAdvancedQueryRequest wrapped in TabularDataRequest.
func TabularAdvancedQueryRequest_20270101AsTabularDataRequest(v *TabularAdvancedQueryRequest) TabularDataRequest {
	return TabularDataRequest{TabularAdvancedQueryRequest_20270101: v}
}

// TimeseriesFormulaRequest_20270101AsTabularDataRequest is a convenience function that returns TimeseriesFormulaRequest wrapped in TabularDataRequest.
func TimeseriesFormulaRequest_20270101AsTabularDataRequest(v *TimeseriesFormulaRequest) TabularDataRequest {
	return TabularDataRequest{TimeseriesFormulaRequest_20270101: v}
}

// ScalarFormulaRequest_20270101AsTabularDataRequest is a convenience function that returns ScalarFormulaRequest wrapped in TabularDataRequest.
func ScalarFormulaRequest_20270101AsTabularDataRequest(v *ScalarFormulaRequest) TabularDataRequest {
	return TabularDataRequest{ScalarFormulaRequest_20270101: v}
}

// TabularLocalDatasetRequest_20270101AsTabularDataRequest is a convenience function that returns TabularLocalDatasetRequest wrapped in TabularDataRequest.
func TabularLocalDatasetRequest_20270101AsTabularDataRequest(v *TabularLocalDatasetRequest) TabularDataRequest {
	return TabularDataRequest{TabularLocalDatasetRequest_20270101: v}
}

// HistogramRequest_20270101AsTabularDataRequest is a convenience function that returns HistogramRequest wrapped in TabularDataRequest.
func HistogramRequest_20270101AsTabularDataRequest(v *HistogramRequest) TabularDataRequest {
	return TabularDataRequest{HistogramRequest_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TabularDataRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AllListRequestVariants_20270101
	err = datadog.Unmarshal(data, &obj.AllListRequestVariants_20270101)
	if err == nil {
		if obj.AllListRequestVariants_20270101 != nil && obj.AllListRequestVariants_20270101.UnparsedObject == nil {
			jsonAllListRequestVariants_20270101, _ := datadog.Marshal(obj.AllListRequestVariants_20270101)
			if string(jsonAllListRequestVariants_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.AllListRequestVariants_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.AllListRequestVariants_20270101 = nil
		}
	} else {
		obj.AllListRequestVariants_20270101 = nil
	}

	// try to unmarshal data into TabularAdvancedQueryRequest_20270101
	err = datadog.Unmarshal(data, &obj.TabularAdvancedQueryRequest_20270101)
	if err == nil {
		if obj.TabularAdvancedQueryRequest_20270101 != nil && obj.TabularAdvancedQueryRequest_20270101.UnparsedObject == nil {
			jsonTabularAdvancedQueryRequest_20270101, _ := datadog.Marshal(obj.TabularAdvancedQueryRequest_20270101)
			if string(jsonTabularAdvancedQueryRequest_20270101) == "{}" { // empty struct
				obj.TabularAdvancedQueryRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TabularAdvancedQueryRequest_20270101 = nil
		}
	} else {
		obj.TabularAdvancedQueryRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesFormulaRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesFormulaRequest_20270101)
	if err == nil {
		if obj.TimeseriesFormulaRequest_20270101 != nil && obj.TimeseriesFormulaRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesFormulaRequest_20270101, _ := datadog.Marshal(obj.TimeseriesFormulaRequest_20270101)
			if string(jsonTimeseriesFormulaRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesFormulaRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesFormulaRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesFormulaRequest_20270101 = nil
	}

	// try to unmarshal data into ScalarFormulaRequest_20270101
	err = datadog.Unmarshal(data, &obj.ScalarFormulaRequest_20270101)
	if err == nil {
		if obj.ScalarFormulaRequest_20270101 != nil && obj.ScalarFormulaRequest_20270101.UnparsedObject == nil {
			jsonScalarFormulaRequest_20270101, _ := datadog.Marshal(obj.ScalarFormulaRequest_20270101)
			if string(jsonScalarFormulaRequest_20270101) == "{}" { // empty struct
				obj.ScalarFormulaRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ScalarFormulaRequest_20270101 = nil
		}
	} else {
		obj.ScalarFormulaRequest_20270101 = nil
	}

	// try to unmarshal data into TabularLocalDatasetRequest_20270101
	err = datadog.Unmarshal(data, &obj.TabularLocalDatasetRequest_20270101)
	if err == nil {
		if obj.TabularLocalDatasetRequest_20270101 != nil && obj.TabularLocalDatasetRequest_20270101.UnparsedObject == nil {
			jsonTabularLocalDatasetRequest_20270101, _ := datadog.Marshal(obj.TabularLocalDatasetRequest_20270101)
			if string(jsonTabularLocalDatasetRequest_20270101) == "{}" { // empty struct
				obj.TabularLocalDatasetRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TabularLocalDatasetRequest_20270101 = nil
		}
	} else {
		obj.TabularLocalDatasetRequest_20270101 = nil
	}

	// try to unmarshal data into HistogramRequest_20270101
	err = datadog.Unmarshal(data, &obj.HistogramRequest_20270101)
	if err == nil {
		if obj.HistogramRequest_20270101 != nil && obj.HistogramRequest_20270101.UnparsedObject == nil {
			jsonHistogramRequest_20270101, _ := datadog.Marshal(obj.HistogramRequest_20270101)
			if string(jsonHistogramRequest_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.HistogramRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.HistogramRequest_20270101 = nil
		}
	} else {
		obj.HistogramRequest_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AllListRequestVariants_20270101 = nil
		obj.TabularAdvancedQueryRequest_20270101 = nil
		obj.TimeseriesFormulaRequest_20270101 = nil
		obj.ScalarFormulaRequest_20270101 = nil
		obj.TabularLocalDatasetRequest_20270101 = nil
		obj.HistogramRequest_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TabularDataRequest) MarshalJSON() ([]byte, error) {
	if obj.AllListRequestVariants_20270101 != nil {
		return datadog.Marshal(&obj.AllListRequestVariants_20270101)
	}

	if obj.TabularAdvancedQueryRequest_20270101 != nil {
		return datadog.Marshal(&obj.TabularAdvancedQueryRequest_20270101)
	}

	if obj.TimeseriesFormulaRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesFormulaRequest_20270101)
	}

	if obj.ScalarFormulaRequest_20270101 != nil {
		return datadog.Marshal(&obj.ScalarFormulaRequest_20270101)
	}

	if obj.TabularLocalDatasetRequest_20270101 != nil {
		return datadog.Marshal(&obj.TabularLocalDatasetRequest_20270101)
	}

	if obj.HistogramRequest_20270101 != nil {
		return datadog.Marshal(&obj.HistogramRequest_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TabularDataRequest) GetActualInstance() interface{} {
	if obj.AllListRequestVariants_20270101 != nil {
		return obj.AllListRequestVariants_20270101
	}

	if obj.TabularAdvancedQueryRequest_20270101 != nil {
		return obj.TabularAdvancedQueryRequest_20270101
	}

	if obj.TimeseriesFormulaRequest_20270101 != nil {
		return obj.TimeseriesFormulaRequest_20270101
	}

	if obj.ScalarFormulaRequest_20270101 != nil {
		return obj.ScalarFormulaRequest_20270101
	}

	if obj.TabularLocalDatasetRequest_20270101 != nil {
		return obj.TabularLocalDatasetRequest_20270101
	}

	if obj.HistogramRequest_20270101 != nil {
		return obj.HistogramRequest_20270101
	}

	// all schemas are nil
	return nil
}

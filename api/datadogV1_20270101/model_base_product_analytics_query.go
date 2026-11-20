// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BaseProductAnalyticsQuery - Test-only generated schema for the dashboard versioning proof of concept.
type BaseProductAnalyticsQuery struct {
	ProductAnalyticsEventQuery_20270101           *ProductAnalyticsEventQuery
	ProductAnalyticsEventOccurrenceQuery_20270101 *ProductAnalyticsEventOccurrenceQuery

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ProductAnalyticsEventQuery_20270101AsBaseProductAnalyticsQuery is a convenience function that returns ProductAnalyticsEventQuery wrapped in BaseProductAnalyticsQuery.
func ProductAnalyticsEventQuery_20270101AsBaseProductAnalyticsQuery(v *ProductAnalyticsEventQuery) BaseProductAnalyticsQuery {
	return BaseProductAnalyticsQuery{ProductAnalyticsEventQuery_20270101: v}
}

// ProductAnalyticsEventOccurrenceQuery_20270101AsBaseProductAnalyticsQuery is a convenience function that returns ProductAnalyticsEventOccurrenceQuery wrapped in BaseProductAnalyticsQuery.
func ProductAnalyticsEventOccurrenceQuery_20270101AsBaseProductAnalyticsQuery(v *ProductAnalyticsEventOccurrenceQuery) BaseProductAnalyticsQuery {
	return BaseProductAnalyticsQuery{ProductAnalyticsEventOccurrenceQuery_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *BaseProductAnalyticsQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ProductAnalyticsEventQuery_20270101
	err = datadog.Unmarshal(data, &obj.ProductAnalyticsEventQuery_20270101)
	if err == nil {
		if obj.ProductAnalyticsEventQuery_20270101 != nil && obj.ProductAnalyticsEventQuery_20270101.UnparsedObject == nil {
			jsonProductAnalyticsEventQuery_20270101, _ := datadog.Marshal(obj.ProductAnalyticsEventQuery_20270101)
			if string(jsonProductAnalyticsEventQuery_20270101) == "{}" { // empty struct
				obj.ProductAnalyticsEventQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ProductAnalyticsEventQuery_20270101 = nil
		}
	} else {
		obj.ProductAnalyticsEventQuery_20270101 = nil
	}

	// try to unmarshal data into ProductAnalyticsEventOccurrenceQuery_20270101
	err = datadog.Unmarshal(data, &obj.ProductAnalyticsEventOccurrenceQuery_20270101)
	if err == nil {
		if obj.ProductAnalyticsEventOccurrenceQuery_20270101 != nil && obj.ProductAnalyticsEventOccurrenceQuery_20270101.UnparsedObject == nil {
			jsonProductAnalyticsEventOccurrenceQuery_20270101, _ := datadog.Marshal(obj.ProductAnalyticsEventOccurrenceQuery_20270101)
			if string(jsonProductAnalyticsEventOccurrenceQuery_20270101) == "{}" { // empty struct
				obj.ProductAnalyticsEventOccurrenceQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ProductAnalyticsEventOccurrenceQuery_20270101 = nil
		}
	} else {
		obj.ProductAnalyticsEventOccurrenceQuery_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ProductAnalyticsEventQuery_20270101 = nil
		obj.ProductAnalyticsEventOccurrenceQuery_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj BaseProductAnalyticsQuery) MarshalJSON() ([]byte, error) {
	if obj.ProductAnalyticsEventQuery_20270101 != nil {
		return datadog.Marshal(&obj.ProductAnalyticsEventQuery_20270101)
	}

	if obj.ProductAnalyticsEventOccurrenceQuery_20270101 != nil {
		return datadog.Marshal(&obj.ProductAnalyticsEventOccurrenceQuery_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *BaseProductAnalyticsQuery) GetActualInstance() interface{} {
	if obj.ProductAnalyticsEventQuery_20270101 != nil {
		return obj.ProductAnalyticsEventQuery_20270101
	}

	if obj.ProductAnalyticsEventOccurrenceQuery_20270101 != nil {
		return obj.ProductAnalyticsEventOccurrenceQuery_20270101
	}

	// all schemas are nil
	return nil
}

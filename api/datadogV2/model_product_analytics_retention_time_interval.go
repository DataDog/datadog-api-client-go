// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionTimeInterval - A retention interval, either aligned to calendar boundaries or of a fixed length.
// Cohort criteria use calendar intervals; return criteria use fixed intervals.
type ProductAnalyticsRetentionTimeInterval struct {
	ProductAnalyticsRetentionCalendarTimeInterval *ProductAnalyticsRetentionCalendarTimeInterval
	ProductAnalyticsRetentionFixedTimeInterval    *ProductAnalyticsRetentionFixedTimeInterval

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ProductAnalyticsRetentionCalendarTimeIntervalAsProductAnalyticsRetentionTimeInterval is a convenience function that returns ProductAnalyticsRetentionCalendarTimeInterval wrapped in ProductAnalyticsRetentionTimeInterval.
func ProductAnalyticsRetentionCalendarTimeIntervalAsProductAnalyticsRetentionTimeInterval(v *ProductAnalyticsRetentionCalendarTimeInterval) ProductAnalyticsRetentionTimeInterval {
	return ProductAnalyticsRetentionTimeInterval{ProductAnalyticsRetentionCalendarTimeInterval: v}
}

// ProductAnalyticsRetentionFixedTimeIntervalAsProductAnalyticsRetentionTimeInterval is a convenience function that returns ProductAnalyticsRetentionFixedTimeInterval wrapped in ProductAnalyticsRetentionTimeInterval.
func ProductAnalyticsRetentionFixedTimeIntervalAsProductAnalyticsRetentionTimeInterval(v *ProductAnalyticsRetentionFixedTimeInterval) ProductAnalyticsRetentionTimeInterval {
	return ProductAnalyticsRetentionTimeInterval{ProductAnalyticsRetentionFixedTimeInterval: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ProductAnalyticsRetentionTimeInterval) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ProductAnalyticsRetentionCalendarTimeInterval
	err = datadog.Unmarshal(data, &obj.ProductAnalyticsRetentionCalendarTimeInterval)
	if err == nil {
		if obj.ProductAnalyticsRetentionCalendarTimeInterval != nil && obj.ProductAnalyticsRetentionCalendarTimeInterval.UnparsedObject == nil {
			jsonProductAnalyticsRetentionCalendarTimeInterval, _ := datadog.Marshal(obj.ProductAnalyticsRetentionCalendarTimeInterval)
			if string(jsonProductAnalyticsRetentionCalendarTimeInterval) == "{}" { // empty struct
				obj.ProductAnalyticsRetentionCalendarTimeInterval = nil
			} else {
				match++
			}
		} else {
			obj.ProductAnalyticsRetentionCalendarTimeInterval = nil
		}
	} else {
		obj.ProductAnalyticsRetentionCalendarTimeInterval = nil
	}

	// try to unmarshal data into ProductAnalyticsRetentionFixedTimeInterval
	err = datadog.Unmarshal(data, &obj.ProductAnalyticsRetentionFixedTimeInterval)
	if err == nil {
		if obj.ProductAnalyticsRetentionFixedTimeInterval != nil && obj.ProductAnalyticsRetentionFixedTimeInterval.UnparsedObject == nil {
			jsonProductAnalyticsRetentionFixedTimeInterval, _ := datadog.Marshal(obj.ProductAnalyticsRetentionFixedTimeInterval)
			if string(jsonProductAnalyticsRetentionFixedTimeInterval) == "{}" { // empty struct
				obj.ProductAnalyticsRetentionFixedTimeInterval = nil
			} else {
				match++
			}
		} else {
			obj.ProductAnalyticsRetentionFixedTimeInterval = nil
		}
	} else {
		obj.ProductAnalyticsRetentionFixedTimeInterval = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ProductAnalyticsRetentionCalendarTimeInterval = nil
		obj.ProductAnalyticsRetentionFixedTimeInterval = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ProductAnalyticsRetentionTimeInterval) MarshalJSON() ([]byte, error) {
	if obj.ProductAnalyticsRetentionCalendarTimeInterval != nil {
		return datadog.Marshal(&obj.ProductAnalyticsRetentionCalendarTimeInterval)
	}

	if obj.ProductAnalyticsRetentionFixedTimeInterval != nil {
		return datadog.Marshal(&obj.ProductAnalyticsRetentionFixedTimeInterval)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ProductAnalyticsRetentionTimeInterval) GetActualInstance() interface{} {
	if obj.ProductAnalyticsRetentionCalendarTimeInterval != nil {
		return obj.ProductAnalyticsRetentionCalendarTimeInterval
	}

	if obj.ProductAnalyticsRetentionFixedTimeInterval != nil {
		return obj.ProductAnalyticsRetentionFixedTimeInterval
	}

	// all schemas are nil
	return nil
}

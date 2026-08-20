// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionCohortTarget - Selects a cohort, either by index or by the aggregation that rolls all cohorts together.
type ProductAnalyticsRetentionCohortTarget struct {
	ProductAnalyticsRetentionIndexTarget       *ProductAnalyticsRetentionIndexTarget
	ProductAnalyticsRetentionAggregationTarget *ProductAnalyticsRetentionAggregationTarget

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ProductAnalyticsRetentionIndexTargetAsProductAnalyticsRetentionCohortTarget is a convenience function that returns ProductAnalyticsRetentionIndexTarget wrapped in ProductAnalyticsRetentionCohortTarget.
func ProductAnalyticsRetentionIndexTargetAsProductAnalyticsRetentionCohortTarget(v *ProductAnalyticsRetentionIndexTarget) ProductAnalyticsRetentionCohortTarget {
	return ProductAnalyticsRetentionCohortTarget{ProductAnalyticsRetentionIndexTarget: v}
}

// ProductAnalyticsRetentionAggregationTargetAsProductAnalyticsRetentionCohortTarget is a convenience function that returns ProductAnalyticsRetentionAggregationTarget wrapped in ProductAnalyticsRetentionCohortTarget.
func ProductAnalyticsRetentionAggregationTargetAsProductAnalyticsRetentionCohortTarget(v *ProductAnalyticsRetentionAggregationTarget) ProductAnalyticsRetentionCohortTarget {
	return ProductAnalyticsRetentionCohortTarget{ProductAnalyticsRetentionAggregationTarget: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ProductAnalyticsRetentionCohortTarget) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ProductAnalyticsRetentionIndexTarget
	err = datadog.Unmarshal(data, &obj.ProductAnalyticsRetentionIndexTarget)
	if err == nil {
		if obj.ProductAnalyticsRetentionIndexTarget != nil && obj.ProductAnalyticsRetentionIndexTarget.UnparsedObject == nil {
			jsonProductAnalyticsRetentionIndexTarget, _ := datadog.Marshal(obj.ProductAnalyticsRetentionIndexTarget)
			if string(jsonProductAnalyticsRetentionIndexTarget) == "{}" { // empty struct
				obj.ProductAnalyticsRetentionIndexTarget = nil
			} else {
				match++
			}
		} else {
			obj.ProductAnalyticsRetentionIndexTarget = nil
		}
	} else {
		obj.ProductAnalyticsRetentionIndexTarget = nil
	}

	// try to unmarshal data into ProductAnalyticsRetentionAggregationTarget
	err = datadog.Unmarshal(data, &obj.ProductAnalyticsRetentionAggregationTarget)
	if err == nil {
		if obj.ProductAnalyticsRetentionAggregationTarget != nil && obj.ProductAnalyticsRetentionAggregationTarget.UnparsedObject == nil {
			jsonProductAnalyticsRetentionAggregationTarget, _ := datadog.Marshal(obj.ProductAnalyticsRetentionAggregationTarget)
			if string(jsonProductAnalyticsRetentionAggregationTarget) == "{}" { // empty struct
				obj.ProductAnalyticsRetentionAggregationTarget = nil
			} else {
				match++
			}
		} else {
			obj.ProductAnalyticsRetentionAggregationTarget = nil
		}
	} else {
		obj.ProductAnalyticsRetentionAggregationTarget = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ProductAnalyticsRetentionIndexTarget = nil
		obj.ProductAnalyticsRetentionAggregationTarget = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ProductAnalyticsRetentionCohortTarget) MarshalJSON() ([]byte, error) {
	if obj.ProductAnalyticsRetentionIndexTarget != nil {
		return datadog.Marshal(&obj.ProductAnalyticsRetentionIndexTarget)
	}

	if obj.ProductAnalyticsRetentionAggregationTarget != nil {
		return datadog.Marshal(&obj.ProductAnalyticsRetentionAggregationTarget)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ProductAnalyticsRetentionCohortTarget) GetActualInstance() interface{} {
	if obj.ProductAnalyticsRetentionIndexTarget != nil {
		return obj.ProductAnalyticsRetentionIndexTarget
	}

	if obj.ProductAnalyticsRetentionAggregationTarget != nil {
		return obj.ProductAnalyticsRetentionAggregationTarget
	}

	// all schemas are nil
	return nil
}

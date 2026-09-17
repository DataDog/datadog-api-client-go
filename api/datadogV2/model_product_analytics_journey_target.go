// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyTarget - A reference to a step, or a range of steps, in the journey.
// Use a `node` target to name a single step, or a `path` target to name the range
// between two steps.
type ProductAnalyticsJourneyTarget struct {
	ProductAnalyticsJourneyNodeTarget *ProductAnalyticsJourneyNodeTarget
	ProductAnalyticsJourneyPathTarget *ProductAnalyticsJourneyPathTarget

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ProductAnalyticsJourneyNodeTargetAsProductAnalyticsJourneyTarget is a convenience function that returns ProductAnalyticsJourneyNodeTarget wrapped in ProductAnalyticsJourneyTarget.
func ProductAnalyticsJourneyNodeTargetAsProductAnalyticsJourneyTarget(v *ProductAnalyticsJourneyNodeTarget) ProductAnalyticsJourneyTarget {
	return ProductAnalyticsJourneyTarget{ProductAnalyticsJourneyNodeTarget: v}
}

// ProductAnalyticsJourneyPathTargetAsProductAnalyticsJourneyTarget is a convenience function that returns ProductAnalyticsJourneyPathTarget wrapped in ProductAnalyticsJourneyTarget.
func ProductAnalyticsJourneyPathTargetAsProductAnalyticsJourneyTarget(v *ProductAnalyticsJourneyPathTarget) ProductAnalyticsJourneyTarget {
	return ProductAnalyticsJourneyTarget{ProductAnalyticsJourneyPathTarget: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ProductAnalyticsJourneyTarget) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ProductAnalyticsJourneyNodeTarget
	err = datadog.Unmarshal(data, &obj.ProductAnalyticsJourneyNodeTarget)
	if err == nil {
		if obj.ProductAnalyticsJourneyNodeTarget != nil && obj.ProductAnalyticsJourneyNodeTarget.UnparsedObject == nil {
			jsonProductAnalyticsJourneyNodeTarget, _ := datadog.Marshal(obj.ProductAnalyticsJourneyNodeTarget)
			if string(jsonProductAnalyticsJourneyNodeTarget) == "{}" { // empty struct
				obj.ProductAnalyticsJourneyNodeTarget = nil
			} else {
				match++
			}
		} else {
			obj.ProductAnalyticsJourneyNodeTarget = nil
		}
	} else {
		obj.ProductAnalyticsJourneyNodeTarget = nil
	}

	// try to unmarshal data into ProductAnalyticsJourneyPathTarget
	err = datadog.Unmarshal(data, &obj.ProductAnalyticsJourneyPathTarget)
	if err == nil {
		if obj.ProductAnalyticsJourneyPathTarget != nil && obj.ProductAnalyticsJourneyPathTarget.UnparsedObject == nil {
			jsonProductAnalyticsJourneyPathTarget, _ := datadog.Marshal(obj.ProductAnalyticsJourneyPathTarget)
			if string(jsonProductAnalyticsJourneyPathTarget) == "{}" { // empty struct
				obj.ProductAnalyticsJourneyPathTarget = nil
			} else {
				match++
			}
		} else {
			obj.ProductAnalyticsJourneyPathTarget = nil
		}
	} else {
		obj.ProductAnalyticsJourneyPathTarget = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ProductAnalyticsJourneyNodeTarget = nil
		obj.ProductAnalyticsJourneyPathTarget = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ProductAnalyticsJourneyTarget) MarshalJSON() ([]byte, error) {
	if obj.ProductAnalyticsJourneyNodeTarget != nil {
		return datadog.Marshal(&obj.ProductAnalyticsJourneyNodeTarget)
	}

	if obj.ProductAnalyticsJourneyPathTarget != nil {
		return datadog.Marshal(&obj.ProductAnalyticsJourneyPathTarget)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ProductAnalyticsJourneyTarget) GetActualInstance() interface{} {
	if obj.ProductAnalyticsJourneyNodeTarget != nil {
		return obj.ProductAnalyticsJourneyNodeTarget
	}

	if obj.ProductAnalyticsJourneyPathTarget != nil {
		return obj.ProductAnalyticsJourneyPathTarget
	}

	// all schemas are nil
	return nil
}

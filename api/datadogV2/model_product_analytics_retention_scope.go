// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionScope - Restricts a retention query to part of the grid, so that results can be examined in detail.
// Omit it to compute the whole grid.
type ProductAnalyticsRetentionScope struct {
	ProductAnalyticsRetentionCohortScope       *ProductAnalyticsRetentionCohortScope
	ProductAnalyticsRetentionReturnPeriodScope *ProductAnalyticsRetentionReturnPeriodScope
	ProductAnalyticsRetentionCellScope         *ProductAnalyticsRetentionCellScope

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ProductAnalyticsRetentionCohortScopeAsProductAnalyticsRetentionScope is a convenience function that returns ProductAnalyticsRetentionCohortScope wrapped in ProductAnalyticsRetentionScope.
func ProductAnalyticsRetentionCohortScopeAsProductAnalyticsRetentionScope(v *ProductAnalyticsRetentionCohortScope) ProductAnalyticsRetentionScope {
	return ProductAnalyticsRetentionScope{ProductAnalyticsRetentionCohortScope: v}
}

// ProductAnalyticsRetentionReturnPeriodScopeAsProductAnalyticsRetentionScope is a convenience function that returns ProductAnalyticsRetentionReturnPeriodScope wrapped in ProductAnalyticsRetentionScope.
func ProductAnalyticsRetentionReturnPeriodScopeAsProductAnalyticsRetentionScope(v *ProductAnalyticsRetentionReturnPeriodScope) ProductAnalyticsRetentionScope {
	return ProductAnalyticsRetentionScope{ProductAnalyticsRetentionReturnPeriodScope: v}
}

// ProductAnalyticsRetentionCellScopeAsProductAnalyticsRetentionScope is a convenience function that returns ProductAnalyticsRetentionCellScope wrapped in ProductAnalyticsRetentionScope.
func ProductAnalyticsRetentionCellScopeAsProductAnalyticsRetentionScope(v *ProductAnalyticsRetentionCellScope) ProductAnalyticsRetentionScope {
	return ProductAnalyticsRetentionScope{ProductAnalyticsRetentionCellScope: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ProductAnalyticsRetentionScope) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ProductAnalyticsRetentionCohortScope
	err = datadog.Unmarshal(data, &obj.ProductAnalyticsRetentionCohortScope)
	if err == nil {
		if obj.ProductAnalyticsRetentionCohortScope != nil && obj.ProductAnalyticsRetentionCohortScope.UnparsedObject == nil {
			jsonProductAnalyticsRetentionCohortScope, _ := datadog.Marshal(obj.ProductAnalyticsRetentionCohortScope)
			if string(jsonProductAnalyticsRetentionCohortScope) == "{}" { // empty struct
				obj.ProductAnalyticsRetentionCohortScope = nil
			} else {
				match++
			}
		} else {
			obj.ProductAnalyticsRetentionCohortScope = nil
		}
	} else {
		obj.ProductAnalyticsRetentionCohortScope = nil
	}

	// try to unmarshal data into ProductAnalyticsRetentionReturnPeriodScope
	err = datadog.Unmarshal(data, &obj.ProductAnalyticsRetentionReturnPeriodScope)
	if err == nil {
		if obj.ProductAnalyticsRetentionReturnPeriodScope != nil && obj.ProductAnalyticsRetentionReturnPeriodScope.UnparsedObject == nil {
			jsonProductAnalyticsRetentionReturnPeriodScope, _ := datadog.Marshal(obj.ProductAnalyticsRetentionReturnPeriodScope)
			if string(jsonProductAnalyticsRetentionReturnPeriodScope) == "{}" { // empty struct
				obj.ProductAnalyticsRetentionReturnPeriodScope = nil
			} else {
				match++
			}
		} else {
			obj.ProductAnalyticsRetentionReturnPeriodScope = nil
		}
	} else {
		obj.ProductAnalyticsRetentionReturnPeriodScope = nil
	}

	// try to unmarshal data into ProductAnalyticsRetentionCellScope
	err = datadog.Unmarshal(data, &obj.ProductAnalyticsRetentionCellScope)
	if err == nil {
		if obj.ProductAnalyticsRetentionCellScope != nil && obj.ProductAnalyticsRetentionCellScope.UnparsedObject == nil {
			jsonProductAnalyticsRetentionCellScope, _ := datadog.Marshal(obj.ProductAnalyticsRetentionCellScope)
			if string(jsonProductAnalyticsRetentionCellScope) == "{}" { // empty struct
				obj.ProductAnalyticsRetentionCellScope = nil
			} else {
				match++
			}
		} else {
			obj.ProductAnalyticsRetentionCellScope = nil
		}
	} else {
		obj.ProductAnalyticsRetentionCellScope = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ProductAnalyticsRetentionCohortScope = nil
		obj.ProductAnalyticsRetentionReturnPeriodScope = nil
		obj.ProductAnalyticsRetentionCellScope = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ProductAnalyticsRetentionScope) MarshalJSON() ([]byte, error) {
	if obj.ProductAnalyticsRetentionCohortScope != nil {
		return datadog.Marshal(&obj.ProductAnalyticsRetentionCohortScope)
	}

	if obj.ProductAnalyticsRetentionReturnPeriodScope != nil {
		return datadog.Marshal(&obj.ProductAnalyticsRetentionReturnPeriodScope)
	}

	if obj.ProductAnalyticsRetentionCellScope != nil {
		return datadog.Marshal(&obj.ProductAnalyticsRetentionCellScope)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ProductAnalyticsRetentionScope) GetActualInstance() interface{} {
	if obj.ProductAnalyticsRetentionCohortScope != nil {
		return obj.ProductAnalyticsRetentionCohortScope
	}

	if obj.ProductAnalyticsRetentionReturnPeriodScope != nil {
		return obj.ProductAnalyticsRetentionReturnPeriodScope
	}

	if obj.ProductAnalyticsRetentionCellScope != nil {
		return obj.ProductAnalyticsRetentionCellScope
	}

	// all schemas are nil
	return nil
}

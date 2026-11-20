// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmRecommendationsStreamRequestColumnsItems20270101Width - Test-only generated schema for the dashboard versioning proof of concept.
type ApmRecommendationsStreamRequestColumnsItems20270101Width struct {
	Float64                          *float64
	DashboardInlineSchema12_20270101 *DashboardInlineSchema12_20270101

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// Float64AsApmRecommendationsStreamRequestColumnsItems20270101Width is a convenience function that returns float64 wrapped in ApmRecommendationsStreamRequestColumnsItems20270101Width.
func Float64AsApmRecommendationsStreamRequestColumnsItems20270101Width(v *float64) ApmRecommendationsStreamRequestColumnsItems20270101Width {
	return ApmRecommendationsStreamRequestColumnsItems20270101Width{Float64: v}
}

// DashboardInlineSchema12_20270101AsApmRecommendationsStreamRequestColumnsItems20270101Width is a convenience function that returns DashboardInlineSchema12_20270101 wrapped in ApmRecommendationsStreamRequestColumnsItems20270101Width.
func DashboardInlineSchema12_20270101AsApmRecommendationsStreamRequestColumnsItems20270101Width(v *DashboardInlineSchema12_20270101) ApmRecommendationsStreamRequestColumnsItems20270101Width {
	return ApmRecommendationsStreamRequestColumnsItems20270101Width{DashboardInlineSchema12_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ApmRecommendationsStreamRequestColumnsItems20270101Width) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Float64
	err = datadog.Unmarshal(data, &obj.Float64)
	if err == nil {
		if obj.Float64 != nil {
			jsonFloat64, _ := datadog.Marshal(obj.Float64)
			if string(jsonFloat64) == "{}" { // empty struct
				obj.Float64 = nil
			} else {
				match++
			}
		} else {
			obj.Float64 = nil
		}
	} else {
		obj.Float64 = nil
	}

	// try to unmarshal data into DashboardInlineSchema12_20270101
	err = datadog.Unmarshal(data, &obj.DashboardInlineSchema12_20270101)
	if err == nil {
		if obj.DashboardInlineSchema12_20270101 != nil {
			jsonDashboardInlineSchema12_20270101, _ := datadog.Marshal(obj.DashboardInlineSchema12_20270101)
			if string(jsonDashboardInlineSchema12_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.DashboardInlineSchema12_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DashboardInlineSchema12_20270101 = nil
		}
	} else {
		obj.DashboardInlineSchema12_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.Float64 = nil
		obj.DashboardInlineSchema12_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ApmRecommendationsStreamRequestColumnsItems20270101Width) MarshalJSON() ([]byte, error) {
	if obj.Float64 != nil {
		return datadog.Marshal(&obj.Float64)
	}

	if obj.DashboardInlineSchema12_20270101 != nil {
		return datadog.Marshal(&obj.DashboardInlineSchema12_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ApmRecommendationsStreamRequestColumnsItems20270101Width) GetActualInstance() interface{} {
	if obj.Float64 != nil {
		return obj.Float64
	}

	if obj.DashboardInlineSchema12_20270101 != nil {
		return obj.DashboardInlineSchema12_20270101
	}

	// all schemas are nil
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation - Measure and calculation settings for a numerator or denominator aggregation. Supply exactly one non-null measure.
type ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation struct {
	ExperimentsWarehouseMetricAggregationInput *ExperimentsWarehouseMetricAggregationInput
	ExperimentsDatadogMetricAggregationInput   *ExperimentsDatadogMetricAggregationInput

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ExperimentsWarehouseMetricAggregationInputAsExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation is a convenience function that returns ExperimentsWarehouseMetricAggregationInput wrapped in ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation.
func ExperimentsWarehouseMetricAggregationInputAsExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation(v *ExperimentsWarehouseMetricAggregationInput) ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation {
	return ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation{ExperimentsWarehouseMetricAggregationInput: v}
}

// ExperimentsDatadogMetricAggregationInputAsExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation is a convenience function that returns ExperimentsDatadogMetricAggregationInput wrapped in ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation.
func ExperimentsDatadogMetricAggregationInputAsExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation(v *ExperimentsDatadogMetricAggregationInput) ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation {
	return ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation{ExperimentsDatadogMetricAggregationInput: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ExperimentsWarehouseMetricAggregationInput
	err = datadog.Unmarshal(data, &obj.ExperimentsWarehouseMetricAggregationInput)
	if err == nil {
		if obj.ExperimentsWarehouseMetricAggregationInput != nil && obj.ExperimentsWarehouseMetricAggregationInput.UnparsedObject == nil {
			jsonExperimentsWarehouseMetricAggregationInput, _ := datadog.Marshal(obj.ExperimentsWarehouseMetricAggregationInput)
			if string(jsonExperimentsWarehouseMetricAggregationInput) == "{}" { // empty struct
				obj.ExperimentsWarehouseMetricAggregationInput = nil
			} else {
				match++
			}
		} else {
			obj.ExperimentsWarehouseMetricAggregationInput = nil
		}
	} else {
		obj.ExperimentsWarehouseMetricAggregationInput = nil
	}

	// try to unmarshal data into ExperimentsDatadogMetricAggregationInput
	err = datadog.Unmarshal(data, &obj.ExperimentsDatadogMetricAggregationInput)
	if err == nil {
		if obj.ExperimentsDatadogMetricAggregationInput != nil && obj.ExperimentsDatadogMetricAggregationInput.UnparsedObject == nil {
			jsonExperimentsDatadogMetricAggregationInput, _ := datadog.Marshal(obj.ExperimentsDatadogMetricAggregationInput)
			if string(jsonExperimentsDatadogMetricAggregationInput) == "{}" { // empty struct
				obj.ExperimentsDatadogMetricAggregationInput = nil
			} else {
				match++
			}
		} else {
			obj.ExperimentsDatadogMetricAggregationInput = nil
		}
	} else {
		obj.ExperimentsDatadogMetricAggregationInput = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ExperimentsWarehouseMetricAggregationInput = nil
		obj.ExperimentsDatadogMetricAggregationInput = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation) MarshalJSON() ([]byte, error) {
	if obj.ExperimentsWarehouseMetricAggregationInput != nil {
		return datadog.Marshal(&obj.ExperimentsWarehouseMetricAggregationInput)
	}

	if obj.ExperimentsDatadogMetricAggregationInput != nil {
		return datadog.Marshal(&obj.ExperimentsDatadogMetricAggregationInput)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation) GetActualInstance() interface{} {
	if obj.ExperimentsWarehouseMetricAggregationInput != nil {
		return obj.ExperimentsWarehouseMetricAggregationInput
	}

	if obj.ExperimentsDatadogMetricAggregationInput != nil {
		return obj.ExperimentsDatadogMetricAggregationInput
	}

	// all schemas are nil
	return nil
}

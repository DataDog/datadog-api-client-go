// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation - Measure and percentile to calculate for the metric. Supply exactly one non-null measure.
type ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation struct {
	ExperimentsWarehousePercentileAggregationInput *ExperimentsWarehousePercentileAggregationInput
	ExperimentsDatadogPercentileAggregationInput   *ExperimentsDatadogPercentileAggregationInput

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ExperimentsWarehousePercentileAggregationInputAsExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation is a convenience function that returns ExperimentsWarehousePercentileAggregationInput wrapped in ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation.
func ExperimentsWarehousePercentileAggregationInputAsExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation(v *ExperimentsWarehousePercentileAggregationInput) ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation {
	return ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation{ExperimentsWarehousePercentileAggregationInput: v}
}

// ExperimentsDatadogPercentileAggregationInputAsExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation is a convenience function that returns ExperimentsDatadogPercentileAggregationInput wrapped in ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation.
func ExperimentsDatadogPercentileAggregationInputAsExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation(v *ExperimentsDatadogPercentileAggregationInput) ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation {
	return ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation{ExperimentsDatadogPercentileAggregationInput: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ExperimentsWarehousePercentileAggregationInput
	err = datadog.Unmarshal(data, &obj.ExperimentsWarehousePercentileAggregationInput)
	if err == nil {
		if obj.ExperimentsWarehousePercentileAggregationInput != nil && obj.ExperimentsWarehousePercentileAggregationInput.UnparsedObject == nil {
			jsonExperimentsWarehousePercentileAggregationInput, _ := datadog.Marshal(obj.ExperimentsWarehousePercentileAggregationInput)
			if string(jsonExperimentsWarehousePercentileAggregationInput) == "{}" { // empty struct
				obj.ExperimentsWarehousePercentileAggregationInput = nil
			} else {
				match++
			}
		} else {
			obj.ExperimentsWarehousePercentileAggregationInput = nil
		}
	} else {
		obj.ExperimentsWarehousePercentileAggregationInput = nil
	}

	// try to unmarshal data into ExperimentsDatadogPercentileAggregationInput
	err = datadog.Unmarshal(data, &obj.ExperimentsDatadogPercentileAggregationInput)
	if err == nil {
		if obj.ExperimentsDatadogPercentileAggregationInput != nil && obj.ExperimentsDatadogPercentileAggregationInput.UnparsedObject == nil {
			jsonExperimentsDatadogPercentileAggregationInput, _ := datadog.Marshal(obj.ExperimentsDatadogPercentileAggregationInput)
			if string(jsonExperimentsDatadogPercentileAggregationInput) == "{}" { // empty struct
				obj.ExperimentsDatadogPercentileAggregationInput = nil
			} else {
				match++
			}
		} else {
			obj.ExperimentsDatadogPercentileAggregationInput = nil
		}
	} else {
		obj.ExperimentsDatadogPercentileAggregationInput = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ExperimentsWarehousePercentileAggregationInput = nil
		obj.ExperimentsDatadogPercentileAggregationInput = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation) MarshalJSON() ([]byte, error) {
	if obj.ExperimentsWarehousePercentileAggregationInput != nil {
		return datadog.Marshal(&obj.ExperimentsWarehousePercentileAggregationInput)
	}

	if obj.ExperimentsDatadogPercentileAggregationInput != nil {
		return datadog.Marshal(&obj.ExperimentsDatadogPercentileAggregationInput)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation) GetActualInstance() interface{} {
	if obj.ExperimentsWarehousePercentileAggregationInput != nil {
		return obj.ExperimentsWarehousePercentileAggregationInput
	}

	if obj.ExperimentsDatadogPercentileAggregationInput != nil {
		return obj.ExperimentsDatadogPercentileAggregationInput
	}

	// all schemas are nil
	return nil
}

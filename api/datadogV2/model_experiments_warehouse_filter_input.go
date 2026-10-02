// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsWarehouseFilterInput - A property or measure comparison for a Warehouse numerator or denominator. Set exactly one target ID that is not blank. Measure comparisons require numeric measures and numeric values. BETWEEN bounds must be in ascending order.
type ExperimentsWarehouseFilterInput struct {
	ExperimentsPropertyFilterInput          *ExperimentsPropertyFilterInput
	ExperimentsPropertyNullFilterInput      *ExperimentsPropertyNullFilterInput
	ExperimentsMeasureComparisonFilterInput *ExperimentsMeasureComparisonFilterInput
	ExperimentsMeasureRangeFilterInput      *ExperimentsMeasureRangeFilterInput
	ExperimentsMeasureNullFilterInput       *ExperimentsMeasureNullFilterInput

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ExperimentsPropertyFilterInputAsExperimentsWarehouseFilterInput is a convenience function that returns ExperimentsPropertyFilterInput wrapped in ExperimentsWarehouseFilterInput.
func ExperimentsPropertyFilterInputAsExperimentsWarehouseFilterInput(v *ExperimentsPropertyFilterInput) ExperimentsWarehouseFilterInput {
	return ExperimentsWarehouseFilterInput{ExperimentsPropertyFilterInput: v}
}

// ExperimentsPropertyNullFilterInputAsExperimentsWarehouseFilterInput is a convenience function that returns ExperimentsPropertyNullFilterInput wrapped in ExperimentsWarehouseFilterInput.
func ExperimentsPropertyNullFilterInputAsExperimentsWarehouseFilterInput(v *ExperimentsPropertyNullFilterInput) ExperimentsWarehouseFilterInput {
	return ExperimentsWarehouseFilterInput{ExperimentsPropertyNullFilterInput: v}
}

// ExperimentsMeasureComparisonFilterInputAsExperimentsWarehouseFilterInput is a convenience function that returns ExperimentsMeasureComparisonFilterInput wrapped in ExperimentsWarehouseFilterInput.
func ExperimentsMeasureComparisonFilterInputAsExperimentsWarehouseFilterInput(v *ExperimentsMeasureComparisonFilterInput) ExperimentsWarehouseFilterInput {
	return ExperimentsWarehouseFilterInput{ExperimentsMeasureComparisonFilterInput: v}
}

// ExperimentsMeasureRangeFilterInputAsExperimentsWarehouseFilterInput is a convenience function that returns ExperimentsMeasureRangeFilterInput wrapped in ExperimentsWarehouseFilterInput.
func ExperimentsMeasureRangeFilterInputAsExperimentsWarehouseFilterInput(v *ExperimentsMeasureRangeFilterInput) ExperimentsWarehouseFilterInput {
	return ExperimentsWarehouseFilterInput{ExperimentsMeasureRangeFilterInput: v}
}

// ExperimentsMeasureNullFilterInputAsExperimentsWarehouseFilterInput is a convenience function that returns ExperimentsMeasureNullFilterInput wrapped in ExperimentsWarehouseFilterInput.
func ExperimentsMeasureNullFilterInputAsExperimentsWarehouseFilterInput(v *ExperimentsMeasureNullFilterInput) ExperimentsWarehouseFilterInput {
	return ExperimentsWarehouseFilterInput{ExperimentsMeasureNullFilterInput: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ExperimentsWarehouseFilterInput) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ExperimentsPropertyFilterInput
	err = datadog.Unmarshal(data, &obj.ExperimentsPropertyFilterInput)
	if err == nil {
		if obj.ExperimentsPropertyFilterInput != nil && obj.ExperimentsPropertyFilterInput.UnparsedObject == nil {
			jsonExperimentsPropertyFilterInput, _ := datadog.Marshal(obj.ExperimentsPropertyFilterInput)
			if string(jsonExperimentsPropertyFilterInput) == "{}" { // empty struct
				obj.ExperimentsPropertyFilterInput = nil
			} else {
				match++
			}
		} else {
			obj.ExperimentsPropertyFilterInput = nil
		}
	} else {
		obj.ExperimentsPropertyFilterInput = nil
	}

	// try to unmarshal data into ExperimentsPropertyNullFilterInput
	err = datadog.Unmarshal(data, &obj.ExperimentsPropertyNullFilterInput)
	if err == nil {
		if obj.ExperimentsPropertyNullFilterInput != nil && obj.ExperimentsPropertyNullFilterInput.UnparsedObject == nil {
			jsonExperimentsPropertyNullFilterInput, _ := datadog.Marshal(obj.ExperimentsPropertyNullFilterInput)
			if string(jsonExperimentsPropertyNullFilterInput) == "{}" { // empty struct
				obj.ExperimentsPropertyNullFilterInput = nil
			} else {
				match++
			}
		} else {
			obj.ExperimentsPropertyNullFilterInput = nil
		}
	} else {
		obj.ExperimentsPropertyNullFilterInput = nil
	}

	// try to unmarshal data into ExperimentsMeasureComparisonFilterInput
	err = datadog.Unmarshal(data, &obj.ExperimentsMeasureComparisonFilterInput)
	if err == nil {
		if obj.ExperimentsMeasureComparisonFilterInput != nil && obj.ExperimentsMeasureComparisonFilterInput.UnparsedObject == nil {
			jsonExperimentsMeasureComparisonFilterInput, _ := datadog.Marshal(obj.ExperimentsMeasureComparisonFilterInput)
			if string(jsonExperimentsMeasureComparisonFilterInput) == "{}" { // empty struct
				obj.ExperimentsMeasureComparisonFilterInput = nil
			} else {
				match++
			}
		} else {
			obj.ExperimentsMeasureComparisonFilterInput = nil
		}
	} else {
		obj.ExperimentsMeasureComparisonFilterInput = nil
	}

	// try to unmarshal data into ExperimentsMeasureRangeFilterInput
	err = datadog.Unmarshal(data, &obj.ExperimentsMeasureRangeFilterInput)
	if err == nil {
		if obj.ExperimentsMeasureRangeFilterInput != nil && obj.ExperimentsMeasureRangeFilterInput.UnparsedObject == nil {
			jsonExperimentsMeasureRangeFilterInput, _ := datadog.Marshal(obj.ExperimentsMeasureRangeFilterInput)
			if string(jsonExperimentsMeasureRangeFilterInput) == "{}" { // empty struct
				obj.ExperimentsMeasureRangeFilterInput = nil
			} else {
				match++
			}
		} else {
			obj.ExperimentsMeasureRangeFilterInput = nil
		}
	} else {
		obj.ExperimentsMeasureRangeFilterInput = nil
	}

	// try to unmarshal data into ExperimentsMeasureNullFilterInput
	err = datadog.Unmarshal(data, &obj.ExperimentsMeasureNullFilterInput)
	if err == nil {
		if obj.ExperimentsMeasureNullFilterInput != nil && obj.ExperimentsMeasureNullFilterInput.UnparsedObject == nil {
			jsonExperimentsMeasureNullFilterInput, _ := datadog.Marshal(obj.ExperimentsMeasureNullFilterInput)
			if string(jsonExperimentsMeasureNullFilterInput) == "{}" { // empty struct
				obj.ExperimentsMeasureNullFilterInput = nil
			} else {
				match++
			}
		} else {
			obj.ExperimentsMeasureNullFilterInput = nil
		}
	} else {
		obj.ExperimentsMeasureNullFilterInput = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ExperimentsPropertyFilterInput = nil
		obj.ExperimentsPropertyNullFilterInput = nil
		obj.ExperimentsMeasureComparisonFilterInput = nil
		obj.ExperimentsMeasureRangeFilterInput = nil
		obj.ExperimentsMeasureNullFilterInput = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ExperimentsWarehouseFilterInput) MarshalJSON() ([]byte, error) {
	if obj.ExperimentsPropertyFilterInput != nil {
		return datadog.Marshal(&obj.ExperimentsPropertyFilterInput)
	}

	if obj.ExperimentsPropertyNullFilterInput != nil {
		return datadog.Marshal(&obj.ExperimentsPropertyNullFilterInput)
	}

	if obj.ExperimentsMeasureComparisonFilterInput != nil {
		return datadog.Marshal(&obj.ExperimentsMeasureComparisonFilterInput)
	}

	if obj.ExperimentsMeasureRangeFilterInput != nil {
		return datadog.Marshal(&obj.ExperimentsMeasureRangeFilterInput)
	}

	if obj.ExperimentsMeasureNullFilterInput != nil {
		return datadog.Marshal(&obj.ExperimentsMeasureNullFilterInput)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ExperimentsWarehouseFilterInput) GetActualInstance() interface{} {
	if obj.ExperimentsPropertyFilterInput != nil {
		return obj.ExperimentsPropertyFilterInput
	}

	if obj.ExperimentsPropertyNullFilterInput != nil {
		return obj.ExperimentsPropertyNullFilterInput
	}

	if obj.ExperimentsMeasureComparisonFilterInput != nil {
		return obj.ExperimentsMeasureComparisonFilterInput
	}

	if obj.ExperimentsMeasureRangeFilterInput != nil {
		return obj.ExperimentsMeasureRangeFilterInput
	}

	if obj.ExperimentsMeasureNullFilterInput != nil {
		return obj.ExperimentsMeasureNullFilterInput
	}

	// all schemas are nil
	return nil
}

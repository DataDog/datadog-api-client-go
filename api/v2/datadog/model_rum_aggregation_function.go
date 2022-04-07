/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
	"fmt"
)

// RUMAggregationFunction An aggregation function.
type RUMAggregationFunction string

// List of RUMAggregationFunction
const (
	RUMAGGREGATIONFUNCTION_COUNT         RUMAggregationFunction = "count"
	RUMAGGREGATIONFUNCTION_CARDINALITY   RUMAggregationFunction = "cardinality"
	RUMAGGREGATIONFUNCTION_PERCENTILE_75 RUMAggregationFunction = "pc75"
	RUMAGGREGATIONFUNCTION_PERCENTILE_90 RUMAggregationFunction = "pc90"
	RUMAGGREGATIONFUNCTION_PERCENTILE_95 RUMAggregationFunction = "pc95"
	RUMAGGREGATIONFUNCTION_PERCENTILE_98 RUMAggregationFunction = "pc98"
	RUMAGGREGATIONFUNCTION_PERCENTILE_99 RUMAggregationFunction = "pc99"
	RUMAGGREGATIONFUNCTION_SUM           RUMAggregationFunction = "sum"
	RUMAGGREGATIONFUNCTION_MIN           RUMAggregationFunction = "min"
	RUMAGGREGATIONFUNCTION_MAX           RUMAggregationFunction = "max"
	RUMAGGREGATIONFUNCTION_AVG           RUMAggregationFunction = "avg"
	RUMAGGREGATIONFUNCTION_MEDIAN        RUMAggregationFunction = "median"
)

var allowedRUMAggregationFunctionEnumValues = []RUMAggregationFunction{
	RUMAGGREGATIONFUNCTION_COUNT,
	RUMAGGREGATIONFUNCTION_CARDINALITY,
	RUMAGGREGATIONFUNCTION_PERCENTILE_75,
	RUMAGGREGATIONFUNCTION_PERCENTILE_90,
	RUMAGGREGATIONFUNCTION_PERCENTILE_95,
	RUMAGGREGATIONFUNCTION_PERCENTILE_98,
	RUMAGGREGATIONFUNCTION_PERCENTILE_99,
	RUMAGGREGATIONFUNCTION_SUM,
	RUMAGGREGATIONFUNCTION_MIN,
	RUMAGGREGATIONFUNCTION_MAX,
	RUMAGGREGATIONFUNCTION_AVG,
	RUMAGGREGATIONFUNCTION_MEDIAN,
}

func (w *RUMAggregationFunction) GetAllowedValues() []RUMAggregationFunction {
	return allowedRUMAggregationFunctionEnumValues
}

func (v *RUMAggregationFunction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMAggregationFunction(value)
	return nil
}

// NewRUMAggregationFunctionFromValue returns a pointer to a valid RUMAggregationFunction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRUMAggregationFunctionFromValue(v string) (*RUMAggregationFunction, error) {
	ev := RUMAggregationFunction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RUMAggregationFunction: valid values are %v", v, allowedRUMAggregationFunctionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RUMAggregationFunction) IsValid() bool {
	for _, existing := range allowedRUMAggregationFunctionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMAggregationFunction value
func (v RUMAggregationFunction) Ptr() *RUMAggregationFunction {
	return &v
}

type NullableRUMAggregationFunction struct {
	value *RUMAggregationFunction
	isSet bool
}

func (v NullableRUMAggregationFunction) Get() *RUMAggregationFunction {
	return v.value
}

func (v *NullableRUMAggregationFunction) Set(val *RUMAggregationFunction) {
	v.value = val
	v.isSet = true
}

func (v NullableRUMAggregationFunction) IsSet() bool {
	return v.isSet
}

func (v *NullableRUMAggregationFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRUMAggregationFunction(val *RUMAggregationFunction) *NullableRUMAggregationFunction {
	return &NullableRUMAggregationFunction{value: val, isSet: true}
}

func (v NullableRUMAggregationFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRUMAggregationFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

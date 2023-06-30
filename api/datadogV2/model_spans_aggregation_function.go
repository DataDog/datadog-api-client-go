// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// SpansAggregationFunction An aggregation function.
type SpansAggregationFunction string

// List of SpansAggregationFunction.
const (
	SPANSAGGREGATIONFUNCTION_COUNT         SpansAggregationFunction = "count"
	SPANSAGGREGATIONFUNCTION_CARDINALITY   SpansAggregationFunction = "cardinality"
	SPANSAGGREGATIONFUNCTION_PERCENTILE_75 SpansAggregationFunction = "pc75"
	SPANSAGGREGATIONFUNCTION_PERCENTILE_90 SpansAggregationFunction = "pc90"
	SPANSAGGREGATIONFUNCTION_PERCENTILE_95 SpansAggregationFunction = "pc95"
	SPANSAGGREGATIONFUNCTION_PERCENTILE_98 SpansAggregationFunction = "pc98"
	SPANSAGGREGATIONFUNCTION_PERCENTILE_99 SpansAggregationFunction = "pc99"
	SPANSAGGREGATIONFUNCTION_SUM           SpansAggregationFunction = "sum"
	SPANSAGGREGATIONFUNCTION_MIN           SpansAggregationFunction = "min"
	SPANSAGGREGATIONFUNCTION_MAX           SpansAggregationFunction = "max"
	SPANSAGGREGATIONFUNCTION_AVG           SpansAggregationFunction = "avg"
	SPANSAGGREGATIONFUNCTION_MEDIAN        SpansAggregationFunction = "median"
)

var allowedSpansAggregationFunctionEnumValues = []SpansAggregationFunction{
	SPANSAGGREGATIONFUNCTION_COUNT,
	SPANSAGGREGATIONFUNCTION_CARDINALITY,
	SPANSAGGREGATIONFUNCTION_PERCENTILE_75,
	SPANSAGGREGATIONFUNCTION_PERCENTILE_90,
	SPANSAGGREGATIONFUNCTION_PERCENTILE_95,
	SPANSAGGREGATIONFUNCTION_PERCENTILE_98,
	SPANSAGGREGATIONFUNCTION_PERCENTILE_99,
	SPANSAGGREGATIONFUNCTION_SUM,
	SPANSAGGREGATIONFUNCTION_MIN,
	SPANSAGGREGATIONFUNCTION_MAX,
	SPANSAGGREGATIONFUNCTION_AVG,
	SPANSAGGREGATIONFUNCTION_MEDIAN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpansAggregationFunction) GetAllowedValues() []SpansAggregationFunction {
	return allowedSpansAggregationFunctionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpansAggregationFunction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpansAggregationFunction(value)
	return nil
}

// NewSpansAggregationFunctionFromValue returns a pointer to a valid SpansAggregationFunction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpansAggregationFunctionFromValue(v string) (*SpansAggregationFunction, error) {
	ev := SpansAggregationFunction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpansAggregationFunction: valid values are %v", v, allowedSpansAggregationFunctionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpansAggregationFunction) IsValid() bool {
	for _, existing := range allowedSpansAggregationFunctionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpansAggregationFunction value.
func (v SpansAggregationFunction) Ptr() *SpansAggregationFunction {
	return &v
}

// NullableSpansAggregationFunction handles when a null is used for SpansAggregationFunction.
type NullableSpansAggregationFunction struct {
	value *SpansAggregationFunction
	isSet bool
}

// Get returns the associated value.
func (v NullableSpansAggregationFunction) Get() *SpansAggregationFunction {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableSpansAggregationFunction) Set(val *SpansAggregationFunction) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableSpansAggregationFunction) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableSpansAggregationFunction) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSpansAggregationFunction initializes the struct as if Set has been called.
func NewNullableSpansAggregationFunction(val *SpansAggregationFunction) *NullableSpansAggregationFunction {
	return &NullableSpansAggregationFunction{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableSpansAggregationFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableSpansAggregationFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

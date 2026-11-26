// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScalarAggregator Test-only generated schema for the dashboard versioning proof of concept.
type ScalarAggregator string

// List of ScalarAggregator.
const (
	SCALARAGGREGATOR_SUM          ScalarAggregator = "sum"
	SCALARAGGREGATOR_MIN          ScalarAggregator = "min"
	SCALARAGGREGATOR_MAX          ScalarAggregator = "max"
	SCALARAGGREGATOR_AVG          ScalarAggregator = "avg"
	SCALARAGGREGATOR_LAST         ScalarAggregator = "last"
	SCALARAGGREGATOR_AREA         ScalarAggregator = "area"
	SCALARAGGREGATOR_L2NORM       ScalarAggregator = "l2norm"
	SCALARAGGREGATOR_PERCENTILE   ScalarAggregator = "percentile"
	SCALARAGGREGATOR_STDDEV       ScalarAggregator = "stddev"
	SCALARAGGREGATOR_COUNT_UNIQUE ScalarAggregator = "count_unique"
)

var allowedScalarAggregatorEnumValues = []ScalarAggregator{
	SCALARAGGREGATOR_SUM,
	SCALARAGGREGATOR_MIN,
	SCALARAGGREGATOR_MAX,
	SCALARAGGREGATOR_AVG,
	SCALARAGGREGATOR_LAST,
	SCALARAGGREGATOR_AREA,
	SCALARAGGREGATOR_L2NORM,
	SCALARAGGREGATOR_PERCENTILE,
	SCALARAGGREGATOR_STDDEV,
	SCALARAGGREGATOR_COUNT_UNIQUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScalarAggregator) GetAllowedValues() []ScalarAggregator {
	return allowedScalarAggregatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScalarAggregator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScalarAggregator(value)
	return nil
}

// NewScalarAggregatorFromValue returns a pointer to a valid ScalarAggregator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScalarAggregatorFromValue(v string) (*ScalarAggregator, error) {
	ev := ScalarAggregator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScalarAggregator: valid values are %v", v, allowedScalarAggregatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScalarAggregator) IsValid() bool {
	for _, existing := range allowedScalarAggregatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScalarAggregator value.
func (v ScalarAggregator) Ptr() *ScalarAggregator {
	return &v
}

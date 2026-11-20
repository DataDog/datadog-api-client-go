// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RunChartValueDimensionDimension Test-only generated schema for the dashboard versioning proof of concept.
type RunChartValueDimensionDimension string

// List of RunChartValueDimensionDimension.
const (
	RUNCHARTVALUEDIMENSIONDIMENSION_VALUE RunChartValueDimensionDimension = "value"
)

var allowedRunChartValueDimensionDimensionEnumValues = []RunChartValueDimensionDimension{
	RUNCHARTVALUEDIMENSIONDIMENSION_VALUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RunChartValueDimensionDimension) GetAllowedValues() []RunChartValueDimensionDimension {
	return allowedRunChartValueDimensionDimensionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RunChartValueDimensionDimension) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RunChartValueDimensionDimension(value)
	return nil
}

// NewRunChartValueDimensionDimensionFromValue returns a pointer to a valid RunChartValueDimensionDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRunChartValueDimensionDimensionFromValue(v string) (*RunChartValueDimensionDimension, error) {
	ev := RunChartValueDimensionDimension(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RunChartValueDimensionDimension: valid values are %v", v, allowedRunChartValueDimensionDimensionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RunChartValueDimensionDimension) IsValid() bool {
	for _, existing := range allowedRunChartValueDimensionDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RunChartValueDimensionDimension value.
func (v RunChartValueDimensionDimension) Ptr() *RunChartValueDimensionDimension {
	return &v
}

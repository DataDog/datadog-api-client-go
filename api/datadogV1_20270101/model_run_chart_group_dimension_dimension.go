// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RunChartGroupDimensionDimension Test-only generated schema for the dashboard versioning proof of concept.
type RunChartGroupDimensionDimension string

// List of RunChartGroupDimensionDimension.
const (
	RUNCHARTGROUPDIMENSIONDIMENSION_GROUP RunChartGroupDimensionDimension = "group"
)

var allowedRunChartGroupDimensionDimensionEnumValues = []RunChartGroupDimensionDimension{
	RUNCHARTGROUPDIMENSIONDIMENSION_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RunChartGroupDimensionDimension) GetAllowedValues() []RunChartGroupDimensionDimension {
	return allowedRunChartGroupDimensionDimensionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RunChartGroupDimensionDimension) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RunChartGroupDimensionDimension(value)
	return nil
}

// NewRunChartGroupDimensionDimensionFromValue returns a pointer to a valid RunChartGroupDimensionDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRunChartGroupDimensionDimensionFromValue(v string) (*RunChartGroupDimensionDimension, error) {
	ev := RunChartGroupDimensionDimension(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RunChartGroupDimensionDimension: valid values are %v", v, allowedRunChartGroupDimensionDimensionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RunChartGroupDimensionDimension) IsValid() bool {
	for _, existing := range allowedRunChartGroupDimensionDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RunChartGroupDimensionDimension value.
func (v RunChartGroupDimensionDimension) Ptr() *RunChartGroupDimensionDimension {
	return &v
}

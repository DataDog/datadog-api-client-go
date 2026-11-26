// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RunChartProjectionType Test-only generated schema for the dashboard versioning proof of concept.
type RunChartProjectionType string

// List of RunChartProjectionType.
const (
	RUNCHARTPROJECTIONTYPE_RUN_CHART RunChartProjectionType = "run_chart"
)

var allowedRunChartProjectionTypeEnumValues = []RunChartProjectionType{
	RUNCHARTPROJECTIONTYPE_RUN_CHART,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RunChartProjectionType) GetAllowedValues() []RunChartProjectionType {
	return allowedRunChartProjectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RunChartProjectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RunChartProjectionType(value)
	return nil
}

// NewRunChartProjectionTypeFromValue returns a pointer to a valid RunChartProjectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRunChartProjectionTypeFromValue(v string) (*RunChartProjectionType, error) {
	ev := RunChartProjectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RunChartProjectionType: valid values are %v", v, allowedRunChartProjectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RunChartProjectionType) IsValid() bool {
	for _, existing := range allowedRunChartProjectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RunChartProjectionType value.
func (v RunChartProjectionType) Ptr() *RunChartProjectionType {
	return &v
}

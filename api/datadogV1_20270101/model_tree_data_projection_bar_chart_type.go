// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreeDataProjectionBarChartType Test-only generated schema for the dashboard versioning proof of concept.
type TreeDataProjectionBarChartType string

// List of TreeDataProjectionBarChartType.
const (
	TREEDATAPROJECTIONBARCHARTTYPE_BAR_CHART TreeDataProjectionBarChartType = "bar_chart"
)

var allowedTreeDataProjectionBarChartTypeEnumValues = []TreeDataProjectionBarChartType{
	TREEDATAPROJECTIONBARCHARTTYPE_BAR_CHART,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreeDataProjectionBarChartType) GetAllowedValues() []TreeDataProjectionBarChartType {
	return allowedTreeDataProjectionBarChartTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreeDataProjectionBarChartType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreeDataProjectionBarChartType(value)
	return nil
}

// NewTreeDataProjectionBarChartTypeFromValue returns a pointer to a valid TreeDataProjectionBarChartType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreeDataProjectionBarChartTypeFromValue(v string) (*TreeDataProjectionBarChartType, error) {
	ev := TreeDataProjectionBarChartType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreeDataProjectionBarChartType: valid values are %v", v, allowedTreeDataProjectionBarChartTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreeDataProjectionBarChartType) IsValid() bool {
	for _, existing := range allowedTreeDataProjectionBarChartTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreeDataProjectionBarChartType value.
func (v TreeDataProjectionBarChartType) Ptr() *TreeDataProjectionBarChartType {
	return &v
}

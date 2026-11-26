// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreeDataProjectionBarChartDimensionsItemsDimension Test-only generated schema for the dashboard versioning proof of concept.
type TreeDataProjectionBarChartDimensionsItemsDimension string

// List of TreeDataProjectionBarChartDimensionsItemsDimension.
const (
	TREEDATAPROJECTIONBARCHARTDIMENSIONSITEMSDIMENSION_GROUP TreeDataProjectionBarChartDimensionsItemsDimension = "group"
	TREEDATAPROJECTIONBARCHARTDIMENSIONSITEMSDIMENSION_VALUE TreeDataProjectionBarChartDimensionsItemsDimension = "value"
)

var allowedTreeDataProjectionBarChartDimensionsItemsDimensionEnumValues = []TreeDataProjectionBarChartDimensionsItemsDimension{
	TREEDATAPROJECTIONBARCHARTDIMENSIONSITEMSDIMENSION_GROUP,
	TREEDATAPROJECTIONBARCHARTDIMENSIONSITEMSDIMENSION_VALUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreeDataProjectionBarChartDimensionsItemsDimension) GetAllowedValues() []TreeDataProjectionBarChartDimensionsItemsDimension {
	return allowedTreeDataProjectionBarChartDimensionsItemsDimensionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreeDataProjectionBarChartDimensionsItemsDimension) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreeDataProjectionBarChartDimensionsItemsDimension(value)
	return nil
}

// NewTreeDataProjectionBarChartDimensionsItemsDimensionFromValue returns a pointer to a valid TreeDataProjectionBarChartDimensionsItemsDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreeDataProjectionBarChartDimensionsItemsDimensionFromValue(v string) (*TreeDataProjectionBarChartDimensionsItemsDimension, error) {
	ev := TreeDataProjectionBarChartDimensionsItemsDimension(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreeDataProjectionBarChartDimensionsItemsDimension: valid values are %v", v, allowedTreeDataProjectionBarChartDimensionsItemsDimensionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreeDataProjectionBarChartDimensionsItemsDimension) IsValid() bool {
	for _, existing := range allowedTreeDataProjectionBarChartDimensionsItemsDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreeDataProjectionBarChartDimensionsItemsDimension value.
func (v TreeDataProjectionBarChartDimensionsItemsDimension) Ptr() *TreeDataProjectionBarChartDimensionsItemsDimension {
	return &v
}

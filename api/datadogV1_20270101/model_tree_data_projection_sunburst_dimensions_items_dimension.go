// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreeDataProjectionSunburstDimensionsItemsDimension Test-only generated schema for the dashboard versioning proof of concept.
type TreeDataProjectionSunburstDimensionsItemsDimension string

// List of TreeDataProjectionSunburstDimensionsItemsDimension.
const (
	TREEDATAPROJECTIONSUNBURSTDIMENSIONSITEMSDIMENSION_GROUP TreeDataProjectionSunburstDimensionsItemsDimension = "group"
	TREEDATAPROJECTIONSUNBURSTDIMENSIONSITEMSDIMENSION_VALUE TreeDataProjectionSunburstDimensionsItemsDimension = "value"
)

var allowedTreeDataProjectionSunburstDimensionsItemsDimensionEnumValues = []TreeDataProjectionSunburstDimensionsItemsDimension{
	TREEDATAPROJECTIONSUNBURSTDIMENSIONSITEMSDIMENSION_GROUP,
	TREEDATAPROJECTIONSUNBURSTDIMENSIONSITEMSDIMENSION_VALUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreeDataProjectionSunburstDimensionsItemsDimension) GetAllowedValues() []TreeDataProjectionSunburstDimensionsItemsDimension {
	return allowedTreeDataProjectionSunburstDimensionsItemsDimensionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreeDataProjectionSunburstDimensionsItemsDimension) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreeDataProjectionSunburstDimensionsItemsDimension(value)
	return nil
}

// NewTreeDataProjectionSunburstDimensionsItemsDimensionFromValue returns a pointer to a valid TreeDataProjectionSunburstDimensionsItemsDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreeDataProjectionSunburstDimensionsItemsDimensionFromValue(v string) (*TreeDataProjectionSunburstDimensionsItemsDimension, error) {
	ev := TreeDataProjectionSunburstDimensionsItemsDimension(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreeDataProjectionSunburstDimensionsItemsDimension: valid values are %v", v, allowedTreeDataProjectionSunburstDimensionsItemsDimensionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreeDataProjectionSunburstDimensionsItemsDimension) IsValid() bool {
	for _, existing := range allowedTreeDataProjectionSunburstDimensionsItemsDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreeDataProjectionSunburstDimensionsItemsDimension value.
func (v TreeDataProjectionSunburstDimensionsItemsDimension) Ptr() *TreeDataProjectionSunburstDimensionsItemsDimension {
	return &v
}

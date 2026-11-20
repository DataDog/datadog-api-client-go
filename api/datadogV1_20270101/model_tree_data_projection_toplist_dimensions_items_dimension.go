// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreeDataProjectionToplistDimensionsItemsDimension Test-only generated schema for the dashboard versioning proof of concept.
type TreeDataProjectionToplistDimensionsItemsDimension string

// List of TreeDataProjectionToplistDimensionsItemsDimension.
const (
	TREEDATAPROJECTIONTOPLISTDIMENSIONSITEMSDIMENSION_GROUP TreeDataProjectionToplistDimensionsItemsDimension = "group"
	TREEDATAPROJECTIONTOPLISTDIMENSIONSITEMSDIMENSION_VALUE TreeDataProjectionToplistDimensionsItemsDimension = "value"
)

var allowedTreeDataProjectionToplistDimensionsItemsDimensionEnumValues = []TreeDataProjectionToplistDimensionsItemsDimension{
	TREEDATAPROJECTIONTOPLISTDIMENSIONSITEMSDIMENSION_GROUP,
	TREEDATAPROJECTIONTOPLISTDIMENSIONSITEMSDIMENSION_VALUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreeDataProjectionToplistDimensionsItemsDimension) GetAllowedValues() []TreeDataProjectionToplistDimensionsItemsDimension {
	return allowedTreeDataProjectionToplistDimensionsItemsDimensionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreeDataProjectionToplistDimensionsItemsDimension) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreeDataProjectionToplistDimensionsItemsDimension(value)
	return nil
}

// NewTreeDataProjectionToplistDimensionsItemsDimensionFromValue returns a pointer to a valid TreeDataProjectionToplistDimensionsItemsDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreeDataProjectionToplistDimensionsItemsDimensionFromValue(v string) (*TreeDataProjectionToplistDimensionsItemsDimension, error) {
	ev := TreeDataProjectionToplistDimensionsItemsDimension(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreeDataProjectionToplistDimensionsItemsDimension: valid values are %v", v, allowedTreeDataProjectionToplistDimensionsItemsDimensionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreeDataProjectionToplistDimensionsItemsDimension) IsValid() bool {
	for _, existing := range allowedTreeDataProjectionToplistDimensionsItemsDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreeDataProjectionToplistDimensionsItemsDimension value.
func (v TreeDataProjectionToplistDimensionsItemsDimension) Ptr() *TreeDataProjectionToplistDimensionsItemsDimension {
	return &v
}

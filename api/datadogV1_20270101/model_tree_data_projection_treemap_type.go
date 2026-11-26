// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreeDataProjectionTreemapType Test-only generated schema for the dashboard versioning proof of concept.
type TreeDataProjectionTreemapType string

// List of TreeDataProjectionTreemapType.
const (
	TREEDATAPROJECTIONTREEMAPTYPE_TREEMAP TreeDataProjectionTreemapType = "treemap"
)

var allowedTreeDataProjectionTreemapTypeEnumValues = []TreeDataProjectionTreemapType{
	TREEDATAPROJECTIONTREEMAPTYPE_TREEMAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreeDataProjectionTreemapType) GetAllowedValues() []TreeDataProjectionTreemapType {
	return allowedTreeDataProjectionTreemapTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreeDataProjectionTreemapType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreeDataProjectionTreemapType(value)
	return nil
}

// NewTreeDataProjectionTreemapTypeFromValue returns a pointer to a valid TreeDataProjectionTreemapType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreeDataProjectionTreemapTypeFromValue(v string) (*TreeDataProjectionTreemapType, error) {
	ev := TreeDataProjectionTreemapType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreeDataProjectionTreemapType: valid values are %v", v, allowedTreeDataProjectionTreemapTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreeDataProjectionTreemapType) IsValid() bool {
	for _, existing := range allowedTreeDataProjectionTreemapTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreeDataProjectionTreemapType value.
func (v TreeDataProjectionTreemapType) Ptr() *TreeDataProjectionTreemapType {
	return &v
}

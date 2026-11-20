// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreeDataProjectionToplistType Test-only generated schema for the dashboard versioning proof of concept.
type TreeDataProjectionToplistType string

// List of TreeDataProjectionToplistType.
const (
	TREEDATAPROJECTIONTOPLISTTYPE_TOPLIST TreeDataProjectionToplistType = "toplist"
)

var allowedTreeDataProjectionToplistTypeEnumValues = []TreeDataProjectionToplistType{
	TREEDATAPROJECTIONTOPLISTTYPE_TOPLIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreeDataProjectionToplistType) GetAllowedValues() []TreeDataProjectionToplistType {
	return allowedTreeDataProjectionToplistTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreeDataProjectionToplistType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreeDataProjectionToplistType(value)
	return nil
}

// NewTreeDataProjectionToplistTypeFromValue returns a pointer to a valid TreeDataProjectionToplistType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreeDataProjectionToplistTypeFromValue(v string) (*TreeDataProjectionToplistType, error) {
	ev := TreeDataProjectionToplistType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreeDataProjectionToplistType: valid values are %v", v, allowedTreeDataProjectionToplistTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreeDataProjectionToplistType) IsValid() bool {
	for _, existing := range allowedTreeDataProjectionToplistTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreeDataProjectionToplistType value.
func (v TreeDataProjectionToplistType) Ptr() *TreeDataProjectionToplistType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GraphTargetPathType Test-only generated schema for the dashboard versioning proof of concept.
type GraphTargetPathType string

// List of GraphTargetPathType.
const (
	GRAPHTARGETPATHTYPE_PATH GraphTargetPathType = "path"
)

var allowedGraphTargetPathTypeEnumValues = []GraphTargetPathType{
	GRAPHTARGETPATHTYPE_PATH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GraphTargetPathType) GetAllowedValues() []GraphTargetPathType {
	return allowedGraphTargetPathTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GraphTargetPathType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GraphTargetPathType(value)
	return nil
}

// NewGraphTargetPathTypeFromValue returns a pointer to a valid GraphTargetPathType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGraphTargetPathTypeFromValue(v string) (*GraphTargetPathType, error) {
	ev := GraphTargetPathType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GraphTargetPathType: valid values are %v", v, allowedGraphTargetPathTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GraphTargetPathType) IsValid() bool {
	for _, existing := range allowedGraphTargetPathTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GraphTargetPathType value.
func (v GraphTargetPathType) Ptr() *GraphTargetPathType {
	return &v
}

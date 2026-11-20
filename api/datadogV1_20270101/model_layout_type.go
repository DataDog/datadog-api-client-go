// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LayoutType Test-only generated schema for the dashboard versioning proof of concept.
type LayoutType string

// List of LayoutType.
const (
	LAYOUTTYPE_ORDERED LayoutType = "ordered"
	LAYOUTTYPE_FREE    LayoutType = "free"
)

var allowedLayoutTypeEnumValues = []LayoutType{
	LAYOUTTYPE_ORDERED,
	LAYOUTTYPE_FREE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LayoutType) GetAllowedValues() []LayoutType {
	return allowedLayoutTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LayoutType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LayoutType(value)
	return nil
}

// NewLayoutTypeFromValue returns a pointer to a valid LayoutType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLayoutTypeFromValue(v string) (*LayoutType, error) {
	ev := LayoutType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LayoutType: valid values are %v", v, allowedLayoutTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LayoutType) IsValid() bool {
	for _, existing := range allowedLayoutTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LayoutType value.
func (v LayoutType) Ptr() *LayoutType {
	return &v
}

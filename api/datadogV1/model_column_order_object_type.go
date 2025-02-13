// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ColumnOrderObjectType type of column
type ColumnOrderObjectType string

// List of ColumnOrderObjectType.
const (
	COLUMNORDEROBJECTTYPE_FORMULA ColumnOrderObjectType = "formula"
	COLUMNORDEROBJECTTYPE_GROUP   ColumnOrderObjectType = "group"
)

var allowedColumnOrderObjectTypeEnumValues = []ColumnOrderObjectType{
	COLUMNORDEROBJECTTYPE_FORMULA,
	COLUMNORDEROBJECTTYPE_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ColumnOrderObjectType) GetAllowedValues() []ColumnOrderObjectType {
	return allowedColumnOrderObjectTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ColumnOrderObjectType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ColumnOrderObjectType(value)
	return nil
}

// NewColumnOrderObjectTypeFromValue returns a pointer to a valid ColumnOrderObjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewColumnOrderObjectTypeFromValue(v string) (*ColumnOrderObjectType, error) {
	ev := ColumnOrderObjectType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ColumnOrderObjectType: valid values are %v", v, allowedColumnOrderObjectTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ColumnOrderObjectType) IsValid() bool {
	for _, existing := range allowedColumnOrderObjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ColumnOrderObjectType value.
func (v ColumnOrderObjectType) Ptr() *ColumnOrderObjectType {
	return &v
}

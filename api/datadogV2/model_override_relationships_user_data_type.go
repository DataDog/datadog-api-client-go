// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OverrideRelationshipsUserDataType Indicates that the related resource is of type 'users'.
type OverrideRelationshipsUserDataType string

// List of OverrideRelationshipsUserDataType.
const (
	OVERRIDERELATIONSHIPSUSERDATATYPE_USERS OverrideRelationshipsUserDataType = "users"
)

var allowedOverrideRelationshipsUserDataTypeEnumValues = []OverrideRelationshipsUserDataType{
	OVERRIDERELATIONSHIPSUSERDATATYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OverrideRelationshipsUserDataType) GetAllowedValues() []OverrideRelationshipsUserDataType {
	return allowedOverrideRelationshipsUserDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OverrideRelationshipsUserDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OverrideRelationshipsUserDataType(value)
	return nil
}

// NewOverrideRelationshipsUserDataTypeFromValue returns a pointer to a valid OverrideRelationshipsUserDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOverrideRelationshipsUserDataTypeFromValue(v string) (*OverrideRelationshipsUserDataType, error) {
	ev := OverrideRelationshipsUserDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OverrideRelationshipsUserDataType: valid values are %v", v, allowedOverrideRelationshipsUserDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OverrideRelationshipsUserDataType) IsValid() bool {
	for _, existing := range allowedOverrideRelationshipsUserDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OverrideRelationshipsUserDataType value.
func (v OverrideRelationshipsUserDataType) Ptr() *OverrideRelationshipsUserDataType {
	return &v
}

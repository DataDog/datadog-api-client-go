// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VariantDataType The resource type for variants.
type VariantDataType string

// List of VariantDataType.
const (
	VARIANTDATATYPE_VARIANTS VariantDataType = "variants"
)

var allowedVariantDataTypeEnumValues = []VariantDataType{
	VARIANTDATATYPE_VARIANTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *VariantDataType) GetAllowedValues() []VariantDataType {
	return allowedVariantDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *VariantDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = VariantDataType(value)
	return nil
}

// NewVariantDataTypeFromValue returns a pointer to a valid VariantDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewVariantDataTypeFromValue(v string) (*VariantDataType, error) {
	ev := VariantDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for VariantDataType: valid values are %v", v, allowedVariantDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v VariantDataType) IsValid() bool {
	for _, existing := range allowedVariantDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VariantDataType value.
func (v VariantDataType) Ptr() *VariantDataType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ASMExclusionFilterType The type of the resource. The value should always be `exclusion_filter`.
type ASMExclusionFilterType string

// List of ASMExclusionFilterType.
const (
	ASMEXCLUSIONFILTERTYPE_EXCLUSION_FILTER ASMExclusionFilterType = "exclusion_filter"
)

var allowedASMExclusionFilterTypeEnumValues = []ASMExclusionFilterType{
	ASMEXCLUSIONFILTERTYPE_EXCLUSION_FILTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ASMExclusionFilterType) GetAllowedValues() []ASMExclusionFilterType {
	return allowedASMExclusionFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ASMExclusionFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ASMExclusionFilterType(value)
	return nil
}

// NewASMExclusionFilterTypeFromValue returns a pointer to a valid ASMExclusionFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewASMExclusionFilterTypeFromValue(v string) (*ASMExclusionFilterType, error) {
	ev := ASMExclusionFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ASMExclusionFilterType: valid values are %v", v, allowedASMExclusionFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ASMExclusionFilterType) IsValid() bool {
	for _, existing := range allowedASMExclusionFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ASMExclusionFilterType value.
func (v ASMExclusionFilterType) Ptr() *ASMExclusionFilterType {
	return &v
}

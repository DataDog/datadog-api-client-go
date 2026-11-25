// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMExclusionFilterType The resource type. The value must be `exclusion_filters`.
type RUMExclusionFilterType string

// List of RUMExclusionFilterType.
const (
	RUMEXCLUSIONFILTERTYPE_EXCLUSION_FILTERS RUMExclusionFilterType = "exclusion_filters"
)

var allowedRUMExclusionFilterTypeEnumValues = []RUMExclusionFilterType{
	RUMEXCLUSIONFILTERTYPE_EXCLUSION_FILTERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMExclusionFilterType) GetAllowedValues() []RUMExclusionFilterType {
	return allowedRUMExclusionFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMExclusionFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMExclusionFilterType(value)
	return nil
}

// NewRUMExclusionFilterTypeFromValue returns a pointer to a valid RUMExclusionFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMExclusionFilterTypeFromValue(v string) (*RUMExclusionFilterType, error) {
	ev := RUMExclusionFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMExclusionFilterType: valid values are %v", v, allowedRUMExclusionFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMExclusionFilterType) IsValid() bool {
	for _, existing := range allowedRUMExclusionFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMExclusionFilterType value.
func (v RUMExclusionFilterType) Ptr() *RUMExclusionFilterType {
	return &v
}

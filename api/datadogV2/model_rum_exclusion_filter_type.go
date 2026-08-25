// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumExclusionFilterType The resource type. The value must be `exclusion_filters`.
type RumExclusionFilterType string

// List of RumExclusionFilterType.
const (
	RUMEXCLUSIONFILTERTYPE_EXCLUSION_FILTERS RumExclusionFilterType = "exclusion_filters"
)

var allowedRumExclusionFilterTypeEnumValues = []RumExclusionFilterType{
	RUMEXCLUSIONFILTERTYPE_EXCLUSION_FILTERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumExclusionFilterType) GetAllowedValues() []RumExclusionFilterType {
	return allowedRumExclusionFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumExclusionFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumExclusionFilterType(value)
	return nil
}

// NewRumExclusionFilterTypeFromValue returns a pointer to a valid RumExclusionFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumExclusionFilterTypeFromValue(v string) (*RumExclusionFilterType, error) {
	ev := RumExclusionFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumExclusionFilterType: valid values are %v", v, allowedRumExclusionFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumExclusionFilterType) IsValid() bool {
	for _, existing := range allowedRumExclusionFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumExclusionFilterType value.
func (v RumExclusionFilterType) Ptr() *RumExclusionFilterType {
	return &v
}

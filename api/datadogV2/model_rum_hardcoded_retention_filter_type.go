// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumHardcodedRetentionFilterType The resource type. The value must be `hardcoded_retention_filters`.
type RumHardcodedRetentionFilterType string

// List of RumHardcodedRetentionFilterType.
const (
	RUMHARDCODEDRETENTIONFILTERTYPE_HARDCODED_RETENTION_FILTERS RumHardcodedRetentionFilterType = "hardcoded_retention_filters"
)

var allowedRumHardcodedRetentionFilterTypeEnumValues = []RumHardcodedRetentionFilterType{
	RUMHARDCODEDRETENTIONFILTERTYPE_HARDCODED_RETENTION_FILTERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumHardcodedRetentionFilterType) GetAllowedValues() []RumHardcodedRetentionFilterType {
	return allowedRumHardcodedRetentionFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumHardcodedRetentionFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumHardcodedRetentionFilterType(value)
	return nil
}

// NewRumHardcodedRetentionFilterTypeFromValue returns a pointer to a valid RumHardcodedRetentionFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumHardcodedRetentionFilterTypeFromValue(v string) (*RumHardcodedRetentionFilterType, error) {
	ev := RumHardcodedRetentionFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumHardcodedRetentionFilterType: valid values are %v", v, allowedRumHardcodedRetentionFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumHardcodedRetentionFilterType) IsValid() bool {
	for _, existing := range allowedRumHardcodedRetentionFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumHardcodedRetentionFilterType value.
func (v RumHardcodedRetentionFilterType) Ptr() *RumHardcodedRetentionFilterType {
	return &v
}

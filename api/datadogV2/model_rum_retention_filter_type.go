// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRetentionFilterType The type of the resource. The value should always be retention_filters.
type RUMRetentionFilterType string

// List of RUMRetentionFilterType.
const (
	RUMRETENTIONFILTERTYPE_RETENTION_FILTERS RUMRetentionFilterType = "retention_filters"
)

var allowedRUMRetentionFilterTypeEnumValues = []RUMRetentionFilterType{
	RUMRETENTIONFILTERTYPE_RETENTION_FILTERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRetentionFilterType) GetAllowedValues() []RUMRetentionFilterType {
	return allowedRUMRetentionFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRetentionFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRetentionFilterType(value)
	return nil
}

// NewRUMRetentionFilterTypeFromValue returns a pointer to a valid RUMRetentionFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRetentionFilterTypeFromValue(v string) (*RUMRetentionFilterType, error) {
	ev := RUMRetentionFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRetentionFilterType: valid values are %v", v, allowedRUMRetentionFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRetentionFilterType) IsValid() bool {
	for _, existing := range allowedRUMRetentionFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRetentionFilterType value.
func (v RUMRetentionFilterType) Ptr() *RUMRetentionFilterType {
	return &v
}

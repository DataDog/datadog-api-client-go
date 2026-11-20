// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMPermanentRetentionFilterType The type of the resource. The value should always be `permanent_retention_filters`.
type RUMPermanentRetentionFilterType string

// List of RUMPermanentRetentionFilterType.
const (
	RUMPERMANENTRETENTIONFILTERTYPE_PERMANENT_RETENTION_FILTERS RUMPermanentRetentionFilterType = "permanent_retention_filters"
)

var allowedRUMPermanentRetentionFilterTypeEnumValues = []RUMPermanentRetentionFilterType{
	RUMPERMANENTRETENTIONFILTERTYPE_PERMANENT_RETENTION_FILTERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMPermanentRetentionFilterType) GetAllowedValues() []RUMPermanentRetentionFilterType {
	return allowedRUMPermanentRetentionFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMPermanentRetentionFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMPermanentRetentionFilterType(value)
	return nil
}

// NewRUMPermanentRetentionFilterTypeFromValue returns a pointer to a valid RUMPermanentRetentionFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMPermanentRetentionFilterTypeFromValue(v string) (*RUMPermanentRetentionFilterType, error) {
	ev := RUMPermanentRetentionFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMPermanentRetentionFilterType: valid values are %v", v, allowedRUMPermanentRetentionFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMPermanentRetentionFilterType) IsValid() bool {
	for _, existing := range allowedRUMPermanentRetentionFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMPermanentRetentionFilterType value.
func (v RUMPermanentRetentionFilterType) Ptr() *RUMPermanentRetentionFilterType {
	return &v
}

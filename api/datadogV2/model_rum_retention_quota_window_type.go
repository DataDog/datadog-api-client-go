// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRetentionQuotaWindowType The window type over which the session limit is enforced.
type RUMRetentionQuotaWindowType string

// List of RUMRetentionQuotaWindowType.
const (
	RUMRETENTIONQUOTAWINDOWTYPE_DAILY RUMRetentionQuotaWindowType = "daily"
)

var allowedRUMRetentionQuotaWindowTypeEnumValues = []RUMRetentionQuotaWindowType{
	RUMRETENTIONQUOTAWINDOWTYPE_DAILY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRetentionQuotaWindowType) GetAllowedValues() []RUMRetentionQuotaWindowType {
	return allowedRUMRetentionQuotaWindowTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRetentionQuotaWindowType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRetentionQuotaWindowType(value)
	return nil
}

// NewRUMRetentionQuotaWindowTypeFromValue returns a pointer to a valid RUMRetentionQuotaWindowType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRetentionQuotaWindowTypeFromValue(v string) (*RUMRetentionQuotaWindowType, error) {
	ev := RUMRetentionQuotaWindowType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRetentionQuotaWindowType: valid values are %v", v, allowedRUMRetentionQuotaWindowTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRetentionQuotaWindowType) IsValid() bool {
	for _, existing := range allowedRUMRetentionQuotaWindowTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRetentionQuotaWindowType value.
func (v RUMRetentionQuotaWindowType) Ptr() *RUMRetentionQuotaWindowType {
	return &v
}

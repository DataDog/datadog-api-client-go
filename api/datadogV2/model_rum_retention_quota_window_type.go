// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionQuotaWindowType The window type over which the session limit is enforced.
type RumRetentionQuotaWindowType string

// List of RumRetentionQuotaWindowType.
const (
	RUMRETENTIONQUOTAWINDOWTYPE_DAILY RumRetentionQuotaWindowType = "daily"
)

var allowedRumRetentionQuotaWindowTypeEnumValues = []RumRetentionQuotaWindowType{
	RUMRETENTIONQUOTAWINDOWTYPE_DAILY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRetentionQuotaWindowType) GetAllowedValues() []RumRetentionQuotaWindowType {
	return allowedRumRetentionQuotaWindowTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRetentionQuotaWindowType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRetentionQuotaWindowType(value)
	return nil
}

// NewRumRetentionQuotaWindowTypeFromValue returns a pointer to a valid RumRetentionQuotaWindowType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRetentionQuotaWindowTypeFromValue(v string) (*RumRetentionQuotaWindowType, error) {
	ev := RumRetentionQuotaWindowType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRetentionQuotaWindowType: valid values are %v", v, allowedRumRetentionQuotaWindowTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRetentionQuotaWindowType) IsValid() bool {
	for _, existing := range allowedRumRetentionQuotaWindowTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRetentionQuotaWindowType value.
func (v RumRetentionQuotaWindowType) Ptr() *RumRetentionQuotaWindowType {
	return &v
}

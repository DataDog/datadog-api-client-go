// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionQuotaMode The retention quota mode. `custom` enforces a fixed session limit, while
// `adaptive` dynamically adjusts retention.
type RumRetentionQuotaMode string

// List of RumRetentionQuotaMode.
const (
	RUMRETENTIONQUOTAMODE_CUSTOM   RumRetentionQuotaMode = "custom"
	RUMRETENTIONQUOTAMODE_ADAPTIVE RumRetentionQuotaMode = "adaptive"
)

var allowedRumRetentionQuotaModeEnumValues = []RumRetentionQuotaMode{
	RUMRETENTIONQUOTAMODE_CUSTOM,
	RUMRETENTIONQUOTAMODE_ADAPTIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRetentionQuotaMode) GetAllowedValues() []RumRetentionQuotaMode {
	return allowedRumRetentionQuotaModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRetentionQuotaMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRetentionQuotaMode(value)
	return nil
}

// NewRumRetentionQuotaModeFromValue returns a pointer to a valid RumRetentionQuotaMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRetentionQuotaModeFromValue(v string) (*RumRetentionQuotaMode, error) {
	ev := RumRetentionQuotaMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRetentionQuotaMode: valid values are %v", v, allowedRumRetentionQuotaModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRetentionQuotaMode) IsValid() bool {
	for _, existing := range allowedRumRetentionQuotaModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRetentionQuotaMode value.
func (v RumRetentionQuotaMode) Ptr() *RumRetentionQuotaMode {
	return &v
}

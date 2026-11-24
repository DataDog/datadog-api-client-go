// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRetentionQuotaMode The retention quota mode. `custom` enforces a fixed session limit.
// `custom` is the only supported mode.
type RUMRetentionQuotaMode string

// List of RUMRetentionQuotaMode.
const (
	RUMRETENTIONQUOTAMODE_CUSTOM RUMRetentionQuotaMode = "custom"
)

var allowedRUMRetentionQuotaModeEnumValues = []RUMRetentionQuotaMode{
	RUMRETENTIONQUOTAMODE_CUSTOM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRetentionQuotaMode) GetAllowedValues() []RUMRetentionQuotaMode {
	return allowedRUMRetentionQuotaModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRetentionQuotaMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRetentionQuotaMode(value)
	return nil
}

// NewRUMRetentionQuotaModeFromValue returns a pointer to a valid RUMRetentionQuotaMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRetentionQuotaModeFromValue(v string) (*RUMRetentionQuotaMode, error) {
	ev := RUMRetentionQuotaMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRetentionQuotaMode: valid values are %v", v, allowedRUMRetentionQuotaModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRetentionQuotaMode) IsValid() bool {
	for _, existing := range allowedRUMRetentionQuotaModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRetentionQuotaMode value.
func (v RUMRetentionQuotaMode) Ptr() *RUMRetentionQuotaMode {
	return &v
}

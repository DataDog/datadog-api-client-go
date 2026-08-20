// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionQuotaConfigType The type of the resource, always `rum_quota_config`.
type RumRetentionQuotaConfigType string

// List of RumRetentionQuotaConfigType.
const (
	RUMRETENTIONQUOTACONFIGTYPE_RUM_QUOTA_CONFIG RumRetentionQuotaConfigType = "rum_quota_config"
)

var allowedRumRetentionQuotaConfigTypeEnumValues = []RumRetentionQuotaConfigType{
	RUMRETENTIONQUOTACONFIGTYPE_RUM_QUOTA_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRetentionQuotaConfigType) GetAllowedValues() []RumRetentionQuotaConfigType {
	return allowedRumRetentionQuotaConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRetentionQuotaConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRetentionQuotaConfigType(value)
	return nil
}

// NewRumRetentionQuotaConfigTypeFromValue returns a pointer to a valid RumRetentionQuotaConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRetentionQuotaConfigTypeFromValue(v string) (*RumRetentionQuotaConfigType, error) {
	ev := RumRetentionQuotaConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRetentionQuotaConfigType: valid values are %v", v, allowedRumRetentionQuotaConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRetentionQuotaConfigType) IsValid() bool {
	for _, existing := range allowedRumRetentionQuotaConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRetentionQuotaConfigType value.
func (v RumRetentionQuotaConfigType) Ptr() *RumRetentionQuotaConfigType {
	return &v
}

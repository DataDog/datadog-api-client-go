// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRetentionQuotaConfigType The type of the resource, always `rum_quota_config`.
type RUMRetentionQuotaConfigType string

// List of RUMRetentionQuotaConfigType.
const (
	RUMRETENTIONQUOTACONFIGTYPE_RUM_QUOTA_CONFIG RUMRetentionQuotaConfigType = "rum_quota_config"
)

var allowedRUMRetentionQuotaConfigTypeEnumValues = []RUMRetentionQuotaConfigType{
	RUMRETENTIONQUOTACONFIGTYPE_RUM_QUOTA_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRetentionQuotaConfigType) GetAllowedValues() []RUMRetentionQuotaConfigType {
	return allowedRUMRetentionQuotaConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRetentionQuotaConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRetentionQuotaConfigType(value)
	return nil
}

// NewRUMRetentionQuotaConfigTypeFromValue returns a pointer to a valid RUMRetentionQuotaConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRetentionQuotaConfigTypeFromValue(v string) (*RUMRetentionQuotaConfigType, error) {
	ev := RUMRetentionQuotaConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRetentionQuotaConfigType: valid values are %v", v, allowedRUMRetentionQuotaConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRetentionQuotaConfigType) IsValid() bool {
	for _, existing := range allowedRUMRetentionQuotaConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRetentionQuotaConfigType value.
func (v RUMRetentionQuotaConfigType) Ptr() *RUMRetentionQuotaConfigType {
	return &v
}

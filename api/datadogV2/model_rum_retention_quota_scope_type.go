// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionQuotaScopeType The type of scope the retention quota configuration applies to.
type RumRetentionQuotaScopeType string

// List of RumRetentionQuotaScopeType.
const (
	RUMRETENTIONQUOTASCOPETYPE_APPLICATION RumRetentionQuotaScopeType = "application"
)

var allowedRumRetentionQuotaScopeTypeEnumValues = []RumRetentionQuotaScopeType{
	RUMRETENTIONQUOTASCOPETYPE_APPLICATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRetentionQuotaScopeType) GetAllowedValues() []RumRetentionQuotaScopeType {
	return allowedRumRetentionQuotaScopeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRetentionQuotaScopeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRetentionQuotaScopeType(value)
	return nil
}

// NewRumRetentionQuotaScopeTypeFromValue returns a pointer to a valid RumRetentionQuotaScopeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRetentionQuotaScopeTypeFromValue(v string) (*RumRetentionQuotaScopeType, error) {
	ev := RumRetentionQuotaScopeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRetentionQuotaScopeType: valid values are %v", v, allowedRumRetentionQuotaScopeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRetentionQuotaScopeType) IsValid() bool {
	for _, existing := range allowedRumRetentionQuotaScopeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRetentionQuotaScopeType value.
func (v RumRetentionQuotaScopeType) Ptr() *RumRetentionQuotaScopeType {
	return &v
}

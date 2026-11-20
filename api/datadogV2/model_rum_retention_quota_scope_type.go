// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRetentionQuotaScopeType The type of scope the retention quota configuration applies to.
// `application` is the only supported scope type.
type RUMRetentionQuotaScopeType string

// List of RUMRetentionQuotaScopeType.
const (
	RUMRETENTIONQUOTASCOPETYPE_APPLICATION RUMRetentionQuotaScopeType = "application"
)

var allowedRUMRetentionQuotaScopeTypeEnumValues = []RUMRetentionQuotaScopeType{
	RUMRETENTIONQUOTASCOPETYPE_APPLICATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRetentionQuotaScopeType) GetAllowedValues() []RUMRetentionQuotaScopeType {
	return allowedRUMRetentionQuotaScopeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRetentionQuotaScopeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRetentionQuotaScopeType(value)
	return nil
}

// NewRUMRetentionQuotaScopeTypeFromValue returns a pointer to a valid RUMRetentionQuotaScopeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRetentionQuotaScopeTypeFromValue(v string) (*RUMRetentionQuotaScopeType, error) {
	ev := RUMRetentionQuotaScopeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRetentionQuotaScopeType: valid values are %v", v, allowedRUMRetentionQuotaScopeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRetentionQuotaScopeType) IsValid() bool {
	for _, existing := range allowedRUMRetentionQuotaScopeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRetentionQuotaScopeType value.
func (v RUMRetentionQuotaScopeType) Ptr() *RUMRetentionQuotaScopeType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPageEmailAccessPolicyType Status page email access policy resource type.
type StatusPageEmailAccessPolicyType string

// List of StatusPageEmailAccessPolicyType.
const (
	STATUSPAGEEMAILACCESSPOLICYTYPE_EMAIL_ACCESS_POLICIES StatusPageEmailAccessPolicyType = "email_access_policies"
)

var allowedStatusPageEmailAccessPolicyTypeEnumValues = []StatusPageEmailAccessPolicyType{
	STATUSPAGEEMAILACCESSPOLICYTYPE_EMAIL_ACCESS_POLICIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StatusPageEmailAccessPolicyType) GetAllowedValues() []StatusPageEmailAccessPolicyType {
	return allowedStatusPageEmailAccessPolicyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StatusPageEmailAccessPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StatusPageEmailAccessPolicyType(value)
	return nil
}

// NewStatusPageEmailAccessPolicyTypeFromValue returns a pointer to a valid StatusPageEmailAccessPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStatusPageEmailAccessPolicyTypeFromValue(v string) (*StatusPageEmailAccessPolicyType, error) {
	ev := StatusPageEmailAccessPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StatusPageEmailAccessPolicyType: valid values are %v", v, allowedStatusPageEmailAccessPolicyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StatusPageEmailAccessPolicyType) IsValid() bool {
	for _, existing := range allowedStatusPageEmailAccessPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusPageEmailAccessPolicyType value.
func (v StatusPageEmailAccessPolicyType) Ptr() *StatusPageEmailAccessPolicyType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InboxRuleType The JSON:API type for inbox rules.
type InboxRuleType string

// List of InboxRuleType.
const (
	INBOXRULETYPE_INBOX_RULES InboxRuleType = "inbox_rules"
)

var allowedInboxRuleTypeEnumValues = []InboxRuleType{
	INBOXRULETYPE_INBOX_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *InboxRuleType) GetAllowedValues() []InboxRuleType {
	return allowedInboxRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InboxRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InboxRuleType(value)
	return nil
}

// NewInboxRuleTypeFromValue returns a pointer to a valid InboxRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInboxRuleTypeFromValue(v string) (*InboxRuleType, error) {
	ev := InboxRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InboxRuleType: valid values are %v", v, allowedInboxRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InboxRuleType) IsValid() bool {
	for _, existing := range allowedInboxRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InboxRuleType value.
func (v InboxRuleType) Ptr() *InboxRuleType {
	return &v
}

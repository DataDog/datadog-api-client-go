// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InboxRulesType The pipeline rule type associated to inbox rules
type InboxRulesType string

// List of InboxRulesType.
const (
	INBOXRULESTYPE_INBOX_RULES InboxRulesType = "inbox_rules"
)

var allowedInboxRulesTypeEnumValues = []InboxRulesType{
	INBOXRULESTYPE_INBOX_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *InboxRulesType) GetAllowedValues() []InboxRulesType {
	return allowedInboxRulesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InboxRulesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InboxRulesType(value)
	return nil
}

// NewInboxRulesTypeFromValue returns a pointer to a valid InboxRulesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInboxRulesTypeFromValue(v string) (*InboxRulesType, error) {
	ev := InboxRulesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InboxRulesType: valid values are %v", v, allowedInboxRulesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InboxRulesType) IsValid() bool {
	for _, existing := range allowedInboxRulesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InboxRulesType value.
func (v InboxRulesType) Ptr() *InboxRulesType {
	return &v
}

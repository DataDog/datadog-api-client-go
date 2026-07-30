// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMOperationJourneyCompositeRuleKind The rule used to combine the composite rule's predicates. `all_of` requires every predicate
// to match, in any order. `in_order` requires every predicate to match in the given order.
type RUMOperationJourneyCompositeRuleKind string

// List of RUMOperationJourneyCompositeRuleKind.
const (
	RUMOPERATIONJOURNEYCOMPOSITERULEKIND_ALL_OF   RUMOperationJourneyCompositeRuleKind = "all_of"
	RUMOPERATIONJOURNEYCOMPOSITERULEKIND_IN_ORDER RUMOperationJourneyCompositeRuleKind = "in_order"
)

var allowedRUMOperationJourneyCompositeRuleKindEnumValues = []RUMOperationJourneyCompositeRuleKind{
	RUMOPERATIONJOURNEYCOMPOSITERULEKIND_ALL_OF,
	RUMOPERATIONJOURNEYCOMPOSITERULEKIND_IN_ORDER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMOperationJourneyCompositeRuleKind) GetAllowedValues() []RUMOperationJourneyCompositeRuleKind {
	return allowedRUMOperationJourneyCompositeRuleKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMOperationJourneyCompositeRuleKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMOperationJourneyCompositeRuleKind(value)
	return nil
}

// NewRUMOperationJourneyCompositeRuleKindFromValue returns a pointer to a valid RUMOperationJourneyCompositeRuleKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMOperationJourneyCompositeRuleKindFromValue(v string) (*RUMOperationJourneyCompositeRuleKind, error) {
	ev := RUMOperationJourneyCompositeRuleKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMOperationJourneyCompositeRuleKind: valid values are %v", v, allowedRUMOperationJourneyCompositeRuleKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMOperationJourneyCompositeRuleKind) IsValid() bool {
	for _, existing := range allowedRUMOperationJourneyCompositeRuleKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMOperationJourneyCompositeRuleKind value.
func (v RUMOperationJourneyCompositeRuleKind) Ptr() *RUMOperationJourneyCompositeRuleKind {
	return &v
}

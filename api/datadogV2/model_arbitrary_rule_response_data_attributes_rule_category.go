// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArbitraryRuleResponseDataAttributesRuleCategory Rule category. Omitted for custom rules.
type ArbitraryRuleResponseDataAttributesRuleCategory string

// List of ArbitraryRuleResponseDataAttributesRuleCategory.
const (
	ARBITRARYRULERESPONSEDATAATTRIBUTESRULECATEGORY_CUSTOM     ArbitraryRuleResponseDataAttributesRuleCategory = "custom"
	ARBITRARYRULERESPONSEDATAATTRIBUTESRULECATEGORY_AI_GATEWAY ArbitraryRuleResponseDataAttributesRuleCategory = "ai_gateway"
)

var allowedArbitraryRuleResponseDataAttributesRuleCategoryEnumValues = []ArbitraryRuleResponseDataAttributesRuleCategory{
	ARBITRARYRULERESPONSEDATAATTRIBUTESRULECATEGORY_CUSTOM,
	ARBITRARYRULERESPONSEDATAATTRIBUTESRULECATEGORY_AI_GATEWAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ArbitraryRuleResponseDataAttributesRuleCategory) GetAllowedValues() []ArbitraryRuleResponseDataAttributesRuleCategory {
	return allowedArbitraryRuleResponseDataAttributesRuleCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ArbitraryRuleResponseDataAttributesRuleCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ArbitraryRuleResponseDataAttributesRuleCategory(value)
	return nil
}

// NewArbitraryRuleResponseDataAttributesRuleCategoryFromValue returns a pointer to a valid ArbitraryRuleResponseDataAttributesRuleCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewArbitraryRuleResponseDataAttributesRuleCategoryFromValue(v string) (*ArbitraryRuleResponseDataAttributesRuleCategory, error) {
	ev := ArbitraryRuleResponseDataAttributesRuleCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ArbitraryRuleResponseDataAttributesRuleCategory: valid values are %v", v, allowedArbitraryRuleResponseDataAttributesRuleCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ArbitraryRuleResponseDataAttributesRuleCategory) IsValid() bool {
	for _, existing := range allowedArbitraryRuleResponseDataAttributesRuleCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ArbitraryRuleResponseDataAttributesRuleCategory value.
func (v ArbitraryRuleResponseDataAttributesRuleCategory) Ptr() *ArbitraryRuleResponseDataAttributesRuleCategory {
	return &v
}

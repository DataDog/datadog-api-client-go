// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationCustomRuleType Custom recommendation rule resource type.
type CostRecommendationCustomRuleType string

// List of CostRecommendationCustomRuleType.
const (
	COSTRECOMMENDATIONCUSTOMRULETYPE_CUSTOM_RULES CostRecommendationCustomRuleType = "custom_rules"
)

var allowedCostRecommendationCustomRuleTypeEnumValues = []CostRecommendationCustomRuleType{
	COSTRECOMMENDATIONCUSTOMRULETYPE_CUSTOM_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostRecommendationCustomRuleType) GetAllowedValues() []CostRecommendationCustomRuleType {
	return allowedCostRecommendationCustomRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostRecommendationCustomRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostRecommendationCustomRuleType(value)
	return nil
}

// NewCostRecommendationCustomRuleTypeFromValue returns a pointer to a valid CostRecommendationCustomRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostRecommendationCustomRuleTypeFromValue(v string) (*CostRecommendationCustomRuleType, error) {
	ev := CostRecommendationCustomRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostRecommendationCustomRuleType: valid values are %v", v, allowedCostRecommendationCustomRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostRecommendationCustomRuleType) IsValid() bool {
	for _, existing := range allowedCostRecommendationCustomRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostRecommendationCustomRuleType value.
func (v CostRecommendationCustomRuleType) Ptr() *CostRecommendationCustomRuleType {
	return &v
}

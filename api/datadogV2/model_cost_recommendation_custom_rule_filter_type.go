// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationCustomRuleFilterType How the filter compares the tag. `is`, `is_not`, `contains`, `does_not_contain`, `like` (glob), and `not_like`
// compare the tag values with `values`. `all` matches resources that have the tag and `untagged` matches resources
// that do not, so both ignore `values`.
type CostRecommendationCustomRuleFilterType string

// List of CostRecommendationCustomRuleFilterType.
const (
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_IS               CostRecommendationCustomRuleFilterType = "is"
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_IS_NOT           CostRecommendationCustomRuleFilterType = "is_not"
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_CONTAINS         CostRecommendationCustomRuleFilterType = "contains"
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_DOES_NOT_CONTAIN CostRecommendationCustomRuleFilterType = "does_not_contain"
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_LIKE             CostRecommendationCustomRuleFilterType = "like"
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_NOT_LIKE         CostRecommendationCustomRuleFilterType = "not_like"
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_ALL              CostRecommendationCustomRuleFilterType = "all"
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_UNTAGGED         CostRecommendationCustomRuleFilterType = "untagged"
)

var allowedCostRecommendationCustomRuleFilterTypeEnumValues = []CostRecommendationCustomRuleFilterType{
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_IS,
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_IS_NOT,
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_CONTAINS,
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_DOES_NOT_CONTAIN,
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_LIKE,
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_NOT_LIKE,
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_ALL,
	COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_UNTAGGED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostRecommendationCustomRuleFilterType) GetAllowedValues() []CostRecommendationCustomRuleFilterType {
	return allowedCostRecommendationCustomRuleFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostRecommendationCustomRuleFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostRecommendationCustomRuleFilterType(value)
	return nil
}

// NewCostRecommendationCustomRuleFilterTypeFromValue returns a pointer to a valid CostRecommendationCustomRuleFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostRecommendationCustomRuleFilterTypeFromValue(v string) (*CostRecommendationCustomRuleFilterType, error) {
	ev := CostRecommendationCustomRuleFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostRecommendationCustomRuleFilterType: valid values are %v", v, allowedCostRecommendationCustomRuleFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostRecommendationCustomRuleFilterType) IsValid() bool {
	for _, existing := range allowedCostRecommendationCustomRuleFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostRecommendationCustomRuleFilterType value.
func (v CostRecommendationCustomRuleFilterType) Ptr() *CostRecommendationCustomRuleFilterType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationNotificationRuleType Recommendation notification rule resource type.
type CostRecommendationNotificationRuleType string

// List of CostRecommendationNotificationRuleType.
const (
	COSTRECOMMENDATIONNOTIFICATIONRULETYPE_NOTIFICATION_RULES CostRecommendationNotificationRuleType = "notification_rules"
)

var allowedCostRecommendationNotificationRuleTypeEnumValues = []CostRecommendationNotificationRuleType{
	COSTRECOMMENDATIONNOTIFICATIONRULETYPE_NOTIFICATION_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostRecommendationNotificationRuleType) GetAllowedValues() []CostRecommendationNotificationRuleType {
	return allowedCostRecommendationNotificationRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostRecommendationNotificationRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostRecommendationNotificationRuleType(value)
	return nil
}

// NewCostRecommendationNotificationRuleTypeFromValue returns a pointer to a valid CostRecommendationNotificationRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostRecommendationNotificationRuleTypeFromValue(v string) (*CostRecommendationNotificationRuleType, error) {
	ev := CostRecommendationNotificationRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostRecommendationNotificationRuleType: valid values are %v", v, allowedCostRecommendationNotificationRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostRecommendationNotificationRuleType) IsValid() bool {
	for _, existing := range allowedCostRecommendationNotificationRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostRecommendationNotificationRuleType value.
func (v CostRecommendationNotificationRuleType) Ptr() *CostRecommendationNotificationRuleType {
	return &v
}

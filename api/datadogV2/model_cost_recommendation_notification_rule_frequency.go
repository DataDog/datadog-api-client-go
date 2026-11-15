// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationNotificationRuleFrequency How often the digest is sent. `weekly` and `biweekly` use `weekly_execution_day` and `weekly_execution_time`;
// `every_30` and `every_90` (days) use `interval_execution_date` and `interval_execution_time`.
type CostRecommendationNotificationRuleFrequency string

// List of CostRecommendationNotificationRuleFrequency.
const (
	COSTRECOMMENDATIONNOTIFICATIONRULEFREQUENCY_WEEKLY   CostRecommendationNotificationRuleFrequency = "weekly"
	COSTRECOMMENDATIONNOTIFICATIONRULEFREQUENCY_BIWEEKLY CostRecommendationNotificationRuleFrequency = "biweekly"
	COSTRECOMMENDATIONNOTIFICATIONRULEFREQUENCY_EVERY_30 CostRecommendationNotificationRuleFrequency = "every_30"
	COSTRECOMMENDATIONNOTIFICATIONRULEFREQUENCY_EVERY_90 CostRecommendationNotificationRuleFrequency = "every_90"
)

var allowedCostRecommendationNotificationRuleFrequencyEnumValues = []CostRecommendationNotificationRuleFrequency{
	COSTRECOMMENDATIONNOTIFICATIONRULEFREQUENCY_WEEKLY,
	COSTRECOMMENDATIONNOTIFICATIONRULEFREQUENCY_BIWEEKLY,
	COSTRECOMMENDATIONNOTIFICATIONRULEFREQUENCY_EVERY_30,
	COSTRECOMMENDATIONNOTIFICATIONRULEFREQUENCY_EVERY_90,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostRecommendationNotificationRuleFrequency) GetAllowedValues() []CostRecommendationNotificationRuleFrequency {
	return allowedCostRecommendationNotificationRuleFrequencyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostRecommendationNotificationRuleFrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostRecommendationNotificationRuleFrequency(value)
	return nil
}

// NewCostRecommendationNotificationRuleFrequencyFromValue returns a pointer to a valid CostRecommendationNotificationRuleFrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostRecommendationNotificationRuleFrequencyFromValue(v string) (*CostRecommendationNotificationRuleFrequency, error) {
	ev := CostRecommendationNotificationRuleFrequency(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostRecommendationNotificationRuleFrequency: valid values are %v", v, allowedCostRecommendationNotificationRuleFrequencyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostRecommendationNotificationRuleFrequency) IsValid() bool {
	for _, existing := range allowedCostRecommendationNotificationRuleFrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostRecommendationNotificationRuleFrequency value.
func (v CostRecommendationNotificationRuleFrequency) Ptr() *CostRecommendationNotificationRuleFrequency {
	return &v
}

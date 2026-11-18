// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationNotificationRuleMessageCustomizationMode How the message is customized.
type CostRecommendationNotificationRuleMessageCustomizationMode string

// List of CostRecommendationNotificationRuleMessageCustomizationMode.
const (
	COSTRECOMMENDATIONNOTIFICATIONRULEMESSAGECUSTOMIZATIONMODE_STRUCTURED CostRecommendationNotificationRuleMessageCustomizationMode = "structured"
)

var allowedCostRecommendationNotificationRuleMessageCustomizationModeEnumValues = []CostRecommendationNotificationRuleMessageCustomizationMode{
	COSTRECOMMENDATIONNOTIFICATIONRULEMESSAGECUSTOMIZATIONMODE_STRUCTURED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostRecommendationNotificationRuleMessageCustomizationMode) GetAllowedValues() []CostRecommendationNotificationRuleMessageCustomizationMode {
	return allowedCostRecommendationNotificationRuleMessageCustomizationModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostRecommendationNotificationRuleMessageCustomizationMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostRecommendationNotificationRuleMessageCustomizationMode(value)
	return nil
}

// NewCostRecommendationNotificationRuleMessageCustomizationModeFromValue returns a pointer to a valid CostRecommendationNotificationRuleMessageCustomizationMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostRecommendationNotificationRuleMessageCustomizationModeFromValue(v string) (*CostRecommendationNotificationRuleMessageCustomizationMode, error) {
	ev := CostRecommendationNotificationRuleMessageCustomizationMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostRecommendationNotificationRuleMessageCustomizationMode: valid values are %v", v, allowedCostRecommendationNotificationRuleMessageCustomizationModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostRecommendationNotificationRuleMessageCustomizationMode) IsValid() bool {
	for _, existing := range allowedCostRecommendationNotificationRuleMessageCustomizationModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostRecommendationNotificationRuleMessageCustomizationMode value.
func (v CostRecommendationNotificationRuleMessageCustomizationMode) Ptr() *CostRecommendationNotificationRuleMessageCustomizationMode {
	return &v
}

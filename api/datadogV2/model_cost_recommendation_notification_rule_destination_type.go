// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationNotificationRuleDestinationType Where the digest is sent. `slack` requires `workspace_id` and `channel_id`; `teams` requires the
// `teams_*` attributes. Defaults to `slack`.
type CostRecommendationNotificationRuleDestinationType string

// List of CostRecommendationNotificationRuleDestinationType.
const (
	COSTRECOMMENDATIONNOTIFICATIONRULEDESTINATIONTYPE_SLACK CostRecommendationNotificationRuleDestinationType = "slack"
	COSTRECOMMENDATIONNOTIFICATIONRULEDESTINATIONTYPE_TEAMS CostRecommendationNotificationRuleDestinationType = "teams"
)

var allowedCostRecommendationNotificationRuleDestinationTypeEnumValues = []CostRecommendationNotificationRuleDestinationType{
	COSTRECOMMENDATIONNOTIFICATIONRULEDESTINATIONTYPE_SLACK,
	COSTRECOMMENDATIONNOTIFICATIONRULEDESTINATIONTYPE_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostRecommendationNotificationRuleDestinationType) GetAllowedValues() []CostRecommendationNotificationRuleDestinationType {
	return allowedCostRecommendationNotificationRuleDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostRecommendationNotificationRuleDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostRecommendationNotificationRuleDestinationType(value)
	return nil
}

// NewCostRecommendationNotificationRuleDestinationTypeFromValue returns a pointer to a valid CostRecommendationNotificationRuleDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostRecommendationNotificationRuleDestinationTypeFromValue(v string) (*CostRecommendationNotificationRuleDestinationType, error) {
	ev := CostRecommendationNotificationRuleDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostRecommendationNotificationRuleDestinationType: valid values are %v", v, allowedCostRecommendationNotificationRuleDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostRecommendationNotificationRuleDestinationType) IsValid() bool {
	for _, existing := range allowedCostRecommendationNotificationRuleDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostRecommendationNotificationRuleDestinationType value.
func (v CostRecommendationNotificationRuleDestinationType) Ptr() *CostRecommendationNotificationRuleDestinationType {
	return &v
}

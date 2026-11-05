// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayBudgetWindowPolicy UTC accounting period for a budget.
type AiGatewayBudgetWindowPolicy string

// List of AiGatewayBudgetWindowPolicy.
const (
	AIGATEWAYBUDGETWINDOWPOLICY_DAILY   AiGatewayBudgetWindowPolicy = "DAILY"
	AIGATEWAYBUDGETWINDOWPOLICY_WEEKLY  AiGatewayBudgetWindowPolicy = "WEEKLY"
	AIGATEWAYBUDGETWINDOWPOLICY_MONTHLY AiGatewayBudgetWindowPolicy = "MONTHLY"
)

var allowedAiGatewayBudgetWindowPolicyEnumValues = []AiGatewayBudgetWindowPolicy{
	AIGATEWAYBUDGETWINDOWPOLICY_DAILY,
	AIGATEWAYBUDGETWINDOWPOLICY_WEEKLY,
	AIGATEWAYBUDGETWINDOWPOLICY_MONTHLY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayBudgetWindowPolicy) GetAllowedValues() []AiGatewayBudgetWindowPolicy {
	return allowedAiGatewayBudgetWindowPolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayBudgetWindowPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayBudgetWindowPolicy(value)
	return nil
}

// NewAiGatewayBudgetWindowPolicyFromValue returns a pointer to a valid AiGatewayBudgetWindowPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayBudgetWindowPolicyFromValue(v string) (*AiGatewayBudgetWindowPolicy, error) {
	ev := AiGatewayBudgetWindowPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayBudgetWindowPolicy: valid values are %v", v, allowedAiGatewayBudgetWindowPolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayBudgetWindowPolicy) IsValid() bool {
	for _, existing := range allowedAiGatewayBudgetWindowPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayBudgetWindowPolicy value.
func (v AiGatewayBudgetWindowPolicy) Ptr() *AiGatewayBudgetWindowPolicy {
	return &v
}

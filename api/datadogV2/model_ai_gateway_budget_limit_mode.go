// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayBudgetLimitMode How the budget limit applies. `shared` is one limit for all matching traffic. `per_instance` is one
// limit for each counter of a template budget that splits by value.
type AiGatewayBudgetLimitMode string

// List of AiGatewayBudgetLimitMode.
const (
	AIGATEWAYBUDGETLIMITMODE_SHARED       AiGatewayBudgetLimitMode = "shared"
	AIGATEWAYBUDGETLIMITMODE_PER_INSTANCE AiGatewayBudgetLimitMode = "per_instance"
)

var allowedAiGatewayBudgetLimitModeEnumValues = []AiGatewayBudgetLimitMode{
	AIGATEWAYBUDGETLIMITMODE_SHARED,
	AIGATEWAYBUDGETLIMITMODE_PER_INSTANCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayBudgetLimitMode) GetAllowedValues() []AiGatewayBudgetLimitMode {
	return allowedAiGatewayBudgetLimitModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayBudgetLimitMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayBudgetLimitMode(value)
	return nil
}

// NewAiGatewayBudgetLimitModeFromValue returns a pointer to a valid AiGatewayBudgetLimitMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayBudgetLimitModeFromValue(v string) (*AiGatewayBudgetLimitMode, error) {
	ev := AiGatewayBudgetLimitMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayBudgetLimitMode: valid values are %v", v, allowedAiGatewayBudgetLimitModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayBudgetLimitMode) IsValid() bool {
	for _, existing := range allowedAiGatewayBudgetLimitModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayBudgetLimitMode value.
func (v AiGatewayBudgetLimitMode) Ptr() *AiGatewayBudgetLimitMode {
	return &v
}

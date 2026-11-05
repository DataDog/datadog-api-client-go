// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayBudgetType AI Gateway budget resource type.
type AiGatewayBudgetType string

// List of AiGatewayBudgetType.
const (
	AIGATEWAYBUDGETTYPE_BUDGETS AiGatewayBudgetType = "budgets"
)

var allowedAiGatewayBudgetTypeEnumValues = []AiGatewayBudgetType{
	AIGATEWAYBUDGETTYPE_BUDGETS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayBudgetType) GetAllowedValues() []AiGatewayBudgetType {
	return allowedAiGatewayBudgetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayBudgetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayBudgetType(value)
	return nil
}

// NewAiGatewayBudgetTypeFromValue returns a pointer to a valid AiGatewayBudgetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayBudgetTypeFromValue(v string) (*AiGatewayBudgetType, error) {
	ev := AiGatewayBudgetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayBudgetType: valid values are %v", v, allowedAiGatewayBudgetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayBudgetType) IsValid() bool {
	for _, existing := range allowedAiGatewayBudgetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayBudgetType value.
func (v AiGatewayBudgetType) Ptr() *AiGatewayBudgetType {
	return &v
}

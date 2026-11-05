// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayBudgetInstanceDimensionType Budget split dimension values resource type.
type AiGatewayBudgetInstanceDimensionType string

// List of AiGatewayBudgetInstanceDimensionType.
const (
	AIGATEWAYBUDGETINSTANCEDIMENSIONTYPE_BUDGET_INSTANCE_DIMENSIONS AiGatewayBudgetInstanceDimensionType = "budget-instance-dimensions"
)

var allowedAiGatewayBudgetInstanceDimensionTypeEnumValues = []AiGatewayBudgetInstanceDimensionType{
	AIGATEWAYBUDGETINSTANCEDIMENSIONTYPE_BUDGET_INSTANCE_DIMENSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayBudgetInstanceDimensionType) GetAllowedValues() []AiGatewayBudgetInstanceDimensionType {
	return allowedAiGatewayBudgetInstanceDimensionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayBudgetInstanceDimensionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayBudgetInstanceDimensionType(value)
	return nil
}

// NewAiGatewayBudgetInstanceDimensionTypeFromValue returns a pointer to a valid AiGatewayBudgetInstanceDimensionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayBudgetInstanceDimensionTypeFromValue(v string) (*AiGatewayBudgetInstanceDimensionType, error) {
	ev := AiGatewayBudgetInstanceDimensionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayBudgetInstanceDimensionType: valid values are %v", v, allowedAiGatewayBudgetInstanceDimensionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayBudgetInstanceDimensionType) IsValid() bool {
	for _, existing := range allowedAiGatewayBudgetInstanceDimensionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayBudgetInstanceDimensionType value.
func (v AiGatewayBudgetInstanceDimensionType) Ptr() *AiGatewayBudgetInstanceDimensionType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayBudgetInstancesSummaryType Budget instance summary resource type.
type AiGatewayBudgetInstancesSummaryType string

// List of AiGatewayBudgetInstancesSummaryType.
const (
	AIGATEWAYBUDGETINSTANCESSUMMARYTYPE_BUDGET_INSTANCES_SUMMARY AiGatewayBudgetInstancesSummaryType = "budget-instances-summary"
)

var allowedAiGatewayBudgetInstancesSummaryTypeEnumValues = []AiGatewayBudgetInstancesSummaryType{
	AIGATEWAYBUDGETINSTANCESSUMMARYTYPE_BUDGET_INSTANCES_SUMMARY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayBudgetInstancesSummaryType) GetAllowedValues() []AiGatewayBudgetInstancesSummaryType {
	return allowedAiGatewayBudgetInstancesSummaryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayBudgetInstancesSummaryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayBudgetInstancesSummaryType(value)
	return nil
}

// NewAiGatewayBudgetInstancesSummaryTypeFromValue returns a pointer to a valid AiGatewayBudgetInstancesSummaryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayBudgetInstancesSummaryTypeFromValue(v string) (*AiGatewayBudgetInstancesSummaryType, error) {
	ev := AiGatewayBudgetInstancesSummaryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayBudgetInstancesSummaryType: valid values are %v", v, allowedAiGatewayBudgetInstancesSummaryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayBudgetInstancesSummaryType) IsValid() bool {
	for _, existing := range allowedAiGatewayBudgetInstancesSummaryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayBudgetInstancesSummaryType value.
func (v AiGatewayBudgetInstancesSummaryType) Ptr() *AiGatewayBudgetInstancesSummaryType {
	return &v
}

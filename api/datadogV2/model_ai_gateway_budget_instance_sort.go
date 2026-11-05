// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayBudgetInstanceSort Sort order for budget instance counters.
type AiGatewayBudgetInstanceSort string

// List of AiGatewayBudgetInstanceSort.
const (
	AIGATEWAYBUDGETINSTANCESORT_SPEND_DESC AiGatewayBudgetInstanceSort = "spend_desc"
	AIGATEWAYBUDGETINSTANCESORT_SPEND_ASC  AiGatewayBudgetInstanceSort = "spend_asc"
	AIGATEWAYBUDGETINSTANCESORT_KEY_ASC    AiGatewayBudgetInstanceSort = "key_asc"
)

var allowedAiGatewayBudgetInstanceSortEnumValues = []AiGatewayBudgetInstanceSort{
	AIGATEWAYBUDGETINSTANCESORT_SPEND_DESC,
	AIGATEWAYBUDGETINSTANCESORT_SPEND_ASC,
	AIGATEWAYBUDGETINSTANCESORT_KEY_ASC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayBudgetInstanceSort) GetAllowedValues() []AiGatewayBudgetInstanceSort {
	return allowedAiGatewayBudgetInstanceSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayBudgetInstanceSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayBudgetInstanceSort(value)
	return nil
}

// NewAiGatewayBudgetInstanceSortFromValue returns a pointer to a valid AiGatewayBudgetInstanceSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayBudgetInstanceSortFromValue(v string) (*AiGatewayBudgetInstanceSort, error) {
	ev := AiGatewayBudgetInstanceSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayBudgetInstanceSort: valid values are %v", v, allowedAiGatewayBudgetInstanceSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayBudgetInstanceSort) IsValid() bool {
	for _, existing := range allowedAiGatewayBudgetInstanceSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayBudgetInstanceSort value.
func (v AiGatewayBudgetInstanceSort) Ptr() *AiGatewayBudgetInstanceSort {
	return &v
}

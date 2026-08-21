// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneySearchGraphFilterOperator Comparison operator applied to the graph filter value.
type ProductAnalyticsJourneySearchGraphFilterOperator string

// List of ProductAnalyticsJourneySearchGraphFilterOperator.
const (
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTEROPERATOR_EQUAL                 ProductAnalyticsJourneySearchGraphFilterOperator = "="
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTEROPERATOR_LESS_THAN             ProductAnalyticsJourneySearchGraphFilterOperator = "<"
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTEROPERATOR_GREATER_THAN          ProductAnalyticsJourneySearchGraphFilterOperator = ">"
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTEROPERATOR_LESS_THAN_OR_EQUAL    ProductAnalyticsJourneySearchGraphFilterOperator = "<="
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTEROPERATOR_GREATER_THAN_OR_EQUAL ProductAnalyticsJourneySearchGraphFilterOperator = ">="
)

var allowedProductAnalyticsJourneySearchGraphFilterOperatorEnumValues = []ProductAnalyticsJourneySearchGraphFilterOperator{
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTEROPERATOR_EQUAL,
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTEROPERATOR_LESS_THAN,
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTEROPERATOR_GREATER_THAN,
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTEROPERATOR_LESS_THAN_OR_EQUAL,
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTEROPERATOR_GREATER_THAN_OR_EQUAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneySearchGraphFilterOperator) GetAllowedValues() []ProductAnalyticsJourneySearchGraphFilterOperator {
	return allowedProductAnalyticsJourneySearchGraphFilterOperatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneySearchGraphFilterOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneySearchGraphFilterOperator(value)
	return nil
}

// NewProductAnalyticsJourneySearchGraphFilterOperatorFromValue returns a pointer to a valid ProductAnalyticsJourneySearchGraphFilterOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneySearchGraphFilterOperatorFromValue(v string) (*ProductAnalyticsJourneySearchGraphFilterOperator, error) {
	ev := ProductAnalyticsJourneySearchGraphFilterOperator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneySearchGraphFilterOperator: valid values are %v", v, allowedProductAnalyticsJourneySearchGraphFilterOperatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneySearchGraphFilterOperator) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneySearchGraphFilterOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneySearchGraphFilterOperator value.
func (v ProductAnalyticsJourneySearchGraphFilterOperator) Ptr() *ProductAnalyticsJourneySearchGraphFilterOperator {
	return &v
}

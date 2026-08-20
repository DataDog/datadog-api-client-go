// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionReturnCondition When an entity counts as having returned. Use `conversion_on` to count only entities that
// returned during the period itself, or `conversion_on_or_after` to also count later returns.
type ProductAnalyticsRetentionReturnCondition string

// List of ProductAnalyticsRetentionReturnCondition.
const (
	PRODUCTANALYTICSRETENTIONRETURNCONDITION_CONVERSION_ON          ProductAnalyticsRetentionReturnCondition = "conversion_on"
	PRODUCTANALYTICSRETENTIONRETURNCONDITION_CONVERSION_ON_OR_AFTER ProductAnalyticsRetentionReturnCondition = "conversion_on_or_after"
)

var allowedProductAnalyticsRetentionReturnConditionEnumValues = []ProductAnalyticsRetentionReturnCondition{
	PRODUCTANALYTICSRETENTIONRETURNCONDITION_CONVERSION_ON,
	PRODUCTANALYTICSRETENTIONRETURNCONDITION_CONVERSION_ON_OR_AFTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsRetentionReturnCondition) GetAllowedValues() []ProductAnalyticsRetentionReturnCondition {
	return allowedProductAnalyticsRetentionReturnConditionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsRetentionReturnCondition) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsRetentionReturnCondition(value)
	return nil
}

// NewProductAnalyticsRetentionReturnConditionFromValue returns a pointer to a valid ProductAnalyticsRetentionReturnCondition
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsRetentionReturnConditionFromValue(v string) (*ProductAnalyticsRetentionReturnCondition, error) {
	ev := ProductAnalyticsRetentionReturnCondition(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsRetentionReturnCondition: valid values are %v", v, allowedProductAnalyticsRetentionReturnConditionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsRetentionReturnCondition) IsValid() bool {
	for _, existing := range allowedProductAnalyticsRetentionReturnConditionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsRetentionReturnCondition value.
func (v ProductAnalyticsRetentionReturnCondition) Ptr() *ProductAnalyticsRetentionReturnCondition {
	return &v
}

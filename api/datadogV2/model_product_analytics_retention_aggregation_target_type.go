// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionAggregationTargetType The discriminator identifying a target selected by aggregation.
type ProductAnalyticsRetentionAggregationTargetType string

// List of ProductAnalyticsRetentionAggregationTargetType.
const (
	PRODUCTANALYTICSRETENTIONAGGREGATIONTARGETTYPE_AGGREGATION ProductAnalyticsRetentionAggregationTargetType = "aggregation"
)

var allowedProductAnalyticsRetentionAggregationTargetTypeEnumValues = []ProductAnalyticsRetentionAggregationTargetType{
	PRODUCTANALYTICSRETENTIONAGGREGATIONTARGETTYPE_AGGREGATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsRetentionAggregationTargetType) GetAllowedValues() []ProductAnalyticsRetentionAggregationTargetType {
	return allowedProductAnalyticsRetentionAggregationTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsRetentionAggregationTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsRetentionAggregationTargetType(value)
	return nil
}

// NewProductAnalyticsRetentionAggregationTargetTypeFromValue returns a pointer to a valid ProductAnalyticsRetentionAggregationTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsRetentionAggregationTargetTypeFromValue(v string) (*ProductAnalyticsRetentionAggregationTargetType, error) {
	ev := ProductAnalyticsRetentionAggregationTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsRetentionAggregationTargetType: valid values are %v", v, allowedProductAnalyticsRetentionAggregationTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsRetentionAggregationTargetType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsRetentionAggregationTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsRetentionAggregationTargetType value.
func (v ProductAnalyticsRetentionAggregationTargetType) Ptr() *ProductAnalyticsRetentionAggregationTargetType {
	return &v
}

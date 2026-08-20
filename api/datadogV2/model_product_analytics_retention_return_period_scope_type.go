// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionReturnPeriodScopeType The discriminator identifying a scope narrowed to one return period.
type ProductAnalyticsRetentionReturnPeriodScopeType string

// List of ProductAnalyticsRetentionReturnPeriodScopeType.
const (
	PRODUCTANALYTICSRETENTIONRETURNPERIODSCOPETYPE_RETURN_PERIOD ProductAnalyticsRetentionReturnPeriodScopeType = "return_period"
)

var allowedProductAnalyticsRetentionReturnPeriodScopeTypeEnumValues = []ProductAnalyticsRetentionReturnPeriodScopeType{
	PRODUCTANALYTICSRETENTIONRETURNPERIODSCOPETYPE_RETURN_PERIOD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsRetentionReturnPeriodScopeType) GetAllowedValues() []ProductAnalyticsRetentionReturnPeriodScopeType {
	return allowedProductAnalyticsRetentionReturnPeriodScopeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsRetentionReturnPeriodScopeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsRetentionReturnPeriodScopeType(value)
	return nil
}

// NewProductAnalyticsRetentionReturnPeriodScopeTypeFromValue returns a pointer to a valid ProductAnalyticsRetentionReturnPeriodScopeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsRetentionReturnPeriodScopeTypeFromValue(v string) (*ProductAnalyticsRetentionReturnPeriodScopeType, error) {
	ev := ProductAnalyticsRetentionReturnPeriodScopeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsRetentionReturnPeriodScopeType: valid values are %v", v, allowedProductAnalyticsRetentionReturnPeriodScopeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsRetentionReturnPeriodScopeType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsRetentionReturnPeriodScopeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsRetentionReturnPeriodScopeType value.
func (v ProductAnalyticsRetentionReturnPeriodScopeType) Ptr() *ProductAnalyticsRetentionReturnPeriodScopeType {
	return &v
}

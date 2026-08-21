// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionCohortScopeType The discriminator identifying a scope narrowed to one cohort.
type ProductAnalyticsRetentionCohortScopeType string

// List of ProductAnalyticsRetentionCohortScopeType.
const (
	PRODUCTANALYTICSRETENTIONCOHORTSCOPETYPE_COHORT ProductAnalyticsRetentionCohortScopeType = "cohort"
)

var allowedProductAnalyticsRetentionCohortScopeTypeEnumValues = []ProductAnalyticsRetentionCohortScopeType{
	PRODUCTANALYTICSRETENTIONCOHORTSCOPETYPE_COHORT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsRetentionCohortScopeType) GetAllowedValues() []ProductAnalyticsRetentionCohortScopeType {
	return allowedProductAnalyticsRetentionCohortScopeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsRetentionCohortScopeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsRetentionCohortScopeType(value)
	return nil
}

// NewProductAnalyticsRetentionCohortScopeTypeFromValue returns a pointer to a valid ProductAnalyticsRetentionCohortScopeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsRetentionCohortScopeTypeFromValue(v string) (*ProductAnalyticsRetentionCohortScopeType, error) {
	ev := ProductAnalyticsRetentionCohortScopeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsRetentionCohortScopeType: valid values are %v", v, allowedProductAnalyticsRetentionCohortScopeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsRetentionCohortScopeType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsRetentionCohortScopeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsRetentionCohortScopeType value.
func (v ProductAnalyticsRetentionCohortScopeType) Ptr() *ProductAnalyticsRetentionCohortScopeType {
	return &v
}

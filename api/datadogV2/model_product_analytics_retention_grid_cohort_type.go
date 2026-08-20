// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionGridCohortType Whether the row holds one cohort's own numbers, or the weighted roll-up across every cohort.
type ProductAnalyticsRetentionGridCohortType string

// List of ProductAnalyticsRetentionGridCohortType.
const (
	PRODUCTANALYTICSRETENTIONGRIDCOHORTTYPE_RAW        ProductAnalyticsRetentionGridCohortType = "raw"
	PRODUCTANALYTICSRETENTIONGRIDCOHORTTYPE_AGGREGATED ProductAnalyticsRetentionGridCohortType = "aggregated"
)

var allowedProductAnalyticsRetentionGridCohortTypeEnumValues = []ProductAnalyticsRetentionGridCohortType{
	PRODUCTANALYTICSRETENTIONGRIDCOHORTTYPE_RAW,
	PRODUCTANALYTICSRETENTIONGRIDCOHORTTYPE_AGGREGATED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsRetentionGridCohortType) GetAllowedValues() []ProductAnalyticsRetentionGridCohortType {
	return allowedProductAnalyticsRetentionGridCohortTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsRetentionGridCohortType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsRetentionGridCohortType(value)
	return nil
}

// NewProductAnalyticsRetentionGridCohortTypeFromValue returns a pointer to a valid ProductAnalyticsRetentionGridCohortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsRetentionGridCohortTypeFromValue(v string) (*ProductAnalyticsRetentionGridCohortType, error) {
	ev := ProductAnalyticsRetentionGridCohortType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsRetentionGridCohortType: valid values are %v", v, allowedProductAnalyticsRetentionGridCohortTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsRetentionGridCohortType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsRetentionGridCohortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsRetentionGridCohortType value.
func (v ProductAnalyticsRetentionGridCohortType) Ptr() *ProductAnalyticsRetentionGridCohortType {
	return &v
}

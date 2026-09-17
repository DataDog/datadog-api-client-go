// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionGroupByTarget Which axis of the retention grid a group-by applies to.
type ProductAnalyticsRetentionGroupByTarget string

// List of ProductAnalyticsRetentionGroupByTarget.
const (
	PRODUCTANALYTICSRETENTIONGROUPBYTARGET_COHORT        ProductAnalyticsRetentionGroupByTarget = "cohort"
	PRODUCTANALYTICSRETENTIONGROUPBYTARGET_RETURN_PERIOD ProductAnalyticsRetentionGroupByTarget = "return_period"
)

var allowedProductAnalyticsRetentionGroupByTargetEnumValues = []ProductAnalyticsRetentionGroupByTarget{
	PRODUCTANALYTICSRETENTIONGROUPBYTARGET_COHORT,
	PRODUCTANALYTICSRETENTIONGROUPBYTARGET_RETURN_PERIOD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsRetentionGroupByTarget) GetAllowedValues() []ProductAnalyticsRetentionGroupByTarget {
	return allowedProductAnalyticsRetentionGroupByTargetEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsRetentionGroupByTarget) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsRetentionGroupByTarget(value)
	return nil
}

// NewProductAnalyticsRetentionGroupByTargetFromValue returns a pointer to a valid ProductAnalyticsRetentionGroupByTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsRetentionGroupByTargetFromValue(v string) (*ProductAnalyticsRetentionGroupByTarget, error) {
	ev := ProductAnalyticsRetentionGroupByTarget(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsRetentionGroupByTarget: valid values are %v", v, allowedProductAnalyticsRetentionGroupByTargetEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsRetentionGroupByTarget) IsValid() bool {
	for _, existing := range allowedProductAnalyticsRetentionGroupByTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsRetentionGroupByTarget value.
func (v ProductAnalyticsRetentionGroupByTarget) Ptr() *ProductAnalyticsRetentionGroupByTarget {
	return &v
}

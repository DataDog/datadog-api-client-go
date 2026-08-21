// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionFixedTimeIntervalUnit Time unit for a fixed-length retention interval.
type ProductAnalyticsRetentionFixedTimeIntervalUnit string

// List of ProductAnalyticsRetentionFixedTimeIntervalUnit.
const (
	PRODUCTANALYTICSRETENTIONFIXEDTIMEINTERVALUNIT_DAY   ProductAnalyticsRetentionFixedTimeIntervalUnit = "day"
	PRODUCTANALYTICSRETENTIONFIXEDTIMEINTERVALUNIT_WEEK  ProductAnalyticsRetentionFixedTimeIntervalUnit = "week"
	PRODUCTANALYTICSRETENTIONFIXEDTIMEINTERVALUNIT_MONTH ProductAnalyticsRetentionFixedTimeIntervalUnit = "month"
)

var allowedProductAnalyticsRetentionFixedTimeIntervalUnitEnumValues = []ProductAnalyticsRetentionFixedTimeIntervalUnit{
	PRODUCTANALYTICSRETENTIONFIXEDTIMEINTERVALUNIT_DAY,
	PRODUCTANALYTICSRETENTIONFIXEDTIMEINTERVALUNIT_WEEK,
	PRODUCTANALYTICSRETENTIONFIXEDTIMEINTERVALUNIT_MONTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsRetentionFixedTimeIntervalUnit) GetAllowedValues() []ProductAnalyticsRetentionFixedTimeIntervalUnit {
	return allowedProductAnalyticsRetentionFixedTimeIntervalUnitEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsRetentionFixedTimeIntervalUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsRetentionFixedTimeIntervalUnit(value)
	return nil
}

// NewProductAnalyticsRetentionFixedTimeIntervalUnitFromValue returns a pointer to a valid ProductAnalyticsRetentionFixedTimeIntervalUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsRetentionFixedTimeIntervalUnitFromValue(v string) (*ProductAnalyticsRetentionFixedTimeIntervalUnit, error) {
	ev := ProductAnalyticsRetentionFixedTimeIntervalUnit(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsRetentionFixedTimeIntervalUnit: valid values are %v", v, allowedProductAnalyticsRetentionFixedTimeIntervalUnitEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsRetentionFixedTimeIntervalUnit) IsValid() bool {
	for _, existing := range allowedProductAnalyticsRetentionFixedTimeIntervalUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsRetentionFixedTimeIntervalUnit value.
func (v ProductAnalyticsRetentionFixedTimeIntervalUnit) Ptr() *ProductAnalyticsRetentionFixedTimeIntervalUnit {
	return &v
}

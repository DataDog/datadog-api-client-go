// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsCalendarIntervalType Calendar unit used to bucket cohorts.
type ProductAnalyticsCalendarIntervalType string

// List of ProductAnalyticsCalendarIntervalType.
const (
	PRODUCTANALYTICSCALENDARINTERVALTYPE_MINUTE  ProductAnalyticsCalendarIntervalType = "minute"
	PRODUCTANALYTICSCALENDARINTERVALTYPE_HOUR    ProductAnalyticsCalendarIntervalType = "hour"
	PRODUCTANALYTICSCALENDARINTERVALTYPE_DAY     ProductAnalyticsCalendarIntervalType = "day"
	PRODUCTANALYTICSCALENDARINTERVALTYPE_WEEK    ProductAnalyticsCalendarIntervalType = "week"
	PRODUCTANALYTICSCALENDARINTERVALTYPE_MONTH   ProductAnalyticsCalendarIntervalType = "month"
	PRODUCTANALYTICSCALENDARINTERVALTYPE_QUARTER ProductAnalyticsCalendarIntervalType = "quarter"
	PRODUCTANALYTICSCALENDARINTERVALTYPE_YEAR    ProductAnalyticsCalendarIntervalType = "year"
)

var allowedProductAnalyticsCalendarIntervalTypeEnumValues = []ProductAnalyticsCalendarIntervalType{
	PRODUCTANALYTICSCALENDARINTERVALTYPE_MINUTE,
	PRODUCTANALYTICSCALENDARINTERVALTYPE_HOUR,
	PRODUCTANALYTICSCALENDARINTERVALTYPE_DAY,
	PRODUCTANALYTICSCALENDARINTERVALTYPE_WEEK,
	PRODUCTANALYTICSCALENDARINTERVALTYPE_MONTH,
	PRODUCTANALYTICSCALENDARINTERVALTYPE_QUARTER,
	PRODUCTANALYTICSCALENDARINTERVALTYPE_YEAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsCalendarIntervalType) GetAllowedValues() []ProductAnalyticsCalendarIntervalType {
	return allowedProductAnalyticsCalendarIntervalTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsCalendarIntervalType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsCalendarIntervalType(value)
	return nil
}

// NewProductAnalyticsCalendarIntervalTypeFromValue returns a pointer to a valid ProductAnalyticsCalendarIntervalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsCalendarIntervalTypeFromValue(v string) (*ProductAnalyticsCalendarIntervalType, error) {
	ev := ProductAnalyticsCalendarIntervalType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsCalendarIntervalType: valid values are %v", v, allowedProductAnalyticsCalendarIntervalTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsCalendarIntervalType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsCalendarIntervalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsCalendarIntervalType value.
func (v ProductAnalyticsCalendarIntervalType) Ptr() *ProductAnalyticsCalendarIntervalType {
	return &v
}

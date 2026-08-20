// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionCalendarTimeIntervalType The discriminator identifying a calendar-aligned retention interval.
type ProductAnalyticsRetentionCalendarTimeIntervalType string

// List of ProductAnalyticsRetentionCalendarTimeIntervalType.
const (
	PRODUCTANALYTICSRETENTIONCALENDARTIMEINTERVALTYPE_CALENDAR ProductAnalyticsRetentionCalendarTimeIntervalType = "calendar"
)

var allowedProductAnalyticsRetentionCalendarTimeIntervalTypeEnumValues = []ProductAnalyticsRetentionCalendarTimeIntervalType{
	PRODUCTANALYTICSRETENTIONCALENDARTIMEINTERVALTYPE_CALENDAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsRetentionCalendarTimeIntervalType) GetAllowedValues() []ProductAnalyticsRetentionCalendarTimeIntervalType {
	return allowedProductAnalyticsRetentionCalendarTimeIntervalTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsRetentionCalendarTimeIntervalType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsRetentionCalendarTimeIntervalType(value)
	return nil
}

// NewProductAnalyticsRetentionCalendarTimeIntervalTypeFromValue returns a pointer to a valid ProductAnalyticsRetentionCalendarTimeIntervalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsRetentionCalendarTimeIntervalTypeFromValue(v string) (*ProductAnalyticsRetentionCalendarTimeIntervalType, error) {
	ev := ProductAnalyticsRetentionCalendarTimeIntervalType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsRetentionCalendarTimeIntervalType: valid values are %v", v, allowedProductAnalyticsRetentionCalendarTimeIntervalTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsRetentionCalendarTimeIntervalType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsRetentionCalendarTimeIntervalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsRetentionCalendarTimeIntervalType value.
func (v ProductAnalyticsRetentionCalendarTimeIntervalType) Ptr() *ProductAnalyticsRetentionCalendarTimeIntervalType {
	return &v
}

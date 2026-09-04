// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetCalendarAlignedSpanType Calendar-aligned time span type.
type WidgetCalendarAlignedSpanType string

// List of WidgetCalendarAlignedSpanType.
const (
	WIDGETCALENDARALIGNEDSPANTYPE_DAILY   WidgetCalendarAlignedSpanType = "daily"
	WIDGETCALENDARALIGNEDSPANTYPE_WEEKLY  WidgetCalendarAlignedSpanType = "weekly"
	WIDGETCALENDARALIGNEDSPANTYPE_MONTHLY WidgetCalendarAlignedSpanType = "monthly"
	WIDGETCALENDARALIGNEDSPANTYPE_YEARLY  WidgetCalendarAlignedSpanType = "yearly"
)

var allowedWidgetCalendarAlignedSpanTypeEnumValues = []WidgetCalendarAlignedSpanType{
	WIDGETCALENDARALIGNEDSPANTYPE_DAILY,
	WIDGETCALENDARALIGNEDSPANTYPE_WEEKLY,
	WIDGETCALENDARALIGNEDSPANTYPE_MONTHLY,
	WIDGETCALENDARALIGNEDSPANTYPE_YEARLY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetCalendarAlignedSpanType) GetAllowedValues() []WidgetCalendarAlignedSpanType {
	return allowedWidgetCalendarAlignedSpanTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetCalendarAlignedSpanType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetCalendarAlignedSpanType(value)
	return nil
}

// NewWidgetCalendarAlignedSpanTypeFromValue returns a pointer to a valid WidgetCalendarAlignedSpanType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetCalendarAlignedSpanTypeFromValue(v string) (*WidgetCalendarAlignedSpanType, error) {
	ev := WidgetCalendarAlignedSpanType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetCalendarAlignedSpanType: valid values are %v", v, allowedWidgetCalendarAlignedSpanTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetCalendarAlignedSpanType) IsValid() bool {
	for _, existing := range allowedWidgetCalendarAlignedSpanTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetCalendarAlignedSpanType value.
func (v WidgetCalendarAlignedSpanType) Ptr() *WidgetCalendarAlignedSpanType {
	return &v
}

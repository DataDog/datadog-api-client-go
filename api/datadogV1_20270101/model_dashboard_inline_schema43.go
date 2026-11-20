// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema43 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema43 string

// List of DashboardInlineSchema43.
const (
	DASHBOARDINLINESCHEMA43_GLOBAL        DashboardInlineSchema43 = "global"
	DASHBOARDINLINESCHEMA43_ALERT         DashboardInlineSchema43 = "alert"
	DASHBOARDINLINESCHEMA43_FULL_WEEK     DashboardInlineSchema43 = "full_week"
	DASHBOARDINLINESCHEMA43_FULL_MONTH    DashboardInlineSchema43 = "full_month"
	DASHBOARDINLINESCHEMA43_WEEK_TO_DATE  DashboardInlineSchema43 = "week_to_date"
	DASHBOARDINLINESCHEMA43_MONTH_TO_DATE DashboardInlineSchema43 = "month_to_date"
	DASHBOARDINLINESCHEMA43_YEAR_TO_DATE  DashboardInlineSchema43 = "year_to_date"
)

var allowedDashboardInlineSchema43EnumValues = []DashboardInlineSchema43{
	DASHBOARDINLINESCHEMA43_GLOBAL,
	DASHBOARDINLINESCHEMA43_ALERT,
	DASHBOARDINLINESCHEMA43_FULL_WEEK,
	DASHBOARDINLINESCHEMA43_FULL_MONTH,
	DASHBOARDINLINESCHEMA43_WEEK_TO_DATE,
	DASHBOARDINLINESCHEMA43_MONTH_TO_DATE,
	DASHBOARDINLINESCHEMA43_YEAR_TO_DATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema43) GetAllowedValues() []DashboardInlineSchema43 {
	return allowedDashboardInlineSchema43EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema43) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema43(value)
	return nil
}

// NewDashboardInlineSchema43FromValue returns a pointer to a valid DashboardInlineSchema43
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema43FromValue(v string) (*DashboardInlineSchema43, error) {
	ev := DashboardInlineSchema43(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema43: valid values are %v", v, allowedDashboardInlineSchema43EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema43) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema43EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema43 value.
func (v DashboardInlineSchema43) Ptr() *DashboardInlineSchema43 {
	return &v
}

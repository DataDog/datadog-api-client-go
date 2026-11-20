// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema46 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema46 string

// List of DashboardInlineSchema46.
const (
	DASHBOARDINLINESCHEMA46_BURN_RATE              DashboardInlineSchema46 = "burn_rate"
	DASHBOARDINLINESCHEMA46_ERROR_BUDGET_BURNDOWN  DashboardInlineSchema46 = "error_budget_burndown"
	DASHBOARDINLINESCHEMA46_ERROR_BUDGET_REMAINING DashboardInlineSchema46 = "error_budget_remaining"
	DASHBOARDINLINESCHEMA46_SLO_STATUS             DashboardInlineSchema46 = "slo_status"
)

var allowedDashboardInlineSchema46EnumValues = []DashboardInlineSchema46{
	DASHBOARDINLINESCHEMA46_BURN_RATE,
	DASHBOARDINLINESCHEMA46_ERROR_BUDGET_BURNDOWN,
	DASHBOARDINLINESCHEMA46_ERROR_BUDGET_REMAINING,
	DASHBOARDINLINESCHEMA46_SLO_STATUS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema46) GetAllowedValues() []DashboardInlineSchema46 {
	return allowedDashboardInlineSchema46EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema46) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema46(value)
	return nil
}

// NewDashboardInlineSchema46FromValue returns a pointer to a valid DashboardInlineSchema46
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema46FromValue(v string) (*DashboardInlineSchema46, error) {
	ev := DashboardInlineSchema46(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema46: valid values are %v", v, allowedDashboardInlineSchema46EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema46) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema46EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema46 value.
func (v DashboardInlineSchema46) Ptr() *DashboardInlineSchema46 {
	return &v
}

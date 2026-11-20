// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema26 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema26 string

// List of DashboardInlineSchema26.
const (
	DASHBOARDINLINESCHEMA26_AUTO        DashboardInlineSchema26 = "auto"
	DASHBOARDINLINESCHEMA26_COMPACT     DashboardInlineSchema26 = "compact"
	DASHBOARDINLINESCHEMA26_FULL        DashboardInlineSchema26 = "full"
	DASHBOARDINLINESCHEMA26_FIT_CONTENT DashboardInlineSchema26 = "fit-content"
)

var allowedDashboardInlineSchema26EnumValues = []DashboardInlineSchema26{
	DASHBOARDINLINESCHEMA26_AUTO,
	DASHBOARDINLINESCHEMA26_COMPACT,
	DASHBOARDINLINESCHEMA26_FULL,
	DASHBOARDINLINESCHEMA26_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema26) GetAllowedValues() []DashboardInlineSchema26 {
	return allowedDashboardInlineSchema26EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema26) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema26(value)
	return nil
}

// NewDashboardInlineSchema26FromValue returns a pointer to a valid DashboardInlineSchema26
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema26FromValue(v string) (*DashboardInlineSchema26, error) {
	ev := DashboardInlineSchema26(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema26: valid values are %v", v, allowedDashboardInlineSchema26EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema26) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema26EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema26 value.
func (v DashboardInlineSchema26) Ptr() *DashboardInlineSchema26 {
	return &v
}

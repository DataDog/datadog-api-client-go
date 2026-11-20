// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema9 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema9 string

// List of DashboardInlineSchema9.
const (
	DASHBOARDINLINESCHEMA9_AUTO        DashboardInlineSchema9 = "auto"
	DASHBOARDINLINESCHEMA9_COMPACT     DashboardInlineSchema9 = "compact"
	DASHBOARDINLINESCHEMA9_FULL        DashboardInlineSchema9 = "full"
	DASHBOARDINLINESCHEMA9_FIT_CONTENT DashboardInlineSchema9 = "fit-content"
)

var allowedDashboardInlineSchema9EnumValues = []DashboardInlineSchema9{
	DASHBOARDINLINESCHEMA9_AUTO,
	DASHBOARDINLINESCHEMA9_COMPACT,
	DASHBOARDINLINESCHEMA9_FULL,
	DASHBOARDINLINESCHEMA9_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema9) GetAllowedValues() []DashboardInlineSchema9 {
	return allowedDashboardInlineSchema9EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema9) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema9(value)
	return nil
}

// NewDashboardInlineSchema9FromValue returns a pointer to a valid DashboardInlineSchema9
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema9FromValue(v string) (*DashboardInlineSchema9, error) {
	ev := DashboardInlineSchema9(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema9: valid values are %v", v, allowedDashboardInlineSchema9EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema9) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema9EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema9 value.
func (v DashboardInlineSchema9) Ptr() *DashboardInlineSchema9 {
	return &v
}

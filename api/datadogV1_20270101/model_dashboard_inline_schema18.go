// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema18 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema18 string

// List of DashboardInlineSchema18.
const (
	DASHBOARDINLINESCHEMA18_AUTO        DashboardInlineSchema18 = "auto"
	DASHBOARDINLINESCHEMA18_COMPACT     DashboardInlineSchema18 = "compact"
	DASHBOARDINLINESCHEMA18_FULL        DashboardInlineSchema18 = "full"
	DASHBOARDINLINESCHEMA18_FIT_CONTENT DashboardInlineSchema18 = "fit-content"
)

var allowedDashboardInlineSchema18EnumValues = []DashboardInlineSchema18{
	DASHBOARDINLINESCHEMA18_AUTO,
	DASHBOARDINLINESCHEMA18_COMPACT,
	DASHBOARDINLINESCHEMA18_FULL,
	DASHBOARDINLINESCHEMA18_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema18) GetAllowedValues() []DashboardInlineSchema18 {
	return allowedDashboardInlineSchema18EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema18) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema18(value)
	return nil
}

// NewDashboardInlineSchema18FromValue returns a pointer to a valid DashboardInlineSchema18
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema18FromValue(v string) (*DashboardInlineSchema18, error) {
	ev := DashboardInlineSchema18(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema18: valid values are %v", v, allowedDashboardInlineSchema18EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema18) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema18EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema18 value.
func (v DashboardInlineSchema18) Ptr() *DashboardInlineSchema18 {
	return &v
}

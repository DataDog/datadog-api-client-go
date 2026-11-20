// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema29 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema29 string

// List of DashboardInlineSchema29.
const (
	DASHBOARDINLINESCHEMA29_AUTO        DashboardInlineSchema29 = "auto"
	DASHBOARDINLINESCHEMA29_COMPACT     DashboardInlineSchema29 = "compact"
	DASHBOARDINLINESCHEMA29_FULL        DashboardInlineSchema29 = "full"
	DASHBOARDINLINESCHEMA29_FIT_CONTENT DashboardInlineSchema29 = "fit-content"
)

var allowedDashboardInlineSchema29EnumValues = []DashboardInlineSchema29{
	DASHBOARDINLINESCHEMA29_AUTO,
	DASHBOARDINLINESCHEMA29_COMPACT,
	DASHBOARDINLINESCHEMA29_FULL,
	DASHBOARDINLINESCHEMA29_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema29) GetAllowedValues() []DashboardInlineSchema29 {
	return allowedDashboardInlineSchema29EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema29) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema29(value)
	return nil
}

// NewDashboardInlineSchema29FromValue returns a pointer to a valid DashboardInlineSchema29
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema29FromValue(v string) (*DashboardInlineSchema29, error) {
	ev := DashboardInlineSchema29(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema29: valid values are %v", v, allowedDashboardInlineSchema29EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema29) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema29EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema29 value.
func (v DashboardInlineSchema29) Ptr() *DashboardInlineSchema29 {
	return &v
}

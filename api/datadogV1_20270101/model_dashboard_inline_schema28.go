// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema28 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema28 string

// List of DashboardInlineSchema28.
const (
	DASHBOARDINLINESCHEMA28_AUTO        DashboardInlineSchema28 = "auto"
	DASHBOARDINLINESCHEMA28_COMPACT     DashboardInlineSchema28 = "compact"
	DASHBOARDINLINESCHEMA28_FULL        DashboardInlineSchema28 = "full"
	DASHBOARDINLINESCHEMA28_FIT_CONTENT DashboardInlineSchema28 = "fit-content"
)

var allowedDashboardInlineSchema28EnumValues = []DashboardInlineSchema28{
	DASHBOARDINLINESCHEMA28_AUTO,
	DASHBOARDINLINESCHEMA28_COMPACT,
	DASHBOARDINLINESCHEMA28_FULL,
	DASHBOARDINLINESCHEMA28_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema28) GetAllowedValues() []DashboardInlineSchema28 {
	return allowedDashboardInlineSchema28EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema28) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema28(value)
	return nil
}

// NewDashboardInlineSchema28FromValue returns a pointer to a valid DashboardInlineSchema28
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema28FromValue(v string) (*DashboardInlineSchema28, error) {
	ev := DashboardInlineSchema28(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema28: valid values are %v", v, allowedDashboardInlineSchema28EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema28) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema28EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema28 value.
func (v DashboardInlineSchema28) Ptr() *DashboardInlineSchema28 {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema15 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema15 string

// List of DashboardInlineSchema15.
const (
	DASHBOARDINLINESCHEMA15_AUTO        DashboardInlineSchema15 = "auto"
	DASHBOARDINLINESCHEMA15_COMPACT     DashboardInlineSchema15 = "compact"
	DASHBOARDINLINESCHEMA15_FULL        DashboardInlineSchema15 = "full"
	DASHBOARDINLINESCHEMA15_FIT_CONTENT DashboardInlineSchema15 = "fit-content"
)

var allowedDashboardInlineSchema15EnumValues = []DashboardInlineSchema15{
	DASHBOARDINLINESCHEMA15_AUTO,
	DASHBOARDINLINESCHEMA15_COMPACT,
	DASHBOARDINLINESCHEMA15_FULL,
	DASHBOARDINLINESCHEMA15_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema15) GetAllowedValues() []DashboardInlineSchema15 {
	return allowedDashboardInlineSchema15EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema15) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema15(value)
	return nil
}

// NewDashboardInlineSchema15FromValue returns a pointer to a valid DashboardInlineSchema15
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema15FromValue(v string) (*DashboardInlineSchema15, error) {
	ev := DashboardInlineSchema15(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema15: valid values are %v", v, allowedDashboardInlineSchema15EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema15) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema15EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema15 value.
func (v DashboardInlineSchema15) Ptr() *DashboardInlineSchema15 {
	return &v
}

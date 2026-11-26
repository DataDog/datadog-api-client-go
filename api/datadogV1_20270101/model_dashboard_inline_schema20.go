// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema20 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema20 string

// List of DashboardInlineSchema20.
const (
	DASHBOARDINLINESCHEMA20_AUTO        DashboardInlineSchema20 = "auto"
	DASHBOARDINLINESCHEMA20_COMPACT     DashboardInlineSchema20 = "compact"
	DASHBOARDINLINESCHEMA20_FULL        DashboardInlineSchema20 = "full"
	DASHBOARDINLINESCHEMA20_FIT_CONTENT DashboardInlineSchema20 = "fit-content"
)

var allowedDashboardInlineSchema20EnumValues = []DashboardInlineSchema20{
	DASHBOARDINLINESCHEMA20_AUTO,
	DASHBOARDINLINESCHEMA20_COMPACT,
	DASHBOARDINLINESCHEMA20_FULL,
	DASHBOARDINLINESCHEMA20_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema20) GetAllowedValues() []DashboardInlineSchema20 {
	return allowedDashboardInlineSchema20EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema20) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema20(value)
	return nil
}

// NewDashboardInlineSchema20FromValue returns a pointer to a valid DashboardInlineSchema20
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema20FromValue(v string) (*DashboardInlineSchema20, error) {
	ev := DashboardInlineSchema20(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema20: valid values are %v", v, allowedDashboardInlineSchema20EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema20) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema20EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema20 value.
func (v DashboardInlineSchema20) Ptr() *DashboardInlineSchema20 {
	return &v
}

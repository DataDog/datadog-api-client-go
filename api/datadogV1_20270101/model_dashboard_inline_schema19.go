// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema19 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema19 string

// List of DashboardInlineSchema19.
const (
	DASHBOARDINLINESCHEMA19_AUTO        DashboardInlineSchema19 = "auto"
	DASHBOARDINLINESCHEMA19_COMPACT     DashboardInlineSchema19 = "compact"
	DASHBOARDINLINESCHEMA19_FULL        DashboardInlineSchema19 = "full"
	DASHBOARDINLINESCHEMA19_FIT_CONTENT DashboardInlineSchema19 = "fit-content"
)

var allowedDashboardInlineSchema19EnumValues = []DashboardInlineSchema19{
	DASHBOARDINLINESCHEMA19_AUTO,
	DASHBOARDINLINESCHEMA19_COMPACT,
	DASHBOARDINLINESCHEMA19_FULL,
	DASHBOARDINLINESCHEMA19_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema19) GetAllowedValues() []DashboardInlineSchema19 {
	return allowedDashboardInlineSchema19EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema19) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema19(value)
	return nil
}

// NewDashboardInlineSchema19FromValue returns a pointer to a valid DashboardInlineSchema19
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema19FromValue(v string) (*DashboardInlineSchema19, error) {
	ev := DashboardInlineSchema19(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema19: valid values are %v", v, allowedDashboardInlineSchema19EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema19) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema19EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema19 value.
func (v DashboardInlineSchema19) Ptr() *DashboardInlineSchema19 {
	return &v
}

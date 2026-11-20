// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema21 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema21 string

// List of DashboardInlineSchema21.
const (
	DASHBOARDINLINESCHEMA21_AUTO        DashboardInlineSchema21 = "auto"
	DASHBOARDINLINESCHEMA21_COMPACT     DashboardInlineSchema21 = "compact"
	DASHBOARDINLINESCHEMA21_FULL        DashboardInlineSchema21 = "full"
	DASHBOARDINLINESCHEMA21_FIT_CONTENT DashboardInlineSchema21 = "fit-content"
)

var allowedDashboardInlineSchema21EnumValues = []DashboardInlineSchema21{
	DASHBOARDINLINESCHEMA21_AUTO,
	DASHBOARDINLINESCHEMA21_COMPACT,
	DASHBOARDINLINESCHEMA21_FULL,
	DASHBOARDINLINESCHEMA21_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema21) GetAllowedValues() []DashboardInlineSchema21 {
	return allowedDashboardInlineSchema21EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema21) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema21(value)
	return nil
}

// NewDashboardInlineSchema21FromValue returns a pointer to a valid DashboardInlineSchema21
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema21FromValue(v string) (*DashboardInlineSchema21, error) {
	ev := DashboardInlineSchema21(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema21: valid values are %v", v, allowedDashboardInlineSchema21EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema21) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema21EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema21 value.
func (v DashboardInlineSchema21) Ptr() *DashboardInlineSchema21 {
	return &v
}

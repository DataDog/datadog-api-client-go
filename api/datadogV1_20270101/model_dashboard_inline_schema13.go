// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema13 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema13 string

// List of DashboardInlineSchema13.
const (
	DASHBOARDINLINESCHEMA13_AUTO        DashboardInlineSchema13 = "auto"
	DASHBOARDINLINESCHEMA13_COMPACT     DashboardInlineSchema13 = "compact"
	DASHBOARDINLINESCHEMA13_FULL        DashboardInlineSchema13 = "full"
	DASHBOARDINLINESCHEMA13_FIT_CONTENT DashboardInlineSchema13 = "fit-content"
)

var allowedDashboardInlineSchema13EnumValues = []DashboardInlineSchema13{
	DASHBOARDINLINESCHEMA13_AUTO,
	DASHBOARDINLINESCHEMA13_COMPACT,
	DASHBOARDINLINESCHEMA13_FULL,
	DASHBOARDINLINESCHEMA13_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema13) GetAllowedValues() []DashboardInlineSchema13 {
	return allowedDashboardInlineSchema13EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema13) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema13(value)
	return nil
}

// NewDashboardInlineSchema13FromValue returns a pointer to a valid DashboardInlineSchema13
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema13FromValue(v string) (*DashboardInlineSchema13, error) {
	ev := DashboardInlineSchema13(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema13: valid values are %v", v, allowedDashboardInlineSchema13EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema13) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema13EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema13 value.
func (v DashboardInlineSchema13) Ptr() *DashboardInlineSchema13 {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema16 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema16 string

// List of DashboardInlineSchema16.
const (
	DASHBOARDINLINESCHEMA16_AUTO        DashboardInlineSchema16 = "auto"
	DASHBOARDINLINESCHEMA16_COMPACT     DashboardInlineSchema16 = "compact"
	DASHBOARDINLINESCHEMA16_FULL        DashboardInlineSchema16 = "full"
	DASHBOARDINLINESCHEMA16_FIT_CONTENT DashboardInlineSchema16 = "fit-content"
)

var allowedDashboardInlineSchema16EnumValues = []DashboardInlineSchema16{
	DASHBOARDINLINESCHEMA16_AUTO,
	DASHBOARDINLINESCHEMA16_COMPACT,
	DASHBOARDINLINESCHEMA16_FULL,
	DASHBOARDINLINESCHEMA16_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema16) GetAllowedValues() []DashboardInlineSchema16 {
	return allowedDashboardInlineSchema16EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema16) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema16(value)
	return nil
}

// NewDashboardInlineSchema16FromValue returns a pointer to a valid DashboardInlineSchema16
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema16FromValue(v string) (*DashboardInlineSchema16, error) {
	ev := DashboardInlineSchema16(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema16: valid values are %v", v, allowedDashboardInlineSchema16EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema16) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema16EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema16 value.
func (v DashboardInlineSchema16) Ptr() *DashboardInlineSchema16 {
	return &v
}

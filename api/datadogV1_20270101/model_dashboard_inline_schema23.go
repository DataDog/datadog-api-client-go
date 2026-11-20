// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema23 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema23 string

// List of DashboardInlineSchema23.
const (
	DASHBOARDINLINESCHEMA23_AUTO        DashboardInlineSchema23 = "auto"
	DASHBOARDINLINESCHEMA23_COMPACT     DashboardInlineSchema23 = "compact"
	DASHBOARDINLINESCHEMA23_FULL        DashboardInlineSchema23 = "full"
	DASHBOARDINLINESCHEMA23_FIT_CONTENT DashboardInlineSchema23 = "fit-content"
)

var allowedDashboardInlineSchema23EnumValues = []DashboardInlineSchema23{
	DASHBOARDINLINESCHEMA23_AUTO,
	DASHBOARDINLINESCHEMA23_COMPACT,
	DASHBOARDINLINESCHEMA23_FULL,
	DASHBOARDINLINESCHEMA23_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema23) GetAllowedValues() []DashboardInlineSchema23 {
	return allowedDashboardInlineSchema23EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema23) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema23(value)
	return nil
}

// NewDashboardInlineSchema23FromValue returns a pointer to a valid DashboardInlineSchema23
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema23FromValue(v string) (*DashboardInlineSchema23, error) {
	ev := DashboardInlineSchema23(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema23: valid values are %v", v, allowedDashboardInlineSchema23EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema23) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema23EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema23 value.
func (v DashboardInlineSchema23) Ptr() *DashboardInlineSchema23 {
	return &v
}

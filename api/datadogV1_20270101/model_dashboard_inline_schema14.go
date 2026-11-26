// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema14 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema14 string

// List of DashboardInlineSchema14.
const (
	DASHBOARDINLINESCHEMA14_AUTO        DashboardInlineSchema14 = "auto"
	DASHBOARDINLINESCHEMA14_COMPACT     DashboardInlineSchema14 = "compact"
	DASHBOARDINLINESCHEMA14_FULL        DashboardInlineSchema14 = "full"
	DASHBOARDINLINESCHEMA14_FIT_CONTENT DashboardInlineSchema14 = "fit-content"
)

var allowedDashboardInlineSchema14EnumValues = []DashboardInlineSchema14{
	DASHBOARDINLINESCHEMA14_AUTO,
	DASHBOARDINLINESCHEMA14_COMPACT,
	DASHBOARDINLINESCHEMA14_FULL,
	DASHBOARDINLINESCHEMA14_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema14) GetAllowedValues() []DashboardInlineSchema14 {
	return allowedDashboardInlineSchema14EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema14) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema14(value)
	return nil
}

// NewDashboardInlineSchema14FromValue returns a pointer to a valid DashboardInlineSchema14
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema14FromValue(v string) (*DashboardInlineSchema14, error) {
	ev := DashboardInlineSchema14(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema14: valid values are %v", v, allowedDashboardInlineSchema14EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema14) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema14EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema14 value.
func (v DashboardInlineSchema14) Ptr() *DashboardInlineSchema14 {
	return &v
}

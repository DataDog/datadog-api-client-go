// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema38 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema38 string

// List of DashboardInlineSchema38.
const (
	DASHBOARDINLINESCHEMA38_ID       DashboardInlineSchema38 = "id"
	DASHBOARDINLINESCHEMA38_TITLE    DashboardInlineSchema38 = "title"
	DASHBOARDINLINESCHEMA38_SUBTITLE DashboardInlineSchema38 = "subtitle"
)

var allowedDashboardInlineSchema38EnumValues = []DashboardInlineSchema38{
	DASHBOARDINLINESCHEMA38_ID,
	DASHBOARDINLINESCHEMA38_TITLE,
	DASHBOARDINLINESCHEMA38_SUBTITLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema38) GetAllowedValues() []DashboardInlineSchema38 {
	return allowedDashboardInlineSchema38EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema38) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema38(value)
	return nil
}

// NewDashboardInlineSchema38FromValue returns a pointer to a valid DashboardInlineSchema38
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema38FromValue(v string) (*DashboardInlineSchema38, error) {
	ev := DashboardInlineSchema38(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema38: valid values are %v", v, allowedDashboardInlineSchema38EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema38) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema38EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema38 value.
func (v DashboardInlineSchema38) Ptr() *DashboardInlineSchema38 {
	return &v
}

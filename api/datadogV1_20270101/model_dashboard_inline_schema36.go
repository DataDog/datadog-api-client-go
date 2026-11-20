// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema36 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema36 string

// List of DashboardInlineSchema36.
const (
	DASHBOARDINLINESCHEMA36_AUTO        DashboardInlineSchema36 = "auto"
	DASHBOARDINLINESCHEMA36_COMPACT     DashboardInlineSchema36 = "compact"
	DASHBOARDINLINESCHEMA36_FULL        DashboardInlineSchema36 = "full"
	DASHBOARDINLINESCHEMA36_FIT_CONTENT DashboardInlineSchema36 = "fit-content"
)

var allowedDashboardInlineSchema36EnumValues = []DashboardInlineSchema36{
	DASHBOARDINLINESCHEMA36_AUTO,
	DASHBOARDINLINESCHEMA36_COMPACT,
	DASHBOARDINLINESCHEMA36_FULL,
	DASHBOARDINLINESCHEMA36_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema36) GetAllowedValues() []DashboardInlineSchema36 {
	return allowedDashboardInlineSchema36EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema36) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema36(value)
	return nil
}

// NewDashboardInlineSchema36FromValue returns a pointer to a valid DashboardInlineSchema36
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema36FromValue(v string) (*DashboardInlineSchema36, error) {
	ev := DashboardInlineSchema36(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema36: valid values are %v", v, allowedDashboardInlineSchema36EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema36) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema36EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema36 value.
func (v DashboardInlineSchema36) Ptr() *DashboardInlineSchema36 {
	return &v
}

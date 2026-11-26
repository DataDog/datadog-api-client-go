// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema34 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema34 string

// List of DashboardInlineSchema34.
const (
	DASHBOARDINLINESCHEMA34_AUTO        DashboardInlineSchema34 = "auto"
	DASHBOARDINLINESCHEMA34_COMPACT     DashboardInlineSchema34 = "compact"
	DASHBOARDINLINESCHEMA34_FULL        DashboardInlineSchema34 = "full"
	DASHBOARDINLINESCHEMA34_FIT_CONTENT DashboardInlineSchema34 = "fit-content"
)

var allowedDashboardInlineSchema34EnumValues = []DashboardInlineSchema34{
	DASHBOARDINLINESCHEMA34_AUTO,
	DASHBOARDINLINESCHEMA34_COMPACT,
	DASHBOARDINLINESCHEMA34_FULL,
	DASHBOARDINLINESCHEMA34_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema34) GetAllowedValues() []DashboardInlineSchema34 {
	return allowedDashboardInlineSchema34EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema34) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema34(value)
	return nil
}

// NewDashboardInlineSchema34FromValue returns a pointer to a valid DashboardInlineSchema34
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema34FromValue(v string) (*DashboardInlineSchema34, error) {
	ev := DashboardInlineSchema34(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema34: valid values are %v", v, allowedDashboardInlineSchema34EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema34) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema34EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema34 value.
func (v DashboardInlineSchema34) Ptr() *DashboardInlineSchema34 {
	return &v
}

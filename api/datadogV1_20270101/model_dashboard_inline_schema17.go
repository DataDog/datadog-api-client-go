// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema17 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema17 string

// List of DashboardInlineSchema17.
const (
	DASHBOARDINLINESCHEMA17_AUTO        DashboardInlineSchema17 = "auto"
	DASHBOARDINLINESCHEMA17_COMPACT     DashboardInlineSchema17 = "compact"
	DASHBOARDINLINESCHEMA17_FULL        DashboardInlineSchema17 = "full"
	DASHBOARDINLINESCHEMA17_FIT_CONTENT DashboardInlineSchema17 = "fit-content"
)

var allowedDashboardInlineSchema17EnumValues = []DashboardInlineSchema17{
	DASHBOARDINLINESCHEMA17_AUTO,
	DASHBOARDINLINESCHEMA17_COMPACT,
	DASHBOARDINLINESCHEMA17_FULL,
	DASHBOARDINLINESCHEMA17_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema17) GetAllowedValues() []DashboardInlineSchema17 {
	return allowedDashboardInlineSchema17EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema17) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema17(value)
	return nil
}

// NewDashboardInlineSchema17FromValue returns a pointer to a valid DashboardInlineSchema17
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema17FromValue(v string) (*DashboardInlineSchema17, error) {
	ev := DashboardInlineSchema17(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema17: valid values are %v", v, allowedDashboardInlineSchema17EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema17) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema17EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema17 value.
func (v DashboardInlineSchema17) Ptr() *DashboardInlineSchema17 {
	return &v
}

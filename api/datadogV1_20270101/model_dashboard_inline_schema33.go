// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema33 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema33 string

// List of DashboardInlineSchema33.
const (
	DASHBOARDINLINESCHEMA33_AUTO        DashboardInlineSchema33 = "auto"
	DASHBOARDINLINESCHEMA33_COMPACT     DashboardInlineSchema33 = "compact"
	DASHBOARDINLINESCHEMA33_FULL        DashboardInlineSchema33 = "full"
	DASHBOARDINLINESCHEMA33_FIT_CONTENT DashboardInlineSchema33 = "fit-content"
)

var allowedDashboardInlineSchema33EnumValues = []DashboardInlineSchema33{
	DASHBOARDINLINESCHEMA33_AUTO,
	DASHBOARDINLINESCHEMA33_COMPACT,
	DASHBOARDINLINESCHEMA33_FULL,
	DASHBOARDINLINESCHEMA33_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema33) GetAllowedValues() []DashboardInlineSchema33 {
	return allowedDashboardInlineSchema33EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema33) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema33(value)
	return nil
}

// NewDashboardInlineSchema33FromValue returns a pointer to a valid DashboardInlineSchema33
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema33FromValue(v string) (*DashboardInlineSchema33, error) {
	ev := DashboardInlineSchema33(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema33: valid values are %v", v, allowedDashboardInlineSchema33EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema33) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema33EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema33 value.
func (v DashboardInlineSchema33) Ptr() *DashboardInlineSchema33 {
	return &v
}

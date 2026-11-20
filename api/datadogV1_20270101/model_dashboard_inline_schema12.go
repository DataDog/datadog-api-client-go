// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema12 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema12 string

// List of DashboardInlineSchema12.
const (
	DASHBOARDINLINESCHEMA12_AUTO        DashboardInlineSchema12 = "auto"
	DASHBOARDINLINESCHEMA12_COMPACT     DashboardInlineSchema12 = "compact"
	DASHBOARDINLINESCHEMA12_FULL        DashboardInlineSchema12 = "full"
	DASHBOARDINLINESCHEMA12_FIT_CONTENT DashboardInlineSchema12 = "fit-content"
)

var allowedDashboardInlineSchema12EnumValues = []DashboardInlineSchema12{
	DASHBOARDINLINESCHEMA12_AUTO,
	DASHBOARDINLINESCHEMA12_COMPACT,
	DASHBOARDINLINESCHEMA12_FULL,
	DASHBOARDINLINESCHEMA12_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema12) GetAllowedValues() []DashboardInlineSchema12 {
	return allowedDashboardInlineSchema12EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema12) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema12(value)
	return nil
}

// NewDashboardInlineSchema12FromValue returns a pointer to a valid DashboardInlineSchema12
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema12FromValue(v string) (*DashboardInlineSchema12, error) {
	ev := DashboardInlineSchema12(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema12: valid values are %v", v, allowedDashboardInlineSchema12EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema12) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema12EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema12 value.
func (v DashboardInlineSchema12) Ptr() *DashboardInlineSchema12 {
	return &v
}

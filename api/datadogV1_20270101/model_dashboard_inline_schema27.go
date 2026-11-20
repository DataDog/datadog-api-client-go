// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema27 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema27 string

// List of DashboardInlineSchema27.
const (
	DASHBOARDINLINESCHEMA27_AUTO        DashboardInlineSchema27 = "auto"
	DASHBOARDINLINESCHEMA27_COMPACT     DashboardInlineSchema27 = "compact"
	DASHBOARDINLINESCHEMA27_FULL        DashboardInlineSchema27 = "full"
	DASHBOARDINLINESCHEMA27_FIT_CONTENT DashboardInlineSchema27 = "fit-content"
)

var allowedDashboardInlineSchema27EnumValues = []DashboardInlineSchema27{
	DASHBOARDINLINESCHEMA27_AUTO,
	DASHBOARDINLINESCHEMA27_COMPACT,
	DASHBOARDINLINESCHEMA27_FULL,
	DASHBOARDINLINESCHEMA27_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema27) GetAllowedValues() []DashboardInlineSchema27 {
	return allowedDashboardInlineSchema27EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema27) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema27(value)
	return nil
}

// NewDashboardInlineSchema27FromValue returns a pointer to a valid DashboardInlineSchema27
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema27FromValue(v string) (*DashboardInlineSchema27, error) {
	ev := DashboardInlineSchema27(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema27: valid values are %v", v, allowedDashboardInlineSchema27EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema27) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema27EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema27 value.
func (v DashboardInlineSchema27) Ptr() *DashboardInlineSchema27 {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema25 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema25 string

// List of DashboardInlineSchema25.
const (
	DASHBOARDINLINESCHEMA25_AUTO        DashboardInlineSchema25 = "auto"
	DASHBOARDINLINESCHEMA25_COMPACT     DashboardInlineSchema25 = "compact"
	DASHBOARDINLINESCHEMA25_FULL        DashboardInlineSchema25 = "full"
	DASHBOARDINLINESCHEMA25_FIT_CONTENT DashboardInlineSchema25 = "fit-content"
)

var allowedDashboardInlineSchema25EnumValues = []DashboardInlineSchema25{
	DASHBOARDINLINESCHEMA25_AUTO,
	DASHBOARDINLINESCHEMA25_COMPACT,
	DASHBOARDINLINESCHEMA25_FULL,
	DASHBOARDINLINESCHEMA25_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema25) GetAllowedValues() []DashboardInlineSchema25 {
	return allowedDashboardInlineSchema25EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema25) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema25(value)
	return nil
}

// NewDashboardInlineSchema25FromValue returns a pointer to a valid DashboardInlineSchema25
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema25FromValue(v string) (*DashboardInlineSchema25, error) {
	ev := DashboardInlineSchema25(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema25: valid values are %v", v, allowedDashboardInlineSchema25EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema25) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema25EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema25 value.
func (v DashboardInlineSchema25) Ptr() *DashboardInlineSchema25 {
	return &v
}

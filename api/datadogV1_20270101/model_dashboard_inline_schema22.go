// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema22 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema22 string

// List of DashboardInlineSchema22.
const (
	DASHBOARDINLINESCHEMA22_AUTO        DashboardInlineSchema22 = "auto"
	DASHBOARDINLINESCHEMA22_COMPACT     DashboardInlineSchema22 = "compact"
	DASHBOARDINLINESCHEMA22_FULL        DashboardInlineSchema22 = "full"
	DASHBOARDINLINESCHEMA22_FIT_CONTENT DashboardInlineSchema22 = "fit-content"
)

var allowedDashboardInlineSchema22EnumValues = []DashboardInlineSchema22{
	DASHBOARDINLINESCHEMA22_AUTO,
	DASHBOARDINLINESCHEMA22_COMPACT,
	DASHBOARDINLINESCHEMA22_FULL,
	DASHBOARDINLINESCHEMA22_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema22) GetAllowedValues() []DashboardInlineSchema22 {
	return allowedDashboardInlineSchema22EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema22) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema22(value)
	return nil
}

// NewDashboardInlineSchema22FromValue returns a pointer to a valid DashboardInlineSchema22
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema22FromValue(v string) (*DashboardInlineSchema22, error) {
	ev := DashboardInlineSchema22(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema22: valid values are %v", v, allowedDashboardInlineSchema22EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema22) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema22EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema22 value.
func (v DashboardInlineSchema22) Ptr() *DashboardInlineSchema22 {
	return &v
}

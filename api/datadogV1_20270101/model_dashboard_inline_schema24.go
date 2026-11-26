// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema24 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema24 string

// List of DashboardInlineSchema24.
const (
	DASHBOARDINLINESCHEMA24_AUTO        DashboardInlineSchema24 = "auto"
	DASHBOARDINLINESCHEMA24_COMPACT     DashboardInlineSchema24 = "compact"
	DASHBOARDINLINESCHEMA24_FULL        DashboardInlineSchema24 = "full"
	DASHBOARDINLINESCHEMA24_FIT_CONTENT DashboardInlineSchema24 = "fit-content"
)

var allowedDashboardInlineSchema24EnumValues = []DashboardInlineSchema24{
	DASHBOARDINLINESCHEMA24_AUTO,
	DASHBOARDINLINESCHEMA24_COMPACT,
	DASHBOARDINLINESCHEMA24_FULL,
	DASHBOARDINLINESCHEMA24_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema24) GetAllowedValues() []DashboardInlineSchema24 {
	return allowedDashboardInlineSchema24EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema24) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema24(value)
	return nil
}

// NewDashboardInlineSchema24FromValue returns a pointer to a valid DashboardInlineSchema24
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema24FromValue(v string) (*DashboardInlineSchema24, error) {
	ev := DashboardInlineSchema24(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema24: valid values are %v", v, allowedDashboardInlineSchema24EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema24) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema24EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema24 value.
func (v DashboardInlineSchema24) Ptr() *DashboardInlineSchema24 {
	return &v
}

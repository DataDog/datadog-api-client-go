// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema30 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema30 string

// List of DashboardInlineSchema30.
const (
	DASHBOARDINLINESCHEMA30_AUTO        DashboardInlineSchema30 = "auto"
	DASHBOARDINLINESCHEMA30_COMPACT     DashboardInlineSchema30 = "compact"
	DASHBOARDINLINESCHEMA30_FULL        DashboardInlineSchema30 = "full"
	DASHBOARDINLINESCHEMA30_FIT_CONTENT DashboardInlineSchema30 = "fit-content"
)

var allowedDashboardInlineSchema30EnumValues = []DashboardInlineSchema30{
	DASHBOARDINLINESCHEMA30_AUTO,
	DASHBOARDINLINESCHEMA30_COMPACT,
	DASHBOARDINLINESCHEMA30_FULL,
	DASHBOARDINLINESCHEMA30_FIT_CONTENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema30) GetAllowedValues() []DashboardInlineSchema30 {
	return allowedDashboardInlineSchema30EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema30) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema30(value)
	return nil
}

// NewDashboardInlineSchema30FromValue returns a pointer to a valid DashboardInlineSchema30
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema30FromValue(v string) (*DashboardInlineSchema30, error) {
	ev := DashboardInlineSchema30(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema30: valid values are %v", v, allowedDashboardInlineSchema30EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema30) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema30EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema30 value.
func (v DashboardInlineSchema30) Ptr() *DashboardInlineSchema30 {
	return &v
}

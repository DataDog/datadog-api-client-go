// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema44 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema44 string

// List of DashboardInlineSchema44.
const (
	DASHBOARDINLINESCHEMA44_ALL     DashboardInlineSchema44 = "*"
	DASHBOARDINLINESCHEMA44_INTEGER DashboardInlineSchema44 = "integer"
)

var allowedDashboardInlineSchema44EnumValues = []DashboardInlineSchema44{
	DASHBOARDINLINESCHEMA44_ALL,
	DASHBOARDINLINESCHEMA44_INTEGER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema44) GetAllowedValues() []DashboardInlineSchema44 {
	return allowedDashboardInlineSchema44EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema44) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema44(value)
	return nil
}

// NewDashboardInlineSchema44FromValue returns a pointer to a valid DashboardInlineSchema44
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema44FromValue(v string) (*DashboardInlineSchema44, error) {
	ev := DashboardInlineSchema44(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema44: valid values are %v", v, allowedDashboardInlineSchema44EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema44) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema44EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema44 value.
func (v DashboardInlineSchema44) Ptr() *DashboardInlineSchema44 {
	return &v
}

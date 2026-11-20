// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema37 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema37 string

// List of DashboardInlineSchema37.
const (
	DASHBOARDINLINESCHEMA37_PROCESS   DashboardInlineSchema37 = "process"
	DASHBOARDINLINESCHEMA37_CONTAINER DashboardInlineSchema37 = "container"
)

var allowedDashboardInlineSchema37EnumValues = []DashboardInlineSchema37{
	DASHBOARDINLINESCHEMA37_PROCESS,
	DASHBOARDINLINESCHEMA37_CONTAINER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema37) GetAllowedValues() []DashboardInlineSchema37 {
	return allowedDashboardInlineSchema37EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema37) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema37(value)
	return nil
}

// NewDashboardInlineSchema37FromValue returns a pointer to a valid DashboardInlineSchema37
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema37FromValue(v string) (*DashboardInlineSchema37, error) {
	ev := DashboardInlineSchema37(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema37: valid values are %v", v, allowedDashboardInlineSchema37EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema37) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema37EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema37 value.
func (v DashboardInlineSchema37) Ptr() *DashboardInlineSchema37 {
	return &v
}

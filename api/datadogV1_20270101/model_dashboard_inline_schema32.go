// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema32 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema32 string

// List of DashboardInlineSchema32.
const (
	DASHBOARDINLINESCHEMA32_GROUP DashboardInlineSchema32 = "group"
)

var allowedDashboardInlineSchema32EnumValues = []DashboardInlineSchema32{
	DASHBOARDINLINESCHEMA32_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema32) GetAllowedValues() []DashboardInlineSchema32 {
	return allowedDashboardInlineSchema32EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema32) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema32(value)
	return nil
}

// NewDashboardInlineSchema32FromValue returns a pointer to a valid DashboardInlineSchema32
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema32FromValue(v string) (*DashboardInlineSchema32, error) {
	ev := DashboardInlineSchema32(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema32: valid values are %v", v, allowedDashboardInlineSchema32EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema32) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema32EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema32 value.
func (v DashboardInlineSchema32) Ptr() *DashboardInlineSchema32 {
	return &v
}

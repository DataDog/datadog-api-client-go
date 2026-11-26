// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema10 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema10 string

// List of DashboardInlineSchema10.
const (
	DASHBOARDINLINESCHEMA10_GROUP DashboardInlineSchema10 = "group"
	DASHBOARDINLINESCHEMA10_NONE  DashboardInlineSchema10 = "none"
)

var allowedDashboardInlineSchema10EnumValues = []DashboardInlineSchema10{
	DASHBOARDINLINESCHEMA10_GROUP,
	DASHBOARDINLINESCHEMA10_NONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema10) GetAllowedValues() []DashboardInlineSchema10 {
	return allowedDashboardInlineSchema10EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema10) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema10(value)
	return nil
}

// NewDashboardInlineSchema10FromValue returns a pointer to a valid DashboardInlineSchema10
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema10FromValue(v string) (*DashboardInlineSchema10, error) {
	ev := DashboardInlineSchema10(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema10: valid values are %v", v, allowedDashboardInlineSchema10EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema10) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema10EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema10 value.
func (v DashboardInlineSchema10) Ptr() *DashboardInlineSchema10 {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema49 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema49 string

// List of DashboardInlineSchema49.
const (
	DASHBOARDINLINESCHEMA49_CUSTOM_BG    DashboardInlineSchema49 = "custom_bg"
	DASHBOARDINLINESCHEMA49_CUSTOM_TEXT  DashboardInlineSchema49 = "custom_text"
	DASHBOARDINLINESCHEMA49_CUSTOM_IMAGE DashboardInlineSchema49 = "custom_image"
)

var allowedDashboardInlineSchema49EnumValues = []DashboardInlineSchema49{
	DASHBOARDINLINESCHEMA49_CUSTOM_BG,
	DASHBOARDINLINESCHEMA49_CUSTOM_TEXT,
	DASHBOARDINLINESCHEMA49_CUSTOM_IMAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema49) GetAllowedValues() []DashboardInlineSchema49 {
	return allowedDashboardInlineSchema49EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema49) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema49(value)
	return nil
}

// NewDashboardInlineSchema49FromValue returns a pointer to a valid DashboardInlineSchema49
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema49FromValue(v string) (*DashboardInlineSchema49, error) {
	ev := DashboardInlineSchema49(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema49: valid values are %v", v, allowedDashboardInlineSchema49EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema49) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema49EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema49 value.
func (v DashboardInlineSchema49) Ptr() *DashboardInlineSchema49 {
	return &v
}

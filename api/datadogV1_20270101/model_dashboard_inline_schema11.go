// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema11 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema11 string

// List of DashboardInlineSchema11.
const (
	DASHBOARDINLINESCHEMA11_NONE DashboardInlineSchema11 = "none"
)

var allowedDashboardInlineSchema11EnumValues = []DashboardInlineSchema11{
	DASHBOARDINLINESCHEMA11_NONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema11) GetAllowedValues() []DashboardInlineSchema11 {
	return allowedDashboardInlineSchema11EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema11) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema11(value)
	return nil
}

// NewDashboardInlineSchema11FromValue returns a pointer to a valid DashboardInlineSchema11
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema11FromValue(v string) (*DashboardInlineSchema11, error) {
	ev := DashboardInlineSchema11(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema11: valid values are %v", v, allowedDashboardInlineSchema11EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema11) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema11EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema11 value.
func (v DashboardInlineSchema11) Ptr() *DashboardInlineSchema11 {
	return &v
}

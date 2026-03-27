// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardReflowStyle Reflow type for a **new dashboard layout** dashboard. Set this only when layout type is 'ordered'.
// If set to 'fixed', the dashboard expects all widgets to have a layout, and if it's set to 'auto',
// widgets should not have layouts.
type DashboardReflowStyle string

// List of DashboardReflowStyle.
const (
	DASHBOARDREFLOWTYPE_AUTO  DashboardReflowStyle = "auto"
	DASHBOARDREFLOWTYPE_FIXED DashboardReflowStyle = "fixed"
)

var allowedDashboardReflowStyleEnumValues = []DashboardReflowStyle{
	DASHBOARDREFLOWTYPE_AUTO,
	DASHBOARDREFLOWTYPE_FIXED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardReflowStyle) GetAllowedValues() []DashboardReflowStyle {
	return allowedDashboardReflowStyleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardReflowStyle) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardReflowStyle(value)
	return nil
}

// NewDashboardReflowStyleFromValue returns a pointer to a valid DashboardReflowStyle
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardReflowStyleFromValue(v string) (*DashboardReflowStyle, error) {
	ev := DashboardReflowStyle(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardReflowStyle: valid values are %v", v, allowedDashboardReflowStyleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardReflowStyle) IsValid() bool {
	for _, existing := range allowedDashboardReflowStyleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardReflowStyle value.
func (v DashboardReflowStyle) Ptr() *DashboardReflowStyle {
	return &v
}

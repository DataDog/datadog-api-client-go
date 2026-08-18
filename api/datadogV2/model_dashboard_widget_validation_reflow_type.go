// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardWidgetValidationReflowType Reflow behavior used for an ordered dashboard.
type DashboardWidgetValidationReflowType string

// List of DashboardWidgetValidationReflowType.
const (
	DASHBOARDWIDGETVALIDATIONREFLOWTYPE_AUTO  DashboardWidgetValidationReflowType = "auto"
	DASHBOARDWIDGETVALIDATIONREFLOWTYPE_FIXED DashboardWidgetValidationReflowType = "fixed"
)

var allowedDashboardWidgetValidationReflowTypeEnumValues = []DashboardWidgetValidationReflowType{
	DASHBOARDWIDGETVALIDATIONREFLOWTYPE_AUTO,
	DASHBOARDWIDGETVALIDATIONREFLOWTYPE_FIXED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardWidgetValidationReflowType) GetAllowedValues() []DashboardWidgetValidationReflowType {
	return allowedDashboardWidgetValidationReflowTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardWidgetValidationReflowType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardWidgetValidationReflowType(value)
	return nil
}

// NewDashboardWidgetValidationReflowTypeFromValue returns a pointer to a valid DashboardWidgetValidationReflowType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardWidgetValidationReflowTypeFromValue(v string) (*DashboardWidgetValidationReflowType, error) {
	ev := DashboardWidgetValidationReflowType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardWidgetValidationReflowType: valid values are %v", v, allowedDashboardWidgetValidationReflowTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardWidgetValidationReflowType) IsValid() bool {
	for _, existing := range allowedDashboardWidgetValidationReflowTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardWidgetValidationReflowType value.
func (v DashboardWidgetValidationReflowType) Ptr() *DashboardWidgetValidationReflowType {
	return &v
}

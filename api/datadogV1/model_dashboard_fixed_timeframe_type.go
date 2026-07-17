// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardFixedTimeframeType Type of fixed timeframe.
type DashboardFixedTimeframeType string

// List of DashboardFixedTimeframeType.
const (
	DASHBOARDFIXEDTIMEFRAMETYPE_FIXED DashboardFixedTimeframeType = "fixed"
)

var allowedDashboardFixedTimeframeTypeEnumValues = []DashboardFixedTimeframeType{
	DASHBOARDFIXEDTIMEFRAMETYPE_FIXED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardFixedTimeframeType) GetAllowedValues() []DashboardFixedTimeframeType {
	return allowedDashboardFixedTimeframeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardFixedTimeframeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardFixedTimeframeType(value)
	return nil
}

// NewDashboardFixedTimeframeTypeFromValue returns a pointer to a valid DashboardFixedTimeframeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardFixedTimeframeTypeFromValue(v string) (*DashboardFixedTimeframeType, error) {
	ev := DashboardFixedTimeframeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardFixedTimeframeType: valid values are %v", v, allowedDashboardFixedTimeframeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardFixedTimeframeType) IsValid() bool {
	for _, existing := range allowedDashboardFixedTimeframeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardFixedTimeframeType value.
func (v DashboardFixedTimeframeType) Ptr() *DashboardFixedTimeframeType {
	return &v
}

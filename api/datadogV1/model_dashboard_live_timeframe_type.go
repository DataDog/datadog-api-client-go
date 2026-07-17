// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardLiveTimeframeType Type of live timeframe.
type DashboardLiveTimeframeType string

// List of DashboardLiveTimeframeType.
const (
	DASHBOARDLIVETIMEFRAMETYPE_LIVE DashboardLiveTimeframeType = "live"
)

var allowedDashboardLiveTimeframeTypeEnumValues = []DashboardLiveTimeframeType{
	DASHBOARDLIVETIMEFRAMETYPE_LIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardLiveTimeframeType) GetAllowedValues() []DashboardLiveTimeframeType {
	return allowedDashboardLiveTimeframeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardLiveTimeframeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardLiveTimeframeType(value)
	return nil
}

// NewDashboardLiveTimeframeTypeFromValue returns a pointer to a valid DashboardLiveTimeframeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardLiveTimeframeTypeFromValue(v string) (*DashboardLiveTimeframeType, error) {
	ev := DashboardLiveTimeframeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardLiveTimeframeType: valid values are %v", v, allowedDashboardLiveTimeframeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardLiveTimeframeType) IsValid() bool {
	for _, existing := range allowedDashboardLiveTimeframeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardLiveTimeframeType value.
func (v DashboardLiveTimeframeType) Ptr() *DashboardLiveTimeframeType {
	return &v
}

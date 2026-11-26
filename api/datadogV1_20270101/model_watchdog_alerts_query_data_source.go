// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WatchdogAlertsQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type WatchdogAlertsQueryDataSource string

// List of WatchdogAlertsQueryDataSource.
const (
	WATCHDOGALERTSQUERYDATASOURCE_WATCHDOG_ALERTS WatchdogAlertsQueryDataSource = "watchdog_alerts"
)

var allowedWatchdogAlertsQueryDataSourceEnumValues = []WatchdogAlertsQueryDataSource{
	WATCHDOGALERTSQUERYDATASOURCE_WATCHDOG_ALERTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WatchdogAlertsQueryDataSource) GetAllowedValues() []WatchdogAlertsQueryDataSource {
	return allowedWatchdogAlertsQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WatchdogAlertsQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WatchdogAlertsQueryDataSource(value)
	return nil
}

// NewWatchdogAlertsQueryDataSourceFromValue returns a pointer to a valid WatchdogAlertsQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWatchdogAlertsQueryDataSourceFromValue(v string) (*WatchdogAlertsQueryDataSource, error) {
	ev := WatchdogAlertsQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WatchdogAlertsQueryDataSource: valid values are %v", v, allowedWatchdogAlertsQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WatchdogAlertsQueryDataSource) IsValid() bool {
	for _, existing := range allowedWatchdogAlertsQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WatchdogAlertsQueryDataSource value.
func (v WatchdogAlertsQueryDataSource) Ptr() *WatchdogAlertsQueryDataSource {
	return &v
}

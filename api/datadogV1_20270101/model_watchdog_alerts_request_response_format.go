// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WatchdogAlertsRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type WatchdogAlertsRequestResponseFormat string

// List of WatchdogAlertsRequestResponseFormat.
const (
	WATCHDOGALERTSREQUESTRESPONSEFORMAT_EVENT_LIST WatchdogAlertsRequestResponseFormat = "event_list"
)

var allowedWatchdogAlertsRequestResponseFormatEnumValues = []WatchdogAlertsRequestResponseFormat{
	WATCHDOGALERTSREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WatchdogAlertsRequestResponseFormat) GetAllowedValues() []WatchdogAlertsRequestResponseFormat {
	return allowedWatchdogAlertsRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WatchdogAlertsRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WatchdogAlertsRequestResponseFormat(value)
	return nil
}

// NewWatchdogAlertsRequestResponseFormatFromValue returns a pointer to a valid WatchdogAlertsRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWatchdogAlertsRequestResponseFormatFromValue(v string) (*WatchdogAlertsRequestResponseFormat, error) {
	ev := WatchdogAlertsRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WatchdogAlertsRequestResponseFormat: valid values are %v", v, allowedWatchdogAlertsRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WatchdogAlertsRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedWatchdogAlertsRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WatchdogAlertsRequestResponseFormat value.
func (v WatchdogAlertsRequestResponseFormat) Ptr() *WatchdogAlertsRequestResponseFormat {
	return &v
}

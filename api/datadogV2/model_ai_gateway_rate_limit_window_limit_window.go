// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRateLimitWindowLimitWindow The rolling window that a limit applies to. For per-minute limits, use the `*_per_minute` fields instead.
type AiGatewayRateLimitWindowLimitWindow string

// List of AiGatewayRateLimitWindowLimitWindow.
const (
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_FIVE_MINUTES      AiGatewayRateLimitWindowLimitWindow = "5m"
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_FIFTEEN_MINUTES   AiGatewayRateLimitWindowLimitWindow = "15m"
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_THIRTY_MINUTES    AiGatewayRateLimitWindowLimitWindow = "30m"
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_ONE_HOUR          AiGatewayRateLimitWindowLimitWindow = "1h"
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_THREE_HOURS       AiGatewayRateLimitWindowLimitWindow = "3h"
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_SIX_HOURS         AiGatewayRateLimitWindowLimitWindow = "6h"
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_TWELVE_HOURS      AiGatewayRateLimitWindowLimitWindow = "12h"
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_TWENTY_FOUR_HOURS AiGatewayRateLimitWindowLimitWindow = "24h"
)

var allowedAiGatewayRateLimitWindowLimitWindowEnumValues = []AiGatewayRateLimitWindowLimitWindow{
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_FIVE_MINUTES,
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_FIFTEEN_MINUTES,
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_THIRTY_MINUTES,
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_ONE_HOUR,
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_THREE_HOURS,
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_SIX_HOURS,
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_TWELVE_HOURS,
	AIGATEWAYRATELIMITWINDOWLIMITWINDOW_TWENTY_FOUR_HOURS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayRateLimitWindowLimitWindow) GetAllowedValues() []AiGatewayRateLimitWindowLimitWindow {
	return allowedAiGatewayRateLimitWindowLimitWindowEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayRateLimitWindowLimitWindow) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayRateLimitWindowLimitWindow(value)
	return nil
}

// NewAiGatewayRateLimitWindowLimitWindowFromValue returns a pointer to a valid AiGatewayRateLimitWindowLimitWindow
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayRateLimitWindowLimitWindowFromValue(v string) (*AiGatewayRateLimitWindowLimitWindow, error) {
	ev := AiGatewayRateLimitWindowLimitWindow(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayRateLimitWindowLimitWindow: valid values are %v", v, allowedAiGatewayRateLimitWindowLimitWindowEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayRateLimitWindowLimitWindow) IsValid() bool {
	for _, existing := range allowedAiGatewayRateLimitWindowLimitWindowEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayRateLimitWindowLimitWindow value.
func (v AiGatewayRateLimitWindowLimitWindow) Ptr() *AiGatewayRateLimitWindowLimitWindow {
	return &v
}

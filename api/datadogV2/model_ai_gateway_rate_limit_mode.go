// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRateLimitMode Action taken when request rate exceeds the limit.
type AiGatewayRateLimitMode string

// List of AiGatewayRateLimitMode.
const (
	AIGATEWAYRATELIMITMODE_ALERT   AiGatewayRateLimitMode = "ALERT"
	AIGATEWAYRATELIMITMODE_ENFORCE AiGatewayRateLimitMode = "ENFORCE"
)

var allowedAiGatewayRateLimitModeEnumValues = []AiGatewayRateLimitMode{
	AIGATEWAYRATELIMITMODE_ALERT,
	AIGATEWAYRATELIMITMODE_ENFORCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayRateLimitMode) GetAllowedValues() []AiGatewayRateLimitMode {
	return allowedAiGatewayRateLimitModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayRateLimitMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayRateLimitMode(value)
	return nil
}

// NewAiGatewayRateLimitModeFromValue returns a pointer to a valid AiGatewayRateLimitMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayRateLimitModeFromValue(v string) (*AiGatewayRateLimitMode, error) {
	ev := AiGatewayRateLimitMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayRateLimitMode: valid values are %v", v, allowedAiGatewayRateLimitModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayRateLimitMode) IsValid() bool {
	for _, existing := range allowedAiGatewayRateLimitModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayRateLimitMode value.
func (v AiGatewayRateLimitMode) Ptr() *AiGatewayRateLimitMode {
	return &v
}

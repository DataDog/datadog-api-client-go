// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayMaxReasoningLevel Maximum reasoning level permitted by a model group.
type AiGatewayMaxReasoningLevel string

// List of AiGatewayMaxReasoningLevel.
const (
	AIGATEWAYMAXREASONINGLEVEL_NONE    AiGatewayMaxReasoningLevel = "none"
	AIGATEWAYMAXREASONINGLEVEL_MINIMAL AiGatewayMaxReasoningLevel = "minimal"
	AIGATEWAYMAXREASONINGLEVEL_LOW     AiGatewayMaxReasoningLevel = "low"
	AIGATEWAYMAXREASONINGLEVEL_MEDIUM  AiGatewayMaxReasoningLevel = "medium"
	AIGATEWAYMAXREASONINGLEVEL_HIGH    AiGatewayMaxReasoningLevel = "high"
	AIGATEWAYMAXREASONINGLEVEL_XHIGH   AiGatewayMaxReasoningLevel = "xhigh"
	AIGATEWAYMAXREASONINGLEVEL_MAX     AiGatewayMaxReasoningLevel = "max"
)

var allowedAiGatewayMaxReasoningLevelEnumValues = []AiGatewayMaxReasoningLevel{
	AIGATEWAYMAXREASONINGLEVEL_NONE,
	AIGATEWAYMAXREASONINGLEVEL_MINIMAL,
	AIGATEWAYMAXREASONINGLEVEL_LOW,
	AIGATEWAYMAXREASONINGLEVEL_MEDIUM,
	AIGATEWAYMAXREASONINGLEVEL_HIGH,
	AIGATEWAYMAXREASONINGLEVEL_XHIGH,
	AIGATEWAYMAXREASONINGLEVEL_MAX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayMaxReasoningLevel) GetAllowedValues() []AiGatewayMaxReasoningLevel {
	return allowedAiGatewayMaxReasoningLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayMaxReasoningLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayMaxReasoningLevel(value)
	return nil
}

// NewAiGatewayMaxReasoningLevelFromValue returns a pointer to a valid AiGatewayMaxReasoningLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayMaxReasoningLevelFromValue(v string) (*AiGatewayMaxReasoningLevel, error) {
	ev := AiGatewayMaxReasoningLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayMaxReasoningLevel: valid values are %v", v, allowedAiGatewayMaxReasoningLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayMaxReasoningLevel) IsValid() bool {
	for _, existing := range allowedAiGatewayMaxReasoningLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayMaxReasoningLevel value.
func (v AiGatewayMaxReasoningLevel) Ptr() *AiGatewayMaxReasoningLevel {
	return &v
}

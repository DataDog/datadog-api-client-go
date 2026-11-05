// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingReasoningLevel Optional reasoning effort preference for this model destination. Must be supported by the selected model and provider. For provider "*", every eligible provider must support the level. An explicit value overrides the request's reasoning setting, subject to model-group policy. Omitted or null preserves request and provider behavior; none is an explicit level, not the default. Omitted or null clears an override on full replacement. Routing-rule references cannot set this field; each referenced model destination owns its override. If an override cannot be applied at runtime, the gateway keeps the destination and restores the caller's original reasoning settings, or preserves the provider default when omitted. It rechecks all applicable route and model-group policies against those settings before sending. Model-group max_reasoning_level enforces hard limits; the override is not an additional cap. Normal fallback applies if governance denies the resulting request or the destination cannot otherwise serve it. Metrics and span tags record the override outcome and incompatibility reason. When response headers have not already been committed, `X-Ai-Gateway-Reasoning-Override` reports applied or not_applied for the returned attempt, and `X-Ai-Gateway-Reasoning-Override-Reason` reports the reason when not applied.
type AiGatewayRoutingReasoningLevel string

// List of AiGatewayRoutingReasoningLevel.
const (
	AIGATEWAYROUTINGREASONINGLEVEL_NONE    AiGatewayRoutingReasoningLevel = "none"
	AIGATEWAYROUTINGREASONINGLEVEL_MINIMAL AiGatewayRoutingReasoningLevel = "minimal"
	AIGATEWAYROUTINGREASONINGLEVEL_LOW     AiGatewayRoutingReasoningLevel = "low"
	AIGATEWAYROUTINGREASONINGLEVEL_MEDIUM  AiGatewayRoutingReasoningLevel = "medium"
	AIGATEWAYROUTINGREASONINGLEVEL_HIGH    AiGatewayRoutingReasoningLevel = "high"
	AIGATEWAYROUTINGREASONINGLEVEL_XHIGH   AiGatewayRoutingReasoningLevel = "xhigh"
	AIGATEWAYROUTINGREASONINGLEVEL_MAX     AiGatewayRoutingReasoningLevel = "max"
)

var allowedAiGatewayRoutingReasoningLevelEnumValues = []AiGatewayRoutingReasoningLevel{
	AIGATEWAYROUTINGREASONINGLEVEL_NONE,
	AIGATEWAYROUTINGREASONINGLEVEL_MINIMAL,
	AIGATEWAYROUTINGREASONINGLEVEL_LOW,
	AIGATEWAYROUTINGREASONINGLEVEL_MEDIUM,
	AIGATEWAYROUTINGREASONINGLEVEL_HIGH,
	AIGATEWAYROUTINGREASONINGLEVEL_XHIGH,
	AIGATEWAYROUTINGREASONINGLEVEL_MAX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayRoutingReasoningLevel) GetAllowedValues() []AiGatewayRoutingReasoningLevel {
	return allowedAiGatewayRoutingReasoningLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayRoutingReasoningLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayRoutingReasoningLevel(value)
	return nil
}

// NewAiGatewayRoutingReasoningLevelFromValue returns a pointer to a valid AiGatewayRoutingReasoningLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayRoutingReasoningLevelFromValue(v string) (*AiGatewayRoutingReasoningLevel, error) {
	ev := AiGatewayRoutingReasoningLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayRoutingReasoningLevel: valid values are %v", v, allowedAiGatewayRoutingReasoningLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayRoutingReasoningLevel) IsValid() bool {
	for _, existing := range allowedAiGatewayRoutingReasoningLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayRoutingReasoningLevel value.
func (v AiGatewayRoutingReasoningLevel) Ptr() *AiGatewayRoutingReasoningLevel {
	return &v
}

// NullableAiGatewayRoutingReasoningLevel handles when a null is used for AiGatewayRoutingReasoningLevel.
type NullableAiGatewayRoutingReasoningLevel struct {
	value *AiGatewayRoutingReasoningLevel
	isSet bool
}

// Get returns the associated value.
func (v NullableAiGatewayRoutingReasoningLevel) Get() *AiGatewayRoutingReasoningLevel {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableAiGatewayRoutingReasoningLevel) Set(val *AiGatewayRoutingReasoningLevel) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableAiGatewayRoutingReasoningLevel) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableAiGatewayRoutingReasoningLevel) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAiGatewayRoutingReasoningLevel initializes the struct as if Set has been called.
func NewNullableAiGatewayRoutingReasoningLevel(val *AiGatewayRoutingReasoningLevel) *NullableAiGatewayRoutingReasoningLevel {
	return &NullableAiGatewayRoutingReasoningLevel{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableAiGatewayRoutingReasoningLevel) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableAiGatewayRoutingReasoningLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayNullableMaxReasoningLevel Maximum reasoning level permitted by a model group.
type AiGatewayNullableMaxReasoningLevel string

// List of AiGatewayNullableMaxReasoningLevel.
const (
	AIGATEWAYNULLABLEMAXREASONINGLEVEL_NONE    AiGatewayNullableMaxReasoningLevel = "none"
	AIGATEWAYNULLABLEMAXREASONINGLEVEL_MINIMAL AiGatewayNullableMaxReasoningLevel = "minimal"
	AIGATEWAYNULLABLEMAXREASONINGLEVEL_LOW     AiGatewayNullableMaxReasoningLevel = "low"
	AIGATEWAYNULLABLEMAXREASONINGLEVEL_MEDIUM  AiGatewayNullableMaxReasoningLevel = "medium"
	AIGATEWAYNULLABLEMAXREASONINGLEVEL_HIGH    AiGatewayNullableMaxReasoningLevel = "high"
	AIGATEWAYNULLABLEMAXREASONINGLEVEL_XHIGH   AiGatewayNullableMaxReasoningLevel = "xhigh"
	AIGATEWAYNULLABLEMAXREASONINGLEVEL_MAX     AiGatewayNullableMaxReasoningLevel = "max"
)

var allowedAiGatewayNullableMaxReasoningLevelEnumValues = []AiGatewayNullableMaxReasoningLevel{
	AIGATEWAYNULLABLEMAXREASONINGLEVEL_NONE,
	AIGATEWAYNULLABLEMAXREASONINGLEVEL_MINIMAL,
	AIGATEWAYNULLABLEMAXREASONINGLEVEL_LOW,
	AIGATEWAYNULLABLEMAXREASONINGLEVEL_MEDIUM,
	AIGATEWAYNULLABLEMAXREASONINGLEVEL_HIGH,
	AIGATEWAYNULLABLEMAXREASONINGLEVEL_XHIGH,
	AIGATEWAYNULLABLEMAXREASONINGLEVEL_MAX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayNullableMaxReasoningLevel) GetAllowedValues() []AiGatewayNullableMaxReasoningLevel {
	return allowedAiGatewayNullableMaxReasoningLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayNullableMaxReasoningLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayNullableMaxReasoningLevel(value)
	return nil
}

// NewAiGatewayNullableMaxReasoningLevelFromValue returns a pointer to a valid AiGatewayNullableMaxReasoningLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayNullableMaxReasoningLevelFromValue(v string) (*AiGatewayNullableMaxReasoningLevel, error) {
	ev := AiGatewayNullableMaxReasoningLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayNullableMaxReasoningLevel: valid values are %v", v, allowedAiGatewayNullableMaxReasoningLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayNullableMaxReasoningLevel) IsValid() bool {
	for _, existing := range allowedAiGatewayNullableMaxReasoningLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayNullableMaxReasoningLevel value.
func (v AiGatewayNullableMaxReasoningLevel) Ptr() *AiGatewayNullableMaxReasoningLevel {
	return &v
}

// NullableAiGatewayNullableMaxReasoningLevel handles when a null is used for AiGatewayNullableMaxReasoningLevel.
type NullableAiGatewayNullableMaxReasoningLevel struct {
	value *AiGatewayNullableMaxReasoningLevel
	isSet bool
}

// Get returns the associated value.
func (v NullableAiGatewayNullableMaxReasoningLevel) Get() *AiGatewayNullableMaxReasoningLevel {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableAiGatewayNullableMaxReasoningLevel) Set(val *AiGatewayNullableMaxReasoningLevel) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableAiGatewayNullableMaxReasoningLevel) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableAiGatewayNullableMaxReasoningLevel) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAiGatewayNullableMaxReasoningLevel initializes the struct as if Set has been called.
func NewNullableAiGatewayNullableMaxReasoningLevel(val *AiGatewayNullableMaxReasoningLevel) *NullableAiGatewayNullableMaxReasoningLevel {
	return &NullableAiGatewayNullableMaxReasoningLevel{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableAiGatewayNullableMaxReasoningLevel) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableAiGatewayNullableMaxReasoningLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}

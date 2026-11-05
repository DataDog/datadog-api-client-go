// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayNullableRoutingRuleTierOrdering How a tier orders its targets. Omitted or null defaults to configured, which preserves target order. The weighted_affinity ordering distributes first-choice assignments by relative target weights and accepts only model targets with an explicit provider, without repeated provider and model pairs in the same tier.
type AiGatewayNullableRoutingRuleTierOrdering string

// List of AiGatewayNullableRoutingRuleTierOrdering.
const (
	AIGATEWAYNULLABLEROUTINGRULETIERORDERING_CONFIGURED        AiGatewayNullableRoutingRuleTierOrdering = "configured"
	AIGATEWAYNULLABLEROUTINGRULETIERORDERING_WEIGHTED_AFFINITY AiGatewayNullableRoutingRuleTierOrdering = "weighted_affinity"
)

var allowedAiGatewayNullableRoutingRuleTierOrderingEnumValues = []AiGatewayNullableRoutingRuleTierOrdering{
	AIGATEWAYNULLABLEROUTINGRULETIERORDERING_CONFIGURED,
	AIGATEWAYNULLABLEROUTINGRULETIERORDERING_WEIGHTED_AFFINITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayNullableRoutingRuleTierOrdering) GetAllowedValues() []AiGatewayNullableRoutingRuleTierOrdering {
	return allowedAiGatewayNullableRoutingRuleTierOrderingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayNullableRoutingRuleTierOrdering) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayNullableRoutingRuleTierOrdering(value)
	return nil
}

// NewAiGatewayNullableRoutingRuleTierOrderingFromValue returns a pointer to a valid AiGatewayNullableRoutingRuleTierOrdering
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayNullableRoutingRuleTierOrderingFromValue(v string) (*AiGatewayNullableRoutingRuleTierOrdering, error) {
	ev := AiGatewayNullableRoutingRuleTierOrdering(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayNullableRoutingRuleTierOrdering: valid values are %v", v, allowedAiGatewayNullableRoutingRuleTierOrderingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayNullableRoutingRuleTierOrdering) IsValid() bool {
	for _, existing := range allowedAiGatewayNullableRoutingRuleTierOrderingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayNullableRoutingRuleTierOrdering value.
func (v AiGatewayNullableRoutingRuleTierOrdering) Ptr() *AiGatewayNullableRoutingRuleTierOrdering {
	return &v
}

// NullableAiGatewayNullableRoutingRuleTierOrdering handles when a null is used for AiGatewayNullableRoutingRuleTierOrdering.
type NullableAiGatewayNullableRoutingRuleTierOrdering struct {
	value *AiGatewayNullableRoutingRuleTierOrdering
	isSet bool
}

// Get returns the associated value.
func (v NullableAiGatewayNullableRoutingRuleTierOrdering) Get() *AiGatewayNullableRoutingRuleTierOrdering {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableAiGatewayNullableRoutingRuleTierOrdering) Set(val *AiGatewayNullableRoutingRuleTierOrdering) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableAiGatewayNullableRoutingRuleTierOrdering) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableAiGatewayNullableRoutingRuleTierOrdering) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAiGatewayNullableRoutingRuleTierOrdering initializes the struct as if Set has been called.
func NewNullableAiGatewayNullableRoutingRuleTierOrdering(val *AiGatewayNullableRoutingRuleTierOrdering) *NullableAiGatewayNullableRoutingRuleTierOrdering {
	return &NullableAiGatewayNullableRoutingRuleTierOrdering{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableAiGatewayNullableRoutingRuleTierOrdering) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableAiGatewayNullableRoutingRuleTierOrdering) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}

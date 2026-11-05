// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingRuleTierOrdering How a tier orders its targets.
type AiGatewayRoutingRuleTierOrdering string

// List of AiGatewayRoutingRuleTierOrdering.
const (
	AIGATEWAYROUTINGRULETIERORDERING_CONFIGURED        AiGatewayRoutingRuleTierOrdering = "configured"
	AIGATEWAYROUTINGRULETIERORDERING_WEIGHTED_AFFINITY AiGatewayRoutingRuleTierOrdering = "weighted_affinity"
)

var allowedAiGatewayRoutingRuleTierOrderingEnumValues = []AiGatewayRoutingRuleTierOrdering{
	AIGATEWAYROUTINGRULETIERORDERING_CONFIGURED,
	AIGATEWAYROUTINGRULETIERORDERING_WEIGHTED_AFFINITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayRoutingRuleTierOrdering) GetAllowedValues() []AiGatewayRoutingRuleTierOrdering {
	return allowedAiGatewayRoutingRuleTierOrderingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayRoutingRuleTierOrdering) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayRoutingRuleTierOrdering(value)
	return nil
}

// NewAiGatewayRoutingRuleTierOrderingFromValue returns a pointer to a valid AiGatewayRoutingRuleTierOrdering
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayRoutingRuleTierOrderingFromValue(v string) (*AiGatewayRoutingRuleTierOrdering, error) {
	ev := AiGatewayRoutingRuleTierOrdering(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayRoutingRuleTierOrdering: valid values are %v", v, allowedAiGatewayRoutingRuleTierOrderingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayRoutingRuleTierOrdering) IsValid() bool {
	for _, existing := range allowedAiGatewayRoutingRuleTierOrderingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayRoutingRuleTierOrdering value.
func (v AiGatewayRoutingRuleTierOrdering) Ptr() *AiGatewayRoutingRuleTierOrdering {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingRuleTargetType Type of routing target.
type AiGatewayRoutingRuleTargetType string

// List of AiGatewayRoutingRuleTargetType.
const (
	AIGATEWAYROUTINGRULETARGETTYPE_MODEL        AiGatewayRoutingRuleTargetType = "model"
	AIGATEWAYROUTINGRULETARGETTYPE_ROUTING_RULE AiGatewayRoutingRuleTargetType = "routing-rule"
)

var allowedAiGatewayRoutingRuleTargetTypeEnumValues = []AiGatewayRoutingRuleTargetType{
	AIGATEWAYROUTINGRULETARGETTYPE_MODEL,
	AIGATEWAYROUTINGRULETARGETTYPE_ROUTING_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayRoutingRuleTargetType) GetAllowedValues() []AiGatewayRoutingRuleTargetType {
	return allowedAiGatewayRoutingRuleTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayRoutingRuleTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayRoutingRuleTargetType(value)
	return nil
}

// NewAiGatewayRoutingRuleTargetTypeFromValue returns a pointer to a valid AiGatewayRoutingRuleTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayRoutingRuleTargetTypeFromValue(v string) (*AiGatewayRoutingRuleTargetType, error) {
	ev := AiGatewayRoutingRuleTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayRoutingRuleTargetType: valid values are %v", v, allowedAiGatewayRoutingRuleTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayRoutingRuleTargetType) IsValid() bool {
	for _, existing := range allowedAiGatewayRoutingRuleTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayRoutingRuleTargetType value.
func (v AiGatewayRoutingRuleTargetType) Ptr() *AiGatewayRoutingRuleTargetType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingRuleType AI Gateway routing rule resource type.
type AiGatewayRoutingRuleType string

// List of AiGatewayRoutingRuleType.
const (
	AIGATEWAYROUTINGRULETYPE_ROUTING_RULES AiGatewayRoutingRuleType = "routing-rules"
)

var allowedAiGatewayRoutingRuleTypeEnumValues = []AiGatewayRoutingRuleType{
	AIGATEWAYROUTINGRULETYPE_ROUTING_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayRoutingRuleType) GetAllowedValues() []AiGatewayRoutingRuleType {
	return allowedAiGatewayRoutingRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayRoutingRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayRoutingRuleType(value)
	return nil
}

// NewAiGatewayRoutingRuleTypeFromValue returns a pointer to a valid AiGatewayRoutingRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayRoutingRuleTypeFromValue(v string) (*AiGatewayRoutingRuleType, error) {
	ev := AiGatewayRoutingRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayRoutingRuleType: valid values are %v", v, allowedAiGatewayRoutingRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayRoutingRuleType) IsValid() bool {
	for _, existing := range allowedAiGatewayRoutingRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayRoutingRuleType value.
func (v AiGatewayRoutingRuleType) Ptr() *AiGatewayRoutingRuleType {
	return &v
}

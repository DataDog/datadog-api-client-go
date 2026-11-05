// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelReasoningControlType Reasoning-control mechanism.
type AiGatewayModelReasoningControlType string

// List of AiGatewayModelReasoningControlType.
const (
	AIGATEWAYMODELREASONINGCONTROLTYPE_TOGGLE        AiGatewayModelReasoningControlType = "toggle"
	AIGATEWAYMODELREASONINGCONTROLTYPE_EFFORT        AiGatewayModelReasoningControlType = "effort"
	AIGATEWAYMODELREASONINGCONTROLTYPE_BUDGET_TOKENS AiGatewayModelReasoningControlType = "budget_tokens"
)

var allowedAiGatewayModelReasoningControlTypeEnumValues = []AiGatewayModelReasoningControlType{
	AIGATEWAYMODELREASONINGCONTROLTYPE_TOGGLE,
	AIGATEWAYMODELREASONINGCONTROLTYPE_EFFORT,
	AIGATEWAYMODELREASONINGCONTROLTYPE_BUDGET_TOKENS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayModelReasoningControlType) GetAllowedValues() []AiGatewayModelReasoningControlType {
	return allowedAiGatewayModelReasoningControlTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayModelReasoningControlType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayModelReasoningControlType(value)
	return nil
}

// NewAiGatewayModelReasoningControlTypeFromValue returns a pointer to a valid AiGatewayModelReasoningControlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayModelReasoningControlTypeFromValue(v string) (*AiGatewayModelReasoningControlType, error) {
	ev := AiGatewayModelReasoningControlType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayModelReasoningControlType: valid values are %v", v, allowedAiGatewayModelReasoningControlTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayModelReasoningControlType) IsValid() bool {
	for _, existing := range allowedAiGatewayModelReasoningControlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayModelReasoningControlType value.
func (v AiGatewayModelReasoningControlType) Ptr() *AiGatewayModelReasoningControlType {
	return &v
}

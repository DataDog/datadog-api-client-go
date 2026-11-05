// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCustomProviderSupportedOperation Operation supported by a custom provider.
type AiGatewayCustomProviderSupportedOperation string

// List of AiGatewayCustomProviderSupportedOperation.
const (
	AIGATEWAYCUSTOMPROVIDERSUPPORTEDOPERATION_OPENAI_CHAT_COMPLETIONS AiGatewayCustomProviderSupportedOperation = "OPENAI_CHAT_COMPLETIONS"
	AIGATEWAYCUSTOMPROVIDERSUPPORTEDOPERATION_OPENAI_RESPONSES        AiGatewayCustomProviderSupportedOperation = "OPENAI_RESPONSES"
	AIGATEWAYCUSTOMPROVIDERSUPPORTEDOPERATION_ANTHROPIC_MESSAGES      AiGatewayCustomProviderSupportedOperation = "ANTHROPIC_MESSAGES"
)

var allowedAiGatewayCustomProviderSupportedOperationEnumValues = []AiGatewayCustomProviderSupportedOperation{
	AIGATEWAYCUSTOMPROVIDERSUPPORTEDOPERATION_OPENAI_CHAT_COMPLETIONS,
	AIGATEWAYCUSTOMPROVIDERSUPPORTEDOPERATION_OPENAI_RESPONSES,
	AIGATEWAYCUSTOMPROVIDERSUPPORTEDOPERATION_ANTHROPIC_MESSAGES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayCustomProviderSupportedOperation) GetAllowedValues() []AiGatewayCustomProviderSupportedOperation {
	return allowedAiGatewayCustomProviderSupportedOperationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayCustomProviderSupportedOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayCustomProviderSupportedOperation(value)
	return nil
}

// NewAiGatewayCustomProviderSupportedOperationFromValue returns a pointer to a valid AiGatewayCustomProviderSupportedOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayCustomProviderSupportedOperationFromValue(v string) (*AiGatewayCustomProviderSupportedOperation, error) {
	ev := AiGatewayCustomProviderSupportedOperation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayCustomProviderSupportedOperation: valid values are %v", v, allowedAiGatewayCustomProviderSupportedOperationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayCustomProviderSupportedOperation) IsValid() bool {
	for _, existing := range allowedAiGatewayCustomProviderSupportedOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayCustomProviderSupportedOperation value.
func (v AiGatewayCustomProviderSupportedOperation) Ptr() *AiGatewayCustomProviderSupportedOperation {
	return &v
}

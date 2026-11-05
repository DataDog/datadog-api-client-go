// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewaySupportedOperation Operation supported by an AI Gateway model.
type AiGatewaySupportedOperation string

// List of AiGatewaySupportedOperation.
const (
	AIGATEWAYSUPPORTEDOPERATION_ANTHROPIC_MESSAGES      AiGatewaySupportedOperation = "ANTHROPIC_MESSAGES"
	AIGATEWAYSUPPORTEDOPERATION_ANTHROPIC_COUNT_TOKENS  AiGatewaySupportedOperation = "ANTHROPIC_COUNT_TOKENS"
	AIGATEWAYSUPPORTEDOPERATION_OPENAI_CHAT_COMPLETIONS AiGatewaySupportedOperation = "OPENAI_CHAT_COMPLETIONS"
	AIGATEWAYSUPPORTEDOPERATION_OPENAI_RESPONSES        AiGatewaySupportedOperation = "OPENAI_RESPONSES"
	AIGATEWAYSUPPORTEDOPERATION_OPENAI_WEB_SEARCH       AiGatewaySupportedOperation = "OPENAI_WEB_SEARCH"
	AIGATEWAYSUPPORTEDOPERATION_OPENAI_EMBEDDINGS       AiGatewaySupportedOperation = "OPENAI_EMBEDDINGS"
	AIGATEWAYSUPPORTEDOPERATION_GEMINI_GENERATE_CONTENT AiGatewaySupportedOperation = "GEMINI_GENERATE_CONTENT"
	AIGATEWAYSUPPORTEDOPERATION_GEMINI_RAW_PREDICTIONS  AiGatewaySupportedOperation = "GEMINI_RAW_PREDICTIONS"
)

var allowedAiGatewaySupportedOperationEnumValues = []AiGatewaySupportedOperation{
	AIGATEWAYSUPPORTEDOPERATION_ANTHROPIC_MESSAGES,
	AIGATEWAYSUPPORTEDOPERATION_ANTHROPIC_COUNT_TOKENS,
	AIGATEWAYSUPPORTEDOPERATION_OPENAI_CHAT_COMPLETIONS,
	AIGATEWAYSUPPORTEDOPERATION_OPENAI_RESPONSES,
	AIGATEWAYSUPPORTEDOPERATION_OPENAI_WEB_SEARCH,
	AIGATEWAYSUPPORTEDOPERATION_OPENAI_EMBEDDINGS,
	AIGATEWAYSUPPORTEDOPERATION_GEMINI_GENERATE_CONTENT,
	AIGATEWAYSUPPORTEDOPERATION_GEMINI_RAW_PREDICTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewaySupportedOperation) GetAllowedValues() []AiGatewaySupportedOperation {
	return allowedAiGatewaySupportedOperationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewaySupportedOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewaySupportedOperation(value)
	return nil
}

// NewAiGatewaySupportedOperationFromValue returns a pointer to a valid AiGatewaySupportedOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewaySupportedOperationFromValue(v string) (*AiGatewaySupportedOperation, error) {
	ev := AiGatewaySupportedOperation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewaySupportedOperation: valid values are %v", v, allowedAiGatewaySupportedOperationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewaySupportedOperation) IsValid() bool {
	for _, existing := range allowedAiGatewaySupportedOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewaySupportedOperation value.
func (v AiGatewaySupportedOperation) Ptr() *AiGatewaySupportedOperation {
	return &v
}

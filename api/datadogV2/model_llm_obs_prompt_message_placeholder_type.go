// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptMessagePlaceholderType Type of chat-template item.
type LLMObsPromptMessagePlaceholderType string

// List of LLMObsPromptMessagePlaceholderType.
const (
	LLMOBSPROMPTMESSAGEPLACEHOLDERTYPE_PLACEHOLDER LLMObsPromptMessagePlaceholderType = "placeholder"
)

var allowedLLMObsPromptMessagePlaceholderTypeEnumValues = []LLMObsPromptMessagePlaceholderType{
	LLMOBSPROMPTMESSAGEPLACEHOLDERTYPE_PLACEHOLDER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsPromptMessagePlaceholderType) GetAllowedValues() []LLMObsPromptMessagePlaceholderType {
	return allowedLLMObsPromptMessagePlaceholderTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsPromptMessagePlaceholderType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsPromptMessagePlaceholderType(value)
	return nil
}

// NewLLMObsPromptMessagePlaceholderTypeFromValue returns a pointer to a valid LLMObsPromptMessagePlaceholderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsPromptMessagePlaceholderTypeFromValue(v string) (*LLMObsPromptMessagePlaceholderType, error) {
	ev := LLMObsPromptMessagePlaceholderType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsPromptMessagePlaceholderType: valid values are %v", v, allowedLLMObsPromptMessagePlaceholderTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsPromptMessagePlaceholderType) IsValid() bool {
	for _, existing := range allowedLLMObsPromptMessagePlaceholderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsPromptMessagePlaceholderType value.
func (v LLMObsPromptMessagePlaceholderType) Ptr() *LLMObsPromptMessagePlaceholderType {
	return &v
}

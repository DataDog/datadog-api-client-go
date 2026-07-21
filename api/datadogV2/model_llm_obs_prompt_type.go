// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptType Resource type of an LLM Observability prompt.
type LLMObsPromptType string

// List of LLMObsPromptType.
const (
	LLMOBSPROMPTTYPE_PROMPT_TEMPLATES LLMObsPromptType = "prompt-templates"
)

var allowedLLMObsPromptTypeEnumValues = []LLMObsPromptType{
	LLMOBSPROMPTTYPE_PROMPT_TEMPLATES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsPromptType) GetAllowedValues() []LLMObsPromptType {
	return allowedLLMObsPromptTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsPromptType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsPromptType(value)
	return nil
}

// NewLLMObsPromptTypeFromValue returns a pointer to a valid LLMObsPromptType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsPromptTypeFromValue(v string) (*LLMObsPromptType, error) {
	ev := LLMObsPromptType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsPromptType: valid values are %v", v, allowedLLMObsPromptTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsPromptType) IsValid() bool {
	for _, existing := range allowedLLMObsPromptTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsPromptType value.
func (v LLMObsPromptType) Ptr() *LLMObsPromptType {
	return &v
}

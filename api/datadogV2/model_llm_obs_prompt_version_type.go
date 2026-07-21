// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptVersionType Resource type of an LLM Observability prompt version.
type LLMObsPromptVersionType string

// List of LLMObsPromptVersionType.
const (
	LLMOBSPROMPTVERSIONTYPE_PROMPT_TEMPLATE_VERSIONS LLMObsPromptVersionType = "prompt-template-versions"
)

var allowedLLMObsPromptVersionTypeEnumValues = []LLMObsPromptVersionType{
	LLMOBSPROMPTVERSIONTYPE_PROMPT_TEMPLATE_VERSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsPromptVersionType) GetAllowedValues() []LLMObsPromptVersionType {
	return allowedLLMObsPromptVersionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsPromptVersionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsPromptVersionType(value)
	return nil
}

// NewLLMObsPromptVersionTypeFromValue returns a pointer to a valid LLMObsPromptVersionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsPromptVersionTypeFromValue(v string) (*LLMObsPromptVersionType, error) {
	ev := LLMObsPromptVersionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsPromptVersionType: valid values are %v", v, allowedLLMObsPromptVersionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsPromptVersionType) IsValid() bool {
	for _, existing := range allowedLLMObsPromptVersionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsPromptVersionType value.
func (v LLMObsPromptVersionType) Ptr() *LLMObsPromptVersionType {
	return &v
}

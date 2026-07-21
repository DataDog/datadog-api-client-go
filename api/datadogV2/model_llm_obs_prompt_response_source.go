// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptResponseSource Whether the prompt was created from the registry or discovered from observed LLM calls.
type LLMObsPromptResponseSource string

// List of LLMObsPromptResponseSource.
const (
	LLMOBSPROMPTRESPONSESOURCE_REGISTRY LLMObsPromptResponseSource = "registry"
	LLMOBSPROMPTRESPONSESOURCE_CODE     LLMObsPromptResponseSource = "code"
)

var allowedLLMObsPromptResponseSourceEnumValues = []LLMObsPromptResponseSource{
	LLMOBSPROMPTRESPONSESOURCE_REGISTRY,
	LLMOBSPROMPTRESPONSESOURCE_CODE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsPromptResponseSource) GetAllowedValues() []LLMObsPromptResponseSource {
	return allowedLLMObsPromptResponseSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsPromptResponseSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsPromptResponseSource(value)
	return nil
}

// NewLLMObsPromptResponseSourceFromValue returns a pointer to a valid LLMObsPromptResponseSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsPromptResponseSourceFromValue(v string) (*LLMObsPromptResponseSource, error) {
	ev := LLMObsPromptResponseSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsPromptResponseSource: valid values are %v", v, allowedLLMObsPromptResponseSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsPromptResponseSource) IsValid() bool {
	for _, existing := range allowedLLMObsPromptResponseSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsPromptResponseSource value.
func (v LLMObsPromptResponseSource) Ptr() *LLMObsPromptResponseSource {
	return &v
}

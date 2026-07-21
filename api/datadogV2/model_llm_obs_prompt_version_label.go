// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptVersionLabel A label attached to an LLM Observability prompt version.
type LLMObsPromptVersionLabel string

// List of LLMObsPromptVersionLabel.
const (
	LLMOBSPROMPTVERSIONLABEL_PRODUCTION  LLMObsPromptVersionLabel = "production"
	LLMOBSPROMPTVERSIONLABEL_DEVELOPMENT LLMObsPromptVersionLabel = "development"
)

var allowedLLMObsPromptVersionLabelEnumValues = []LLMObsPromptVersionLabel{
	LLMOBSPROMPTVERSIONLABEL_PRODUCTION,
	LLMOBSPROMPTVERSIONLABEL_DEVELOPMENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsPromptVersionLabel) GetAllowedValues() []LLMObsPromptVersionLabel {
	return allowedLLMObsPromptVersionLabelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsPromptVersionLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsPromptVersionLabel(value)
	return nil
}

// NewLLMObsPromptVersionLabelFromValue returns a pointer to a valid LLMObsPromptVersionLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsPromptVersionLabelFromValue(v string) (*LLMObsPromptVersionLabel, error) {
	ev := LLMObsPromptVersionLabel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsPromptVersionLabel: valid values are %v", v, allowedLLMObsPromptVersionLabelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsPromptVersionLabel) IsValid() bool {
	for _, existing := range allowedLLMObsPromptVersionLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsPromptVersionLabel value.
func (v LLMObsPromptVersionLabel) Ptr() *LLMObsPromptVersionLabel {
	return &v
}

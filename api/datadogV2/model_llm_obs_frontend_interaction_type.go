// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsFrontendInteractionType Type discriminator for a `frontend` interaction.
type LLMObsFrontendInteractionType string

// List of LLMObsFrontendInteractionType.
const (
	LLMOBSFRONTENDINTERACTIONTYPE_FRONTEND LLMObsFrontendInteractionType = "frontend"
)

var allowedLLMObsFrontendInteractionTypeEnumValues = []LLMObsFrontendInteractionType{
	LLMOBSFRONTENDINTERACTIONTYPE_FRONTEND,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsFrontendInteractionType) GetAllowedValues() []LLMObsFrontendInteractionType {
	return allowedLLMObsFrontendInteractionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsFrontendInteractionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsFrontendInteractionType(value)
	return nil
}

// NewLLMObsFrontendInteractionTypeFromValue returns a pointer to a valid LLMObsFrontendInteractionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsFrontendInteractionTypeFromValue(v string) (*LLMObsFrontendInteractionType, error) {
	ev := LLMObsFrontendInteractionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsFrontendInteractionType: valid values are %v", v, allowedLLMObsFrontendInteractionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsFrontendInteractionType) IsValid() bool {
	for _, existing := range allowedLLMObsFrontendInteractionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsFrontendInteractionType value.
func (v LLMObsFrontendInteractionType) Ptr() *LLMObsFrontendInteractionType {
	return &v
}

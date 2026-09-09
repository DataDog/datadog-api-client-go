// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsFrontendContentBlockType Type discriminator for a `frontend` display block.
type LLMObsFrontendContentBlockType string

// List of LLMObsFrontendContentBlockType.
const (
	LLMOBSFRONTENDCONTENTBLOCKTYPE_FRONTEND LLMObsFrontendContentBlockType = "frontend"
)

var allowedLLMObsFrontendContentBlockTypeEnumValues = []LLMObsFrontendContentBlockType{
	LLMOBSFRONTENDCONTENTBLOCKTYPE_FRONTEND,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsFrontendContentBlockType) GetAllowedValues() []LLMObsFrontendContentBlockType {
	return allowedLLMObsFrontendContentBlockTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsFrontendContentBlockType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsFrontendContentBlockType(value)
	return nil
}

// NewLLMObsFrontendContentBlockTypeFromValue returns a pointer to a valid LLMObsFrontendContentBlockType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsFrontendContentBlockTypeFromValue(v string) (*LLMObsFrontendContentBlockType, error) {
	ev := LLMObsFrontendContentBlockType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsFrontendContentBlockType: valid values are %v", v, allowedLLMObsFrontendContentBlockTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsFrontendContentBlockType) IsValid() bool {
	for _, existing := range allowedLLMObsFrontendContentBlockTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsFrontendContentBlockType value.
func (v LLMObsFrontendContentBlockType) Ptr() *LLMObsFrontendContentBlockType {
	return &v
}

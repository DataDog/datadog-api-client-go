// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsLegacyContentBlockType Type discriminator for an existing non-frontend display block.
type LLMObsLegacyContentBlockType string

// List of LLMObsLegacyContentBlockType.
const (
	LLMOBSLEGACYCONTENTBLOCKTYPE_MARKDOWN     LLMObsLegacyContentBlockType = "markdown"
	LLMOBSLEGACYCONTENTBLOCKTYPE_HEADER       LLMObsLegacyContentBlockType = "header"
	LLMOBSLEGACYCONTENTBLOCKTYPE_TEXT         LLMObsLegacyContentBlockType = "text"
	LLMOBSLEGACYCONTENTBLOCKTYPE_JSON         LLMObsLegacyContentBlockType = "json"
	LLMOBSLEGACYCONTENTBLOCKTYPE_IMAGE        LLMObsLegacyContentBlockType = "image"
	LLMOBSLEGACYCONTENTBLOCKTYPE_WIDGET       LLMObsLegacyContentBlockType = "widget"
	LLMOBSLEGACYCONTENTBLOCKTYPE_LLMOBS_TRACE LLMObsLegacyContentBlockType = "llmobs_trace"
)

var allowedLLMObsLegacyContentBlockTypeEnumValues = []LLMObsLegacyContentBlockType{
	LLMOBSLEGACYCONTENTBLOCKTYPE_MARKDOWN,
	LLMOBSLEGACYCONTENTBLOCKTYPE_HEADER,
	LLMOBSLEGACYCONTENTBLOCKTYPE_TEXT,
	LLMOBSLEGACYCONTENTBLOCKTYPE_JSON,
	LLMOBSLEGACYCONTENTBLOCKTYPE_IMAGE,
	LLMOBSLEGACYCONTENTBLOCKTYPE_WIDGET,
	LLMOBSLEGACYCONTENTBLOCKTYPE_LLMOBS_TRACE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsLegacyContentBlockType) GetAllowedValues() []LLMObsLegacyContentBlockType {
	return allowedLLMObsLegacyContentBlockTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsLegacyContentBlockType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsLegacyContentBlockType(value)
	return nil
}

// NewLLMObsLegacyContentBlockTypeFromValue returns a pointer to a valid LLMObsLegacyContentBlockType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsLegacyContentBlockTypeFromValue(v string) (*LLMObsLegacyContentBlockType, error) {
	ev := LLMObsLegacyContentBlockType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsLegacyContentBlockType: valid values are %v", v, allowedLLMObsLegacyContentBlockTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsLegacyContentBlockType) IsValid() bool {
	for _, existing := range allowedLLMObsLegacyContentBlockTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsLegacyContentBlockType value.
func (v LLMObsLegacyContentBlockType) Ptr() *LLMObsLegacyContentBlockType {
	return &v
}

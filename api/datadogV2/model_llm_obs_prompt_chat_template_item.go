// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptChatTemplateItem - A chat message or a named message placeholder in a prompt template.
type LLMObsPromptChatTemplateItem struct {
	LLMObsPromptChatMessage        *LLMObsPromptChatMessage
	LLMObsPromptMessagePlaceholder *LLMObsPromptMessagePlaceholder

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LLMObsPromptChatMessageAsLLMObsPromptChatTemplateItem is a convenience function that returns LLMObsPromptChatMessage wrapped in LLMObsPromptChatTemplateItem.
func LLMObsPromptChatMessageAsLLMObsPromptChatTemplateItem(v *LLMObsPromptChatMessage) LLMObsPromptChatTemplateItem {
	return LLMObsPromptChatTemplateItem{LLMObsPromptChatMessage: v}
}

// LLMObsPromptMessagePlaceholderAsLLMObsPromptChatTemplateItem is a convenience function that returns LLMObsPromptMessagePlaceholder wrapped in LLMObsPromptChatTemplateItem.
func LLMObsPromptMessagePlaceholderAsLLMObsPromptChatTemplateItem(v *LLMObsPromptMessagePlaceholder) LLMObsPromptChatTemplateItem {
	return LLMObsPromptChatTemplateItem{LLMObsPromptMessagePlaceholder: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LLMObsPromptChatTemplateItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LLMObsPromptChatMessage
	err = datadog.Unmarshal(data, &obj.LLMObsPromptChatMessage)
	if err == nil {
		if obj.LLMObsPromptChatMessage != nil && obj.LLMObsPromptChatMessage.UnparsedObject == nil {
			jsonLLMObsPromptChatMessage, _ := datadog.Marshal(obj.LLMObsPromptChatMessage)
			if string(jsonLLMObsPromptChatMessage) == "{}" { // empty struct
				obj.LLMObsPromptChatMessage = nil
			} else {
				match++
			}
		} else {
			obj.LLMObsPromptChatMessage = nil
		}
	} else {
		obj.LLMObsPromptChatMessage = nil
	}

	// try to unmarshal data into LLMObsPromptMessagePlaceholder
	err = datadog.Unmarshal(data, &obj.LLMObsPromptMessagePlaceholder)
	if err == nil {
		if obj.LLMObsPromptMessagePlaceholder != nil && obj.LLMObsPromptMessagePlaceholder.UnparsedObject == nil {
			jsonLLMObsPromptMessagePlaceholder, _ := datadog.Marshal(obj.LLMObsPromptMessagePlaceholder)
			if string(jsonLLMObsPromptMessagePlaceholder) == "{}" { // empty struct
				obj.LLMObsPromptMessagePlaceholder = nil
			} else {
				match++
			}
		} else {
			obj.LLMObsPromptMessagePlaceholder = nil
		}
	} else {
		obj.LLMObsPromptMessagePlaceholder = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LLMObsPromptChatMessage = nil
		obj.LLMObsPromptMessagePlaceholder = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LLMObsPromptChatTemplateItem) MarshalJSON() ([]byte, error) {
	if obj.LLMObsPromptChatMessage != nil {
		return datadog.Marshal(&obj.LLMObsPromptChatMessage)
	}

	if obj.LLMObsPromptMessagePlaceholder != nil {
		return datadog.Marshal(&obj.LLMObsPromptMessagePlaceholder)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LLMObsPromptChatTemplateItem) GetActualInstance() interface{} {
	if obj.LLMObsPromptChatMessage != nil {
		return obj.LLMObsPromptChatMessage
	}

	if obj.LLMObsPromptMessagePlaceholder != nil {
		return obj.LLMObsPromptMessagePlaceholder
	}

	// all schemas are nil
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptAuthoringItem - A chat message or an explicitly versioned prompt include.
type LLMObsPromptAuthoringItem struct {
	LLMObsPromptChatMessage *LLMObsPromptChatMessage
	LLMObsPromptIncludeItem *LLMObsPromptIncludeItem

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LLMObsPromptChatMessageAsLLMObsPromptAuthoringItem is a convenience function that returns LLMObsPromptChatMessage wrapped in LLMObsPromptAuthoringItem.
func LLMObsPromptChatMessageAsLLMObsPromptAuthoringItem(v *LLMObsPromptChatMessage) LLMObsPromptAuthoringItem {
	return LLMObsPromptAuthoringItem{LLMObsPromptChatMessage: v}
}

// LLMObsPromptIncludeItemAsLLMObsPromptAuthoringItem is a convenience function that returns LLMObsPromptIncludeItem wrapped in LLMObsPromptAuthoringItem.
func LLMObsPromptIncludeItemAsLLMObsPromptAuthoringItem(v *LLMObsPromptIncludeItem) LLMObsPromptAuthoringItem {
	return LLMObsPromptAuthoringItem{LLMObsPromptIncludeItem: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LLMObsPromptAuthoringItem) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into LLMObsPromptIncludeItem
	err = datadog.Unmarshal(data, &obj.LLMObsPromptIncludeItem)
	if err == nil {
		if obj.LLMObsPromptIncludeItem != nil && obj.LLMObsPromptIncludeItem.UnparsedObject == nil {
			jsonLLMObsPromptIncludeItem, _ := datadog.Marshal(obj.LLMObsPromptIncludeItem)
			if string(jsonLLMObsPromptIncludeItem) == "{}" { // empty struct
				obj.LLMObsPromptIncludeItem = nil
			} else {
				match++
			}
		} else {
			obj.LLMObsPromptIncludeItem = nil
		}
	} else {
		obj.LLMObsPromptIncludeItem = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LLMObsPromptChatMessage = nil
		obj.LLMObsPromptIncludeItem = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LLMObsPromptAuthoringItem) MarshalJSON() ([]byte, error) {
	if obj.LLMObsPromptChatMessage != nil {
		return datadog.Marshal(&obj.LLMObsPromptChatMessage)
	}

	if obj.LLMObsPromptIncludeItem != nil {
		return datadog.Marshal(&obj.LLMObsPromptIncludeItem)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LLMObsPromptAuthoringItem) GetActualInstance() interface{} {
	if obj.LLMObsPromptChatMessage != nil {
		return obj.LLMObsPromptChatMessage
	}

	if obj.LLMObsPromptIncludeItem != nil {
		return obj.LLMObsPromptIncludeItem
	}

	// all schemas are nil
	return nil
}

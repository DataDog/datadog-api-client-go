// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardMessageContent - The content of a message, either a plain string or a list of OpenAI-style content parts.
type AIGuardMessageContent struct {
	AIGuardMessageContentText  *string
	AIGuardMessageContentParts *[]AIGuardContentPart

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AIGuardMessageContentTextAsAIGuardMessageContent is a convenience function that returns string wrapped in AIGuardMessageContent.
func AIGuardMessageContentTextAsAIGuardMessageContent(v *string) AIGuardMessageContent {
	return AIGuardMessageContent{AIGuardMessageContentText: v}
}

// AIGuardMessageContentPartsAsAIGuardMessageContent is a convenience function that returns []AIGuardContentPart wrapped in AIGuardMessageContent.
func AIGuardMessageContentPartsAsAIGuardMessageContent(v *[]AIGuardContentPart) AIGuardMessageContent {
	return AIGuardMessageContent{AIGuardMessageContentParts: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AIGuardMessageContent) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AIGuardMessageContentText
	err = datadog.Unmarshal(data, &obj.AIGuardMessageContentText)
	if err == nil {
		if obj.AIGuardMessageContentText != nil {
			jsonAIGuardMessageContentText, _ := datadog.Marshal(obj.AIGuardMessageContentText)
			if string(jsonAIGuardMessageContentText) == "{}" { // empty struct
				obj.AIGuardMessageContentText = nil
			} else {
				match++
			}
		} else {
			obj.AIGuardMessageContentText = nil
		}
	} else {
		obj.AIGuardMessageContentText = nil
	}

	// try to unmarshal data into AIGuardMessageContentParts
	err = datadog.Unmarshal(data, &obj.AIGuardMessageContentParts)
	if err == nil {
		if obj.AIGuardMessageContentParts != nil {
			jsonAIGuardMessageContentParts, _ := datadog.Marshal(obj.AIGuardMessageContentParts)
			if string(jsonAIGuardMessageContentParts) == "{}" && string(data) != "{}" { // empty struct
				obj.AIGuardMessageContentParts = nil
			} else {
				match++
			}
		} else {
			obj.AIGuardMessageContentParts = nil
		}
	} else {
		obj.AIGuardMessageContentParts = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AIGuardMessageContentText = nil
		obj.AIGuardMessageContentParts = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AIGuardMessageContent) MarshalJSON() ([]byte, error) {
	if obj.AIGuardMessageContentText != nil {
		return datadog.Marshal(&obj.AIGuardMessageContentText)
	}

	if obj.AIGuardMessageContentParts != nil {
		return datadog.Marshal(&obj.AIGuardMessageContentParts)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AIGuardMessageContent) GetActualInstance() interface{} {
	if obj.AIGuardMessageContentText != nil {
		return obj.AIGuardMessageContentText
	}

	if obj.AIGuardMessageContentParts != nil {
		return obj.AIGuardMessageContentParts
	}

	// all schemas are nil
	return nil
}

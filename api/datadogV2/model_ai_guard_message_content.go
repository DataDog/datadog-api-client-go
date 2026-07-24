// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardMessageContent - The message content, either a plain string or an array of content parts.
type AIGuardMessageContent struct {
	String                 *string
	AIGuardContentPartList *[]AIGuardContentPart

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsAIGuardMessageContent is a convenience function that returns string wrapped in AIGuardMessageContent.
func StringAsAIGuardMessageContent(v *string) AIGuardMessageContent {
	return AIGuardMessageContent{String: v}
}

// AIGuardContentPartListAsAIGuardMessageContent is a convenience function that returns []AIGuardContentPart wrapped in AIGuardMessageContent.
func AIGuardContentPartListAsAIGuardMessageContent(v *[]AIGuardContentPart) AIGuardMessageContent {
	return AIGuardMessageContent{AIGuardContentPartList: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AIGuardMessageContent) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = datadog.Unmarshal(data, &obj.String)
	if err == nil {
		if obj.String != nil {
			jsonString, _ := datadog.Marshal(obj.String)
			if string(jsonString) == "{}" { // empty struct
				obj.String = nil
			} else {
				match++
			}
		} else {
			obj.String = nil
		}
	} else {
		obj.String = nil
	}

	// try to unmarshal data into AIGuardContentPartList
	err = datadog.Unmarshal(data, &obj.AIGuardContentPartList)
	if err == nil {
		if obj.AIGuardContentPartList != nil {
			jsonAIGuardContentPartList, _ := datadog.Marshal(obj.AIGuardContentPartList)
			if string(jsonAIGuardContentPartList) == "{}" && string(data) != "{}" { // empty struct
				obj.AIGuardContentPartList = nil
			} else {
				match++
			}
		} else {
			obj.AIGuardContentPartList = nil
		}
	} else {
		obj.AIGuardContentPartList = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.String = nil
		obj.AIGuardContentPartList = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AIGuardMessageContent) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}

	if obj.AIGuardContentPartList != nil {
		return datadog.Marshal(&obj.AIGuardContentPartList)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AIGuardMessageContent) GetActualInstance() interface{} {
	if obj.String != nil {
		return obj.String
	}

	if obj.AIGuardContentPartList != nil {
		return obj.AIGuardContentPartList
	}

	// all schemas are nil
	return nil
}

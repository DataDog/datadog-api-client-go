// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptTemplate - A text template or a list of chat messages.
type LLMObsPromptTemplate struct {
	LLMObsPromptTextTemplate *string
	LLMObsPromptChatTemplate *LLMObsPromptChatTemplate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LLMObsPromptTextTemplateAsLLMObsPromptTemplate is a convenience function that returns string wrapped in LLMObsPromptTemplate.
func LLMObsPromptTextTemplateAsLLMObsPromptTemplate(v *string) LLMObsPromptTemplate {
	return LLMObsPromptTemplate{LLMObsPromptTextTemplate: v}
}

// LLMObsPromptChatTemplateAsLLMObsPromptTemplate is a convenience function that returns LLMObsPromptChatTemplate wrapped in LLMObsPromptTemplate.
func LLMObsPromptChatTemplateAsLLMObsPromptTemplate(v *LLMObsPromptChatTemplate) LLMObsPromptTemplate {
	return LLMObsPromptTemplate{LLMObsPromptChatTemplate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LLMObsPromptTemplate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LLMObsPromptTextTemplate
	err = datadog.Unmarshal(data, &obj.LLMObsPromptTextTemplate)
	if err == nil {
		if obj.LLMObsPromptTextTemplate != nil {
			jsonLLMObsPromptTextTemplate, _ := datadog.Marshal(obj.LLMObsPromptTextTemplate)
			if string(jsonLLMObsPromptTextTemplate) == "{}" { // empty struct
				obj.LLMObsPromptTextTemplate = nil
			} else {
				match++
			}
		} else {
			obj.LLMObsPromptTextTemplate = nil
		}
	} else {
		obj.LLMObsPromptTextTemplate = nil
	}

	// try to unmarshal data into LLMObsPromptChatTemplate
	err = datadog.Unmarshal(data, &obj.LLMObsPromptChatTemplate)
	if err == nil {
		if obj.LLMObsPromptChatTemplate != nil {
			jsonLLMObsPromptChatTemplate, _ := datadog.Marshal(obj.LLMObsPromptChatTemplate)
			if string(jsonLLMObsPromptChatTemplate) == "{}" && string(data) != "{}" { // empty struct
				obj.LLMObsPromptChatTemplate = nil
			} else {
				match++
			}
		} else {
			obj.LLMObsPromptChatTemplate = nil
		}
	} else {
		obj.LLMObsPromptChatTemplate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LLMObsPromptTextTemplate = nil
		obj.LLMObsPromptChatTemplate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LLMObsPromptTemplate) MarshalJSON() ([]byte, error) {
	if obj.LLMObsPromptTextTemplate != nil {
		return datadog.Marshal(&obj.LLMObsPromptTextTemplate)
	}

	if obj.LLMObsPromptChatTemplate != nil {
		return datadog.Marshal(&obj.LLMObsPromptChatTemplate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LLMObsPromptTemplate) GetActualInstance() interface{} {
	if obj.LLMObsPromptTextTemplate != nil {
		return obj.LLMObsPromptTextTemplate
	}

	if obj.LLMObsPromptChatTemplate != nil {
		return obj.LLMObsPromptChatTemplate
	}

	// all schemas are nil
	return nil
}

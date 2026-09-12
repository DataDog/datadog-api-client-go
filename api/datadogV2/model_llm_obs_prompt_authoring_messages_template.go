// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptAuthoringMessagesTemplate A chat prompt whose authored items are stored under `messages`.
type LLMObsPromptAuthoringMessagesTemplate struct {
	// A chat prompt containing messages, pinned includes, or both.
	Messages LLMObsPromptAuthoringChatTemplate `json:"messages"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPromptAuthoringMessagesTemplate instantiates a new LLMObsPromptAuthoringMessagesTemplate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPromptAuthoringMessagesTemplate(messages LLMObsPromptAuthoringChatTemplate) *LLMObsPromptAuthoringMessagesTemplate {
	this := LLMObsPromptAuthoringMessagesTemplate{}
	this.Messages = messages
	return &this
}

// NewLLMObsPromptAuthoringMessagesTemplateWithDefaults instantiates a new LLMObsPromptAuthoringMessagesTemplate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPromptAuthoringMessagesTemplateWithDefaults() *LLMObsPromptAuthoringMessagesTemplate {
	this := LLMObsPromptAuthoringMessagesTemplate{}
	return &this
}

// GetMessages returns the Messages field value.
func (o *LLMObsPromptAuthoringMessagesTemplate) GetMessages() LLMObsPromptAuthoringChatTemplate {
	if o == nil {
		var ret LLMObsPromptAuthoringChatTemplate
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptAuthoringMessagesTemplate) GetMessagesOk() (*LLMObsPromptAuthoringChatTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Messages, true
}

// SetMessages sets field value.
func (o *LLMObsPromptAuthoringMessagesTemplate) SetMessages(v LLMObsPromptAuthoringChatTemplate) {
	o.Messages = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPromptAuthoringMessagesTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["messages"] = o.Messages

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPromptAuthoringMessagesTemplate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Messages *LLMObsPromptAuthoringChatTemplate `json:"messages"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Messages == nil {
		return fmt.Errorf("required field messages missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"messages"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Messages.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Messages = *all.Messages

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptChatMessage A single chat message in a prompt template.
type LLMObsPromptChatMessage struct {
	// Content of the message.
	Content string `json:"content"`
	// Role of the message (for example `system`, `user`, or `assistant`).
	Role string `json:"role"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPromptChatMessage instantiates a new LLMObsPromptChatMessage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPromptChatMessage(content string, role string) *LLMObsPromptChatMessage {
	this := LLMObsPromptChatMessage{}
	this.Content = content
	this.Role = role
	return &this
}

// NewLLMObsPromptChatMessageWithDefaults instantiates a new LLMObsPromptChatMessage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPromptChatMessageWithDefaults() *LLMObsPromptChatMessage {
	this := LLMObsPromptChatMessage{}
	return &this
}

// GetContent returns the Content field value.
func (o *LLMObsPromptChatMessage) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptChatMessage) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *LLMObsPromptChatMessage) SetContent(v string) {
	o.Content = v
}

// GetRole returns the Role field value.
func (o *LLMObsPromptChatMessage) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptChatMessage) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *LLMObsPromptChatMessage) SetRole(v string) {
	o.Role = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPromptChatMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["content"] = o.Content
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPromptChatMessage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content *string `json:"content"`
		Role    *string `json:"role"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content", "role"})
	} else {
		return err
	}
	o.Content = *all.Content
	o.Role = *all.Role

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

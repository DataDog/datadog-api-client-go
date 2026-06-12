// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardMessage A single message in the conversation to evaluate.
type AIGuardMessage struct {
	// The message content, either a plain string or an array of content parts.
	Content *AIGuardMessageContent `json:"content,omitempty"`
	// The role of the message author in the conversation.
	Role AIGuardMessageRole `json:"role"`
	// The ID of the tool call this message is responding to, required for tool messages.
	ToolCallId *string `json:"tool_call_id,omitempty"`
	// Tool calls issued by the assistant in this message.
	ToolCalls []AIGuardToolCall `json:"tool_calls,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAIGuardMessage instantiates a new AIGuardMessage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIGuardMessage(role AIGuardMessageRole) *AIGuardMessage {
	this := AIGuardMessage{}
	this.Role = role
	return &this
}

// NewAIGuardMessageWithDefaults instantiates a new AIGuardMessage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIGuardMessageWithDefaults() *AIGuardMessage {
	this := AIGuardMessage{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AIGuardMessage) GetContent() AIGuardMessageContent {
	if o == nil || o.Content == nil {
		var ret AIGuardMessageContent
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardMessage) GetContentOk() (*AIGuardMessageContent, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AIGuardMessage) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given AIGuardMessageContent and assigns it to the Content field.
func (o *AIGuardMessage) SetContent(v AIGuardMessageContent) {
	o.Content = &v
}

// GetRole returns the Role field value.
func (o *AIGuardMessage) GetRole() AIGuardMessageRole {
	if o == nil {
		var ret AIGuardMessageRole
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *AIGuardMessage) GetRoleOk() (*AIGuardMessageRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *AIGuardMessage) SetRole(v AIGuardMessageRole) {
	o.Role = v
}

// GetToolCallId returns the ToolCallId field value if set, zero value otherwise.
func (o *AIGuardMessage) GetToolCallId() string {
	if o == nil || o.ToolCallId == nil {
		var ret string
		return ret
	}
	return *o.ToolCallId
}

// GetToolCallIdOk returns a tuple with the ToolCallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardMessage) GetToolCallIdOk() (*string, bool) {
	if o == nil || o.ToolCallId == nil {
		return nil, false
	}
	return o.ToolCallId, true
}

// HasToolCallId returns a boolean if a field has been set.
func (o *AIGuardMessage) HasToolCallId() bool {
	return o != nil && o.ToolCallId != nil
}

// SetToolCallId gets a reference to the given string and assigns it to the ToolCallId field.
func (o *AIGuardMessage) SetToolCallId(v string) {
	o.ToolCallId = &v
}

// GetToolCalls returns the ToolCalls field value if set, zero value otherwise.
func (o *AIGuardMessage) GetToolCalls() []AIGuardToolCall {
	if o == nil || o.ToolCalls == nil {
		var ret []AIGuardToolCall
		return ret
	}
	return o.ToolCalls
}

// GetToolCallsOk returns a tuple with the ToolCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardMessage) GetToolCallsOk() (*[]AIGuardToolCall, bool) {
	if o == nil || o.ToolCalls == nil {
		return nil, false
	}
	return &o.ToolCalls, true
}

// HasToolCalls returns a boolean if a field has been set.
func (o *AIGuardMessage) HasToolCalls() bool {
	return o != nil && o.ToolCalls != nil
}

// SetToolCalls gets a reference to the given []AIGuardToolCall and assigns it to the ToolCalls field.
func (o *AIGuardMessage) SetToolCalls(v []AIGuardToolCall) {
	o.ToolCalls = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIGuardMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	toSerialize["role"] = o.Role
	if o.ToolCallId != nil {
		toSerialize["tool_call_id"] = o.ToolCallId
	}
	if o.ToolCalls != nil {
		toSerialize["tool_calls"] = o.ToolCalls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIGuardMessage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content    *AIGuardMessageContent `json:"content,omitempty"`
		Role       *AIGuardMessageRole    `json:"role"`
		ToolCallId *string                `json:"tool_call_id,omitempty"`
		ToolCalls  []AIGuardToolCall      `json:"tool_calls,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content", "role", "tool_call_id", "tool_calls"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Content = all.Content
	if !all.Role.IsValid() {
		hasInvalidField = true
	} else {
		o.Role = *all.Role
	}
	o.ToolCallId = all.ToolCallId
	o.ToolCalls = all.ToolCalls

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

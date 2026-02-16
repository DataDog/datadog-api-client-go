// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChatMessage
type ChatMessage struct {
	// The chat ID.
	ChatId string `json:"chatId"`
	// The message content.
	Content string `json:"content"`
	// The message ID.
	Id string `json:"id"`
	// The role of the message sender.
	Role ChatMessageRole `json:"role"`
	// The UUID of the user who sent the message.
	UserUuid *string `json:"userUuid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChatMessage instantiates a new ChatMessage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChatMessage(chatId string, content string, id string, role ChatMessageRole) *ChatMessage {
	this := ChatMessage{}
	this.ChatId = chatId
	this.Content = content
	this.Id = id
	this.Role = role
	return &this
}

// NewChatMessageWithDefaults instantiates a new ChatMessage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChatMessageWithDefaults() *ChatMessage {
	this := ChatMessage{}
	return &this
}

// GetChatId returns the ChatId field value.
func (o *ChatMessage) GetChatId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *ChatMessage) GetChatIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value.
func (o *ChatMessage) SetChatId(v string) {
	o.ChatId = v
}

// GetContent returns the Content field value.
func (o *ChatMessage) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ChatMessage) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *ChatMessage) SetContent(v string) {
	o.Content = v
}

// GetId returns the Id field value.
func (o *ChatMessage) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChatMessage) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ChatMessage) SetId(v string) {
	o.Id = v
}

// GetRole returns the Role field value.
func (o *ChatMessage) GetRole() ChatMessageRole {
	if o == nil {
		var ret ChatMessageRole
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ChatMessage) GetRoleOk() (*ChatMessageRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *ChatMessage) SetRole(v ChatMessageRole) {
	o.Role = v
}

// GetUserUuid returns the UserUuid field value if set, zero value otherwise.
func (o *ChatMessage) GetUserUuid() string {
	if o == nil || o.UserUuid == nil {
		var ret string
		return ret
	}
	return *o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMessage) GetUserUuidOk() (*string, bool) {
	if o == nil || o.UserUuid == nil {
		return nil, false
	}
	return o.UserUuid, true
}

// HasUserUuid returns a boolean if a field has been set.
func (o *ChatMessage) HasUserUuid() bool {
	return o != nil && o.UserUuid != nil
}

// SetUserUuid gets a reference to the given string and assigns it to the UserUuid field.
func (o *ChatMessage) SetUserUuid(v string) {
	o.UserUuid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChatMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["chatId"] = o.ChatId
	toSerialize["content"] = o.Content
	toSerialize["id"] = o.Id
	toSerialize["role"] = o.Role
	if o.UserUuid != nil {
		toSerialize["userUuid"] = o.UserUuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChatMessage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChatId   *string          `json:"chatId"`
		Content  *string          `json:"content"`
		Id       *string          `json:"id"`
		Role     *ChatMessageRole `json:"role"`
		UserUuid *string          `json:"userUuid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ChatId == nil {
		return fmt.Errorf("required field chatId missing")
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"chatId", "content", "id", "role", "userUuid"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ChatId = *all.ChatId
	o.Content = *all.Content
	o.Id = *all.Id
	if !all.Role.IsValid() {
		hasInvalidField = true
	} else {
		o.Role = *all.Role
	}
	o.UserUuid = all.UserUuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

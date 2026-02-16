// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChatHistoryItem
type ChatHistoryItem struct {
	// The message content.
	Content string `json:"content"`
	// The role of the message sender.
	Role ChatHistoryItemRole `json:"role"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChatHistoryItem instantiates a new ChatHistoryItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChatHistoryItem(content string, role ChatHistoryItemRole) *ChatHistoryItem {
	this := ChatHistoryItem{}
	this.Content = content
	this.Role = role
	return &this
}

// NewChatHistoryItemWithDefaults instantiates a new ChatHistoryItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChatHistoryItemWithDefaults() *ChatHistoryItem {
	this := ChatHistoryItem{}
	return &this
}

// GetContent returns the Content field value.
func (o *ChatHistoryItem) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ChatHistoryItem) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *ChatHistoryItem) SetContent(v string) {
	o.Content = v
}

// GetRole returns the Role field value.
func (o *ChatHistoryItem) GetRole() ChatHistoryItemRole {
	if o == nil {
		var ret ChatHistoryItemRole
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ChatHistoryItem) GetRoleOk() (*ChatHistoryItemRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *ChatHistoryItem) SetRole(v ChatHistoryItemRole) {
	o.Role = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChatHistoryItem) MarshalJSON() ([]byte, error) {
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
func (o *ChatHistoryItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content *string              `json:"content"`
		Role    *ChatHistoryItemRole `json:"role"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content", "role"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Content = *all.Content
	if !all.Role.IsValid() {
		hasInvalidField = true
	} else {
		o.Role = *all.Role
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

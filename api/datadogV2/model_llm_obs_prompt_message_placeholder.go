// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptMessagePlaceholder A named placeholder that inserts a list of messages when a compatible SDK formats the prompt.
// **Preview:** Message placeholders are available in Preview. To request access, contact [Datadog Support](https://www.datadoghq.com/support/) or your Customer Success Manager.
type LLMObsPromptMessagePlaceholder struct {
	// Name used to supply the message list when formatting the prompt.
	Name string `json:"name"`
	// Type of chat-template item.
	Type LLMObsPromptMessagePlaceholderType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewLLMObsPromptMessagePlaceholder instantiates a new LLMObsPromptMessagePlaceholder object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPromptMessagePlaceholder(name string, typeVar LLMObsPromptMessagePlaceholderType) *LLMObsPromptMessagePlaceholder {
	this := LLMObsPromptMessagePlaceholder{}
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewLLMObsPromptMessagePlaceholderWithDefaults instantiates a new LLMObsPromptMessagePlaceholder object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPromptMessagePlaceholderWithDefaults() *LLMObsPromptMessagePlaceholder {
	this := LLMObsPromptMessagePlaceholder{}
	return &this
}

// GetName returns the Name field value.
func (o *LLMObsPromptMessagePlaceholder) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptMessagePlaceholder) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsPromptMessagePlaceholder) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *LLMObsPromptMessagePlaceholder) GetType() LLMObsPromptMessagePlaceholderType {
	if o == nil {
		var ret LLMObsPromptMessagePlaceholderType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptMessagePlaceholder) GetTypeOk() (*LLMObsPromptMessagePlaceholderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsPromptMessagePlaceholder) SetType(v LLMObsPromptMessagePlaceholderType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPromptMessagePlaceholder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPromptMessagePlaceholder) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string                             `json:"name"`
		Type *LLMObsPromptMessagePlaceholderType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Name = *all.Name
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

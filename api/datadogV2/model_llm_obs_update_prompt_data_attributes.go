// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsUpdatePromptDataAttributes Attributes for updating an LLM Observability prompt. At least one of `title` or `description` must be provided; both attributes are optional individually.
type LLMObsUpdatePromptDataAttributes struct {
	// Optional new description for the prompt.
	Description *string `json:"description,omitempty"`
	// Optional new title for the prompt.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewLLMObsUpdatePromptDataAttributes instantiates a new LLMObsUpdatePromptDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsUpdatePromptDataAttributes() *LLMObsUpdatePromptDataAttributes {
	this := LLMObsUpdatePromptDataAttributes{}
	return &this
}

// NewLLMObsUpdatePromptDataAttributesWithDefaults instantiates a new LLMObsUpdatePromptDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsUpdatePromptDataAttributesWithDefaults() *LLMObsUpdatePromptDataAttributes {
	this := LLMObsUpdatePromptDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LLMObsUpdatePromptDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsUpdatePromptDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LLMObsUpdatePromptDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LLMObsUpdatePromptDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *LLMObsUpdatePromptDataAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsUpdatePromptDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *LLMObsUpdatePromptDataAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *LLMObsUpdatePromptDataAttributes) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsUpdatePromptDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsUpdatePromptDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string `json:"description,omitempty"`
		Title       *string `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Description = all.Description
	o.Title = all.Title

	return nil
}

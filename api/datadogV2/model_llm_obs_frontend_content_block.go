// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsFrontendContentBlock Validation requirements for a `frontend` display block.
type LLMObsFrontendContentBlock struct {
	// HTML code rendered by a `frontend` block. Required for `frontend` blocks.
	Code string `json:"code"`
	// Type discriminator for a `frontend` display block.
	Type *LLMObsFrontendContentBlockType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsFrontendContentBlock instantiates a new LLMObsFrontendContentBlock object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsFrontendContentBlock(code string) *LLMObsFrontendContentBlock {
	this := LLMObsFrontendContentBlock{}
	this.Code = code
	return &this
}

// NewLLMObsFrontendContentBlockWithDefaults instantiates a new LLMObsFrontendContentBlock object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsFrontendContentBlockWithDefaults() *LLMObsFrontendContentBlock {
	this := LLMObsFrontendContentBlock{}
	return &this
}

// GetCode returns the Code field value.
func (o *LLMObsFrontendContentBlock) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *LLMObsFrontendContentBlock) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *LLMObsFrontendContentBlock) SetCode(v string) {
	o.Code = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LLMObsFrontendContentBlock) GetType() LLMObsFrontendContentBlockType {
	if o == nil || o.Type == nil {
		var ret LLMObsFrontendContentBlockType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsFrontendContentBlock) GetTypeOk() (*LLMObsFrontendContentBlockType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LLMObsFrontendContentBlock) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given LLMObsFrontendContentBlockType and assigns it to the Type field.
func (o *LLMObsFrontendContentBlock) SetType(v LLMObsFrontendContentBlockType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsFrontendContentBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["code"] = o.Code
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsFrontendContentBlock) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code *string                         `json:"code"`
		Type *LLMObsFrontendContentBlockType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Code = *all.Code
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsFrontendContent Web content that makes up a `frontend` interaction.
type LLMObsFrontendContent struct {
	// Caller-provided web content.
	Code string `json:"code"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsFrontendContent instantiates a new LLMObsFrontendContent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsFrontendContent(code string) *LLMObsFrontendContent {
	this := LLMObsFrontendContent{}
	this.Code = code
	return &this
}

// NewLLMObsFrontendContentWithDefaults instantiates a new LLMObsFrontendContent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsFrontendContentWithDefaults() *LLMObsFrontendContent {
	this := LLMObsFrontendContent{}
	return &this
}

// GetCode returns the Code field value.
func (o *LLMObsFrontendContent) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *LLMObsFrontendContent) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *LLMObsFrontendContent) SetCode(v string) {
	o.Code = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsFrontendContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["code"] = o.Code

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsFrontendContent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code *string `json:"code"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code"})
	} else {
		return err
	}
	o.Code = *all.Code

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsTemplatesResponseAttributes Attributes of a list of Agent Observability patterns templates.
type LLMObsPatternsTemplatesResponseAttributes struct {
	// List of patterns templates.
	Templates []LLMObsPatternsTemplateItem `json:"templates"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPatternsTemplatesResponseAttributes instantiates a new LLMObsPatternsTemplatesResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPatternsTemplatesResponseAttributes(templates []LLMObsPatternsTemplateItem) *LLMObsPatternsTemplatesResponseAttributes {
	this := LLMObsPatternsTemplatesResponseAttributes{}
	this.Templates = templates
	return &this
}

// NewLLMObsPatternsTemplatesResponseAttributesWithDefaults instantiates a new LLMObsPatternsTemplatesResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPatternsTemplatesResponseAttributesWithDefaults() *LLMObsPatternsTemplatesResponseAttributes {
	this := LLMObsPatternsTemplatesResponseAttributes{}
	return &this
}

// GetTemplates returns the Templates field value.
func (o *LLMObsPatternsTemplatesResponseAttributes) GetTemplates() []LLMObsPatternsTemplateItem {
	if o == nil {
		var ret []LLMObsPatternsTemplateItem
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTemplatesResponseAttributes) GetTemplatesOk() (*[]LLMObsPatternsTemplateItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Templates, true
}

// SetTemplates sets field value.
func (o *LLMObsPatternsTemplatesResponseAttributes) SetTemplates(v []LLMObsPatternsTemplateItem) {
	o.Templates = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPatternsTemplatesResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["templates"] = o.Templates

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPatternsTemplatesResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Templates *[]LLMObsPatternsTemplateItem `json:"templates"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Templates == nil {
		return fmt.Errorf("required field templates missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"templates"})
	} else {
		return err
	}
	o.Templates = *all.Templates

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

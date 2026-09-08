// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsLegacyContentBlock Validation requirements for an existing non-frontend display block.
type LLMObsLegacyContentBlock struct {
	// Type discriminator for an existing non-frontend display block.
	Type *LLMObsLegacyContentBlockType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsLegacyContentBlock instantiates a new LLMObsLegacyContentBlock object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsLegacyContentBlock() *LLMObsLegacyContentBlock {
	this := LLMObsLegacyContentBlock{}
	return &this
}

// NewLLMObsLegacyContentBlockWithDefaults instantiates a new LLMObsLegacyContentBlock object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsLegacyContentBlockWithDefaults() *LLMObsLegacyContentBlock {
	this := LLMObsLegacyContentBlock{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LLMObsLegacyContentBlock) GetType() LLMObsLegacyContentBlockType {
	if o == nil || o.Type == nil {
		var ret LLMObsLegacyContentBlockType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsLegacyContentBlock) GetTypeOk() (*LLMObsLegacyContentBlockType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LLMObsLegacyContentBlock) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given LLMObsLegacyContentBlockType and assigns it to the Type field.
func (o *LLMObsLegacyContentBlock) SetType(v LLMObsLegacyContentBlockType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsLegacyContentBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsLegacyContentBlock) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type *LLMObsLegacyContentBlockType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"type"})
	} else {
		return err
	}

	hasInvalidField := false
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

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomCostsFileUpdateAttributes Attributes of a Custom Costs file that can be updated.
type CustomCostsFileUpdateAttributes struct {
	// Whether the Custom Costs file contains the costs of an AI provider.
	AiProvider bool `json:"ai_provider"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomCostsFileUpdateAttributes instantiates a new CustomCostsFileUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomCostsFileUpdateAttributes(aiProvider bool) *CustomCostsFileUpdateAttributes {
	this := CustomCostsFileUpdateAttributes{}
	this.AiProvider = aiProvider
	return &this
}

// NewCustomCostsFileUpdateAttributesWithDefaults instantiates a new CustomCostsFileUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomCostsFileUpdateAttributesWithDefaults() *CustomCostsFileUpdateAttributes {
	this := CustomCostsFileUpdateAttributes{}
	return &this
}

// GetAiProvider returns the AiProvider field value.
func (o *CustomCostsFileUpdateAttributes) GetAiProvider() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AiProvider
}

// GetAiProviderOk returns a tuple with the AiProvider field value
// and a boolean to check if the value has been set.
func (o *CustomCostsFileUpdateAttributes) GetAiProviderOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AiProvider, true
}

// SetAiProvider sets field value.
func (o *CustomCostsFileUpdateAttributes) SetAiProvider(v bool) {
	o.AiProvider = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomCostsFileUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ai_provider"] = o.AiProvider

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomCostsFileUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AiProvider *bool `json:"ai_provider"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AiProvider == nil {
		return fmt.Errorf("required field ai_provider missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ai_provider"})
	} else {
		return err
	}
	o.AiProvider = *all.AiProvider

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

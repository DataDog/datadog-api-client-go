// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingAttributes Attributes of a model pricing resource.
type AiGatewayModelPricingAttributes struct {
	// Effective source of the model pricing data.
	EffectiveSource *AiGatewayModelPricingEffectiveSource `json:"effective_source,omitempty"`
	// Pricing configuration for a model.
	PricingConfig *AiGatewayModelPricingConfig `json:"pricing_config,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelPricingAttributes instantiates a new AiGatewayModelPricingAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelPricingAttributes() *AiGatewayModelPricingAttributes {
	this := AiGatewayModelPricingAttributes{}
	return &this
}

// NewAiGatewayModelPricingAttributesWithDefaults instantiates a new AiGatewayModelPricingAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelPricingAttributesWithDefaults() *AiGatewayModelPricingAttributes {
	this := AiGatewayModelPricingAttributes{}
	return &this
}

// GetEffectiveSource returns the EffectiveSource field value if set, zero value otherwise.
func (o *AiGatewayModelPricingAttributes) GetEffectiveSource() AiGatewayModelPricingEffectiveSource {
	if o == nil || o.EffectiveSource == nil {
		var ret AiGatewayModelPricingEffectiveSource
		return ret
	}
	return *o.EffectiveSource
}

// GetEffectiveSourceOk returns a tuple with the EffectiveSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingAttributes) GetEffectiveSourceOk() (*AiGatewayModelPricingEffectiveSource, bool) {
	if o == nil || o.EffectiveSource == nil {
		return nil, false
	}
	return o.EffectiveSource, true
}

// HasEffectiveSource returns a boolean if a field has been set.
func (o *AiGatewayModelPricingAttributes) HasEffectiveSource() bool {
	return o != nil && o.EffectiveSource != nil
}

// SetEffectiveSource gets a reference to the given AiGatewayModelPricingEffectiveSource and assigns it to the EffectiveSource field.
func (o *AiGatewayModelPricingAttributes) SetEffectiveSource(v AiGatewayModelPricingEffectiveSource) {
	o.EffectiveSource = &v
}

// GetPricingConfig returns the PricingConfig field value if set, zero value otherwise.
func (o *AiGatewayModelPricingAttributes) GetPricingConfig() AiGatewayModelPricingConfig {
	if o == nil || o.PricingConfig == nil {
		var ret AiGatewayModelPricingConfig
		return ret
	}
	return *o.PricingConfig
}

// GetPricingConfigOk returns a tuple with the PricingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingAttributes) GetPricingConfigOk() (*AiGatewayModelPricingConfig, bool) {
	if o == nil || o.PricingConfig == nil {
		return nil, false
	}
	return o.PricingConfig, true
}

// HasPricingConfig returns a boolean if a field has been set.
func (o *AiGatewayModelPricingAttributes) HasPricingConfig() bool {
	return o != nil && o.PricingConfig != nil
}

// SetPricingConfig gets a reference to the given AiGatewayModelPricingConfig and assigns it to the PricingConfig field.
func (o *AiGatewayModelPricingAttributes) SetPricingConfig(v AiGatewayModelPricingConfig) {
	o.PricingConfig = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelPricingAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EffectiveSource != nil {
		toSerialize["effective_source"] = o.EffectiveSource
	}
	if o.PricingConfig != nil {
		toSerialize["pricing_config"] = o.PricingConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelPricingAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EffectiveSource *AiGatewayModelPricingEffectiveSource `json:"effective_source,omitempty"`
		PricingConfig   *AiGatewayModelPricingConfig          `json:"pricing_config,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"effective_source", "pricing_config"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.EffectiveSource != nil && !all.EffectiveSource.IsValid() {
		hasInvalidField = true
	} else {
		o.EffectiveSource = all.EffectiveSource
	}
	if all.PricingConfig != nil && all.PricingConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PricingConfig = all.PricingConfig

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

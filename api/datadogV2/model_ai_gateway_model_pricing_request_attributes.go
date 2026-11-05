// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingRequestAttributes Attributes for a model pricing request.
type AiGatewayModelPricingRequestAttributes struct {
	// Pricing configuration accepted when setting model pricing.
	PricingConfig AiGatewayModelPricingRequestConfig `json:"pricing_config"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelPricingRequestAttributes instantiates a new AiGatewayModelPricingRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelPricingRequestAttributes(pricingConfig AiGatewayModelPricingRequestConfig) *AiGatewayModelPricingRequestAttributes {
	this := AiGatewayModelPricingRequestAttributes{}
	this.PricingConfig = pricingConfig
	return &this
}

// NewAiGatewayModelPricingRequestAttributesWithDefaults instantiates a new AiGatewayModelPricingRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelPricingRequestAttributesWithDefaults() *AiGatewayModelPricingRequestAttributes {
	this := AiGatewayModelPricingRequestAttributes{}
	return &this
}

// GetPricingConfig returns the PricingConfig field value.
func (o *AiGatewayModelPricingRequestAttributes) GetPricingConfig() AiGatewayModelPricingRequestConfig {
	if o == nil {
		var ret AiGatewayModelPricingRequestConfig
		return ret
	}
	return o.PricingConfig
}

// GetPricingConfigOk returns a tuple with the PricingConfig field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingRequestAttributes) GetPricingConfigOk() (*AiGatewayModelPricingRequestConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PricingConfig, true
}

// SetPricingConfig sets field value.
func (o *AiGatewayModelPricingRequestAttributes) SetPricingConfig(v AiGatewayModelPricingRequestConfig) {
	o.PricingConfig = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelPricingRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["pricing_config"] = o.PricingConfig

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelPricingRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PricingConfig *AiGatewayModelPricingRequestConfig `json:"pricing_config"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PricingConfig == nil {
		return fmt.Errorf("required field pricing_config missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"pricing_config"})
	} else {
		return err
	}
	o.PricingConfig = *all.PricingConfig

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
